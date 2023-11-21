package stub

import (
	"go.flow.arcalot.io/pluginsdk/schema"
)

var SchemaStub = schema.NewTypedScopeSchema[*Config](
	schema.NewStructMappedObjectSchema[*Config](
		"Config",
		map[string]*schema.PropertySchema{},
	),
)
