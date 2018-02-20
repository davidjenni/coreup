package main

import (
	"log"

	"github.com/davidjenni/coreup/vm"
	"github.com/urfave/cli"
)

// Create will call docker-machine to create a new VM in the cloud
func Create(c *cli.Context) error {
	config := &vm.Config{
		VMName:        c.String("name"),
		CloudProvider: c.GlobalString("cloudProvider"),
		CloudAPIToken: c.GlobalString("apiToken"),
	}
	m := vm.New(config)
	error := m.CreateMachine()
	if error != nil {
		log.Fatalf("Cannot create VM: %s", error.Error())
	}
	return error
}
