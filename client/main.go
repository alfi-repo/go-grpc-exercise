package main

import (
	"context"
	"fmt"
	"go-grpc-exercise/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	serverAddress := "127.0.0.1:8080"

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddress, opts...)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	res, err := client.Add(ctx, &pb.CalculationRequest{A: 2, B: 1})
	if err != nil {
		log.Fatalln("Failed to call Add:", err)
	}
	fmt.Printf("Add: %d\n", res.Result)

	res, err = client.Divide(ctx, &pb.CalculationRequest{A: 2, B: 1})
	if err != nil {
		log.Fatalln("Failed to call Divide:", err)
	}
	fmt.Printf("Divide: %d\n", res.Result)

	res, err = client.Multiply(ctx, &pb.CalculationRequest{A: 2, B: 1})
	if err != nil {
		log.Fatalln("Failed to call Multiply:", err)
	}
	fmt.Printf("Multiply: %d\n", res.Result)

	res, err = client.Sum(ctx, &pb.NumbersRequest{Numbers: []int64{2, 1}})
	if err != nil {
		log.Fatalln("Failed to call Sum:", err)
	}
	fmt.Printf("Sum: %d\n", res.Result)
}
