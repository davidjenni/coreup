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
