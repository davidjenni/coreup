package digitalocean

import "strconv"

// Options represents the VM options for DO
type Options struct {
	Region            string
	Image             string
	Size              string
	SSHUser           string
	SSHPort           int
	SSHKeyFile        string
	SSHKeyFingerprint string
}

// GetDefaults returns default option values
// func (o *Options) GetDefaults() *DigitalOceanOptions {
func GetDefaults() *Options {
	return &Options{
		Region:  "sfo1",
		Image:   "coreos-stable",
		Size:    "512mb",
		SSHUser: "core",
		SSHPort: 4410,
	}
}

// Render options as argument string array
func (d Options) Render() ([]string, error) {
	args := []string{}
	args = append(args, "--digitalocean-region", d.Region)
	args = append(args, "--digitalocean-image", d.Image)
	args = append(args, "--digitalocean-size", d.Size)
	args = append(args, "--digitalocean-ssh-user", d.SSHUser)
	args = append(args, "--digitalocean-ssh-port", strconv.Itoa(d.SSHPort))
	if d.SSHKeyFile != "" {
		args = append(args, "--digitalocean-ssh-key-path", d.SSHKeyFile)
		args = append(args, "--digitalocean-ssh-key-fingerprint", d.SSHKeyFingerprint)
	}
	args = append(args, "--digitalocean-ipv6")
	args = append(args, "--digitalocean-private-networking")
	return args, nil
}
