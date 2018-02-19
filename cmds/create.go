package cmds

import (
	"log"

	"github.com/davidjenni/coreUp/machine"
	"github.com/urfave/cli"
)

// Create will call docker-machine to create a new VM in the cloud
func Create(c *cli.Context) error {
	config := &machine.Config{
		VMName:        c.String("name"),
		CloudProvider: c.GlobalString("cloudProvider"),
		CloudAPIToken: c.GlobalString("apiToken"),
	}
	m := machine.New(config)
	error := m.CreateMachine()
	if error != nil {
		log.Fatalf("Cannot create VM: %s", error.Error())
	}
	return error
}
