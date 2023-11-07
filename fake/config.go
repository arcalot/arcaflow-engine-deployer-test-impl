package fake

// Config is the configuration structure of the fake connector.
type Config struct{}

// Validate checks the configuration structure for conformance with the schema.
func (c *Config) Validate() error {
	return SchemaFake.ValidateType(c)
}
