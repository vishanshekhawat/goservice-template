package grpc

import (
	"github.com/vishn007/go-service-template/app/services/user-service/service"
	pb "github.com/vishn007/go-service-template/proto"
)

type PostServer struct {
	pb.UnimplementedUserServiceServer
	UserService service.Service
}

func NewGrpcPostServer(UserService service.Service) (*PostServer, error) {
	postServer := &PostServer{
		UserService: UserService,
	}

	return postServer, nil
}
