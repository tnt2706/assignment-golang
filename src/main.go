package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"assignment/src/pkg/pb"
)

var (
	GRPC_PORT = flag.Int("port", 6900, "The server port")
)

type server struct {
	pb.UnimplementedCalculatorServiceServer
}

func (s *server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	fmt.Print(in)
	return &pb.SumResponse{Result: in.FirstNumber + in.SecondNumber}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *GRPC_PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterCalculatorServiceServer(grpcServer, &server{})

	log.Printf("server listening at %v", lis.Addr())
	grpcServer.Serve(lis)
}
