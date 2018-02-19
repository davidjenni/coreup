package machine

import (
	"os"
	"strconv"
	"testing"

	"github.com/davidjenni/coreUp/machine/digitalocean"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDefaultArguments(t *testing.T) {

	config := &Config{
		VMName:        "vm1",
		CloudProvider: "digitalocean",
		CloudAPIToken: "fakeToken",
	}

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

	config := &Config{
		VMName:          "vm1",
		CloudProvider:   "digitalocean",
		OptionsYamlFile: "digitalocean/testdata/doOptions.yaml",
		CloudAPIToken:   "fakeToken",
	}

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

	config := &Config{VMName: "vm1", CloudProvider: "digitalocean", OptionsYamlFile: "digitalocean/testdata/doMissingRequiredOptions.yaml"}

	args, err := config.GetCreateArguments()
	require.Nil(t, args)
	require.NotEmpty(t, err)
	assert.Contains(t, err.Error(), "Missing required option")
}

func TestCreateArgumentsMissingApiToken(t *testing.T) {

	config := &Config{VMName: "vm1", CloudProvider: "digitalocean"}

	// shell/parent process might have this set; unset it for this test fixture
	os.Unsetenv("DIGITALOCEAN_ACCESS_TOKEN")

	var args []string
	var err error
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
