package index_geospatial

import (
	"fmt"
	"math"
)

type Globegridtree struct {
	grids []Globegrid
}

func NewGlobegridtree(n int, num_lat int, num_lon int) *Globegridtree {

	g := &Globegridtree{}
	g.grids = make([]Globegrid, n)

	for i := 0; i < n; i++ {
		id := fmt.Sprintf("R%d", i)
		gg := NewGlobegrid(id, num_lat, num_lon)
		g.grids[i] = gg
		num_lat = num_lat * 2
		num_lon = num_lon * 2
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
