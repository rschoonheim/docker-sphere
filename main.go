package main

import (
	docker_sphere_gateway "docker_sphere/gateway"
	"log/slog"
	"os"
)

func main() {

	/**
	 * Structured logging setup
	 */
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	go docker_sphere_gateway.Initialize()

	// Keep running
	select {}
}
