package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	intcode "github.com/mjm/advent-of-code-2019/pkg/intcode/proto"
	"github.com/mjm/advent-of-code-2019/pkg/intcode/server"
)

var (
	port = flag.Int("p", 8080, "Port to listen on for gRPC requests")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	intcode.RegisterIntcodeServer(grpcServer, server.NewIntcodeServer())

	grpcServer.Serve(lis)
}
