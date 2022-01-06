package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Ibrahimsyah/temasi-go/shared/pb/hello"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := insecure.NewCredentials()

	cc, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(opts))
	if err != nil {
		log.Fatal("[Gateway] GRPC Connection Error: ", err)
	}
	defer cc.Close()

	client := hello.NewHelloClient(cc)

	for {
		name := fmt.Sprint("Fulan ", time.Now().Unix())

		request := &hello.HelloRequest{Name: name}
		resp, _ := client.Greet(context.Background(), request)
		fmt.Println("[Gateway] response: ", resp.Greet)

		time.Sleep(2 * time.Second)
	}
}
