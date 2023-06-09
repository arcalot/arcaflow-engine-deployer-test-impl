package testimpl

// Config is the configuration structure of the Docker connector.
type Config struct {
	// The number of seconds it should wait while deploying to mimic a real deployer.
	DeployTime int32 `json:"deploy_time"`
	Succeed    bool  `json:"succeed"`
}

// Validate checks the configuration structure for conformance with the schema.
func (c *Config) Validate() error {
	return Schema.ValidateType(c)
}
