package main

import (
	"docker_sphere/packages/deployment"
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
	 * Start HTTP 2.0 server for gRPC in
	 * a goroutine
	 */
	go func() {
		listener, err := net.Listen("tcp", ":50091")
		if err != nil {
			panic(err)
		}

		if err := server.Serve(listener); err != nil {
			panic(err)
		}
	}()

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

	select {}
}
