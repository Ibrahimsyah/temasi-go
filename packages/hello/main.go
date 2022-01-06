package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Ibrahimsyah/temasi-go/shared/pb/hello"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, request *hello.HelloRequest) (*hello.HelloResponse, error) {
	fmt.Printf("[Hello] new request with name: %v\n", request.Name)

	name := request.Name
	response := &hello.HelloResponse{
		Greet: fmt.Sprintf("Hello %s!", name),
	}

	return response, nil
}

func main() {
	address := "0.0.0.0:5001"

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("[Hello] Error ", err)
	}

	fmt.Println("[Hello] Service is running on ", address)

	grpcSrv := grpc.NewServer()

	hello.RegisterHelloServer(grpcSrv, &server{})

	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatal("[Hello] GRPC Error: ", err)
	}
}
