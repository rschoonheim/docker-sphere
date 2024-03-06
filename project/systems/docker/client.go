package docker

import (
	"github.com/docker/docker/client"
)

// GetClient - returns client for docker
func GetClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	return cli
}
