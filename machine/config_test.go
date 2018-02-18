package machine

import (
	"strconv"
	"testing"

	"github.com/davidjenni/coreUp/machine/digitalocean"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDefaultArguments(t *testing.T) {

	config := &Config{Name: "vm1", CloudProvider: "digitalocean"}

	args, err := config.GetCreateArguments()
	assert := assert.New(t)
	require.Nil(t, err)
	assert.NotEmpty(args)
	assert.Len(args, 16)
	assert.Contains(args, "create")
	assert.Contains(args, "digitalocean")

	defaults := digitalocean.GetDefaults()
	assert.Contains(args, defaults.Region)
	assert.Contains(args, defaults.Size)
	assert.Contains(args, defaults.Image)
	assert.Contains(args, defaults.SSHUser)
	assert.Contains(args, strconv.Itoa(defaults.SSHPort))
	assert.NotContains(args, defaults.SSHKeyFile)
	assert.NotContains(args, defaults.SSHKeyFingerprint)
}

func TestCreateArgumentsFromFile(t *testing.T) {

	config := &Config{Name: "vm1", CloudProvider: "digitalocean", OptionsYamlFile: "digitalocean/testdata/doOptions.yaml"}

	args, err := config.GetCreateArguments()
	assert := assert.New(t)
	require.Nil(t, err)
	assert.NotEmpty(args)
	assert.Len(args, 20)
	assert.Contains(args, "create")
	assert.Contains(args, "digitalocean")

	assert.Contains(args, "sfo2", "incorrect region")
	assert.Contains(args, "4gb", "incorrect size")
	assert.Contains(args, "debian-8-x64", "incorrect image")
	assert.Contains(args, "foobar", "incorrect sshUser")
	assert.Contains(args, "4410", "incorrect sshPort")
	assert.Contains(args, "myKeyFile", "incorrect sshKeyFile")
	assert.Contains(args, "abcd", "incorrect sshKeyFingerprint")
}
