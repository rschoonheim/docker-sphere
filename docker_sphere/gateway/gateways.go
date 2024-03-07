package docker_sphere_gateway

import (
	authentication_connect "docker_sphere/gateway/pkg/authentication/gateway_authenticationconnect"
	gateway_dockerconnect "docker_sphere/gateway/pkg/docker/dockerconnect"
	gateway_services "docker_sphere/gateway/services"
	"fmt"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
	"net/http"
)

// Initialize - initializes the gateways for Docker Sphere.
func Initialize() {
	slog.Info("Initializing Docker Sphere gateways...")

	/**
	 * Authentication Service
	 */
	greeter := &gateway_services.AuthenticationService{}

	/**
	 * Docker Service
	 */
	docker := &gateway_services.DockerService{}

	/**
	 * HTTP/2 Server
	 */
	mux := http.NewServeMux()

	/**
	 * Register gRPC services to the
	 * HTTP/2 server.
	 */

	// Authentication Service
	path, handler := authentication_connect.NewAuthenticationServiceHandler(greeter)
	mux.Handle(path, cors.AllowAll().Handler(handler))

	// Docker Service
	path, handler = gateway_dockerconnect.NewDockerServiceHandler(docker)
	mux.Handle(path, cors.AllowAll().Handler(handler))

	// Start the HTTP/2 server in a goroutine.
	go func() {
		http.ListenAndServe(
			"localhost:50051",
			h2c.NewHandler(mux, &http2.Server{}),
		)
	}()

	// Start standard go server in a goroutine.
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50052))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		var opts []grpc.ServerOption

		grpcServer := grpc.NewServer(opts...)

		//

		//pb.RegisterRouteGuideServer(grpcServer, newServer())
		grpcServer.Serve(lis)
	}()
}
