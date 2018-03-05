package digitalocean

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

// Options represents the VM options for DO
type Options struct {
	Region            string `yaml:"region"`
	Image             string `yaml:"image"`
	Size              string `yaml:"size"`
	SSHUser           string `yaml:"sshUser"`
	SSHPort           int    `yaml:"sshPort"`
	SSHKeyFile        string `yaml:"sshKeyFile"`
	SSHKeyFingerprint string `yaml:"sshKeyFingerprint"`
}

// GetDefaults returns default option values
func GetDefaults() *Options {
	return &Options{
		Region:  "sfo2",
		Image:   "coreos-stable",
		Size:    "s-1vcpu-1gb",
		SSHUser: "core",
		SSHPort: 22,
	}
}

// LoadOptionsFromFile parses YAML file and return as Options struct
func LoadOptionsFromFile(optionsYamlFile string) (*Options, error) {
	var err error

	buffer, err := ioutil.ReadFile(optionsYamlFile)
	if err != nil {
		return nil, err
	}
	return loadOptions(buffer)
}

func loadOptions(buffer []byte) (*Options, error) {
	var options Options
	yaml.Unmarshal(buffer, &options)
	err := ensureRequiredOptions(options)
	if err != nil {
		return nil, err
	}
	return &options, nil
}

// Render options as argument string array
func (d Options) Render(apiToken string) ([]string, error) {
	defaults := GetDefaults()
	err := ensureRequiredOptions(d)
	if err != nil {
		return nil, err
	}

	args := []string{}
	args = append(args, "--digitalocean-region", d.Region)
	args = append(args, "--digitalocean-image", d.Image)
	args = append(args, "--digitalocean-size", d.Size)
	args = append(args, "--digitalocean-ssh-user", d.SSHUser)
	if d.SSHPort == 0 {
		d.SSHPort = defaults.SSHPort
	}
	args = append(args, "--digitalocean-ssh-port", strconv.Itoa(d.SSHPort))
	if d.SSHKeyFile != "" {
		args = append(args, "--digitalocean-ssh-key-path", d.SSHKeyFile)
		args = append(args, "--digitalocean-ssh-key-fingerprint", d.SSHKeyFingerprint)
	}
	// opinionated: always enable IPv6 and private networks:
	args = append(args, "--digitalocean-ipv6")
	args = append(args, "--digitalocean-private-networking")
	if apiToken == "" {
		apiToken = os.Getenv("DIGITALOCEAN_ACCESS_TOKEN")
	}
	if apiToken == "" {
		return nil, errors.New("Must specify cloud provider's API token, either as config.CloudAPIToken or as env variable DIGITALOCEAN_ACCESS_TOKEN")
	}
	args = append(args, "--digitalocean-access-token", apiToken)
	return args, nil
}

func ensureRequiredOptions(d Options) error {
	if d.Region == "" {
		return missingRequiredOption("Region")
	}
	if d.Image == "" {
		return missingRequiredOption("Image")
	}
	if d.Size == "" {
		return missingRequiredOption("Size")
	}
	if d.SSHUser == "" {
		return missingRequiredOption("SSHUser")
	}
	return nil
}

func missingRequiredOption(optionName string) error {
	return fmt.Errorf("Missing required option: options.%v", optionName)
}
