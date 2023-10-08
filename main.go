package main

import (
	"log"
	"os"

	"github.com/modern-era-devs/lious-healthcare/server"
	"github.com/urfave/cli"
)

func main() {
	clientApp := cli.NewApp()
	clientApp.Name = "lious-healthcare"
	clientApp.Version = "0.0.1"
	clientApp.Action = func(c *cli.Context) error {
		err := server.StartAPIServer()
		return err
	}
	clientApp.Commands = []cli.Command{
		{
			Name:        "start",
			Description: "Start HTTP api server",
			Action: func(c *cli.Context) error {
				err := server.StartAPIServer()
				return err
			},
		},
	}

	if err := clientApp.Run(os.Args); err != nil {
		log.Println(err.Error())
	}
}
