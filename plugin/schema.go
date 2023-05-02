package plugin

import (
	"fmt"
	"go.flow.arcalot.io/pluginsdk/schema"
	"time"
)

type Input struct {
	WaitTime int `json:"wait_time"`
}

// We define a separate scope, so we can add subobjects later.
var inputSchema = schema.NewScopeSchema(
	// Struct-mapped object schemas are object definitions that are mapped to a specific struct (Input)
	schema.NewStructMappedObjectSchema[Input](
		// ID for the object:
		"input",
		// Properties of the object:
		map[string]*schema.PropertySchema{
			"wait_time": schema.NewPropertySchema(
				// Type properties:
				schema.NewIntSchema(schema.PointerTo[int64](1), nil, nil),
				// Display metadata:
				schema.NewDisplayValue(
					schema.PointerTo("Wait Time"),
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

func wait_(input Input) (string, any) {
	time.Sleep(time.Duration(input.WaitTime) * time.Millisecond)
	return "success", Output{
		fmt.Sprintf("Plugin waited for %d ms.", input.WaitTime),
	}
}

var WaitSchema = schema.NewCallableSchema(
	schema.NewCallableStep[Input](
		// ID of the function:
		"wait_",
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
		wait_,
	),
)
