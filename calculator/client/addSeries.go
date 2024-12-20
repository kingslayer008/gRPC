package main

import (
	"context"
	pb "grpc-go-calculator/calculator/proto"
	"io"
	"log"
	"time"
)

func doAddSeries(c pb.CalculatorServiceClient) {
	log.Println("doAddSeries is invoked")

	stream, err := c.AddSeries(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []float64{4, 7, 2, 19, 4, 6, 32}
		for _, number := range numbers {
			log.Printf("Sending number: %d\n", number)
			stream.Send(&pb.AddSeriesRequest{
				Num1: number,
			})

			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Problem while reading server stream: %v\n", err)
				break
			}

			log.Printf("Received a new sum of...: %v\n", res.Sum)
		}
		close(waitc)
	}()
	<-waitc
}
