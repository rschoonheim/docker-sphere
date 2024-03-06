package docker_sphere_gateway

import (
	"connectrpc.com/connect"
	"context"
	"docker_sphere/project/gateway/pkg/server/authentication"
	"docker_sphere/project/gateway/pkg/server/authentication/authenticationconnect"
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

func (s *AuthenticationServer) Login(ctx context.Context, req *connect.Request[authentication.LoginRequest]) (*connect.Response[authentication.LoginResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&authentication.LoginResponse{
		Token: "token",
	})
	res.Header().Set("Authentication-Version", "v1")
	return res, nil
}

func (s *AuthenticationServer) Logout(ctx context.Context, req *connect.Request[authentication.LogoutRequest]) (*connect.Response[authentication.LogoutResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&authentication.LogoutResponse{})
	res.Header().Set("Authentication-Version", "v1")
	return res, nil
}

// Initialize - initializes the gateways for Docker Sphere.
func Initialize() {
	slog.Info("Initializing Docker Sphere gateways...")

	greeter := &AuthenticationServer{}
	mux := http.NewServeMux()
	path, handler := authenticationconnect.NewAuthenticationServiceHandler(greeter)
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
