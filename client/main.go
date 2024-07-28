package main

import (
	"context"
	"fmt"
	"log"

	"buf.build/gen/go/phong/concept/grpc/go/math/v1/mathv1grpc"
	mathv1 "buf.build/gen/go/phong/concept/protocolbuffers/go/math/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Hello, World!")

	conn, err := grpc.NewClient("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println(err)
	}
	client := mathv1grpc.NewMathServiceClient(conn)

	res, err := client.Add(context.Background(), &mathv1.AddRequest{Num1: 5, Num2: 10})

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res.Sum, "Successfully")

}
