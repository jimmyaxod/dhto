package index_geospatial

/**
 * globegridtile defines a tile in gps space.
 * It contains an ID and hash, then min/max lon/lat coords.
 *
 */

import (
	"fmt"
)

type Globegridtile struct {
	id      string
	hash    string
	min_lat float64
	max_lat float64
	min_lon float64
	max_lon float64
}

// NewGlobegridtile creates a new tile
func NewGlobegridtile(id string, hash string, min_lat float64, max_lat float64, min_lon float64, max_lon float64) Globegridtile {
	return Globegridtile{
		id:      id,
		hash:    hash,
		min_lat: min_lat,
		max_lat: max_lat,
		min_lon: min_lon,
		max_lon: max_lon,
	}
}

// ContainsCompletely returns true if the arg is in this tile.
func (ggt Globegridtile) ContainsCompletely(ggt2 Globegridtile) bool {
	if ggt2.min_lat < ggt.min_lat ||
		ggt2.min_lat > ggt.max_lat ||
		ggt2.max_lat < ggt.min_lat ||
		ggt2.max_lat > ggt.max_lat ||
		ggt2.min_lon < ggt.min_lon ||
		ggt2.min_lon > ggt.max_lon ||
		ggt2.max_lon < ggt.min_lon ||
		ggt2.max_lon > ggt.max_lon {
		return false
	}
	return true
}

// Contains returns true if the point is in this tile
func (ggt Globegridtile) Contains(gp Gridpoint) bool {
	return ggt.ContainsLatLon(gp.lat, gp.lon)
}

// Contains returns true if the point is in this tile
func (ggt Globegridtile) ContainsLatLon(lat float64, lon float64) bool {
	return ggt.ContainsLat(lat) &&
		ggt.ContainsLon(lon)
}

// ContainsLat returns true if lat is contained
func (ggt Globegridtile) ContainsLat(lat float64) bool {
	return (lat >= ggt.min_lat && lat < ggt.max_lat)
}

// ContainsLon returns true if lon is contained
func (ggt Globegridtile) ContainsLon(lon float64) bool {
	return (lon >= ggt.min_lon && lon < ggt.max_lon)
}

// ContainedInRange returns true if the tile is contained in a range
func (ggt Globegridtile) ContainedInRange(gp Gridpoint, dist float64) bool {
	d1 := DistanceBetween(gp, NewGridpoint("tlc", ggt.min_lat, ggt.min_lon))
	d2 := DistanceBetween(gp, NewGridpoint("trc", ggt.min_lat, ggt.max_lon))
	d3 := DistanceBetween(gp, NewGridpoint("blc", ggt.max_lat, ggt.min_lon))
	d4 := DistanceBetween(gp, NewGridpoint("brc", ggt.max_lat, ggt.max_lon))

	return (d1 <= dist && d2 <= dist && d3 <= dist && d4 <= dist)
}

// ContainsRange returns true if the range is contained in a tile
func (ggt Globegridtile) ContainsRange(gp Gridpoint, dist float64) bool {
	top := gp.MoveTo(0, dist)
	right := gp.MoveTo(90, dist)
	bottom := gp.MoveTo(180, dist)
	left := gp.MoveTo(270, dist)

	return ggt.ContainsLatLon(top.lat, top.lon) &&
		ggt.ContainsLatLon(right.lat, right.lon) &&
		ggt.ContainsLatLon(bottom.lat, bottom.lon) &&
		ggt.ContainsLatLon(left.lat, left.lon)
}

// IntersectsRange returns true if the range intersects this tile
func (ggt Globegridtile) IntersectsRange(gp Gridpoint, dist float64) bool {
	// First easy bit, if the gp is inside the ggt...
	if ggt.ContainsLatLon(gp.lat, gp.lon) {
		return true
	}

	// OK so the gp is outside the ggt. Either a corner is in range, or a side is in range.

	// Check if any of the corners are...
	d1 := DistanceBetween(gp, NewGridpoint("tlc", ggt.min_lat, ggt.min_lon))
	d2 := DistanceBetween(gp, NewGridpoint("trc", ggt.min_lat, ggt.max_lon))
	d3 := DistanceBetween(gp, NewGridpoint("blc", ggt.max_lat, ggt.min_lon))
	d4 := DistanceBetween(gp, NewGridpoint("brc", ggt.max_lat, ggt.max_lon))

	if d1 <= dist ||
		d2 <= dist ||
		d3 <= dist ||
		d4 <= dist {
		return true
	}

	// Lastly we check sides...

	if gp.lat >= ggt.min_lat && gp.lat < ggt.max_lat {
		// lat is within range, so use gp.lat
		if gp.lon > ggt.max_lon {
			d := DistanceBetween(gp, NewGridpoint("maxlon", gp.lat, ggt.max_lon))
			if d <= dist {
				return true
			}
		} else if gp.lon < ggt.min_lon {
			d := DistanceBetween(gp, NewGridpoint("minlon", gp.lat, ggt.min_lon))
			if d <= dist {
				return true
			}
		}
	}

	if gp.lon >= ggt.min_lon && gp.lon < ggt.max_lon {
		// lon is within range, so use gp.lon
		if gp.lat > ggt.max_lat {
			d := DistanceBetween(gp, NewGridpoint("maxlat", ggt.max_lat, gp.lon))
			if d <= dist {
				return true
			}
		} else if gp.lat < ggt.min_lat {
			d := DistanceBetween(gp, NewGridpoint("minlon", ggt.min_lat, gp.lon))
			if d <= dist {
				return true
			}
		}

	}

	return false
}

func (ggt Globegridtile) GetHash() string {
	return ggt.hash
}

func (ggt Globegridtile) Area() float64 {
	points := make([]Gridpoint, 5)
	points[0] = NewGridpoint("tlc", ggt.min_lat, ggt.min_lon)
	points[1] = NewGridpoint("trc", ggt.min_lat, ggt.max_lon)
	points[2] = NewGridpoint("brc", ggt.max_lat, ggt.max_lon)
	points[3] = NewGridpoint("blc", ggt.max_lat, ggt.min_lon)
	points[4] = NewGridpoint("tlc", ggt.min_lat, ggt.min_lon)
	return Area(points)
}

func (ggt Globegridtile) String() string {
	area := ggt.Area()
	return fmt.Sprintf("Globegridtile id=%s hash=%s (%.3f, %.3f) - (%.3f, %.3f) [area %.2f sqm]", ggt.id, ggt.hash, ggt.min_lat, ggt.min_lon, ggt.max_lat, ggt.max_lon, area)
}
