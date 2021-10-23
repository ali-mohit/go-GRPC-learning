package main

import (
	pb "github.com/ali-mohit/go-GRPC-learning/grpc-services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main(){
	listener, err := net.Listen("tcp", ":9000")
	if err != nil{
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}


	grpcServer := grpc.NewServer()

	userManagementImp := UserManagementImp{}
	pb.RegisterUserManagementServer(grpcServer, &userManagementImp)

	if err:= grpcServer.Serve(listener); err != nil{
		log.Fatalf("Failed to Serve gRPC server over port 9000: %v", err)
	}
}
