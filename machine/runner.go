package machine

// Runner is an abstraction of exec's Cmd & Output func; also useful for mocking
type Runner interface {
	Run(exeName string, args ...string) (string, error)
}
