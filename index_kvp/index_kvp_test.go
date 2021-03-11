package index_kvp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gtihub.com/jimmyaxod/dhto/dht"
)

func TestCanIndexAndLookup(t *testing.T) {
	//	datasources.OSGB_DATA_DIR = "../data/OSGB"

	// TODO: Index all names, and then search by name...

	dht := dht.NewDHTSimple()
	index := NewIndexKVP(dht)

	// Make a couple of test ones for now...

	di1 := NewDataitem()
	di1.AddKVP(*NewDataKVP("name", "jimmy"))
	di1.AddKVP(*NewDataKVP("type", "human"))
	di1.AddKVP(*NewDataKVP("kids", "4"))

	di2 := NewDataitem()
	di2.AddKVP(*NewDataKVP("name", "buddy"))
	di2.AddKVP(*NewDataKVP("type", "dog"))

	di3 := NewDataitem()
	di3.AddKVP(*NewDataKVP("name", "angyal"))
	di3.AddKVP(*NewDataKVP("type", "dog"))

	items := []Dataitem{di1, di2, di3}
	index.Index(items)

	// Now try some lookups...

	res1 := index.Lookup(*NewDataKVP("type", "dog"))

	// 2 dogs...
	assert.Equal(t, len(res1), 2)

	res2 := index.Lookup(*NewDataKVP("type", "human"))

	// 1 human...
	assert.Equal(t, len(res2), 1)

}
