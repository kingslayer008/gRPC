package main

import (
	pb "grpc-go-calculator/calculator/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

func (s *Server) AddSeries(stream grpc.BidiStreamingServer[pb.AddSeriesRequest, pb.AddResponse]) error {
	log.Println("AddSeries is invoked")
	var result float64 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Received the number: %v\n", req.Num1)

		result = result + req.Num1

		err = stream.Send(&pb.AddResponse{
			Sum: result,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}

		log.Printf("Sent the SUM as: %v\n", result)

	}

}
