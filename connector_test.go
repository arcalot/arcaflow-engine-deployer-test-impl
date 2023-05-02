package testimpl_test

import (
	"context"
	"encoding/json"
	"fmt"
	"go.arcalot.io/assert"
	"go.arcalot.io/log/v2"
	"go.flow.arcalot.io/pluginsdk/atp"
	"go.flow.arcalot.io/testdeployer"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

func TestSimpleInOut(t *testing.T) {
	// TODO
	configJSON := `{"deploy_time": 2}`
	var config any
	if err := json.Unmarshal([]byte(configJSON), &config); err != nil {
		t.Fatal(err)
	}

	factory := testimpl.NewFactory()
	schema := factory.ConfigurationSchema()
	unserializedConfig, err := schema.UnserializeType(config)
	err = unserializedConfig.Validate()
	assert.NoError(t, err)
	assert.NoError(t, err)
	connector, err := factory.Create(unserializedConfig, log.NewTestLogger(t))
	assert.NoError(t, err)

	container, err := connector.Deploy(context.Background(), "quay.io/joconnel/io-test-script")
	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, container.Close())
	})
}

func TestE2E(t *testing.T) {
	//logConfig := log.Config{
	//	Level:       log.LevelError,
	//	Destination: log.DestinationStdout,
	//}
	//logger := log.New(
	//	logConfig,
	//)
	image := "image-dummy"
	file := "file-dummy"
	stepID := "wait_"

	d := testimpl.NewFactory()
	configSchema := d.ConfigurationSchema()
	defaultConfig, err := configSchema.UnserializeType(map[string]any{})
	if err != nil {
		panic(err)
	}
	connector, err := d.Create(defaultConfig, log.New(log.Config{
		Level:       log.LevelDebug,
		Destination: log.DestinationStdout,
	}))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	plugin, err := connector.Deploy(ctx, image)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := plugin.Close(); err != nil {
			panic(err)
		}
	}()

	atpClient := atp.NewClient(plugin)
	pluginSchema, err := atpClient.ReadSchema()
	if err != nil {
		panic(err)
	}
	steps := pluginSchema.Steps()
	step, ok := steps[stepID]
	if !ok {
		panic(fmt.Errorf("No such step: %s", stepID))
	}
	inputContents, err := os.ReadFile(file) //nolint:gosec
	if err != nil {
		panic(err)
	}
	input := map[string]any{}
	if err := yaml.Unmarshal(inputContents, &input); err != nil {
		panic(err)
	}
	if _, err := step.Input().Unserialize(input); err != nil {
		panic(err)
	}
	outputID, outputData, debugLogs := atpClient.Execute(ctx, stepID, input)
	output := map[string]any{
		"outputID":   outputID,
		"outputData": outputData,
		"debugLogs":  debugLogs,
	}
	result, err := yaml.Marshal(output)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", result)
}
