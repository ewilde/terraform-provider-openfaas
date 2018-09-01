package openfaas

type Config struct {
}

// Client configures and returns a fully initialized OpenFaaS client
func (c *Config) Client() (interface{}, error) {
	return nil, nil
}