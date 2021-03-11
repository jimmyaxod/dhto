package index_geospatial

import (
	"fmt"
	"os"
)

/**
 * Utils to write to a gmap js file so we can view results
 *
 */

type Gmapdata struct {
	Center   Gridpoint
	Distance float64
	Markers  []Gridpoint
	Tiles    []Globegridtile
}

func GmapWrite(out string, data Gmapdata) error {
	f, err := os.Create(out)
	if err != nil {
		return err
	}

	// Write general header...
	_, err = f.WriteString(`
	function drawShapes(map, shapes, color) {
		for(i=0;i<shapes.length;i++) {
			const shape = new google.maps.Polyline({path: shapes[i],
					geodesic: true,
					strokeColor: color,
					strokeOpacity: 1.0,
					strokeWeight: 2,
				  });
			shape.setMap(map);
		}
	}
	`)
	if err != nil {
		return err
	}

	_, err = f.WriteString(fmt.Sprintf(`
	function initMap() {
		const map = new google.maps.Map(document.getElementById("map"), {
		  zoom: 10,
		  center: { lat: %f, lng: %f },
		  mapTypeId: "terrain",
		});
	`, data.Center.lat, data.Center.lon))
	if err != nil {
		return err
	}

	// Write the tiles
	tile_shapes := "["
	for i, tile := range data.Tiles {
		if i > 0 {
			tile_shapes = tile_shapes + ","
		}
		// Now write the tile...
		tile_shape := fmt.Sprintf("[{lat: %f, lng: %f},"+
			"{lat: %f, lng: %f},"+
			"{lat: %f, lng: %f},"+
			"{lat: %f, lng: %f},"+
			"{lat: %f, lng: %f}]",
			tile.min_lat, tile.min_lon,
			tile.max_lat, tile.min_lon,
			tile.max_lat, tile.max_lon,
			tile.min_lat, tile.max_lon,
			tile.min_lat, tile.min_lon)
		tile_shapes = tile_shapes + tile_shape
	}
	tile_shapes = tile_shapes + "]\n"

	_, err = f.WriteString(fmt.Sprintf(`
	var tile_shapes = %s
	drawShapes(map, tile_shapes, "#0000FF");
	`, tile_shapes))
	if err != nil {
		return err
	}

	for i, marker := range data.Markers {
		_, err = f.WriteString(fmt.Sprintf(`
		// Create marker 
		var marker_%d = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(%f, %f),
			title: "%s"
		});
		`, i, marker.lat, marker.lon, marker.id))
		if err != nil {
			return err
		}
	}

	// For now just applies to the first marker
	_, err = f.WriteString(fmt.Sprintf(`
	  // Add circle overlay and bind to marker
	  var circle = new google.maps.Circle({
		map: map,
		radius: %f,
		fillColor: '#AA0000'
	  });
	  circle.bindTo('center', marker_0, 'position');
	  }
	`, data.Distance))
	if err != nil {
		return err
	}

	f.Close()
	return nil
}
