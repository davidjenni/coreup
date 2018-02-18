package machine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateArguments(t *testing.T) {

	config := &Config{Name: "vm1", CloudProvider: "digitalocean"}

	args, err := config.GetCreateArguments()
	assert.Nil(t, err)
	assert.NotEmpty(t, args)
	assert.Len(t, args, 16)
}
