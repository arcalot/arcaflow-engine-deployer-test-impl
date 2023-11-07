package fake

import (
	log "go.arcalot.io/log/v2"
	"go.flow.arcalot.io/deployer"
	"go.flow.arcalot.io/pluginsdk/schema"
)

// NewFactory creates a new factory for the test fake deployer.
func NewFactory() deployer.ConnectorFactory[*Config] {
	return &factory{}
}

type factory struct {
}

func (f factory) Name() string {
	return "test-fake"
}

func (f factory) DeploymentType() deployer.DeploymentType {
	return "builtin"
}

func (f factory) ConfigurationSchema() *schema.TypedScopeSchema[*Config] {
	return SchemaFake
}

func (f factory) Create(config *Config, logger log.Logger) (deployer.Connector, error) {
	//return &connector{
	//	config,
	//	logger,
	//}, nil
	return nil, nil
}
