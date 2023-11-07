package stub

// Config is the configuration structure of the stub connector.
type Config struct{}

// Validate checks the configuration structure for conformance with the schema.
func (c *Config) Validate() error {
	return SchemaStub.ValidateType(c)
}
