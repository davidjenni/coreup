package cmds

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// AppOptions declares app-wide options
type AppOptions struct {
	Name            string
	Usage           string
	Commands        []cli.Command
	AppFlags        []cli.Flag
	CommandNotFound cli.CommandNotFoundFunc
}

// GetAppOptions initializes app-wide options
func GetAppOptions() *AppOptions {
	return &AppOptions{
		Name:  "coreUp",
		Usage: "collection of commands to create and manage a CoreOS cluster running a docker swarm",
		AppFlags: []cli.Flag{
			cli.StringFlag{
				Name:  "cloudProvider",
				Value: "digitalocean",
				Usage: "Cloud provider to create VMs and swarm; only supported provider is 'digitalocean'",
			},
			cli.StringFlag{
				Name:  "apiToken",
				Usage: "API token of cloud provider",
			},
		},
		Commands:        commands,
		CommandNotFound: notFound,
	}
}

// Commands collection supported by this app
var commands = []cli.Command{
	{
		Name:   "create",
		Usage:  "Create a new cloud VM with CoreOS and join it to a swarm",
		Action: Create,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name",
				Value: "minion",
			},
		},
	},
}

func notFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "'%s' is not a valid command. See command 'help'.\n", command)
	os.Exit(2)
}
