package main

import (
	"log"

	"github.com/davidjenni/coreup/vm"
	"github.com/urfave/cli"
)

// Create will call docker-machine to create all VMs defined in cluster config file
func Create(c *cli.Context) error {
	var error error
	configFile := c.String("clusterConfig")
	cluster, error := vm.LoadClusterConfig(configFile)
	if error != nil {
		log.Fatalf("Cannot load cluster configuration file '%s': %s", configFile, error.Error())
	}
	error = cluster.CreateVMs(c.GlobalString("apiToken"))
	if error != nil {
		log.Fatalf("Cannot create cluster: %s", error.Error())
	}
	return error
}
