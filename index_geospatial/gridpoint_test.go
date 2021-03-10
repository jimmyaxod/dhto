package index_geospatial

/**
 * Some tests for Gridpoint
 *
 *
 */

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	myhouse = NewGridpoint("My house", 52.179413, 0.919274)
)

// Check if 2 values are close enough
func CloseEnough(t *testing.T, v1 float64, v2 float64, maxDiff float64) {
	// Special case
	if v1 == 0 && v2 == 0 {
		return
	}

	// eg 100, 120 then p = 16.67 %
	p := 100.0 * math.Abs(v1-v2) / math.Max(v1, v2)
	assert.True(t, p < maxDiff)
}

func TestGridpoint(t *testing.T) {
	me := NewGridpoint("My house", 52.179413, 0.919274)
	KFC := NewGridpoint("KFC", 52.084232, 1.116237)

	dist := me.DistanceTo(KFC)
	bearing := me.BearingTo(KFC)

	CloseEnough(t, dist, 17110, 1)
	CloseEnough(t, bearing, 128.13, 1)

	// Move
	meMoved := me.MoveTo(bearing, dist)

	// We should be at KFC approx!
	CloseEnough(t, meMoved.lat, KFC.lat, 1)
	CloseEnough(t, meMoved.lon, KFC.lon, 1)
}

// TestGridpoint_area tests that the area calculation is ok
func TestGridpoint_area(t *testing.T) {
	distance := 1000.0

	points := make([]Gridpoint, 5)
	points[0] = myhouse
	points[1] = myhouse.MoveTo(90, distance)
	points[2] = myhouse.MoveTo(90, distance).MoveTo(180, distance)
	points[3] = myhouse.MoveTo(180, distance)
	points[4] = myhouse

	area := Area(points)

	// NB It's not exactly, as it's on the surface of a sphere.
	CloseEnough(t, area, distance*distance, 1)
}
