package vm

import (
	"errors"

	"github.com/davidjenni/coreup/vm/digitalocean"
)

// Config represents the CoreOS VM configuration
type Config struct {
	VMName         string
	CloudProvider  string
	CloudAPIToken  string
	providerConfig ConfigRenderer
}

// NewConfig parses the options file and returns a new VM configuration
func NewConfig(vmName string, cloudProvider string, optionsYamlFile string) (*Config, error) {
	// TODO: add cloud provider factory
	if cloudProvider != "digitalocean" {
		return nil, errors.New("Currently, the sole supported cloud provider is 'digitalocean'")
	}

	providerConfig, err := digitalocean.NewConfig(optionsYamlFile)
	if err != nil {
		return nil, err
	}
	return NewConfigFrom(vmName, cloudProvider, providerConfig)
}

// NewConfigFrom returns a new VM configuration from an already existing providerConfig
func NewConfigFrom(vmName string, cloudProvider string, providerConfig ConfigRenderer) (*Config, error) {
	return &Config{
		VMName:         vmName,
		CloudProvider:  cloudProvider,
		providerConfig: providerConfig,
	}, nil
}

// GetDefaultProviderConfig returns the default config renderer
func GetDefaultProviderConfig(cloudProvider string) (ConfigRenderer, error) {
	return LoadProviderConfig(cloudProvider, "")
}

// LoadProviderConfig parses and returns a cloud provider config renderer
func LoadProviderConfig(cloudProvider string, optionsYamlFile string) (ConfigRenderer, error) {
	// TODO: add cloud provider factory
	if cloudProvider != "digitalocean" {
		return nil, errors.New("Currently, the sole supported cloud provider is 'digitalocean'")
	}

	return digitalocean.NewConfig(optionsYamlFile)
}

// GetCreateArguments builds and returns a list of cmd line arguments to pass to docker-machine's create command
func (c Config) GetCreateArguments() ([]string, error) {
	cmdArgs := []string{"create", "--driver", c.CloudProvider}
	args, err := c.providerConfig.Render(c.CloudAPIToken)
	if err != nil {
		return nil, err
	}
	cmdArgs = append(cmdArgs, args...)
	cmdArgs = append(cmdArgs, c.VMName)
	return cmdArgs, nil
}
