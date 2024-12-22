package main

import (
	"context"
	pb "grpc-go-calculator/calculator/proto"
	"io"
	"log"
)

func doPrime(c pb.CalculatorServiceClient) {
	log.Println("---DoPrime is invoked-----")
	req := &pb.PrimeRequest{
		Number: 12390392840,
	}

	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Primes: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res.Result)
	}
}
