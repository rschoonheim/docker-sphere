package main

import (
	"github.com/rschoonheim/docker-sphere/internal/deployment"
	"google.golang.org/grpc"
	"net"
)

func main() {

	/**
	 * Create gRPC server with services
	 * registered.
	 */
	server := grpc.NewServer()
	deployment.RegisterDeploymentServiceServer(server, &deployment.DeploymentService{})

	/**
	 * Start gRPC server in a goroutine
	 */
	go func() {
		listener, err := net.Listen("tcp", ":8080")
		if err != nil {
			panic(err)
		}

		if err := server.Serve(listener); err != nil {
			panic(err)
		}
	}()

	/**
	 * Start gRPC-web server in a goroutine
	 */
	go func() {

	}()

	select {}
}
