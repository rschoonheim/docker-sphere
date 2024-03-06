package orchestrators

import (
	"context"
	"docker_sphere/project/systems/docker"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

// DockerListContainers - list all containers
func DockerListContainers() *[]types.Container {
	docker := docker.GetClient()

	// list all containers
	containers, err := docker.ContainerList(context.Background(), container.ListOptions{
		All: true,
	})
	if err != nil {
		panic(err)
	}

	return &containers
}
