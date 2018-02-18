package machine

import (
	"errors"

	"github.com/davidjenni/coreUp/machine/digitalocean"
)

// Config represents the CoreOS VM configuration
type Config struct {
	Name            string
	CloudProvider   string
	OptionsYamlFile string
}

// GetCreateArguments builds and returns a list of cmd line arguments to pass to docker-machine's create command
func (c Config) GetCreateArguments() ([]string, error) {
	if c.CloudProvider != "digitalocean" {
		return nil, errors.New("Currently, the sole supported cloud provider is 'digitalocean'")
	}

	cmdArgs := []string{"create", "--driver", c.CloudProvider}
	config, err := digitalocean.NewConfig(c.OptionsYamlFile)
	if err != nil {
		return nil, err
	}
	args, err := config.Render()
	if err != nil {
		return nil, err
	}
	cmdArgs = append(cmdArgs, args...)
	cmdArgs = append(cmdArgs, c.Name)
	return cmdArgs, nil
}
