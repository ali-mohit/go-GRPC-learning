package main


import (
	"log"
	"math/rand"
	"golang.org/x/net/context"
	pb "github.com/ali-mohit/go-GRPC-learning/grpc-services"

)

type UserManagementImp struct{
	pb.UnimplementedUserManagementServer
}


func (s *UserManagementImp) CreateNewUser(ctx context.Context,
	message *pb.CreateNewUserRequest) (*pb.CreateNewUserResponse, error){

	log.Println("Request Create New User Recived: ", message.GetName())
	resp := pb.CreateNewUserResponse{
		Id: int32(rand.Intn(1000)),
		Name: message.Name,
		Age: message.Age,
	}

	return &resp, nil
}
