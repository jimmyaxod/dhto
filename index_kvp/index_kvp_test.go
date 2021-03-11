package index_kvp

import (
	"testing"

	"gtihub.com/jimmyaxod/dhto/datasources"
)

func TestCanIndexAndLookup(t *testing.T) {
	datasources.OSGB_DATA_DIR = "../data/OSGB"

	// TODO: Index all names, and then search by name...
}
