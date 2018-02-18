package machine

import (
	"errors"

	"github.com/davidjenni/coreUp/machine/digitalocean"
)

// Config represents the CoreOS VM configuration
type Config struct {
	Name            string
	CloudProvider   string
	CloudConfigFile string
}

// GetCreateArguments builds and returns a list of cmd line arguments to pass to docker-machine's create command
func (c Config) GetCreateArguments() ([]string, error) {
	if c.CloudProvider != "digitalocean" {
		return nil, errors.New("Currently, the sole supported cloud provider is 'digitalocean'")
	}

	var cmdArgs = []string{"--driver", c.CloudProvider}
	config := digitalocean.NewConfig(nil)
	args, err := config.Render()
	if err != nil {
		return nil, err
	}
	cmdArgs = append(args)
	cmdArgs = append(cmdArgs, c.Name)
	return cmdArgs, nil
}
