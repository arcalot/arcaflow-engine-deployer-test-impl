package testimpl

import (
	"go.flow.arcalot.io/pluginsdk/schema"
)

// Schema describes the deployment options of the Docker deployment mechanism.
var Schema = schema.NewTypedScopeSchema[*Config](
	schema.NewStructMappedObjectSchema[*Config](
		"Config",
		map[string]*schema.PropertySchema{
			"deploy_time": schema.NewPropertySchema(
				schema.NewIntSchema(schema.PointerTo(int64(0)), schema.PointerTo(int64(3600000)), nil),
				schema.NewDisplayValue(schema.PointerTo("Deploy Time"),
					schema.PointerTo("How long to wait when fake deploying"), nil),
				false,
				nil,
				nil,
				nil,
				schema.PointerTo("0"),
				nil,
			),
			"deploy_succeed": schema.NewPropertySchema(
				schema.NewBoolSchema(),
				schema.NewDisplayValue(schema.PointerTo("DeploySucceed"),
					schema.PointerTo("Should the deployment succeed?"), nil),
				false,
				nil,
				nil,
				nil,
				schema.PointerTo("true"),
				nil),
			"disable_plugin_writes": schema.NewPropertySchema(
				schema.NewBoolSchema(),
				schema.NewDisplayValue(schema.PointerTo("Disable Plugin Writes"),
					schema.PointerTo("Disable Plugin's ability to write?"), nil),
				false,
				nil,
				nil,
				nil,
				schema.PointerTo("false"),
				nil),
		},
	),
)
