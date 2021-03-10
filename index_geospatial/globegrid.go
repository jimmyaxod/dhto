package index_geospatial

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"
)

type Globegrid struct {
	id      string
	num_lat int
	max_lon int

	// These are created in CreateGrid
	num_tiles int
	num_lon   []int
}

// NewGlobegrid creates a new globegrid
func NewGlobegrid(id string, lat int, lon int) Globegrid {
	gg := Globegrid{
		id:      id,
		num_lat: lat,
		max_lon: lon,
	}
	gg.CreateGrid()
	return gg
}

// String convert to usable string
func (gg *Globegrid) String() string {
	avg := float64(gg.GetAverageGridArea() / 1000000)
	return fmt.Sprintf("[GG %s [%d * <%d] tiles=%d tile_area=%.2f sqkm", gg.id, gg.num_lat, gg.max_lon, gg.num_tiles, avg)
}

// GetAverageGridArea returns the average grid area based on the number of tiles.
func (gg *Globegrid) GetAverageGridArea() float64 {
	return 4 * math.Pi * GLOBE_RADIUS * GLOBE_RADIUS / float64(gg.num_tiles)
}

// CreateGrid creates the grid
func (gg *Globegrid) CreateGrid() {
	// Check if it's in the cache...
	data, total, err := GetGridData(gg.num_lat, gg.max_lon)
	if err == nil {
		gg.num_lon = data
		gg.num_tiles = total
		return
	} else {
		fmt.Printf("Cache err %v\n", err)
	}

	delta_lat := 180.0 / float64(gg.num_lat)
	delta_lon_eq := 360.0 / float64(gg.max_lon)

	eq_points := make([]Gridpoint, 5)
	eq_points[0] = NewGridpoint("eq_tlc", 0, 0)
	eq_points[1] = NewGridpoint("eq_blc", delta_lat, 0)
	eq_points[2] = NewGridpoint("eq_brc", delta_lat, delta_lon_eq)
	eq_points[3] = NewGridpoint("eq_trc", 0, delta_lon_eq)
	eq_points[4] = NewGridpoint("eq_tlc", 0, 0)

	eq_tile_area := Area(eq_points)

	gg.num_lon = make([]int, gg.num_lat)
	lat_i := 0

	for lat := -90.0; lat < 90; lat += delta_lat {
		// Now see how many we can divide this ring into...
		n_grids := 0
		tile_area := float64(0)

		for n_grids = gg.max_lon; n_grids > 1; n_grids-- {
			delta_lon := 360.0 / float64(n_grids)

			points := make([]Gridpoint, 5)
			points[0] = NewGridpoint("tlc", lat, 0)
			points[1] = NewGridpoint("blc", lat+delta_lat, 0)
			points[2] = NewGridpoint("brc", lat+delta_lat, delta_lon)
			points[3] = NewGridpoint("trc", lat, delta_lon)
			points[4] = NewGridpoint("tlc", lat, 0)

			tile_area = Area(points)
			if tile_area >= eq_tile_area {
				break
			}
		}

		gg.num_lon[lat_i] = n_grids
		lat_i++
		gg.num_tiles += n_grids
	}

	// Write it to the cache so we don't have to calculate it all next time
	err = PutGridData(gg.num_lat, gg.max_lon, gg.num_lon)
	if err != nil {
		fmt.Printf("Error writing to cache %v\n", err)
	}
}

// Find finds the given point in this grid
func (gg *Globegrid) Find(p Gridpoint) (Globegridtile, error) {
	delta_lat := 180.0 / float64(gg.num_lat)
	lat_i := 0

	// NB This is suboptimal
	// TODO: Improve this to speed up if needed...
	for lat := -90.0; lat < 90; lat += delta_lat {
		if p.lat >= lat && p.lat < (lat+delta_lat) {
			// Ok we're on the right lat tile
			delta_lon := 360.0 / float64(gg.num_lon[lat_i])

			lon_i := 0
			for lon := -180.0; lon < 180; lon += delta_lon {
				if p.lon >= lon && p.lon < (lon+delta_lon) {
					id_ggt := gg.MapToID(lat_i, lon_i)
					return NewGlobegridtile(gg.id, id_ggt, lat, lat+delta_lat, lon, lon+delta_lon), nil
				}
				lon_i++
			}
		}
		lat_i++
	}
	return Globegridtile{}, errors.New("Not found")
}

// MapToID converts a map position to a hash
func (gg *Globegrid) MapToID(lat int, lon int) string {
	lat_norm := float64(lat) / float64(gg.num_lat)
	lon_norm := float64(lon) / float64(gg.num_lon[lat])

	// Combine these normalised into some sort of ID (shuffle bits)

	idLen := 16
	idBytes := make([]byte, idLen)

	idPos := 0
	idVal := 128

	for i := 0; i < idLen*8; i++ {
		v := false
		if i%2 == 0 {
			lat_norm = lat_norm * 2
			if lat_norm > 1 {
				lat_norm = lat_norm - 1
				v = true
			}
		} else {
			lon_norm = lon_norm * 2
			if lon_norm > 1 {
				lon_norm = lon_norm - 1
				v = true
			}
		}
		if v {
			// Write a 1...
			idBytes[idPos] = idBytes[idPos] | byte(idVal)
		}
		idVal = idVal >> 1
		if idVal == 0 {
			idVal = 128
			idPos++
		}
	}

	id := fmt.Sprintf("%s-%d-%d", hex.EncodeToString(idBytes), gg.num_lat, gg.max_lon)

	return id
}

// Check if something is in range (overlap)
func inRange(tile_min float64, tile_max float64, val_min float64, val_max float64) bool {
	if val_min > tile_max ||
		val_max < tile_min {
		return false
	}
	return true
}

func (gg *Globegrid) FindRange(p Gridpoint, distance float64, range_inside_tile bool, tile_inside_range bool) ([]Globegridtile, error) {
	results := make([]Globegridtile, 0)

	p_top := p.MoveTo(180, distance)
	p_bottom := p.MoveTo(0, distance)

	tolerance := 1.0 // Doesn't matter too much, just to see if we crossed over a pole

	lat_top := p_top.lat
	if math.Abs(p_top.lon-p.lon) > tolerance {
		lat_top = 90
	}
	lat_bottom := p_bottom.lat
	if math.Abs(p_bottom.lon-p.lon) > tolerance {
		lat_bottom = -90
	}

	check_all_lat := false

	p_left := p.MoveTo(270, distance)
	p_right := p.MoveTo(90, distance)
	lon_left := p_left.lon
	lon_right := p_right.lon
	// Possible that we need to check 2 ranges here, if we are at the crossover point.

	check_all_lon := false

	// If we wrapped around, fallback for now to checking all
	// TODO: Check 2 ranges quickly
	if lon_left > lon_right {
		check_all_lon = true
	}

	delta_lat := 180.0 / float64(gg.num_lat)
	lat_i := 0

	for lat := -90.0; lat < 90; lat += delta_lat {

		// Check if we're in the required range here...
		inRangeLat := inRange(lat, lat+delta_lat, lat_top, lat_bottom)

		if check_all_lat || inRangeLat {

			delta_lon := 360.0 / float64(gg.num_lon[lat_i])

			lon_i := 0
			for lon := -180.0; lon < 180; lon += delta_lon {
				inRangeLon := inRange(lon, lon+delta_lon, lon_left, lon_right)

				if check_all_lon || inRangeLon {

					id_ggt := gg.MapToID(lat_i, lon_i)
					ggt := NewGlobegridtile(gg.id, id_ggt, lat, lat+delta_lat, lon, lon+delta_lon)

					if range_inside_tile && ggt.ContainsRange(p, distance) {
						results = append(results, ggt)
					}

					if tile_inside_range && ggt.ContainedInRange(p, distance) {
						results = append(results, ggt)
					}

					if !tile_inside_range && !range_inside_tile && ggt.IntersectsRange(p, distance) {
						results = append(results, ggt)
					}
				}
				lon_i++
			}
		}
		lat_i++
	}

	return results, nil
}
