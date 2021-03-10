package index_geospatial

import "fmt"

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
