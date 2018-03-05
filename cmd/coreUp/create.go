package main

import (
	"log"

	"github.com/davidjenni/coreup/vm"
	"github.com/urfave/cli"
)

// Create will call docker-machine to create a new VM in the cloud
func Create(c *cli.Context) error {
	var error error
	// TODO: add options file argument
	config, error := vm.NewConfig(c.String("name"), c.GlobalString("cloudProvider"), "")
	if error != nil {
		log.Fatalf("Cannot create VM: %s", error.Error())
	}
	config.CloudAPIToken = c.GlobalString("apiToken")
	m := vm.NewMachine(config)
	error = m.CreateMachine()
	if error != nil {
		log.Fatalf("Cannot create VM: %s", error.Error())
	}
	return error
}
