package tiletemplates

import (
	tiles "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/buildings"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/connection"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
)

func MonasteryWithoutRoads() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.TopLeftEdge,
					connection.TopRightEdge,

					connection.RightTopEdge,
					connection.RightBottomEdge,

					connection.LeftTopEdge,
					connection.LeftBottomEdge,

					connection.BottomLeftEdge,
					connection.BottomRightEdge,
				},
			},
		},
		HasShield: false,
		Building:  buildings.Monastery,
	}
}

/*
returns tiles.Tile having monastery and road going bottom
*/
func MonasteryWithSingleRoad() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{
					connection.Center,
					connection.Bottom,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.TopLeftEdge,
					connection.TopRightEdge,

					connection.RightTopEdge,
					connection.RightBottomEdge,

					connection.LeftTopEdge,
					connection.LeftBottomEdge,

					connection.BottomLeftEdge,
					connection.BottomRightEdge,
				},
			},
		},
		HasShield: false,
		Building:  buildings.Monastery,
	}
}

/*
returns tiles.Tile having road from left to right
*/
func StraightRoads() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{
					connection.Left,
					connection.Right,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftBottomEdge,
					connection.BottomLeftEdge,
					connection.BottomRightEdge,
					connection.RightBottomEdge,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftTopEdge,
					connection.TopLeftEdge,
					connection.TopRightEdge,
					connection.RightTopEdge,
				},
			},
		},
		HasShield: false,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having road from left to bottom
*/
func RoadsTurn() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{
					connection.Left,
					connection.Bottom,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftBottomEdge,
					connection.BottomLeftEdge,
				},
			},

			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftTopEdge,
					connection.TopLeftEdge,
					connection.TopRightEdge,
					connection.RightTopEdge,
					connection.RightBottomEdge,
					connection.BottomRightEdge,
				},
			},
		},
		HasShield: false,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having road from left,bottom,right to center
*/
func TCrossRoad() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{
					connection.Left,
					connection.Center,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{

					connection.Right,
					connection.Center,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{
					connection.Bottom,
					connection.Center,
				},
			},

			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftBottomEdge,
					connection.BottomLeftEdge,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.RightBottomEdge,
					connection.BottomRightEdge,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftTopEdge,
					connection.TopLeftEdge,
					connection.TopRightEdge,
					connection.RightTopEdge,
				},
			},
		},
		HasShield: false,
		Building:  buildings.None,
	}
}

func XCrossRoad() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{
					connection.Left,
					connection.Center,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{
					connection.Bottom,
					connection.Center,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{
					connection.Right,
					connection.Center,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{

					connection.Top,
					connection.Center,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftBottomEdge,
					connection.BottomLeftEdge,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.RightBottomEdge,
					connection.BottomRightEdge,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftTopEdge,
					connection.TopLeftEdge,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.TopRightEdge,
					connection.RightTopEdge,
				},
			},
		},
		HasShield: false,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having single city edge on top
*/
func SingleCityEdgeNoRoads() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Top,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftTopEdge,
					connection.RightTopEdge,
					connection.RightBottomEdge,
					connection.BottomRightEdge,
					connection.LeftBottomEdge,
					connection.BottomLeftEdge,
				},
			},
		},
		HasShield: false,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having single city edge on top and road from left to right
*/
func SingleCityEdgeStraightRoads() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{

					connection.Top,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{
					connection.Right,
					connection.Left,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.RightBottomEdge,
					connection.BottomRightEdge,
					connection.LeftBottomEdge,
					connection.BottomLeftEdge,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftTopEdge,
					connection.RightTopEdge,
				},
			},
		},
		HasShield: false,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having single city edge on top and road from left to bottom
*/
func SingleCityEdgeLeftRoadTurn() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Top,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{
					connection.Left,
					connection.Bottom,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.RightBottomEdge,
					connection.BottomRightEdge,
					connection.LeftTopEdge,
					connection.RightTopEdge,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{

					connection.BottomLeftEdge,
					connection.LeftBottomEdge,
				},
			},
		},
		HasShield: false,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having single city edge on top and road from right to bottom
*/
func SingleCityEdgeRightRoadTurn() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Top,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{
					connection.Right,
					connection.Bottom,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftTopEdge,
					connection.RightTopEdge,
					connection.BottomLeftEdge,
					connection.LeftBottomEdge,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.RightBottomEdge,
					connection.BottomRightEdge,
				},
			},
		},
		HasShield: false,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having single city edge on top and roads on other sides
*/
func SingleCityEdgeCrossRoad() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Top,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{
					connection.Right,
					connection.Center,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{
					connection.Left,
					connection.Center,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{
					connection.Bottom,
					connection.Center,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftTopEdge,
					connection.RightTopEdge,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.RightBottomEdge,
					connection.BottomRightEdge,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.BottomLeftEdge,
					connection.LeftBottomEdge,
				},
			},
		},
		HasShield: false,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having city edges on top and bottom. Not connected
*/
func TwoCityEdgesUpAndDownNotConnected() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Top,
				},
			},
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Bottom,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftTopEdge,
					connection.RightTopEdge,
					connection.LeftBottomEdge,
					connection.RightBottomEdge,
				},
			},
		},
		HasShield: false,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having city edges on top and right. Not connected
*/
func TwoCityEdgesCornerNotConnected() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Top,
				},
			},
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Right,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftTopEdge,
					connection.LeftBottomEdge,
					connection.BottomLeftEdge,
					connection.BottomRightEdge,
				},
			},
		},
		HasShield: false,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having city edges on top and down. Connected
*/
func TwoCityEdgesUpAndDownConnected() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Top,
					connection.Bottom,
					connection.Center,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftTopEdge,
					connection.LeftBottomEdge,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.BottomLeftEdge,
					connection.BottomRightEdge,
				},
			},
		},
		HasShield: false,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having city edges on top and down. Connected and shield
*/
func TwoCityEdgesUpAndDownConnectedShield() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Top,
					connection.Bottom,
					connection.Center,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftTopEdge,
					connection.LeftBottomEdge,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.BottomLeftEdge,
					connection.BottomRightEdge,
				},
			},
		},
		HasShield: true,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having city edges on top and right. Connected
*/
func TwoCityEdgesCornerConnected() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Top,
					connection.Right,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftTopEdge,
					connection.LeftBottomEdge,
					connection.BottomLeftEdge,
					connection.BottomRightEdge,
				},
			},
		},
		HasShield: false,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having city edges on top and right. Connected and shield
*/

func TwoCityEdgesCornerConnectedShield() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Top,
					connection.Right,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftTopEdge,
					connection.LeftBottomEdge,
					connection.BottomLeftEdge,
					connection.BottomRightEdge,
				},
			},
		},
		HasShield: true,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having city edges on top and right. Connected but also road from left to bottom
*/
func TwoCityEdgesCornerConnectedRoadTurn() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Top,
					connection.Right,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{
					connection.Left,
					connection.Bottom,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftBottomEdge,
					connection.BottomLeftEdge,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftTopEdge,
					connection.BottomRightEdge,
				},
			},
		},
		HasShield: false,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having city edges on top and right. Connected, shield but also road from left to bottom
*/
func TwoCityEdgesCornerConnectedRoadTurnShield() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Top,
					connection.Right,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{
					connection.Left,
					connection.Bottom,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftBottomEdge,
					connection.BottomLeftEdge,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftTopEdge,
					connection.BottomRightEdge,
				},
			},
		},
		HasShield: true,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having city edges on top, right and left. Connected
*/
func ThreeCityEdgesConnected() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Top,
					connection.Right,
					connection.Center,
					connection.Left,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftBottomEdge,
					connection.BottomLeftEdge,
				},
			},
		},
		HasShield: false,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having city edges on top, right and left. Connected and shield
*/
func ThreeCityEdgesConnectedShield() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Top,
					connection.Right,
					connection.Center,
					connection.Left,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftBottomEdge,
					connection.BottomLeftEdge,
				},
			},
		},
		HasShield: true,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having city edges on top, right and left. Connected and road at the bottom
*/
func ThreeCityEdgesConnectedRoad() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Top,
					connection.Right,
					connection.Center,
					connection.Left,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{
					connection.Center,
					connection.Bottom,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftBottomEdge,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.BottomLeftEdge,
				},
			},
		},
		HasShield: false,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having city edges on top, right and left. Connected, shield and road at the bottom
*/
func ThreeCityEdgesConnectedRoadShield() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Top,
					connection.Right,
					connection.Center,
					connection.Left,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{
					connection.Center,
					connection.Bottom,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.LeftBottomEdge,
				},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{
					connection.BottomLeftEdge,
				},
			},
		},
		HasShield: true,
		Building:  buildings.None,
	}
}

/*
returns tiles.Tile having 4 city edges. Connected
*/
func FourCityEdgesConnectedShield() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.CITY,
				Connections: []connection.Side{
					connection.Top,
					connection.Right,
					connection.Center,
					connection.Left,
					connection.Bottom,
				},
			},
			{
				FeatureType: feature.ROAD,
				Connections: []connection.Side{},
			},
			{
				FeatureType: feature.FIELD,
				Connections: []connection.Side{},
			},
		},
		HasShield: true,
		Building:  buildings.None,
	}
}
