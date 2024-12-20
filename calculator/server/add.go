package main

import (
	"context"
	pb "grpc-go-calculator/calculator/proto"
)

func (s *Server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {

	return &pb.AddResponse{
		Sum: in.Num1 + in.Num2,
	}, nil
}
