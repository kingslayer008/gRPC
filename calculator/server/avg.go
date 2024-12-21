package main

import (
	pb "grpc-go-calculator/calculator/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

func (s *Server) Avg(stream grpc.ClientStreamingServer[pb.AvgRequest, pb.AvgResponse]) error {

	log.Println("Avg is invoked")

	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while receiving stream %v\n", err)
		}

		sum += req.Number
		count++
	}

}
