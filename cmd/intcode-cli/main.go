package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"text/scanner"

	"google.golang.org/grpc"

	intcode "github.com/mjm/advent-of-code-2019/pkg/intcode/proto"
)

var (
	server = flag.String("server", "0.0.0.0:8080", "Host and port to use to connect to server")
)

var client intcode.IntcodeClient

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	in, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	memory := loadProgram(flag.Arg(0), in)
	client = intcode.NewIntcodeClient(conn)

	id, err := createVM(memory)
	if err != nil {
		log.Fatal(err)
	}

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Printf("Exiting...")
		conn.Close()
		os.Exit(1)
	}()

	runVMRepl(id)
}

func loadProgram(filename string, r io.Reader) []int64 {
	var memory []int64
	var s scanner.Scanner
	s.Init(r)
	s.Filename = flag.Arg(0)
	s.Mode = scanner.ScanInts
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		if tok == scanner.Int {
			n, err := strconv.ParseInt(s.TokenText(), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			memory = append(memory, n)
		}
	}
	return memory
}

func createVM(memory []int64) (string, error) {
	req := &intcode.CreateVMRequest{
		Memory: memory,
	}
	res, err := client.CreateVM(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("error creating intcode VM: %w", err)
	}

	return res.GetId(), nil
}

func runVMRepl(id string) error {
	stream, err := client.RunVM(context.Background())
	if err != nil {
		return fmt.Errorf("error trying to run VM: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Running Intcode program...\n")
	startReq := &intcode.RunVMRequest{
		Command: &intcode.RunVMRequest_Start{
			Start: &intcode.StartVMCommand{Id: id},
		},
	}
	if err := stream.Send(startReq); err != nil {
		return fmt.Errorf("error starting VM: %w", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("error waiting for response from intcode server: %w", err)
		}

		switch res.GetType() {
		case intcode.RunVMResponse_HALT:
			stream.CloseSend()
			return nil
		case intcode.RunVMResponse_OUTPUT:
			fmt.Printf("%d\n", res.GetOutput())
		case intcode.RunVMResponse_NEED_INPUT:
			var n int64
			for {
				fmt.Fprintf(os.Stderr, "> ")
				_, err := fmt.Scanf("%d\n", &n)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
				} else {
					break
				}
			}

			sendReq := &intcode.RunVMRequest{
				Command: &intcode.RunVMRequest_SendInput{
					SendInput: &intcode.SendInputCommand{Value: n},
				},
			}
			if err := stream.Send(sendReq); err != nil {
				return fmt.Errorf("error sending input to the VM: %w", err)
			}
		}
	}
}
