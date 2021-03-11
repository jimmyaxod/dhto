package datasources

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"gtihub.com/jimmyaxod/dhto/index_geospatial"
)

var (
	FOOD_DATA = "data/ukFoodRatings/all_food_by_geo.csv"
)

/**
 * This handles Food rating datasets
 *
 */

func GetFoodRatings() ([]index_geospatial.Gridpoint, error) {

	file, err := os.Open(FOOD_DATA)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var data []index_geospatial.Gridpoint

	// Read the datalines and create gridpoints
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")

		// lon,lat,name

		if len(fields) == 3 {
			lon, lonerr := strconv.ParseFloat(fields[0], 64)
			lat, laterr := strconv.ParseFloat(fields[1], 64)
			name := fields[2]

			gpid := name

			if lonerr == nil && laterr == nil {
				gp := index_geospatial.NewGridpoint(gpid, lat, lon)
				data = append(data, gp)
			}
		}
	}
	return data, nil
}
