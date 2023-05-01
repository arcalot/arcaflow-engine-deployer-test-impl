package testimpl

import (
	"go.flow.arcalot.io/pluginsdk/schema"
)

// Schema describes the deployment options of the Docker deployment mechanism.
var Schema = schema.NewTypedScopeSchema[*Config](
	schema.NewStructMappedObjectSchema[*Config](
		"Config",
		map[string]*schema.PropertySchema{
			"DeployTime": schema.NewPropertySchema(
				schema.NewIntSchema(schema.PointerTo(int64(0)), schema.PointerTo(int64(3600000)), nil),
				schema.NewDisplayValue(schema.PointerTo("Deploy Time"), schema.PointerTo("How long to wait when fake deploying"), nil),
				false,
				nil,
				nil,
				nil,
				schema.PointerTo("0"),
				nil,
			),
		},
	),
)
