package plugin

import (
	"fmt"
	"go.flow.arcalot.io/pluginsdk/schema"
)

type Input struct {
	DeployTime string `json:"deploy_time"`
}

// We define a separate scope, so we can add subobjects later.
var inputSchema = schema.NewScopeSchema(
	// Struct-mapped object schemas are object definitions that are mapped to a specific struct (Input)
	schema.NewStructMappedObjectSchema[Input](
		// ID for the object:
		"input",
		// Properties of the object:
		map[string]*schema.PropertySchema{
			"deploy_time": schema.NewPropertySchema(
				// Type properties:
				schema.NewIntSchema(schema.PointerTo[int64](1), nil, nil),
				// Display metadata:
				schema.NewDisplayValue(
					schema.PointerTo("Deploy Time"),
					schema.PointerTo("How long to wait."),
					nil,
				),
				// Required:
				true,
				// Required if:
				[]string{},
				// Required if not:
				[]string{},
				// Conflicts:
				[]string{},
				// Default value, JSON encoded:
				nil,
				//Examples:
				nil,
			),
		},
	),
)

type Output struct {
	Message string `json:"message"`
}

var outputSchema = schema.NewScopeSchema(
	schema.NewStructMappedObjectSchema[Output](
		"output",
		map[string]*schema.PropertySchema{
			"message": schema.NewPropertySchema(
				schema.NewStringSchema(nil, nil, nil),
				schema.NewDisplayValue(
					schema.PointerTo("Message"),
					schema.PointerTo("The resulting message."),
					nil,
				),
				true,
				nil,
				nil,
				nil,
				nil,
				nil,
			),
		},
	),
)

func greet(input Input) (string, any) {
	return "success", Output{
		fmt.Sprintf("Mimicking deployment of a plugin for %s ms.", input.DeployTime),
	}
}

var GreetSchema = schema.NewCallableSchema(
	schema.NewCallableStep[Input](
		// ID of the function:
		"greet",
		// Add the input schema:
		inputSchema,
		map[string]*schema.StepOutputSchema{
			// Define possible outputs:
			"success": schema.NewStepOutputSchema(
				// Add the output schema:
				outputSchema,
				schema.NewDisplayValue(
					schema.PointerTo("Success"),
					schema.PointerTo("Successfully created message."),
					nil,
				),
				false,
			),
		},
		// Metadata for the function:
		schema.NewDisplayValue(
			schema.PointerTo("Wait"),
			schema.PointerTo("Wait on deployment."),
			nil,
		),
		// Reference the function
		greet,
	),
)
