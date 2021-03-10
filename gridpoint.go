package grid

/**
 * gridpoint is a unique GPS point
 * It has an ID, and lat/lon coordinates
 */

import (
	"fmt"
	"math"
)

const GLOBE_RADIUS = 6371 * 1000

var (
	NP = gridpoint{id: "North Pole", lat: 90, lon: 0}
	SP = gridpoint{id: "South Pole", lat: -90, lon: 0}
)

type gridpoint struct {
	lat float64
	lon float64
	id  string
}

func NewGridpoint(id string, lat float64, lon float64) gridpoint {
	return gridpoint{
		lat: lat,
		lon: lon,
		id:  id,
	}
}

func (gp gridpoint) String() string {
	return fmt.Sprintf("gridpoint id=%s (%.3f, %.3f)", gp.id, gp.lat, gp.lon)
}

// DistanceBetween measures distance between two gridpoints
func DistanceBetween(aa gridpoint, bb gridpoint) float64 {
	latDistance := (bb.lat - aa.lat) * math.Pi / 180
	lonDistance := (bb.lon - aa.lon) * math.Pi / 180

	a := math.Sin(latDistance/2)*math.Sin(latDistance/2) + math.Cos(aa.lat*math.Pi/180)*math.Cos(bb.lat*math.Pi/180)*math.Sin(lonDistance/2)*math.Sin(lonDistance/2)

	c := GLOBE_RADIUS * 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return c
}

// BearingBetween calculates initial bearing between two gridpoints
func BearingBetween(aa gridpoint, bb gridpoint) float64 {
	aa_lat_r := aa.lat * math.Pi / 180
	aa_lon_r := aa.lon * math.Pi / 180
	bb_lat_r := bb.lat * math.Pi / 180
	bb_lon_r := bb.lon * math.Pi / 180
	y := math.Sin(bb_lon_r-aa_lon_r) * math.Cos(bb_lat_r)
	x := math.Cos(aa_lat_r)*math.Sin(bb_lat_r) - math.Sin(aa_lat_r)*math.Cos(bb_lat_r)*math.Cos(bb_lon_r-aa_lon_r)

	b := math.Atan2(y, x)
	return math.Mod((b*180/math.Pi)+360, 360)
}

// DistanceTo measures distance to somewhere else
func (gp gridpoint) DistanceTo(dest gridpoint) float64 {
	return DistanceBetween(gp, dest)
}

func (gp gridpoint) BearingTo(dest gridpoint) float64 {
	return BearingBetween(gp, dest)
}

// MoveTo moves in a given direction and returns a new gridpoint.
func (gp gridpoint) MoveTo(bearing float64, distance float64) gridpoint {
	bearing_r := bearing * math.Pi / 180
	lat_r := gp.lat * math.Pi / 180
	lon_r := gp.lon * math.Pi / 180

	end_lat_r := math.Asin(math.Sin(lat_r)*math.Cos(distance/GLOBE_RADIUS) + math.Cos(lat_r)*math.Sin(distance/GLOBE_RADIUS)*math.Cos(bearing_r))

	end_lon_r := lon_r + math.Atan2(math.Sin(bearing_r)*math.Sin(distance/GLOBE_RADIUS)*math.Cos(lat_r), math.Cos(distance/GLOBE_RADIUS)-math.Sin(lat_r)*math.Sin(end_lat_r))

	nlat := end_lat_r * 180 / math.Pi
	nlon := math.Mod((end_lon_r*180/math.Pi)+540, 360) - 180

	return NewGridpoint(gp.id, nlat, nlon)
}

// Area calculates the area of a set of points. The last point should be first.
func Area(points []gridpoint) float64 {
	var area float64 = 0
	num := len(points)
	if num > 2 {
		for i := 0; i < num-1; i++ {
			p1 := points[i]
			p2 := points[i+1]

			area += ((p2.lon - p1.lon) * math.Pi / 180) * (2 + math.Sin(p1.lat*math.Pi/180) + math.Sin(p2.lat*math.Pi/180))
		}
		area = area * GLOBE_RADIUS * GLOBE_RADIUS / 2
	}

	return math.Abs(area)
}

// AreaRange returns the area of a circle radius l. nb for large values of l this
// isn't strictly correct as it's on the surface of a sphere. We'll see if it matters.
func AreaRange(l float64) float64 {
	return math.Pi * l * l
}
