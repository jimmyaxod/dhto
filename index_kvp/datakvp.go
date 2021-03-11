package index_kvp

import "fmt"

// Represents a single kvp data pair.

type DataKVP struct {
	Key   string
	Value string
}

// NewDataKVP Creates a new KVP
func NewDataKVP(key string, value string) *DataKVP {
	return &DataKVP{
		Key:   key,
		Value: value,
	}
}

// String converts it to a string
func (kvp *DataKVP) String() string {
	return fmt.Sprintf("KVP Key=%s Val=%s", kvp.Key, kvp.Value)
}
