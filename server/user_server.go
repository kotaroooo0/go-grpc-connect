package server

import (
	"context"
	userv1 "go-grpc-connect/gen/user/v1"
	"go-grpc-connect/module/user"
	"log"

	"connectrpc.com/connect"
)

type UserServer struct {
	userUseCase *user.UserUseCase
}

func NewUserServer(userUseCase *user.UserUseCase) *UserServer {
	return &UserServer{
		userUseCase: userUseCase,
	}
}

func (s *UserServer) CreateUser(
	ctx context.Context,
	req *connect.Request[userv1.CreateUserRequest],
) (*connect.Response[userv1.User], error) {
	log.Println("Request headers: ", req.Header())
	u, err := s.userUseCase.CreateUser(ctx, req.Msg.Name, int(req.Msg.Age))
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&userv1.User{
		Id:   int32(u.Id),
		Name: u.Name,
		Age:  int32(u.Age),
	})
	res.Header().Set("User-Version", "v1")
	return res, nil
}

func (s *UserServer) GetUser(
	ctx context.Context,
	req *connect.Request[userv1.GetUserRequest],
) (*connect.Response[userv1.User], error) {
	log.Println("Request headers: ", req.Header())
	u, err := s.userUseCase.GetUser(ctx, int(req.Msg.Id))
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&userv1.User{
		Id:   int32(u.Id),
		Name: u.Name,
		Age:  int32(u.Age),
	})
	res.Header().Set("User-Version", "v1")
	return res, nil
}
