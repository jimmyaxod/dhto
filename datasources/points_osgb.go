package datasources

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fofanov/go-osgb"
	"gtihub.com/jimmyaxod/dhto/index_geospatial"
)

var (
	OSGB_DATA_DIR = "data/OSGB"
)

/**
 * This handles OSGB datasets
 * Data should be in /data/OSGB/
 *
 * hydrography | landcover | landform | other | populatedPlace | transportNetwork
 */

func GetOSNames(datatype string) ([]index_geospatial.Gridpoint, error) {
	trans, err := osgb.NewOSTN15Transformer()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(fmt.Sprintf("%s/%s.csv", OSGB_DATA_DIR, datatype))
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

		// ID,NAMES_URI,NAME1,NAME1_LANG,NAME2,NAME2_LANG,TYPE,LOCAL_TYPE,GEOMETRY_X,GEOMETRY_Y,MOST_DETAIL_VIEW_RES,
		// LEAST_DETAIL_VIEW_RES,MBR_XMIN,MBR_YMIN,MBR_XMAX,MBR_YMAX,POSTCODE_DISTRICT,
		// POSTCODE_DISTRICT_URI,POPULATED_PLACE,POPULATED_PLACE_URI,POPULATED_PLACE_TYPE,
		// DISTRICT_BOROUGH,DISTRICT_BOROUGH_URI,DISTRICT_BOROUGH_TYPE,COUNTY_UNITARY,
		// COUNTY_UNITARY_URI,COUNTY_UNITARY_TYPE,REGION,REGION_URI,COUNTRY,COUNTRY_URI,
		// RELATED_SPATIAL_OBJECT,SAME_AS_DBPEDIA,SAME_AS_GEONAMES

		// 0 ID - osgb4000000074558011,
		// 1 NAMES_URI - http://data.ordnancesurvey.co.uk/id/4000000074558011,
		// 2 NAME1 - Meads,
		// 3 NAME1_LANG - ,
		// 4 NAME2 - ,
		// 5 NAME2_LANG - ,
		// 6 TYPE - populatedPlace,
		// 7 LOCAL_TYPE - Suburban Area,
		// 8 GEOMETRY_X - 560312,
		// 9 GEOMETRY_Y - 97870,
		// 10 MOST_DETAIL_VIEW_RES - 12000,
		// 11 LEAST_DETAIL_VIEW_RES - 25000,
		// 12 MBR_XMIN - 559346,
		// 13 MBR_YMIN - 96935,
		// 14 MBR_XMAX - 561257,
		// 15 MBR_YMAX - 98679,
		// 16 POSTCODE_DISTRICT - BN20,
		// 17 POSTCODE_DISTRICT_URI - http://data.ordnancesurvey.co.uk/id/postcodedistrict/BN20,
		// 18 POPULATED_PLACE - ,
		// 19 POPULATED_PLACE_URI - ,
		// 20 POPULATED_PLACE_TYPE - ,
		// 21 DISTRICT_BOROUGH - Eastbourne,
		// 22 DISTRICT_BOROUGH_URI - http://data.ordnancesurvey.co.uk/id/7000000000002632,
		// 23 DISTRICT_BOROUGH_TYPE - http://data.ordnancesurvey.co.uk/ontology/admingeo/District,
		// 24 COUNTY_UNITARY - East Sussex,
		// 25 COUNTY_UNITARY_URI - http://data.ordnancesurvey.co.uk/id/7000000000002625,
		// 26 COUNTY_UNITARY_TYPE - http://data.ordnancesurvey.co.uk/ontology/admingeo/County,
		// 27 REGION - South East,
		// 28 REGION_URI - http://data.ordnancesurvey.co.uk/id/7000000000041421,
		// 29 COUNTRY - England,
		// 30 COUNTRY_URI - http://data.ordnancesurvey.co.uk/id/country/england,
		// 31 RELATED_SPATIAL_OBJECT - ,
		// 32 SAME_AS_DBPEDIA - http://dbpedia.org/resource/Meads,
		// 33 SAME_AS_GEONAMES

		if len(fields) == 34 {
			//			id := fields[0]
			name := fields[2]

			gpid := name

			eastling, errx := strconv.ParseFloat(fields[8], 64) // eastling
			northing, erry := strconv.ParseFloat(fields[9], 64) // northing
			height := 5.0

			if errx == nil && erry == nil {
				// Convert to gps
				nationalGridCoord := osgb.NewOSGB36Coord(eastling, northing, height)
				gpsCoord, err := trans.FromNationalGrid(nationalGridCoord)
				if err == nil {
					gp := index_geospatial.NewGridpoint(gpid, gpsCoord.Lat, gpsCoord.Lon)
					data = append(data, gp)
				}
			}
		}
	}
	return data, nil
}
