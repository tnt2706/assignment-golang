package calculator

import (
	"context"
	"log"
	"net"

	pb "assignment/pkg/calculator/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServiceServer
}

func (s *server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	result := in.GetFirstNumber() + in.GetSecondNumber()
	return &pb.SumResponse{Result: result}, nil
}

func StartGrpcServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:6900")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &server{})
	log.Printf("🚀 Running listening at %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
