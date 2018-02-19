package vm

// Runner is an abstraction of os/exec's Cmd, Run & Output func; also useful for mocking
type Runner interface {
	Run(exeName string, args ...string) (string, error)
}
