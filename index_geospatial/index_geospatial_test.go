package index_geospatial_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gtihub.com/jimmyaxod/dhto/datasources"
	"gtihub.com/jimmyaxod/dhto/dht"
	"gtihub.com/jimmyaxod/dhto/index_geospatial"
)

func TestCanIndexAndLookup(t *testing.T) {
	datasources.OSGB_DATA_DIR = "../data/OSGB"

	dht := dht.NewDHTSimple()
	opts := index_geospatial.IndexGeospatialOptions{
		Num_grids:      12,
		First_grid_lat: 2,
		First_grid_lon: 4,
	}
	igeo := index_geospatial.NewIndexGeospatial(dht, opts)

	// Load up some OS data to play with
	data, err := datasources.GetOSNames("transportNetwork")

	if err != nil {
		t.Errorf("Can't load %v\n", err)
	}

	// Index all the data points
	igeo.Index(data)

	// Now do a query for some...
	me := index_geospatial.NewGridpoint("My house", 52.179413, 0.919274)
	results := igeo.Query(me, 1000.0)

	fmt.Printf("Got %d results\n", len(results))

	for _, gp := range results {
		dist := gp.DistanceTo(me)
		fmt.Printf("Result %.3f - %v\n", dist, gp)
	}

	assert.True(t, false)
}
