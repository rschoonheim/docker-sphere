package gateway_services

import (
	"context"
	gateway_authentication "docker_sphere/gateway/pkg/authentication"
	"errors"
	"google.golang.org/grpc"
)

type AuthenticationService struct {
	gateway_authentication.UnimplementedAuthenticationServiceServer
}

func (s *AuthenticationService) Login(ctx context.Context, in *gateway_authentication.LoginRequest, opts ...grpc.CallOption) (*gateway_authentication.LoginResponse, error) {
	return nil, errors.New("Not implemented")
}

func (s *AuthenticationService) Logout(ctx context.Context, in *gateway_authentication.LogoutRequest, opts ...grpc.CallOption) (*gateway_authentication.LogoutResponse, error) {
	return nil, errors.New("Not implemented")
}

//func (s *AuthenticationService) Login(ctx context.Context, req *connect.Request[gateway_authentication.LoginRequest]) (*connect.Response[gateway_authentication.LoginResponse], error) {
//	slog.Info("Received login request", "username", req.Msg.Username)
//
//	result, err := orchestrators.Login(req.Msg.Username, req.Msg.Password)
//	if err != nil {
//		return nil, errors.New("Failed to authenticate user")
//	}
//
//	slog.Info("User is authenticated", "username", req.Msg.Username)
//
//	return &connect.Response[gateway_authentication.LoginResponse]{
//		Msg: &gateway_authentication.LoginResponse{
//			Token: result.Token,
//		},
//	}, nil
//}
//
//func (s *AuthenticationService) Logout(ctx context.Context, req *connect.Request[gateway_authentication.LogoutRequest]) (*connect.Response[gateway_authentication.LogoutResponse], error) {
//	return nil, errors.New("Not implemented")
//}
