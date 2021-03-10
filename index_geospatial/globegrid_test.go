package index_geospatial

/**
 * Some tests for globegrid
 *
 *
 */

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlobegrid(t *testing.T) {
	me := NewGridpoint("My house", 52.179413, 0.919274)

	gg := NewGlobegrid("Globe-256-512", 256, 512)

	ggt, err := gg.Find(me)

	assert.Nil(t, err)

	// Check we got the correct tile
	assert.Equal(t, ggt.GetHash(), "b5d7ffffffffffffffffffffffffffff-256-512")

	// Check the tile contains me
	assert.True(t, ggt.Contains(me))
}
