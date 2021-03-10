package index_geospatial

/**
 * Some tests for globegridtree
 *
 *
 */

import (
	"fmt"
	"testing"
)

func TestGlobegridtree(t *testing.T) {
	ggt := NewGlobegridtree(4, 4, 2)

	tiles := ggt.Find(myhouse)

	fmt.Printf("Find %v got %v results\n", myhouse, len(tiles))

	for i := 0; i < len(tiles); i++ {
		fmt.Printf("Result: %v\n", tiles[i])
	}

	tiles = ggt.FindRange(myhouse, 50000)

	fmt.Printf("FindRange %v got %v results\n", myhouse, len(tiles))

	//	for i := 0; i < len(tiles); i++ {
	//		fmt.Printf("Result: %v\n", tiles[i])
	//	}

	//	t.Error("Testing... Hmm")
}
