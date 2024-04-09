package tiletemplates

import (
	tiles "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/buildings"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

func MonasteryWithoutRoads() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.City,
				Sides:       []side.Side{},
			},
			{
				FeatureType: feature.Road,
				Sides:       []side.Side{},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.TopLeftEdge,
					side.TopRightEdge,

					side.RightTopEdge,
					side.RightBottomEdge,

					side.LeftTopEdge,
					side.LeftBottomEdge,

					side.BottomLeftEdge,
					side.BottomRightEdge,
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
				FeatureType: feature.City,
				Sides:       []side.Side{},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{
					side.Center,
					side.Bottom,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.TopLeftEdge,
					side.TopRightEdge,

					side.RightTopEdge,
					side.RightBottomEdge,

					side.LeftTopEdge,
					side.LeftBottomEdge,

					side.BottomLeftEdge,
					side.BottomRightEdge,
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
				FeatureType: feature.City,
				Sides:       []side.Side{},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{
					side.Left,
					side.Right,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftBottomEdge,
					side.BottomLeftEdge,
					side.BottomRightEdge,
					side.RightBottomEdge,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftTopEdge,
					side.TopLeftEdge,
					side.TopRightEdge,
					side.RightTopEdge,
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
				FeatureType: feature.City,
				Sides:       []side.Side{},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{
					side.Left,
					side.Bottom,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftBottomEdge,
					side.BottomLeftEdge,
				},
			},

			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftTopEdge,
					side.TopLeftEdge,
					side.TopRightEdge,
					side.RightTopEdge,
					side.RightBottomEdge,
					side.BottomRightEdge,
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
				FeatureType: feature.City,
				Sides:       []side.Side{},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{
					side.Left,
					side.Center,
				},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{

					side.Right,
					side.Center,
				},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{
					side.Bottom,
					side.Center,
				},
			},

			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftBottomEdge,
					side.BottomLeftEdge,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.RightBottomEdge,
					side.BottomRightEdge,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftTopEdge,
					side.TopLeftEdge,
					side.TopRightEdge,
					side.RightTopEdge,
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
				FeatureType: feature.City,
				Sides:       []side.Side{},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{
					side.Left,
					side.Center,
				},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{
					side.Bottom,
					side.Center,
				},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{
					side.Right,
					side.Center,
				},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{

					side.Top,
					side.Center,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftBottomEdge,
					side.BottomLeftEdge,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.RightBottomEdge,
					side.BottomRightEdge,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftTopEdge,
					side.TopLeftEdge,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.TopRightEdge,
					side.RightTopEdge,
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
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Top,
				},
			},
			{
				FeatureType: feature.Road,
				Sides:       []side.Side{},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftTopEdge,
					side.RightTopEdge,
					side.RightBottomEdge,
					side.BottomRightEdge,
					side.LeftBottomEdge,
					side.BottomLeftEdge,
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
				FeatureType: feature.City,
				Sides: []side.Side{

					side.Top,
				},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{
					side.Right,
					side.Left,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.RightBottomEdge,
					side.BottomRightEdge,
					side.LeftBottomEdge,
					side.BottomLeftEdge,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftTopEdge,
					side.RightTopEdge,
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
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Top,
				},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{
					side.Left,
					side.Bottom,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.RightBottomEdge,
					side.BottomRightEdge,
					side.LeftTopEdge,
					side.RightTopEdge,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{

					side.BottomLeftEdge,
					side.LeftBottomEdge,
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
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Top,
				},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{
					side.Right,
					side.Bottom,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftTopEdge,
					side.RightTopEdge,
					side.BottomLeftEdge,
					side.LeftBottomEdge,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.RightBottomEdge,
					side.BottomRightEdge,
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
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Top,
				},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{
					side.Right,
					side.Center,
				},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{
					side.Left,
					side.Center,
				},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{
					side.Bottom,
					side.Center,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftTopEdge,
					side.RightTopEdge,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.RightBottomEdge,
					side.BottomRightEdge,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.BottomLeftEdge,
					side.LeftBottomEdge,
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
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Top,
				},
			},
			{
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Bottom,
				},
			},
			{
				FeatureType: feature.Road,
				Sides:       []side.Side{},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftTopEdge,
					side.RightTopEdge,
					side.LeftBottomEdge,
					side.RightBottomEdge,
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
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Top,
				},
			},
			{
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Right,
				},
			},
			{
				FeatureType: feature.Road,
				Sides:       []side.Side{},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftTopEdge,
					side.LeftBottomEdge,
					side.BottomLeftEdge,
					side.BottomRightEdge,
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
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Top,
					side.Bottom,
					side.Center,
				},
			},
			{
				FeatureType: feature.Road,
				Sides:       []side.Side{},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftTopEdge,
					side.LeftBottomEdge,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.BottomLeftEdge,
					side.BottomRightEdge,
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
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Top,
					side.Bottom,
					side.Center,
				},
			},
			{
				FeatureType: feature.Road,
				Sides:       []side.Side{},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftTopEdge,
					side.LeftBottomEdge,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.BottomLeftEdge,
					side.BottomRightEdge,
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
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Top,
					side.Right,
				},
			},
			{
				FeatureType: feature.Road,
				Sides:       []side.Side{},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftTopEdge,
					side.LeftBottomEdge,
					side.BottomLeftEdge,
					side.BottomRightEdge,
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
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Top,
					side.Right,
				},
			},
			{
				FeatureType: feature.Road,
				Sides:       []side.Side{},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftTopEdge,
					side.LeftBottomEdge,
					side.BottomLeftEdge,
					side.BottomRightEdge,
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
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Top,
					side.Right,
				},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{
					side.Left,
					side.Bottom,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftBottomEdge,
					side.BottomLeftEdge,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftTopEdge,
					side.BottomRightEdge,
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
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Top,
					side.Right,
				},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{
					side.Left,
					side.Bottom,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftBottomEdge,
					side.BottomLeftEdge,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftTopEdge,
					side.BottomRightEdge,
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
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Top,
					side.Right,
					side.Center,
					side.Left,
				},
			},
			{
				FeatureType: feature.Road,
				Sides:       []side.Side{},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftBottomEdge,
					side.BottomLeftEdge,
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
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Top,
					side.Right,
					side.Center,
					side.Left,
				},
			},
			{
				FeatureType: feature.Road,
				Sides:       []side.Side{},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftBottomEdge,
					side.BottomLeftEdge,
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
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Top,
					side.Right,
					side.Center,
					side.Left,
				},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{
					side.Center,
					side.Bottom,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftBottomEdge,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.BottomLeftEdge,
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
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Top,
					side.Right,
					side.Center,
					side.Left,
				},
			},
			{
				FeatureType: feature.Road,
				Sides: []side.Side{
					side.Center,
					side.Bottom,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.LeftBottomEdge,
				},
			},
			{
				FeatureType: feature.Field,
				Sides: []side.Side{
					side.BottomLeftEdge,
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
				FeatureType: feature.City,
				Sides: []side.Side{
					side.Top,
					side.Right,
					side.Center,
					side.Left,
					side.Bottom,
				},
			},
			{
				FeatureType: feature.Road,
				Sides:       []side.Side{},
			},
			{
				FeatureType: feature.Field,
				Sides:       []side.Side{},
			},
		},
		HasShield: true,
		Building:  buildings.None,
	}
}
