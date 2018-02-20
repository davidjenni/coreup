package vm

import "os/exec"

// Run the docker-machine comand
func Run(verb string, args ...string) (string, error) {
	argsWithVerb := []string{verb}
	argsWithVerb = append(argsWithVerb, args...)
	cmd := exec.Command("docker-machine", argsWithVerb...)
	output, err := cmd.Output()
	return string(output), err
}
