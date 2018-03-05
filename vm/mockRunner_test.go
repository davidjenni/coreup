package vm_test

import "github.com/stretchr/testify/mock"

type mockRunner struct {
	expectedName string
	mock.Mock
}

func (r *mockRunner) Run(exeName string, args ...string) (string, error) {
	result := r.Called("docker-machine", args)
	return result.String(0), result.Error(1)
}
