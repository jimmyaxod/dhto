package index_geospatial_test

import (
	"fmt"
	"testing"
	"time"

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
	distance := 5000.0
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

func TestCanIndexAndLookupPerformance(t *testing.T) {
	datasources.OSGB_DATA_DIR = "../data/OSGB"
	datasources.FOOD_DATA = "../data/ukFoodRatings/all_food_by_geo.csv"

	dht := dht.NewDHTSimple()
	opts := index_geospatial.IndexGeospatialOptions{
		Num_grids:      6,
		First_grid_lat: 1024,
		First_grid_lon: 2048,
	}
	igeo := index_geospatial.NewIndexGeospatial(dht, opts)

	// Load up some food ratings to play with
	data, err := datasources.GetFoodRatings()

	if err != nil {
		t.Errorf("Can't load %v\n", err)
	}

	// Index all the data points
	igeo.Index(data)

	// Now do a query for some...
	me := index_geospatial.NewGridpoint("My house", 52.179413, 0.919274)
	distances := []float64{1000.0, 2000.0, 3000.0, 4000.0, 5000.0, 10000.0, 20000.0, 30000.0}

	// query_time:	60ms (average of 0.0293ms per query)
	// num_queries: 2048
	// total_tiles:	11473 (average of 5.6 per query)
	// total_results: 393

	total_results := 0
	total_tiles := 0
	num_queries := 0
	start_time := time.Now()

	for _, distance := range distances {
		for x := 0; x < 16; x++ {
			for y := 0; y < 16; y++ {
				results, tiles := igeo.Query(me, distance)
				num_queries++
				total_results += len(results)
				total_tiles += len(tiles)
				me = me.MoveTo(0, 1000.0)
			}
			// Move back...
			me = me.MoveTo(180, 8*1000.0)

			me = me.MoveTo(90, 1000.0)
		}
	}

	query_time := time.Since(start_time)

	fmt.Printf("Query time    : %d\n", query_time.Milliseconds())
	fmt.Printf("num_queries   : %d\n", num_queries)
	fmt.Printf("total_tiles   : %d\n", total_tiles)
	fmt.Printf("total_results : %d\n", total_results)

	fmt.Printf("Index %v\n", igeo)

	assert.True(t, false)
}
