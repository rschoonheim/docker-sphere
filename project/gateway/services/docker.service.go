package gateway_services

import (
	"connectrpc.com/connect"
	"context"
	gateway_dockerconnect "docker_sphere/project/gateway/pkg/docker/dockerconnect"
	"docker_sphere/project/orchestrators"
	"docker_sphere/project/systems/docker"
)

type DockerService struct {
	gateway_dockerconnect.UnimplementedDockerServiceHandler
}

func (s *DockerService) ListContainers(ctx context.Context, req *connect.Request[docker.ListContainersRequest]) (*connect.Response[docker.ListContainersResponse], error) {
	orchestrators.DockerListContainers()
	return &connect.Response[docker.ListContainersResponse]{
		Msg: &docker.ListContainersResponse{},
	}, nil
}
