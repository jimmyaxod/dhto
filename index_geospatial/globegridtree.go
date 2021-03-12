package index_geospatial

import (
	"fmt"
	"math"
)

type Globegridtree struct {
	grids []Globegrid
}

func NewGlobegridtree(n int, num_lat int, num_lon int, mul float64) *Globegridtree {

	g := &Globegridtree{}
	g.grids = make([]Globegrid, n)

	n_lat := float64(num_lat)
	n_lon := float64(num_lon)

	for i := 0; i < n; i++ {
		id := fmt.Sprintf("R%d", i)
		gg := NewGlobegrid(id, int(n_lat), int(n_lon))
		g.grids[i] = gg
		n_lat = n_lat * mul
		n_lon = n_lon * mul
	}
	return g
}

func (g *Globegridtree) String() string {
	s := ""
	for _, grid := range g.grids {
		s = fmt.Sprintf("%s\n%s", grid.String(), s)
	}
	return fmt.Sprintf("GGT\n%s", s)
}

// GetGrid gets a specific grid
func (g *Globegridtree) GetGrid(n int) Globegrid {
	return g.grids[n]
}

// Find finds matches for a specific gridpoint
func (g *Globegridtree) Find(p Gridpoint) []Globegridtile {
	results := make([]Globegridtile, 0)

	for i := 0; i < len(g.grids); i++ {
		ggt, err := g.grids[i].Find(p)
		if err == nil {
			results = append(results, ggt)
		}
	}
	return results
}

// FindRange finds matches for a specific gridpoint with a range
func (g *Globegridtree) FindRange(p Gridpoint, distance float64) []Globegridtile {
	results := make([]Globegridtile, 0)

	for i := 0; i < len(g.grids); i++ {
		ggts, err := g.grids[i].FindRange(p, distance, false, false)
		if err == nil {
			// Add them all
			for j := 0; j < len(ggts); j++ {
				results = append(results, ggts[j])
			}
		}
	}
	return results
}

// FindRangeBestGrid finds matches for a specific gridpoint with a range using the single best grid
func (g *Globegridtree) FindRangeBestGrid(p Gridpoint, distance float64) []Globegridtile {
	results := make([]Globegridtile, 0)

	if distance == 0 {
		return results
	}

	// Get the query area
	queryArea := AreaRange(distance)

	bestGrid := g.grids[0]
	bestVal := 0.0

	// Find the best one, then use it...
	for _, grid := range g.grids {
		tileAvgArea := grid.GetAverageGridArea()
		val := math.Abs(queryArea-tileAvgArea) / math.Max(queryArea, tileAvgArea)
		if bestVal == 0 || val < bestVal {
			bestVal = val
			bestGrid = grid
		}
	}

	ggts, err := bestGrid.FindRange(p, distance, false, false)
	if err == nil {
		// Add them all
		for _, ggt := range ggts {
			results = append(results, ggt)
		}
	}

	return results
}

// FindRangeSmallestSize finds matches for a specific gridpoint with a range using the smallest area
func (g *Globegridtree) FindRangeSmallestArea(p Gridpoint, distance float64) []Globegridtile {
	results := make([]Globegridtile, 0)

	if distance == 0 {
		return results
	}

	bestGrid := g.grids[0]
	bestArea := 0.0

	// Find the best one, then use it...
	for _, grid := range g.grids {
		ggts, err := grid.FindRange(p, distance, false, false)
		if err == nil {
			area := 0.0
			for _, tile := range ggts {
				area += tile.Area()
			}

			if bestArea == 0 || area < bestArea {
				bestArea = area
				bestGrid = grid
			}
		}
	}

	ggts, err := bestGrid.FindRange(p, distance, false, false)
	if err == nil {
		// Add them all
		for _, ggt := range ggts {
			results = append(results, ggt)
		}
	}

	return results
}

// FindRangeLeastTiles finds matches for a specific gridpoint with a range using the least tiles, then smallest area
func (g *Globegridtree) FindRangeLeastTilesSmallestArea(p Gridpoint, distance float64) []Globegridtile {
	results := make([]Globegridtile, 0)

	if distance == 0 {
		return results
	}

	bestGrid := g.grids[0]
	bestTiles := 0
	bestArea := 0.0

	// Find the best one, then use it...
	for _, grid := range g.grids {
		ggts, err := grid.FindRange(p, distance, false, false)
		if err == nil {
			area := 0.0
			for _, tile := range ggts {
				area += tile.Area()
			}

			fmt.Printf("DECIDE %v %v %v\n", grid.id, area, len(ggts))

			// If it's better...
			if bestTiles == 0 || len(ggts) < bestTiles ||
				(bestTiles == len(ggts) && area < bestArea) {
				bestTiles = len(ggts)
				bestArea = area
				bestGrid = grid
			}
		}
	}

	fmt.Printf("DECIDE PICKED %v %v %v", bestGrid.id, bestArea, bestTiles)

	ggts, err := bestGrid.FindRange(p, distance, false, false)
	if err == nil {
		// Add them all
		for _, ggt := range ggts {
			results = append(results, ggt)
		}
	}

	return results
}
