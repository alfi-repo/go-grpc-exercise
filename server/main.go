package main

import (
	"context"
	"go-grpc-exercise/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type grpcServer struct {
	pb.UnimplementedCalculatorServer
}

func (s *grpcServer) Add(ctx context.Context, req *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	return &pb.CalculationResponse{
		Result: req.A + req.B,
	}, nil
}

func (s *grpcServer) Divide(ctx context.Context, req *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	if req.B == 0 {
		return nil, status.Error(
			codes.InvalidArgument,
			"cannot divide by zero",
		)
	}

	return &pb.CalculationResponse{
		Result: req.A / req.B,
	}, nil
}

func (s *grpcServer) Multiply(ctx context.Context, req *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	return &pb.CalculationResponse{
		Result: req.A * req.B,
	}, nil
}

func (s *grpcServer) Sum(ctx context.Context, req *pb.NumbersRequest) (*pb.CalculationResponse, error) {
	var sum int64
	for _, num := range req.Numbers {
		sum += num
	}

	return &pb.CalculationResponse{
		Result: sum,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln("Failed to create a listener:", err)
	}

	server := grpc.NewServer()
	pb.RegisterCalculatorServer(server, &grpcServer{})
	if err = server.Serve(listener); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
