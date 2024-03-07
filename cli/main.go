package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	app := &cli.App{
		Commands: make([]*cli.Command, 0),
	}

	// Add your commands here
	app.Commands = append(app.Commands, &cli.Command{
		Name:     "login",
		Usage:    "authenticate to Docker Sphere",
		Category: "authentication",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "username",
				Usage:    "username",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "password",
				Usage:    "password",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {

			username := c.String("username")
			if username == "" {
				log.Fatal("username cannot be empty")
			}

			password := c.String("password")
			if password == "" {
				log.Fatal("password cannot be empty")
			}

			println("login called with username: " + username + " and password: " + password)

			return nil
		},
	})

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
