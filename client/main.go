package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"buf.build/gen/go/phong/concept/connectrpc/go/hello/v1/hellov1connect"
	hello "buf.build/gen/go/phong/concept/protocolbuffers/go/hello/v1"
	connect "connectrpc.com/connect"
)

func main() {
	fmt.Println("Hello, World!")
	client := hellov1connect.NewGreeterServiceClient(
		http.DefaultClient,
		"http://localhost:8888",
	)
	res, err := client.SayHello(
		context.Background(),
		connect.NewRequest(&hello.SayHelloRequest{
			Name: "Phong",
		}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.Message, "Successfully")

}
