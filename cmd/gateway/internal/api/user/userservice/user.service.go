package userservice

import (
	"connectrpc.com/connect"
	"context"
	"docker_sphere_gateway/internal/api/user"
	"docker_sphere_gateway/internal/api/user/userconnect"
)

type UserServer struct {
	userconnect.UnimplementedUserServiceHandler
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (s *UserServer) CreateUser(ctx context.Context, req *connect.Request[user.User]) (*connect.Response[user.User], error) {
	println("CreateUser")
	return nil, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *connect.Request[user.User]) (*connect.Response[user.User], error) {
	return nil, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *connect.Request[user.User]) (*connect.Response[user.User], error) {
	return nil, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, req *connect.Request[user.User]) (*connect.Response[user.User], error) {
	return nil, nil
}
