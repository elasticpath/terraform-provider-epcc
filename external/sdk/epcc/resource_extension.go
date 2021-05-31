package epcc

import (
	"encoding/json"
)

func unmarshalRaw(m map[string]*json.RawMessage, key string, target interface{}) error {
	data, ok := m[key]
	if ok && data != nil {
		return json.Unmarshal(*data, &target)
	}
	return nil
}
