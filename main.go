package main

import (
	"docker_sphere/orchestrators"
	systems_mongodb "docker_sphere/systems/mongodb"
)

func main() {

	/**
	 * Test mongo..
	 */
	systems_mongodb.Initialize()

	orchestrators.CreateUser("test", "test")

	///**
	// * Structured logging setup
	// */
	//logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	//slog.SetDefault(logger)
	//
	//go docker_sphere_gateway.Initialize()
	//
	//// Keep running
	//select {}
}
