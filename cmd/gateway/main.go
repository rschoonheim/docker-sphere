package main

import (
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func main() {

	/**
	 * Create gRPC server with services
	 * registered.
	 */
	server := grpc.NewServer()

	/**
	 * Start HTTP 2.0 server for gRPC in
	 * a goroutine
	 */
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/deployment.v1.DeploymentService/", server)

		http.ListenAndServe(
			"localhost:8080",
			h2c.NewHandler(mux, &http2.Server{}),
		)
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
