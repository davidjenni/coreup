package main

import (
	"fmt"
	"os"

	"github.com/davidjenni/coreUp/cmds"
	"github.com/urfave/cli"
)

const appVersion = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "coreUp"
	app.Version = appVersion
	app.Usage = "collection of commands to create and manage a CoreOS cluster running a docker swarm"
	app.Commands = cmds.Commands
	app.CommandNotFound = notFound
	app.Run(os.Args)
}

func notFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "'%s' is not a valid command. See command 'help'.\n", command)
	os.Exit(2)
}
