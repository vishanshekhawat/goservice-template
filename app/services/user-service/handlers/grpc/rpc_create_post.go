package grpc

import (
	"context"
	"fmt"

	models "github.com/vishn007/go-service-template/buisness/repo/userrepo/model"
	pb "github.com/vishn007/go-service-template/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (postServer *PostServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {

	fmt.Println("Request:", req)

	_, err := postServer.UserService.CreateUser(ctx, models.User{
		Name:  "Vishnu",
		Email: "City",
		City:  "City",
	})

	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.UserResponse{
		User: &pb.User{
			Id:    "1",
			Email: "vishnsingh007@gmail.com",
			City:  "City",
		},
	}
	return res, nil
}
