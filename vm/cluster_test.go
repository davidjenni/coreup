package vm_test

import (
	"testing"

	"github.com/davidjenni/coreup/vm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultClusterConfig(t *testing.T) {
	c := vm.GetDefaultCluster()
	require.NotEmpty(t, c)
	assert := assert.New(t)
	assert.NotEmpty(c.CloudProvider)
	assert.Equal(c.CloudProvider, "digitalocean")
	assert.NotEmpty(c.Manager)
	m := c.Manager
	assert.NotEmpty(m.Basename)
	assert.Contains(m.Basename, "%02d")
	assert.NotEmpty(m.NodeCount)
	assert.NotEmpty(c.Worker)
	w := c.Worker
	assert.NotEmpty(w.Basename)
	assert.Contains(w.Basename, "%02d")
	assert.NotEmpty(w.NodeCount)
}

func TestReadSimpleClusterConfig(t *testing.T) {
	c, err := vm.LoadClusterConfig("testdata/simpleCluster.yaml")
	assert.Nil(t, err)
	require.NotEmpty(t, c)
	defaultConfig := vm.GetDefaultCluster()
	assert.Equal(t, defaultConfig, c)
}

func TestReadClusterConfigWithOptionsFile(t *testing.T) {
	c, err := vm.LoadClusterConfig("testdata/ClusterWithOptions.yaml")
	assert.Nil(t, err)
	require.NotEmpty(t, c)
	defCluster := vm.GetDefaultCluster()

	// clear filenames to not trip up .Equal's diff comparison
	defCluster.Manager.VMOptionsFile = ""
	defCluster.Manager.VMOptionsFile = ""
	c.Manager.VMOptionsFile = ""
	c.Worker.VMOptionsFile = ""
	assert.Equal(t, defCluster, c)
	assert.Equal(t, c.Manager.NodeCount+c.Worker.NodeCount, len(c.GetMachines()))
}

func TestCreateVMs(t *testing.T) {
	runner := &mockRunner{expectedName: ""}
	args := []string{"create", "--driver", "digitalocean", "--digitalocean-region", "sfo2",
		"--digitalocean-image", "coreos-stable", "--digitalocean-size", "s-1vcpu-1gb",
		"--digitalocean-ssh-user", "core", "--digitalocean-ssh-port", "22",
		"--digitalocean-ipv6", "--digitalocean-private-networking", "--digitalocean-access-token", "fakeToken"}

	runner.On("Run", "docker-machine", append(args, "minion-mgr-01")).Return("", nil)
	runner.On("Run", "docker-machine", append(args, "minion-mgr-02")).Return("", nil)
	runner.On("Run", "docker-machine", append(args, "minion-mgr-03")).Return("", nil)
	runner.On("Run", "docker-machine", append(args, "minion-01")).Return("", nil)
	runner.On("Run", "docker-machine", append(args, "minion-02")).Return("", nil)

	c := vm.GetDefaultCluster(runner)
	require.NotEmpty(t, c)
	assert := assert.New(t)
	assert.NotZero(len(c.GetMachines()))
	err := c.CreateVMs("fakeToken")
	assert.Nil(err)
}
