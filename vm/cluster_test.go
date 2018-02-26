package vm_test

import (
	"testing"

	"github.com/davidjenni/coreup/vm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadSimpleClusterConfig(t *testing.T) {
	config, err := vm.LoadClusterConfig("testdata/simpleCluster.yaml")
	assert.Nil(t, err)
	require.NotEmpty(t, config)
	defaultConfig := vm.GetDefaultCluster()
	assert.Equal(t, defaultConfig, config)
}

func TestReadClusterConfigWithOptionsFile(t *testing.T) {
	config, err := vm.LoadClusterConfig("testdata/simpleCluster.yaml")
	assert.Nil(t, err)
	require.NotEmpty(t, config)
	assert.NotEmpty(t, config.Manager.VMOptions)
	assert.NotEmpty(t, config.Worker.VMOptions)
	defaultConfig := vm.GetDefaultCluster()
	defaultConfig.Manager.VMOptionsFile = ""
	defaultConfig.Manager.VMOptionsFile = ""
	assert.Equal(t, defaultConfig, config)
}
