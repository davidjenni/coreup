package vm

import (
	"bytes"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

// Machine represents a VM as known to local docker-machine
type Machine struct {
	config *Config
	runner Runner
}

const dockerMachine = "docker-machine"

// NewMachine instance of Machine
func NewMachine(config *Config, runnerOpt ...Runner) *Machine {
	// sure wished golang had optional/default values, but alas...:
	var runner Runner
	if len(runnerOpt) > 0 {
		runner = runnerOpt[0]
	}
	if runner == nil {
		runner = shellRun{}
	}
	return &Machine{config: config, runner: runner}
}

// Exists checks if this machine is already known to local docker-machine
func (m *Machine) Exists() (bool, error) {
	out, err := m.runner.Run(dockerMachine, "ls", "-q")
	if err != nil {
		return false, err
	}
	return strings.Contains(out, m.config.VMName), nil
}

// CreateMachine initiates the creation of a new VM
func (m *Machine) CreateMachine() error {
	cfg, error := m.config.GetCreateArguments()
	if error != nil {
		return error
	}
	_, error = m.runner.Run(dockerMachine, cfg...)
	return error
}

// shellRun: Runner using the OS' exec; captures both stdout and stderr
type shellRun struct{}

func (sr shellRun) Run(exeName string, args ...string) (string, error) {
	log.Printf("exec: %s %v", exeName, args)
	cmd := exec.Command(exeName, args...)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}
	errMsg, _ := ioutil.ReadAll(stderr)
	output, _ := ioutil.ReadAll(stdout)

	if err := cmd.Wait(); err != nil {
		log.Printf("errMsg : %s", errMsg)
		return "", err
	}

	msg := bytes.NewBuffer(output).String()
	log.Printf("success:\n%s", msg)
	return msg, nil
}
