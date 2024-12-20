package main

import (
	"context"
	"fmt"
	pb "grpc-go-calculator/calculator/proto"
	"log"
)

func addNums(c pb.CalculatorServiceClient, num1 float64, num2 float64) {

	log.Println("----addNumbs is invoked----")

	req := &pb.AddRequest{
		Num1: num1,
		Num2: num2,
	}

	res, err := c.Add(context.Background(), req)

	if err != nil {
		log.Fatalf("Error occured while getting respo %v\n", err)
	}

	fmt.Printf("The answer is %v\n", res)

}
