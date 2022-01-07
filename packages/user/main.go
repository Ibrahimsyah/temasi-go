package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Ibrahimsyah/temasi-go/shared/pb/user"
	"google.golang.org/grpc"
)

func NewUserServer() *UserServer {
	return &UserServer{}
}

type UserServer struct {
}

func (*UserServer) GetAllUsers(ctx context.Context, emp *user.EmptyRequest) (*user.AllUserResponse, error) {

	return &user.AllUserResponse{
		Users: make([]*user.User, 0),
	}, nil
}

func (*UserServer) RegisterUser(ctx context.Context, newUser *user.NewUser) (*user.RegisterUserResponse, error) {

	return &user.RegisterUserResponse{}, nil
}

func main() {
	address := "0.0.0.0:5002"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("[User] Error: ", err)
	}

	userServer := NewUserServer()
	grpcSrv := grpc.NewServer()

	user.RegisterUserServiceServer(grpcSrv, userServer)

	fmt.Println("[User] Service is running on ", address)

	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatal("[User] GRPC Error: ", err)
	}
}
