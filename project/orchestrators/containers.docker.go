package orchestrators

import (
	"context"
	"docker_sphere/project/systems/docker"
	"fmt"
	"github.com/docker/docker/api/types"
)

// DockerListContainers - list all containers
func DockerListContainers() {
	docker := docker.GetClient()

	// list all containers
	containers, err := docker.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID)
	}
}
