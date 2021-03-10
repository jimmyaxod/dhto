package main

import (
	"fmt"
	"testing"
	"time"

	"gtihub.com/jimmyaxod/dhto/datasources"
	"gtihub.com/jimmyaxod/dhto/dht"
	"gtihub.com/jimmyaxod/dhto/index_geospatial"
)

var (
	myhouse = index_geospatial.NewGridpoint("My house", 52.179413, 0.919274)
)

// register a list of gridpoints on a dht using a globegridtree
func registerPointsOnDHT(data []index_geospatial.Gridpoint, ggt *index_geospatial.Globegridtree, dht dht.DHT) {
	// Register all the data on the dht using the gridtree
	for i, item := range data {
		itemID := fmt.Sprintf("transportNetwork %d", i)
		// Do a lookup to see what we need to register the point against
		for _, tile := range ggt.Find(item) {
			dht.Put(tile.GetHash(), itemID)
		}
		if i%10000 == 0 {
			fmt.Printf("Loading data %d %v\n", i, dht)
		}
	}
	fmt.Printf("Loaded data %d %v\n", len(data), dht)
}

/*
	2,4				DHT keys=2 	items=934843		1	121448	12	1000	3141593	63758058988724	2029482050	1128	0	203
	4,8				DHT keys=1 	items=934843		1	934843	12	1000	3141593	24899070830007	792562040	1360	0	1492
	8,16			DHT keys=1 	items=934843		1	934843	12	1000	3141593	6142671051605	195527292	1339	0	1434
	16,32			DHT keys=3 	items=934843		1	121448	12	1000	3141593	1585826544460	50478427	1205	0	182
	32,64			DHT keys=5 	items=934843		1	121448	12	1000	3141593	392341316223	12488612	1257	0	190
	64,128			DHT keys=11 items=934843		1	121074	12	1000	3141593	98733696817		3142791		1350	0	188
	128,256			DHT keys=27 items=934843		1	84795	12	1000	3141593	24452003514		778331		1570	0	158
	256,512			DHT keys=74 items=934843		1	21292	12	1000	3141593	6122798376		194895		1874	0	37
	512,1024		DHT keys=236 items=934843		1	6465	12	1000	3141593	1529440796		48684		2804	0	11
	1024,2048		DHT keys=768 items=934843		1	1368	12	1000	3141593	382202186		12166		3572	0	2
	2048,4096		DHT keys=2495 items=934843		1	448		12	1000	3141593	95532341		3041		5851	0	0
	4096,8192		DHT keys=7967 items=934843		2	346		12	1000	3141593	47761600		1520		10679	0	0
	8192,16384		DHT keys=24357 items=934843		3	48		12	1000	3141593	17909698		570			18943	0	0
	16384,32768		DHT keys=64695 items=934843		7	36		12	1000	3141593	10446912		333			35318	0	0
	32768,65536		DHT keys=142355 items=934843	16	13		12	1000	3141593	5969525			190			67644	0	0
	65536,131072	DHT keys=292813 items=934843	48	12		12	1000	3141593	4477117			143			130746	1	0

*/

// TestLookup do some fun lookups and check everything works
func TestLookup(t *testing.T) {

	index_geospatial.CacheDIR = "./gridcache"

	dht := dht.NewDHTSimple()

	// Lets create a gridtree so we can figure out where things are on a few different grids
	ggt := index_geospatial.NewGlobegridtree(1, 65536, 131072)

	// Load up some OS data to play with
	data, err := datasources.GetOSNames("transportNetwork")

	if err != nil {
		t.Errorf("Can't load %v\n", err)
	}

	// Register all the data on the dht using the gridtree
	regCtime := time.Now()
	registerPointsOnDHT(data, ggt, dht)
	regTime := time.Since(regCtime)

	// Now try a range lookup and make sure we get back what we think we should...
	distance := 1000.0

	findCtime := time.Now()
	tiles := ggt.FindRange(myhouse, distance)
	findTime := time.Since(findCtime)

	total_results := 0
	valid_results := 0
	dht_lookups := 0
	queryCtime := time.Now()

	area_query := index_geospatial.AreaRange(distance)
	area_lookup := 0.0

	for _, tile := range tiles {
		area_lookup += tile.Area()
		fmt.Printf("Doing DHT lookup for %v\n", tile)
		items := dht.Get(tile.GetHash())
		dht_lookups++
		for _, item := range items {
			// Now lookup to find the specific data, and double check...
			var v int
			fmt.Sscanf(item, "transportNetwork %d", &v)

			gp := data[v]
			if gp.DistanceTo(myhouse) < distance {
				fmt.Printf(" + %s\n", gp)
				valid_results++
			}
			total_results++
		}
	}

	queryTime := time.Since(queryCtime)

	// GG info
	fmt.Printf("%v\n", ggt)

	// DHT info
	fmt.Printf("%v\t", dht)

	// Lookup results - num, results, valid results
	fmt.Printf("%d\t%d\t%d\t", dht_lookups, total_results, valid_results)

	// distance, area_query, area_lookup, perc
	perc := area_lookup * 100 / area_query
	fmt.Printf("%.0f\t%.0f\t%.0f\t%.0f\t", distance, area_query, area_lookup, perc)

	fmt.Printf("%d\t%d\t%d\t", regTime.Milliseconds(), findTime.Milliseconds(), queryTime.Milliseconds())

	fmt.Printf("\n")
}
