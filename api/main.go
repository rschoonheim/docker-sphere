package main

import (
	"connectrpc.com/connect"
	"context"
	v1 "docker-sphere/gen/authentication/v1"
	"docker-sphere/gen/authentication/v1/authenticationconnect"
	"errors"
	"fmt"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
)

// Create instance of the authentication service and start the server
type authenticationServiceHandler struct {
	authenticationconnect.UnimplementedAuthenticationServiceHandler
}

func (authenticationServiceHandler) Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	fmt.Printf("Login")
	return connect.NewResponse(&v1.LoginResponse{
		Token: "example",
	}), nil
}

func (authenticationServiceHandler) Logout(context.Context, *connect.Request[v1.LogoutRequest]) (*connect.Response[v1.LogoutResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("authentication.v1.AuthenticationService.Logout is not implemented"))
}

func main() {

	greeter := authenticationServiceHandler{}

	mux := http.NewServeMux()
	path, handler := authenticationconnect.NewAuthenticationServiceHandler(greeter)
	mux.Handle(path, handler)

	handlerx := h2c.NewHandler(mux, &http2.Server{})

	http.ListenAndServe(
		"localhost:50051",
		cors.AllowAll().Handler(handlerx),
	)
}
