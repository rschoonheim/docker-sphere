package docker_sphere_gateway

import (
	authentication_connect "docker_sphere/project/gateway/pkg/authentication/gateway_authenticationconnect"
	gateway_dockerconnect "docker_sphere/project/gateway/pkg/docker/dockerconnect"
	gateway_services "docker_sphere/project/gateway/services"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log/slog"
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
}
