package server

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"

	"github.com/mjm/advent-of-code-2019/pkg/intcode"
	pb "github.com/mjm/advent-of-code-2019/pkg/intcode/proto"
)

// IntcodeServer implements a server that can run Intcode VMs, taking
// input from and providing output to the client.
type IntcodeServer struct {
	vms  map[uuid.UUID]*intcode.VM
	lock sync.Mutex
}

// NewIntcodeServer creates a new Intcode server with an empty list of VMs.
func NewIntcodeServer() *IntcodeServer {
	return &IntcodeServer{
		vms: make(map[uuid.UUID]*intcode.VM),
	}
}

// CreateVM creates a new VM in-memory, and returns an ID to the client that
// it can use to interact with it later.
func (s *IntcodeServer) CreateVM(ctx context.Context, req *pb.CreateVMRequest) (*pb.CreateVMResponse, error) {
	id := uuid.New()

	memory := make([]int, 0, len(req.GetMemory()))
	for _, n := range req.GetMemory() {
		memory = append(memory, int(n))
	}
	vm := intcode.NewVM(memory)

	s.lock.Lock()
	s.vms[id] = vm
	s.lock.Unlock()

	return &pb.CreateVMResponse{
		Id: id.String(),
	}, nil
}

// RunVM runs a VM to completion. It uses a bidirectional stream to allow the
// server to send output and prompts for input back to the client and wait
// for inputs from the client.
func (s *IntcodeServer) RunVM(stream pb.Intcode_RunVMServer) error {
	in, err := stream.Recv()
	start := in.GetStart()
	if start == nil {
		return fmt.Errorf("expected first message to RunVM to be a start command")
	}

	id, err := uuid.Parse(start.GetId())
	if err != nil {
		return fmt.Errorf("vm id is not a valid UUID: %w", err)
	}

	var vm *intcode.VM
	s.lock.Lock()
	vm = s.vms[id]
	s.lock.Unlock()

	vm.SetInputFunc(func() int {
		// tell the client we need input
		resp := &pb.RunVMResponse{
			Type: pb.RunVMResponse_NEED_INPUT,
		}
		if err := stream.Send(resp); err != nil {
			// can't return from outer function from here
			log.Printf("error sending request for input: %v", err)
			return 0
		}

		in, err := stream.Recv()
		if err != nil {
			log.Printf("error receiving input value: %v", err)
			return 0
		}

		input := in.GetSendInput()
		if input == nil {
			log.Printf("expected input message to be a send input command")
			return 0
		}

		return int(input.GetValue())
	})

	done := make(chan error)
	go func() {
		done <- vm.Execute()
		s.deleteVM(id)
	}()

	for out := range vm.Output {
		resp := &pb.RunVMResponse{
			Type:   pb.RunVMResponse_OUTPUT,
			Output: int64(out),
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}

	resp := &pb.RunVMResponse{
		Type: pb.RunVMResponse_HALT,
	}
	if err := stream.Send(resp); err != nil {
		return err
	}

	err = <-done
	return err
}

func (s *IntcodeServer) deleteVM(id uuid.UUID) {
	s.lock.Lock()
	delete(s.vms, id)
	s.lock.Unlock()
}
