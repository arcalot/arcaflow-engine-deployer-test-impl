package fake

import (
	"go.flow.arcalot.io/pluginsdk/schema"
)

var SchemaFake = schema.NewTypedScopeSchema[*Config](
	schema.NewStructMappedObjectSchema[*Config](
		"Config",
		map[string]*schema.PropertySchema{},
	),
)
