package digitalocean

// Config represents VM configuration
type Config struct {
	options *Options
}

// NewConfig initializes configuration for DigitalOcean VM
func NewConfig(options *Options) *Config {
	d := &Config{}
	if options == nil {
		d.options = GetDefaults()
	} else {
		d.options = options
	}
	return d
}

// Render a configuration's argument string array
func (d Config) Render() ([]string, error) {
	return d.options.Render()
}
