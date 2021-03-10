package index_geospatial

/**
 * Some tests for globegrid
 *
 *
 */

import (
	"testing"
)

func TestGlobegrid(t *testing.T) {
	NewGlobegrid("Hello", 8, 8)
	NewGlobegrid("Hello", 16, 16)
	NewGlobegrid("Hello", 32, 32)

	//	t.Error("Testing... Hmm")
}
