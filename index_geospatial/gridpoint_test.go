package index_geospatial

/**
 * Some tests for gridpoint
 *
 *
 */

import (
	"fmt"
	"testing"
)

var (
	myhouse = NewGridpoint("My house", 52.179413, 0.919274)
)

func TestGridpoint(t *testing.T) {
	gp1 := NewGridpoint("My house", 52.179413, 0.919274)
	gp2 := NewGridpoint("KFC", 52.084232, 1.116237)

	dist := gp1.DistanceTo(gp2)
	bearing := gp1.BearingTo(gp2)

	fmt.Printf("gp1 = %v\n", gp1)
	fmt.Printf("gp2 = %v\n", gp2)

	fmt.Printf("Distance to KFC = %v, bearing = %v\n", dist, bearing)

	// Move
	gp3 := gp1.MoveTo(bearing, dist)

	fmt.Printf("gp3 = %v\n", gp3)
	// We should be at KFC approx!

	//	t.Error() // to indicate test failed
}

func TestGridpoint_area(t *testing.T) {
	points := make([]Gridpoint, 5)
	points[0] = myhouse
	points[1] = myhouse.MoveTo(90, 1000)
	points[2] = myhouse.MoveTo(90, 1000).MoveTo(180, 1000)
	points[3] = myhouse.MoveTo(180, 1000)
	points[4] = myhouse

	area := Area(points)

	fmt.Printf("Area of points is %v\n", area)
}
