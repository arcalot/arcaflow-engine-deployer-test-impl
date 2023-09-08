package testimpl

// Config is the configuration structure of the Docker connector.
type Config struct {
	// The number of milliseconds seconds it should wait while deploying to mimic a real deployer.
	DeployTime          int64 `json:"deploy_time"`
	DeploySucceed       bool  `json:"deploy_succeed"`
	DisablePluginWrites bool  `json:"disable_plugin_writes"`
}

// Validate checks the configuration structure for conformance with the schema.
func (c *Config) Validate() error {
	return Schema.ValidateType(c)
}
