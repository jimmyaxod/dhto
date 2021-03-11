package index_kvp

import (
	"fmt"

	"github.com/google/uuid"
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
		dht:         dht,
		indexedData: make(map[string]Dataitem),
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
	uuid := uuid.New()
	index.indexedData[uuid.String()] = item

	kvps := item.GetKVPs()
	for _, kvp := range kvps {
		hash := kvp.String() // TODO: Fix this...
		index.dht.Put(hash, uuid.String())
	}
}

// Lookup by kvp
func (index *IndexKVP) Lookup(kvp DataKVP) []Dataitem {
	// Lookup the kvp on the dht...
	hash := kvp.String()
	itemIDs := index.dht.Get(hash)
	results := make([]Dataitem, 0)

	for _, id := range itemIDs {
		// lookup the dataitem
		res, ok := index.indexedData[id]
		if ok {
			results = append(results, res)
		}
	}
	return results
}

func (index *IndexKVP) String() string {
	return fmt.Sprintf("== IndexKVP ==\n"+
		"DHT = %v\n"+
		"IndexedData size = %d\n", index.dht, len(index.indexedData))
}
