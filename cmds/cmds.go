package cmds

import "github.com/urfave/cli"

// Commands collection supported by this app
var Commands = []cli.Command{
	{
		Name:   "create",
		Usage:  "Create a new cloud VM with CoreOS and joined to a swarm",
		Action: Create,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name",
				Value: "minion",
			},
		},
	},
}
