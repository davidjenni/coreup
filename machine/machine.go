package machine

import "strings"

// Machine represents a VM as known to local docker-machine
type Machine struct {
	Name   string
	runner Runner
}

// Exists checks if this machine is already known to local docker-machine
func (m *Machine) Exists() (bool, error) {
	out, err := m.runner.Run("ls", "-q")
	if err != nil {
		return false, err
	}
	return strings.Contains(out, m.Name), nil
}

// New instance of Machine
func New(name string) *Machine {
	return &Machine{Name: name}
}

// CreateMachine initiates the creation of a new VM
func (m *Machine) CreateMachine(config *Config) error {
	config.VMName = m.Name
	cfg, error := config.GetCreateArguments()
	if error != nil {
		return error
	}
	_, error = m.runner.Run("docker-machine", cfg...)
	return error
}
