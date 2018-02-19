package digitalocean

// Config represents VM configuration
type Config struct {
	options *Options
}

// NewConfig initializes configuration for DigitalOcean VM
func NewConfig(optionsYamlFile string) (*Config, error) {
	var err error
	d := &Config{}
	if optionsYamlFile == "" {
		d.options = GetDefaults()
	} else {
		d.options, err = LoadOptionsFromFile(optionsYamlFile)
		if err != nil {
			return nil, err
		}
	}
	return d, nil
}

// Render a configuration's argument string array
func (d Config) Render(apiToken string) ([]string, error) {
	return d.options.Render(apiToken)
}
