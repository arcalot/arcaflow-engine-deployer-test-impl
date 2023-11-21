package stub

import (
	"encoding/json"
	"go.arcalot.io/assert"
	"go.arcalot.io/log/v2"
	"testing"
)

func TestStubFactory(t *testing.T) {
	configJSON := `{}`
	var config any
	err := json.Unmarshal([]byte(configJSON), &config)
	assert.NoError(t, err)

	f := NewFactory()
	schema := f.ConfigurationSchema()
	unserializedConfig, err := schema.UnserializeType(config)
	assert.NoError(t, err)
	assert.NoError(t, unserializedConfig.Validate())

	assert.NotNil(t, f.DeploymentType())
	assert.NotNil(t, f.Name())

	connector, err := f.Create(unserializedConfig, log.NewTestLogger(t))
	assert.NoError(t, err)
	assert.Nil(t, connector)
}
