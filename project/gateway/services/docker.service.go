package gateway_services

import (
	"connectrpc.com/connect"
	"context"
	"docker_sphere/project/gateway/pkg/docker"
	"docker_sphere/project/gateway/pkg/docker/dockerconnect"
	"docker_sphere/project/orchestrators"
)

type DockerService struct {
	dockerconnect.UnimplementedDockerServiceHandler
}

func (s *DockerService) ListContainers(context.Context, *connect.Request[docker.ListContainersRequest]) (*connect.Response[docker.ListContainersResponse], error) {

	containers := orchestrators.DockerListContainers()

	// Create list of containers
	response := make([]*docker.Container, 0)
	for _, container := range *containers {
		response = append(response, &docker.Container{
			Id:     container.ID,
			Name:   container.Names[0],
			Status: container.Status,
		})

	}

	return &connect.Response[docker.ListContainersResponse]{
		Msg: &docker.ListContainersResponse{
			Containers: response,
		},
	}, nil
}
