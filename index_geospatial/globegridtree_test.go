package index_geospatial

/**
 * Some tests for globegridtree
 *
 *
 */

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlobegridtree(t *testing.T) {
	me := NewGridpoint("My house", 52.179413, 0.919274)

	ggt := NewGlobegridtree(4, 4, 2)

	tiles := ggt.Find(me)

	assert.Equal(t, len(tiles), 4) // Should be 1 result for each grid

	// Make sure it's contained in each tile
	for _, tile := range tiles {
		assert.True(t, tile.Contains(me))
	}
}

func TestGlobegridtreeFindRange(t *testing.T) {
	me := NewGridpoint("My house", 52.179413, 0.919274)

	ggt := NewGlobegridtree(4, 4, 2)

	tiles := ggt.FindRange(me, 50000)

	fmt.Printf("FindRange %v got %v results\n", myhouse, len(tiles))

	//	for i := 0; i < len(tiles); i++ {
	//		fmt.Printf("Result: %v\n", tiles[i])
	//	}

	//	t.Error("Testing... Hmm")
}
