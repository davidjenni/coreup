package machine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRunner struct {
	expectedName string
	mock.Mock
}

func (r *mockRunner) Run(verb string, args ...string) (string, error) {
	result := r.Called(verb, args)
	return result.String(0), result.Error(1)
}

func TestExistsForKnownMachine(t *testing.T) {
	machName := "minion"
	runner := &mockRunner{expectedName: machName}
	runner.On("Run", "ls", []string{"-q"}).Return(machName, nil)

	m := &Machine{Name: machName, runner: runner}
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
