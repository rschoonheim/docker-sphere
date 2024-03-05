package docker_sphere_gateway

import (
	"connectrpc.com/connect"
	"context"
	v1 "docker_sphere/gateway/pkg/server/api/authentication/v1"
	"docker_sphere/gateway/pkg/server/api/authentication/v1/v1connect"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
)

type AuthenticationServer struct{}

func (s *AuthenticationServer) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&v1.LoginResponse{
		Token: "token",
	})
	res.Header().Set("Authentication-Version", "v1")
	return res, nil
}

func (s *AuthenticationServer) Logout(ctx context.Context, req *connect.Request[v1.LogoutRequest]) (*connect.Response[v1.LogoutResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&v1.LogoutResponse{})
	res.Header().Set("Authentication-Version", "v1")
	return res, nil
}

// Initialize - initializes the gateways for Docker Sphere.
func Initialize() {
	slog.Info("Initializing Docker Sphere gateways...")

	greeter := &AuthenticationServer{}
	mux := http.NewServeMux()
	path, handler := v1connect.NewAuthenticationServiceHandler(greeter)
	mux.Handle(path, cors.AllowAll().Handler(handler))

	// Start listening on http.
	go func() {
		http.ListenAndServe(
			"localhost:50051",
			h2c.NewHandler(mux, &http2.Server{}),
		)
	}()

	// Start listening on unix socket.
	go func() {
		if _, err := os.Stat("/tmp/docker-sphere-gateway.sock"); err == nil {
			os.Remove("/tmp/docker-sphere-gateway.sock")
		}

		unixListener, err := net.Listen("unix", "/tmp/docker-sphere-gateway.sock")
		if err != nil {
			log.Fatal("Error starting unix listener: ", err)
			os.Exit(1)
		}

		// Start the server
		http.Serve(unixListener, h2c.NewHandler(mux, &http2.Server{}))
	}()

}