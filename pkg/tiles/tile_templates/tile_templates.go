package tile_templates

import (
	tiles "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/buildings"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/connection"
)

func MonasteryWithoutRoads() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{{
					Sides: []connection.Side{
						connection.TOP_LEFT_EDGE,
						connection.TOP_RIGHT_EDGE,

						connection.RIGHT_TOP_EDGE,
						connection.RIGHT_BOTTOM_EDGE,

						connection.LEFT_TOP_EDGE,
						connection.LEFT_BOTTOM_EDGE,

						connection.BOTTOM_LEFT_EDGE,
						connection.BOTTOM_RIGHT_EDGE,
					},
				}},
			},
		},
		HasShield: false,
		Building:  buildings.MONASTERY,
	}
}

/*
returns tiles.Tile having monastery and road going bottom
*/
func MonasteryWithSingleRoad() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{
					{Sides: []connection.Side{
						connection.CENTER,
						connection.BOTTOM,
					},
					},
				},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP_LEFT_EDGE,
							connection.TOP_RIGHT_EDGE,

							connection.RIGHT_TOP_EDGE,
							connection.RIGHT_BOTTOM_EDGE,

							connection.LEFT_TOP_EDGE,
							connection.LEFT_BOTTOM_EDGE,

							connection.BOTTOM_LEFT_EDGE,
							connection.BOTTOM_RIGHT_EDGE,
						},
					},
				},
			},
		},
		HasShield: false,
		Building:  buildings.MONASTERY,
	}
}

/*
returns tiles.Tile having road from left to right
*/
func StraightRoads() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{
					{Sides: []connection.Side{
						connection.LEFT,
						connection.RIGHT,
					},
					},
				},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_BOTTOM_EDGE,
							connection.BOTTOM_LEFT_EDGE,
							connection.BOTTOM_RIGHT_EDGE,
							connection.RIGHT_BOTTOM_EDGE,
						},
					},
					{
						Sides: []connection.Side{
							connection.LEFT_TOP_EDGE,
							connection.TOP_LEFT_EDGE,
							connection.TOP_RIGHT_EDGE,
							connection.RIGHT_TOP_EDGE,
						},
					},
				},
			},
		},
		HasShield: false,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having road from left to bottom
*/
func RoadsTurn() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{
					{Sides: []connection.Side{
						connection.LEFT,
						connection.BOTTOM,
					},
					},
				},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_BOTTOM_EDGE,
							connection.BOTTOM_LEFT_EDGE,
						},
					},
					{
						Sides: []connection.Side{
							connection.LEFT_TOP_EDGE,
							connection.TOP_LEFT_EDGE,
							connection.TOP_RIGHT_EDGE,
							connection.RIGHT_TOP_EDGE,
							connection.RIGHT_BOTTOM_EDGE,
							connection.BOTTOM_RIGHT_EDGE,
						},
					},
				},
			},
		},
		HasShield: false,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having road from left to bottom
*/
func TCrossRoad() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT,
							connection.CENTER,
						},
					},
					{
						Sides: []connection.Side{
							connection.BOTTOM,
							connection.CENTER,
						},
					},
					{
						Sides: []connection.Side{
							connection.RIGHT,
							connection.CENTER,
						},
					},
				},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_BOTTOM_EDGE,
							connection.BOTTOM_LEFT_EDGE,
						},
					},
					{
						Sides: []connection.Side{
							connection.RIGHT_BOTTOM_EDGE,
							connection.BOTTOM_RIGHT_EDGE,
						},
					},
					{
						Sides: []connection.Side{
							connection.LEFT_TOP_EDGE,
							connection.TOP_LEFT_EDGE,
							connection.TOP_RIGHT_EDGE,
							connection.RIGHT_TOP_EDGE,
						},
					},
				},
			},
		},
		HasShield: false,
		Building:  buildings.NONE_BULDING,
	}
}

func XCrossRoad() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT,
							connection.CENTER,
						},
					},
					{
						Sides: []connection.Side{
							connection.BOTTOM,
							connection.CENTER,
						},
					},

					{
						Sides: []connection.Side{
							connection.RIGHT,
							connection.CENTER,
						},
					},
					{
						Sides: []connection.Side{
							connection.TOP,
							connection.CENTER,
						},
					},
				},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_BOTTOM_EDGE,
							connection.BOTTOM_LEFT_EDGE,
						},
					},
					{
						Sides: []connection.Side{
							connection.RIGHT_BOTTOM_EDGE,
							connection.BOTTOM_RIGHT_EDGE,
						},
					},
					{
						Sides: []connection.Side{
							connection.LEFT_TOP_EDGE,
							connection.TOP_LEFT_EDGE,
						},
					},

					{
						Sides: []connection.Side{
							connection.TOP_RIGHT_EDGE,
							connection.RIGHT_TOP_EDGE,
						},
					},
				},
			},
		},
		HasShield: false,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having single city edge on top
*/
func SingleCityEdgeNoRoads() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP,
						},
					},
				},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_TOP_EDGE,
							connection.RIGHT_TOP_EDGE,
							connection.RIGHT_BOTTOM_EDGE,
							connection.BOTTOM_RIGHT_EDGE,
							connection.LEFT_BOTTOM_EDGE,
							connection.BOTTOM_LEFT_EDGE,
						},
					},
				},
			},
		},
		HasShield: false,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having single city edge on top and road from left to right
*/
func SingleCityEdgeStraightRoads() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP,
						},
					},
				},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.RIGHT,
							connection.LEFT,
						},
					},
				},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.RIGHT_BOTTOM_EDGE,
							connection.BOTTOM_RIGHT_EDGE,
							connection.LEFT_BOTTOM_EDGE,
							connection.BOTTOM_LEFT_EDGE,
						},
					},
					{
						Sides: []connection.Side{
							connection.LEFT_TOP_EDGE,
							connection.RIGHT_TOP_EDGE,
						},
					},
				},
			},
		},
		HasShield: false,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having single city edge on top and road from left to bottom
*/
func SingleCityEdgeLeftRoadTurn() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP,
						},
					},
				},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT,
							connection.BOTTOM,
						},
					},
				},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.RIGHT_BOTTOM_EDGE,
							connection.BOTTOM_RIGHT_EDGE,
							connection.LEFT_TOP_EDGE,
							connection.RIGHT_TOP_EDGE,
						},
					},
					{
						Sides: []connection.Side{
							connection.BOTTOM_LEFT_EDGE,
							connection.LEFT_BOTTOM_EDGE,
						},
					},
				},
			},
		},
		HasShield: false,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having single city edge on top and road from right to bottom
*/
func SingleCityEdgeRightRoadTurn() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP,
						},
					},
				},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.RIGHT,
							connection.BOTTOM,
						},
					},
				},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_TOP_EDGE,
							connection.RIGHT_TOP_EDGE,
							connection.BOTTOM_LEFT_EDGE,
							connection.LEFT_BOTTOM_EDGE,
						},
					},
					{
						Sides: []connection.Side{
							connection.RIGHT_BOTTOM_EDGE,
							connection.BOTTOM_RIGHT_EDGE,
						},
					},
				},
			},
		},
		HasShield: false,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having single city edge on top and roads on other sides
*/
func SingleCityEdgeCrossRoad() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP,
						},
					},
				},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.RIGHT,
							connection.CENTER,
						},
					},
					{
						Sides: []connection.Side{
							connection.LEFT,
							connection.CENTER,
						},
					},
					{
						Sides: []connection.Side{
							connection.BOTTOM,
							connection.CENTER,
						},
					},
				},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_TOP_EDGE,
							connection.RIGHT_TOP_EDGE,
						},
					},
					{
						Sides: []connection.Side{
							connection.RIGHT_BOTTOM_EDGE,
							connection.BOTTOM_RIGHT_EDGE,
						},
					},
					{
						Sides: []connection.Side{
							connection.BOTTOM_LEFT_EDGE,
							connection.LEFT_BOTTOM_EDGE,
						},
					},
				},
			},
		},
		HasShield: false,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having city edges on top and bottom. Not connected
*/
func TwoCityEdgesUpAndDownNotConnected() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP,
						},
					},
					{
						Sides: []connection.Side{
							connection.BOTTOM,
						},
					},
				},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_TOP_EDGE,
							connection.RIGHT_TOP_EDGE,
							connection.LEFT_BOTTOM_EDGE,
							connection.RIGHT_BOTTOM_EDGE,
						},
					},
				},
			},
		},
		HasShield: false,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having city edges on top and right. Not connected
*/
func TwoCityEdgesCornerNotConnected() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP,
						},
					},
					{
						Sides: []connection.Side{
							connection.RIGHT,
						},
					},
				},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_TOP_EDGE,
							connection.LEFT_BOTTOM_EDGE,
							connection.BOTTOM_LEFT_EDGE,
							connection.BOTTOM_RIGHT_EDGE,
						},
					},
				},
			},
		},
		HasShield: false,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having city edges on top and down. Connected
*/
func TwoCityEdgesUpAndDownConnected() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP,
							connection.BOTTOM,
							connection.CENTER,
						},
					},
				},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_TOP_EDGE,
							connection.LEFT_BOTTOM_EDGE,
						},
					},
					{
						Sides: []connection.Side{
							connection.BOTTOM_LEFT_EDGE,
							connection.BOTTOM_RIGHT_EDGE,
						},
					},
				},
			},
		},
		HasShield: false,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having city edges on top and down. Connected and shield
*/
func TwoCityEdgesUpAndDownConnectedShield() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP,
							connection.BOTTOM,
							connection.CENTER,
						},
					},
				},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_TOP_EDGE,
							connection.LEFT_BOTTOM_EDGE,
						},
					},
					{
						Sides: []connection.Side{
							connection.BOTTOM_LEFT_EDGE,
							connection.BOTTOM_RIGHT_EDGE,
						},
					},
				},
			},
		},
		HasShield: true,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having city edges on top and right. Connected
*/
func TwoCityEdgesCornerConnected() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP,
							connection.RIGHT,
						},
					},
				},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_TOP_EDGE,
							connection.LEFT_BOTTOM_EDGE,
							connection.BOTTOM_LEFT_EDGE,
							connection.BOTTOM_RIGHT_EDGE,
						},
					},
				},
			},
		},
		HasShield: false,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having city edges on top and right. Connected and shield
*/

func TwoCityEdgesCornerConnectedShield() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP,
							connection.RIGHT,
						},
					},
				},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_TOP_EDGE,
							connection.LEFT_BOTTOM_EDGE,
							connection.BOTTOM_LEFT_EDGE,
							connection.BOTTOM_RIGHT_EDGE,
						},
					},
				},
			},
		},
		HasShield: true,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having city edges on top and right. Connected but also road from left to bottom
*/
func TwoCityEdgesCornerConnectedRoadTurn() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP,
							connection.RIGHT,
						},
					},
				},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT,
							connection.BOTTOM,
						},
					},
				},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_BOTTOM_EDGE,
							connection.BOTTOM_LEFT_EDGE,
						},
					},
					{
						Sides: []connection.Side{
							connection.LEFT_TOP_EDGE,
							connection.BOTTOM_RIGHT_EDGE,
						},
					},
				},
			},
		},
		HasShield: false,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having city edges on top and right. Connected, shield but also road from left to bottom
*/
func TwoCityEdgesCornerConnectedRoadTurnShield() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP,
							connection.RIGHT,
						},
					},
				},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT,
							connection.BOTTOM,
						},
					},
				},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_BOTTOM_EDGE,
							connection.BOTTOM_LEFT_EDGE,
						},
					},
					{
						Sides: []connection.Side{
							connection.LEFT_TOP_EDGE,
							connection.BOTTOM_RIGHT_EDGE,
						},
					},
				},
			},
		},
		HasShield: true,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having city edges on top, right and left. Connected
*/
func ThreeCityEdgesConnected() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP,
							connection.RIGHT,
							connection.CENTER,
							connection.LEFT,
						},
					},
				},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_BOTTOM_EDGE,
							connection.BOTTOM_LEFT_EDGE,
						},
					},
				},
			},
		},
		HasShield: false,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having city edges on top, right and left. Connected and shield
*/
func ThreeCityEdgesConnectedShield() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP,
							connection.RIGHT,
							connection.CENTER,
							connection.LEFT,
						},
					},
				},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_BOTTOM_EDGE,
							connection.BOTTOM_LEFT_EDGE,
						},
					},
				},
			},
		},
		HasShield: true,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having city edges on top, right and left. Connected and road at the bottom
*/
func ThreeCityEdgesConnectedRoad() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP,
							connection.RIGHT,
							connection.CENTER,
							connection.LEFT,
						},
					},
				},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.CENTER,
							connection.BOTTOM,
						},
					},
				},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_BOTTOM_EDGE,
						},
					},
					{
						Sides: []connection.Side{
							connection.BOTTOM_LEFT_EDGE,
						},
					},
				},
			},
		},
		HasShield: false,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having city edges on top, right and left. Connected, shield and road at the bottom
*/
func ThreeCityEdgesConnectedRoadShield() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP,
							connection.RIGHT,
							connection.CENTER,
							connection.LEFT,
						},
					},
				},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.CENTER,
							connection.BOTTOM,
						},
					},
				},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.LEFT_BOTTOM_EDGE,
						},
					},
					{
						Sides: []connection.Side{
							connection.BOTTOM_LEFT_EDGE,
						},
					},
				},
			},
		},
		HasShield: true,
		Building:  buildings.NONE_BULDING,
	}
}

/*
returns tiles.Tile having 4 city edges. Connected
*/
func FourCityEdgesConnectedShield() tiles.Tile {
	return tiles.Tile{
		Features: []tiles.Feature{
			{
				FeatureType: tiles.CITIES,
				Connections: []connection.Connection{
					{
						Sides: []connection.Side{
							connection.TOP,
							connection.RIGHT,
							connection.CENTER,
							connection.LEFT,
							connection.BOTTOM,
						},
					},
				},
			},
			{
				FeatureType: tiles.ROADS,
				Connections: []connection.Connection{},
			},
			{
				FeatureType: tiles.FIELDS,
				Connections: []connection.Connection{},
			},
		},
		HasShield: true,
		Building:  buildings.NONE_BULDING,
	}
}
