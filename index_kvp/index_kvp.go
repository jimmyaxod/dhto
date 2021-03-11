package index_kvp

import (
	"fmt"

	"gtihub.com/jimmyaxod/dhto/dht"
)

/**
 * Allows us to index kvp on a dht.
 * This allows us to search by exact kvp value.
 *
 *
 */

type IndexKVP struct {
	dht         dht.DHT
	indexedData map[string]Dataitem
}

func NewIndexKVP(dht dht.DHT) *IndexKVP {
	return &IndexKVP{
		dht: dht,
	}
}

// Index a group of Dataitems
func (index *IndexKVP) Index(items []Dataitem) {
	for _, item := range items {
		index.IndexItem(item)
	}
}

// Index a single Dataitem
func (index *IndexKVP) IndexItem(item Dataitem) {
	// TODO: Index some kvp items
}

// Lookup by kvp
func (index *IndexKVP) Lookup(kvp DataKVP) []Dataitem {
	// TODO
	return nil
}

func (index *IndexKVP) String() string {
	return fmt.Sprintf("== IndexKVP ==\n"+
		"DHT = %v\n"+
		"IndexedData size = %d\n", index.dht, len(index.indexedData))
}
