
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
	
	function initMap() {
		const map = new google.maps.Map(document.getElementById("map"), {
		  zoom: 10,
		  center: { lat: 52.179413, lng: 0.919274 },
		  mapTypeId: "terrain",
		});
	
	var tile_shapes = [[{lat: 52.119141, lng: 0.716561},{lat: 52.207031, lng: 0.716561},{lat: 52.207031, lng: 0.859873},{lat: 52.119141, lng: 0.859873},{lat: 52.119141, lng: 0.716561}],[{lat: 52.119141, lng: 0.859873},{lat: 52.207031, lng: 0.859873},{lat: 52.207031, lng: 1.003185},{lat: 52.119141, lng: 1.003185},{lat: 52.119141, lng: 0.859873}],[{lat: 52.207031, lng: 0.789789},{lat: 52.294922, lng: 0.789789},{lat: 52.294922, lng: 0.933387},{lat: 52.207031, lng: 0.933387},{lat: 52.207031, lng: 0.789789}],[{lat: 52.207031, lng: 0.933387},{lat: 52.294922, lng: 0.933387},{lat: 52.294922, lng: 1.076984},{lat: 52.207031, lng: 1.076984},{lat: 52.207031, lng: 0.933387}]]

	drawShapes(map, tile_shapes, "#0000FF");
	
		// Create marker 
		var marker_0 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.179413, 0.919274),
			title: "My house"
		});
		
		// Create marker 
		var marker_1 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.185126, 0.989352),
			title: "Abbot's Hall School"
		});
		
		// Create marker 
		var marker_2 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.191482, 0.976244),
			title: "Ambleside"
		});
		
		// Create marker 
		var marker_3 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.180843, 0.944895),
			title: "Chestnut Horse Inn"
		});
		
		// Create marker 
		var marker_4 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.193502, 0.987837),
			title: "Chilton Community Primary School"
		});
		
		// Create marker 
		var marker_5 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.192882, 0.972079),
			title: "Chilton Meadows Care Home"
		});
		
		// Create marker 
		var marker_6 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.195167, 0.975699),
			title: "Chilton Sports Club"
		});
		
		// Create marker 
		var marker_7 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.192478, 0.986012),
			title: "China City"
		});
		
		// Create marker 
		var marker_8 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.192478, 0.986012),
			title: "Cinnamon Tree"
		});
		
		// Create marker 
		var marker_9 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.182960, 0.923340),
			title: "Crown Inn"
		});
		
		// Create marker 
		var marker_10 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.184726, 0.948841),
			title: "Finborough School"
		});
		
		// Create marker 
		var marker_11 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.192478, 0.986012),
			title: "Fishermans Basket"
		});
		
		// Create marker 
		var marker_12 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.195168, 0.975699),
			title: "Foxglove Montesorri Nursery"
		});
		
		// Create marker 
		var marker_13 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.193501, 0.987837),
			title: "Foxglove Violet Hill"
		});
		
		// Create marker 
		var marker_14 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.192478, 0.986012),
			title: "Golden Garden"
		});
		
		// Create marker 
		var marker_15 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.188440, 0.932260),
			title: "Granary Crafts"
		});
		
		// Create marker 
		var marker_16 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.177332, 0.942427),
			title: "Great Finborough Primary School"
		});
		
		// Create marker 
		var marker_17 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.168004, 0.961161),
			title: "Jackbridge Beef"
		});
		
		// Create marker 
		var marker_18 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.192478, 0.986012),
			title: "Londis Stores"
		});
		
		// Create marker 
		var marker_19 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.180822, 0.916690),
			title: "Lunch Club"
		});
		
		// Create marker 
		var marker_20 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.195167, 0.975699),
			title: "Olive AP Academy"
		});
		
		// Create marker 
		var marker_21 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.192478, 0.986012),
			title: "Pizza King"
		});
		
		// Create marker 
		var marker_22 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.149960, 0.957230),
			title: "Punch Bowl Community Inn"
		});
		
		// Create marker 
		var marker_23 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.192629, 0.891517),
			title: "Rattlesden Community Shop"
		});
		
		// Create marker 
		var marker_24 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.172287, 0.881597),
			title: "Rattlesden Gliding Club"
		});
		
		// Create marker 
		var marker_25 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.192954, 0.892458),
			title: "Rattlesden Playgroup"
		});
		
		// Create marker 
		var marker_26 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.193673, 0.889332),
			title: "Rattlesden Primary School"
		});
		
		// Create marker 
		var marker_27 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.192294, 0.981551),
			title: "SLM Community Leisure"
		});
		
		// Create marker 
		var marker_28 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.174760, 0.989970),
			title: "St Marys Church Hall"
		});
		
		// Create marker 
		var marker_29 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.195167, 0.975699),
			title: "Stepping Stones Nursery"
		});
		
		// Create marker 
		var marker_30 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.194143, 0.983138),
			title: "Stowcare Ltd @ Chilton Court"
		});
		
		// Create marker 
		var marker_31 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.190668, 0.938509),
			title: "Stowmarket Golf Club"
		});
		
		// Create marker 
		var marker_32 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.190983, 0.938134),
			title: "Stowmarket Golf Course"
		});
		
		// Create marker 
		var marker_33 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.191382, 0.979588),
			title: "Stowmarket High School"
		});
		
		// Create marker 
		var marker_34 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.193081, 0.892542),
			title: "The Brewers"
		});
		
		// Create marker 
		var marker_35 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.193883, 0.891394),
			title: "The Five Bells"
		});
		
		// Create marker 
		var marker_36 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.190536, 0.960860),
			title: "The Lakehouse Cafe"
		});
		
		// Create marker 
		var marker_37 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.190932, 0.988835),
			title: "Trudy's Sandwich Bar"
		});
		
		// Create marker 
		var marker_38 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.187362, 0.987762),
			title: "Ultimate Highseat Co"
		});
		
		// Create marker 
		var marker_39 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.195117, 0.979457),
			title: "Wood Ley Primary School"
		});
		
		// Create marker 
		var marker_40 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.153104, 0.873515),
			title: "Centre Academy East Anglia"
		});
		
		// Create marker 
		var marker_41 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.146962, 0.891875),
			title: "Stanstead Hall B & B"
		});
		
		// Create marker 
		var marker_42 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.209334, 0.866590),
			title: "Drinkstone Village Hall"
		});
		
		// Create marker 
		var marker_43 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.216061, 0.889526),
			title: "R & P Baker Ltd"
		});
		
		// Create marker 
		var marker_44 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.221670, 0.895865),
			title: "Woolpit CP School"
		});
		
		// Create marker 
		var marker_45 = new google.maps.Marker({
			map: map,
			position: new google.maps.LatLng(52.214016, 0.964676),
			title: "Travelodge"
		});
		
	  // Add circle overlay and bind to marker
	  var circle = new google.maps.Circle({
		map: map,
		radius: 5000.000000,
		fillColor: '#AA0000'
	  });
	  circle.bindTo('center', marker_0, 'position');
	  }
	