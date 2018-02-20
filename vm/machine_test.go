package vm_test

import (
	"testing"

	"github.com/davidjenni/coreup/vm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRunner struct {
	expectedName string
	mock.Mock
}

func (r *mockRunner) Run(exeName string, args ...string) (string, error) {
	result := r.Called("docker-machine", args)
	return result.String(0), result.Error(1)
}

func TestExistsForKnownMachine(t *testing.T) {
	machName := "minion"
	runner := &mockRunner{expectedName: machName}
	runner.On("Run", "docker-machine", []string{"ls", "-q"}).Return(machName, nil)

	config := &vm.Config{
		VMName: machName,
	}

	m := vm.New(config, runner)
	result, err := m.Exists()
	assert.Nil(t, err)
	assert.True(t, result, "machine should exist")
}

/*
func TestExistsForUnkownMachine(t *testing.T) {
	m := &Machine{Name: "foo"}
	result := m.Exists()
	assert.False(t, result, "machine should not exist")
}
*/

func TestCreateVM(t *testing.T) {
	machName := "minion"
	runner := &mockRunner{expectedName: machName}
	runner.On("Run", "docker-machine", []string{
		"create", "--driver", "digitalocean", "--digitalocean-region", "sfo2",
		"--digitalocean-image", "coreos-stable", "--digitalocean-size", "s-1vcpu-1gb",
		"--digitalocean-ssh-user", "core", "--digitalocean-ssh-port", "22",
		"--digitalocean-ipv6", "--digitalocean-private-networking",
		"--digitalocean-access-token", "fakeToken", "minion",
	}).Return(`Running pre-create checks...
Creating machine...
(minion) Creating SSH key...
(minion) Creating Digital Ocean droplet...
(minion) Waiting for IP address to be assigned to the Droplet...
Waiting for machine to be running, this may take a few minutes...
Detecting operating system of created instance...
Waiting for SSH to be available...
Detecting the provisioner...
Provisioning with coreOS...
Copying certs to the local machine directory...
Copying certs to the remote machine...
Setting Docker configuration on the remote daemon...
Checking connection to Docker...
Docker is up and running!
To see how to connect your Docker Client to the Docker Engine running on this virtual machine, run: docker-machine env minion
`, nil)

	config := &vm.Config{
		VMName:        machName,
		CloudProvider: "digitalocean",
		CloudAPIToken: "fakeToken",
	}

	m := vm.New(config, runner)
	err := m.CreateMachine()
	assert.Nil(t, err)
}
