package stub

import (
	log "go.arcalot.io/log/v2"
	"go.flow.arcalot.io/deployer"
	"go.flow.arcalot.io/pluginsdk/schema"
)

// NewFactory creates a new factory for the test stub deployer.
func NewFactory() deployer.ConnectorFactory[*Config] {
	return &factory{}
}

type factory struct {
}

func (f factory) Name() string {
	return "test-stub"
}

func (f factory) DeploymentType() deployer.DeploymentType {
	return "test-double"
}

func (f factory) ConfigurationSchema() *schema.TypedScopeSchema[*Config] {
	return SchemaStub
}

func (f factory) Create(config *Config, logger log.Logger) (deployer.Connector, error) {
	return nil, nil
}
