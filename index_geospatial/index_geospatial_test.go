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
	datasources.FOOD_DATA = "../data/ukFoodRatings/all_food_by_geo.csv"

	dht := dht.NewDHTSimple()
	opts := index_geospatial.IndexGeospatialOptions{
		Num_grids:      6,
		First_grid_lat: 1024,
		First_grid_lon: 2048,
	}
	igeo := index_geospatial.NewIndexGeospatial(dht, opts)

	// Load up some OS data to play with
	//data, err := datasources.GetOSNames("transportNetwork")

	// Load up some food ratings to play with
	data, err := datasources.GetFoodRatings()

	if err != nil {
		t.Errorf("Can't load %v\n", err)
	}

	// Index all the data points
	igeo.Index(data)

	// Now do a query for some...
	me := index_geospatial.NewGridpoint("My house", 52.179413, 0.919274)
	distance := 10000.0
	results, tiles := igeo.Query(me, distance)

	fmt.Printf("Got %d results\n", len(results))

	for _, gp := range results {
		dist := gp.DistanceTo(me)
		fmt.Printf("Result %.3f - %v\n", dist, gp)
	}

	fmt.Printf("Index %v\n", igeo)

	markers := make([]index_geospatial.Gridpoint, 0)
	markers = append(markers, me) // First one is me

	for _, gp := range results {
		markers = append(markers, gp)
	}

	gmapdata := index_geospatial.Gmapdata{
		Center:   me,
		Distance: distance,
		Markers:  markers,
		Tiles:    tiles,
	}
	index_geospatial.GmapWrite("../gmap/test_output.js", gmapdata)

	assert.True(t, false)
}
