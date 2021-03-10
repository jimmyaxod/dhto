package index_geospatial

import (
	"fmt"

	"github.com/google/uuid"
	"gtihub.com/jimmyaxod/dhto/dht"
)

/**
 * This allows us to perform geospatial indexing using a DHT
 *
 */

type IndexGeospatialOptions struct {
	Num_grids      int
	First_grid_lat int
	First_grid_lon int
}

type IndexGeospatial struct {
	dht         dht.DHT
	ggt         *Globegridtree
	opts        IndexGeospatialOptions
	indexedData map[string]Gridpoint
}

func NewIndexGeospatial(dht dht.DHT, opts IndexGeospatialOptions) *IndexGeospatial {
	return &IndexGeospatial{
		dht:         dht,
		ggt:         NewGlobegridtree(opts.Num_grids, opts.First_grid_lat, opts.First_grid_lon),
		opts:        opts,
		indexedData: make(map[string]Gridpoint),
	}
}

// String
func (index *IndexGeospatial) String() string {
	return fmt.Sprintf("== IndexGeospatial ==\n"+
		"DHT = %v\n"+
		"GGT = %v\n"+
		"IndexedData size = %d\n", index.dht, index.ggt, len(index.indexedData))
}

// Index indexes a number of gridpoints
func (index *IndexGeospatial) Index(data []Gridpoint) {
	for _, item := range data {
		index.IndexItem(item)
	}
}

// IndexItem indexes a gridpoint
func (index *IndexGeospatial) IndexItem(item Gridpoint) {
	uuid := uuid.New()
	// First save it locally so we can get back to the gridpoint.
	index.indexedData[uuid.String()] = item

	// Now lookup the tiles we should index it against
	for _, tile := range index.ggt.Find(item) {
		// Index it on the dht.
		index.dht.Put(tile.GetHash(), uuid.String())
	}
}

// Query finds any gridpoints within an area
func (index *IndexGeospatial) Query(gp Gridpoint, distance float64) []Gridpoint {
	results := make([]Gridpoint, 0)

	// First we find the relevant tiles
	tiles := index.ggt.FindRangeBestGrid(gp, distance)

	// Now do the dht lookups, and add if in range
	for _, tile := range tiles {
		points := index.dht.Get(tile.GetHash())
		for _, pointID := range points {
			res, ok := index.indexedData[pointID]
			if !ok {
				fmt.Printf("Somehow we got back something we didn't index... %s\n", pointID)
			}
			// now check it's in range...
			if gp.DistanceTo(res) < distance {
				results = append(results, res)
			}
		}
	}

	return results
}
