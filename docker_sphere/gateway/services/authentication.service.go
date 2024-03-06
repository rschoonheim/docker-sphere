package gateway_services

import (
	"connectrpc.com/connect"
	"context"
	authentication "docker_sphere/gateway/pkg/authentication"
	"docker_sphere/gateway/pkg/authentication/gateway_authenticationconnect"
	"docker_sphere/orchestrators"
	"errors"
	"log/slog"
)

type AuthenticationService struct {
	gateway_authenticationconnect.UnimplementedAuthenticationServiceHandler
}

func (s *AuthenticationService) Login(ctx context.Context, req *connect.Request[authentication.LoginRequest]) (*connect.Response[authentication.LoginResponse], error) {
	slog.Info("Received login request", "username", req.Msg.Username)

	result, err := orchestrators.Login(req.Msg.Username, req.Msg.Password)
	if err != nil {
		return nil, errors.New("Failed to authenticate user")
	}

	slog.Info("User is authenticated", "username", req.Msg.Username)

	return &connect.Response[authentication.LoginResponse]{
		Msg: &authentication.LoginResponse{
			Token: result.Token,
		},
	}, nil
}

func (s *AuthenticationService) Logout(ctx context.Context, req *connect.Request[authentication.LogoutRequest]) (*connect.Response[authentication.LogoutResponse], error) {
	return nil, errors.New("Not implemented")
}
