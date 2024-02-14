package grpc

import (
	"context"
	"strconv"

	models "github.com/vishn007/go-service-template/buisness/repo/userrepo/model"
	pb "github.com/vishn007/go-service-template/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (postServer *PostServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {

	id, err := postServer.UserService.CreateUser(ctx, models.User{
		Name:  req.Name,
		Email: req.Email,
		City:  req.City,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.UserResponse{
		User: &pb.User{
			Id:    strconv.Itoa(id),
			Email: req.Email,
			City:  req.City,
		},
	}
	return res, nil
}
