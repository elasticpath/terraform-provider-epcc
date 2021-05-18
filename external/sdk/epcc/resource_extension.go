package epcc

import (
	"encoding/json"
)

func ToJSON(coreData interface{}, extension map[string]interface{}) ([]byte, error) {
	// Convert core data to JSON byte array
	core, err := json.Marshal(coreData)
	if err != nil {
		return nil, err
	}

	// Convert core data JSON to dynamic map
	out := map[string]interface{}{}
	err = json.Unmarshal(core, &out)
	if err != nil {
		return nil, err
	}

	// Merge both JSON data structures
	// Add/Override keys from flows into the core resource properties
	data := out["data"].(map[string]interface{})
	for k, v := range extension {
		data[k] = v
	}

	outputJSON, _ := json.Marshal(out)
	return outputJSON, nil
}
