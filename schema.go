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
				schema.NewIntSchema(0, 3600000, nil),
				schema.NewRefSchema("Deploy Time", "How long to wait when fake deploying"),
				false,
				nil,
				nil,
				nil,
				schema.PointerTo(0),
				nil,
			),
		},
	),
)
