package main

import (
	docker_sphere_gateway "docker_sphere/project/gateway"
	systems_mongodb "docker_sphere/project/systems/mongodb"
	"log/slog"
	"os"
)

func main() {

	/**
	* Structured logging setup
	* ----------------------------------------
	*
	 */
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	/**
	 * MongoDB Initialization
	 * ----------------------------------------
	 *
	 */
	systems_mongodb.Initialize()

	/**
	 * Gateway Initialization
	 * ----------------------------------------
	 *
	 */
	go docker_sphere_gateway.Initialize()

	// Keep running
	select {}
}
