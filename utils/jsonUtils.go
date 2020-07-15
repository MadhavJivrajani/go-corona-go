package utils

import (
	"encoding/json"
	"fmt"
)

// GetKeys returns the keys from a JSON string.
func GetKeys(jsonArray []byte) ([]string, error) {
	unmarshalInto := make(map[string]json.RawMessage)

	err := json.Unmarshal(jsonArray, &unmarshalInto)

	var keys []string
	if err != nil {
		return keys, err
	}

	fmt.Println(string(unmarshalInto["state_wise"]))
	keys = make([]string, len(unmarshalInto))

	counter := 0
	for key := range unmarshalInto {
		keys[counter] = key
		counter++
	}
	return keys, nil
}

// GetSubJSON returns a byte array corresponding to a JSON key.
func GetSubJSON(jsonArray []byte, key string) ([]byte, error) {
	unmarshalInto := make(map[string]json.RawMessage)

	err := json.Unmarshal(jsonArray, &unmarshalInto)

	var subJSON []byte
	if err != nil {
		return subJSON, err
	}

	if _, exists := unmarshalInto[key]; !exists {
		return subJSON, fmt.Errorf("key does not exist")
	}

	return unmarshalInto[key], nil
}
