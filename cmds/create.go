package cmds

import (
	"github.com/davidjenni/coreUp/machine"
	"github.com/urfave/cli"
)

// Create will call docker-machine to create a new VM in the cloud
func Create(c *cli.Context) error {
	name := c.String("name")
	config := &machine.Config{CloudProvider: "digitalocean"}
	m := machine.New(name)
	error := m.CreateMachine(config)
	return error
}
