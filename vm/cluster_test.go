package vm_test

import (
	"testing"

	"github.com/davidjenni/coreup/vm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultClusterConfig(t *testing.T) {
	config := vm.GetDefaultCluster()
	require.NotEmpty(t, config)
	assert := assert.New(t)
	assert.NotEmpty(config.CloudProvider)
	assert.Equal(config.CloudProvider, "digitalocean")
	assert.NotEmpty(config.Manager)
	m := config.Manager
	assert.NotEmpty(m.Basename)
	assert.Contains(m.Basename, "%02d")
	assert.NotEmpty(m.NodeCount)
	assert.NotEmpty(config.Worker)
	w := config.Worker
	assert.NotEmpty(w.Basename)
	assert.Contains(w.Basename, "%02d")
	assert.NotEmpty(w.NodeCount)
}

func TestReadSimpleClusterConfig(t *testing.T) {
	config, err := vm.LoadClusterConfig("testdata/simpleCluster.yaml")
	assert.Nil(t, err)
	require.NotEmpty(t, config)
	defaultConfig := vm.GetDefaultCluster()
	assert.Equal(t, defaultConfig, config)
}

func TestReadClusterConfigWithOptionsFile(t *testing.T) {
	config, err := vm.LoadClusterConfig("testdata/ClusterWithOptions.yaml")
	assert.Nil(t, err)
	require.NotEmpty(t, config)
	defaultConfig := vm.GetDefaultCluster()

	// clear filenames to not trip up .Equal's diff comparison
	defaultConfig.Manager.VMOptionsFile = ""
	defaultConfig.Manager.VMOptionsFile = ""
	config.Manager.VMOptionsFile = ""
	config.Worker.VMOptionsFile = ""
	assert.Equal(t, defaultConfig, config)
}
