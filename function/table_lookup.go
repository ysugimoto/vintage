package function

import (
	"github.com/fastly/compute-sdk-go/configstore"
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Table_lookup_Name = "table.lookup"

// Fastly built-in function implementation of table.lookup
// Arguments may be:
// - TABLE, STRING, STRING
// - TABLE, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup/
func Table_lookup[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	id string, // IDENT
	key string,
	optional ...string,
) (string, error) {

	var defaultValue string
	if len(optional) > 0 {
		defaultValue = optional[0]
	}

	table, ok := ctx.Tables[id]
	if !ok {
		return "", errors.FunctionError(
			Table_lookup_Name,
			"table %s does not exist", id,
		)
	}

	// If table is EdgeDictionary, fetch item from remote
	if table.IsEdgeDictionary() {
		return Table_lookup_EdgeDictionary(table.Name, key, defaultValue)
	}

	// Otherwise, lookup in-memory items
	if v, ok := table.Items[key]; ok {
		if cast, ok := v.(string); ok {
			return cast, nil
		}
		return "", errors.FunctionError(
			Table_lookup_Name,
			"table %s item could not cast to STRING type for key %s", id, key,
		)
	}
	return defaultValue, nil
}

// In-memory cache store EdgeDictionary pointer for key
var storeCaches map[string]*configstore.Store

func Table_lookup_EdgeDictionary(dictName, key, defaultValue string) (string, error) {
	store, ok := storeCaches[dictName]
	if !ok {
		opened, err := configstore.Open(dictName)
		if err != nil {
			return "", errors.FunctionError(
				Table_lookup_Name,
				"Failed to open configstore %s, %w", dictName, err,
			)
		}
		storeCaches[dictName] = opened
		store = opened
	}

	v, err := store.Get(key)
	if err != nil {
		return defaultValue, nil
	}
	return v, nil
}
