package main

import (
	"context"
	"github.com/rschoonheim/docker-sphere/internal/deployment"
	"google.golang.org/grpc"
)

// How can I build a CLI for my gRPC server?
func main() {
	dialer, err := grpc.Dial("localhost:50091", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := deployment.NewDeploymentServiceClient(dialer)

	_, err = client.CreateDeployment(context.Background(), &deployment.CreateDeploymentRequest{})
	if err != nil {
		panic(err)
	}

}
