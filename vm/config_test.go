package vm_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/davidjenni/coreup/vm"
	"github.com/davidjenni/coreup/vm/digitalocean"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDefaultArguments(t *testing.T) {
	var err error
	config, err := vm.NewConfig("vm1", "digitalocean", "")
	require.Nil(t, err)
	config.CloudAPIToken = "fakeToken"

	args, err := config.GetCreateArguments()
	require.Nil(t, err)
	require.NotEmpty(t, args)

	assert := assert.New(t)
	assert.Len(args, 18)
	assert.Contains(args, "create")
	assert.Contains(args, "digitalocean")

	defaults := digitalocean.GetDefaults()
	assert.Contains(args, "fakeToken")
	assert.Contains(args, defaults.Region)
	assert.Contains(args, defaults.Size)
	assert.Contains(args, defaults.Image)
	assert.Contains(args, defaults.SSHUser)
	assert.Contains(args, strconv.Itoa(defaults.SSHPort))
	assert.NotContains(args, defaults.SSHKeyFile)
	assert.NotContains(args, defaults.SSHKeyFingerprint)
}

func TestCreateArgumentsFromFile(t *testing.T) {
	var err error
	config, err := vm.NewConfig("vm1", "digitalocean", "digitalocean/testdata/doOptions.yaml")
	require.Nil(t, err)
	config.CloudAPIToken = "fakeToken"

	args, err := config.GetCreateArguments()
	require.Nil(t, err)
	require.NotEmpty(t, args)

	assert := assert.New(t)
	assert.Len(args, 22)
	assert.Contains(args, "create")
	assert.Contains(args, "digitalocean")

	assert.Contains(args, "fakeToken")
	assert.Contains(args, "sfo2", "incorrect region")
	assert.Contains(args, "4gb", "incorrect size")
	assert.Contains(args, "debian-8-x64", "incorrect image")
	assert.Contains(args, "foobar", "incorrect sshUser")
	assert.Contains(args, "4410", "incorrect sshPort")
	assert.Contains(args, "myKeyFile", "incorrect sshKeyFile")
	assert.Contains(args, "abcd", "incorrect sshKeyFingerprint")
}

func TestCreateArgumentsFlagMissingReqOptions(t *testing.T) {
	var err error
	config, err := vm.NewConfig("vm1", "digitalocean", "digitalocean/testdata/doMissingRequiredOptions.yaml")
	require.Nil(t, config)
	require.NotEmpty(t, err)
	assert.Contains(t, err.Error(), "Missing required option")
}

func TestCreateArgumentsMissingApiToken(t *testing.T) {
	var err error
	config, err := vm.NewConfig("vm1", "digitalocean", "")
	require.Nil(t, err)

	// shell/parent process might have this set; unset it for this test fixture
	os.Unsetenv("DIGITALOCEAN_ACCESS_TOKEN")

	var args []string
	args, err = config.GetCreateArguments()
	require.Nil(t, args)
	require.NotEmpty(t, err)
	assert.Contains(t, err.Error(), "Must specify cloud provider's API token")

	os.Setenv("DIGITALOCEAN_ACCESS_TOKEN", "fakeToken2")
	args, err = config.GetCreateArguments()
	require.Nil(t, err)
	require.NotEmpty(t, args)
	assert.Contains(t, args, "fakeToken2")
}
