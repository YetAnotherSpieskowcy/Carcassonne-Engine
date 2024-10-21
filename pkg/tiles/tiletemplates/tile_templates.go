package tiletemplates

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature/modifier"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

func MonasteryWithoutRoads() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.Field,
				side.TopLeftEdge|
					side.TopRightEdge|

					side.RightTopEdge|
					side.RightBottomEdge|

					side.LeftTopEdge|
					side.LeftBottomEdge|

					side.BottomLeftEdge|
					side.BottomRightEdge,
			),
			feature.New(
				feature.Monastery,
				side.NoSide,
			),
		},
	}
}

/*
returns tiles.Tile having monastery and road going bottom
*/
func MonasteryWithSingleRoad() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.Road,
				side.Bottom,
			),
			feature.New(
				feature.Field,
				side.TopLeftEdge|
					side.TopRightEdge|

					side.RightTopEdge|
					side.RightBottomEdge|

					side.LeftTopEdge|
					side.LeftBottomEdge|

					side.BottomLeftEdge|
					side.BottomRightEdge,
			),
			feature.New(
				feature.Monastery,
				side.NoSide,
			),
		},
	}
}

/*
returns tiles.Tile having road from left to right
*/
func StraightRoads() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.Road,
				side.Left|
					side.Right,
			),
			feature.New(
				feature.Field,
				side.LeftBottomEdge|
					side.BottomLeftEdge|
					side.BottomRightEdge|
					side.RightBottomEdge,
			),
			feature.New(
				feature.Field,
				side.LeftTopEdge|
					side.TopLeftEdge|
					side.TopRightEdge|
					side.RightTopEdge,
			),
		},
	}
}

/*
returns tiles.Tile having road from left to bottom
*/
func RoadsTurn() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.Road,
				side.Left|
					side.Bottom,
			),
			feature.New(
				feature.Field,
				side.LeftBottomEdge|
					side.BottomLeftEdge,
			),

			feature.New(
				feature.Field,
				side.LeftTopEdge|
					side.TopLeftEdge|
					side.TopRightEdge|
					side.RightTopEdge|
					side.RightBottomEdge|
					side.BottomRightEdge,
			),
		},
	}
}

/*
returns tiles.Tile having road from left,bottom,right to center
*/
func TCrossRoad() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.Road,
				side.Left,
			),
			feature.New(
				feature.Road,
				side.Right,
			),
			feature.New(
				feature.Road,
				side.Bottom,
			),

			feature.New(
				feature.Field,
				side.LeftBottomEdge|
					side.BottomLeftEdge,
			),
			feature.New(
				feature.Field,
				side.RightBottomEdge|
					side.BottomRightEdge,
			),
			feature.New(
				feature.Field,
				side.LeftTopEdge|
					side.TopLeftEdge|
					side.TopRightEdge|
					side.RightTopEdge,
			),
		},
	}
}

func XCrossRoad() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.Road,
				side.Left,
			),
			feature.New(
				feature.Road,
				side.Bottom,
			),
			feature.New(
				feature.Road,
				side.Right,
			),
			feature.New(
				feature.Road,
				side.Top,
			),
			feature.New(
				feature.Field,
				side.LeftBottomEdge|
					side.BottomLeftEdge,
			),
			feature.New(
				feature.Field,
				side.RightBottomEdge|
					side.BottomRightEdge,
			),
			feature.New(
				feature.Field,
				side.LeftTopEdge|
					side.TopLeftEdge,
			),
			feature.New(
				feature.Field,
				side.TopRightEdge|
					side.RightTopEdge,
			),
		},
	}
}

/*
returns tiles.Tile having single city edge on top
*/
func SingleCityEdgeNoRoads() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,
				side.Top,
			),
			feature.New(
				feature.Field,
				side.LeftTopEdge|
					side.RightTopEdge|
					side.RightBottomEdge|
					side.BottomRightEdge|
					side.LeftBottomEdge|
					side.BottomLeftEdge,
			),
		},
	}
}

/*
returns tiles.Tile having single city edge on top and road from left to right
*/
func SingleCityEdgeStraightRoads() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,
				side.Top,
			),
			feature.New(
				feature.Road,
				side.Right|
					side.Left,
			),
			feature.New(
				feature.Field,
				side.RightBottomEdge|
					side.BottomRightEdge|
					side.LeftBottomEdge|
					side.BottomLeftEdge,
			),
			feature.New(
				feature.Field,
				side.LeftTopEdge|
					side.RightTopEdge,
			),
		},
	}
}

/*
returns tiles.Tile having single city edge on top and road from left to bottom
*/
func SingleCityEdgeLeftRoadTurn() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,
				side.Top,
			),
			feature.New(
				feature.Road,
				side.Left|
					side.Bottom,
			),
			feature.New(
				feature.Field,
				side.RightBottomEdge|
					side.BottomRightEdge|
					side.LeftTopEdge|
					side.RightTopEdge,
			),
			feature.New(
				feature.Field,
				side.BottomLeftEdge|
					side.LeftBottomEdge,
			),
		},
	}
}

/*
returns tiles.Tile having single city edge on top and road from right to bottom
*/
func SingleCityEdgeRightRoadTurn() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,
				side.Top,
			),
			feature.New(
				feature.Road,
				side.Right|
					side.Bottom,
			),
			feature.New(
				feature.Field,
				side.LeftTopEdge|
					side.RightTopEdge|
					side.BottomLeftEdge|
					side.LeftBottomEdge,
			),
			feature.New(
				feature.Field,
				side.RightBottomEdge|
					side.BottomRightEdge,
			),
		},
	}
}

/*
returns tiles.Tile having single city edge on top and roads on other sides
*/
func SingleCityEdgeCrossRoad() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,
				side.Top,
			),
			feature.New(
				feature.Road,
				side.Right,
			),
			feature.New(
				feature.Road,
				side.Left,
			),
			feature.New(
				feature.Road,
				side.Bottom,
			),
			feature.New(
				feature.Field,
				side.LeftTopEdge|
					side.RightTopEdge,
			),
			feature.New(
				feature.Field,
				side.RightBottomEdge|
					side.BottomRightEdge,
			),
			feature.New(
				feature.Field,
				side.BottomLeftEdge|
					side.LeftBottomEdge,
			),
		},
	}
}

/*
returns tiles.Tile having city edges on top and bottom. Not connected
*/
func TwoCityEdgesUpAndDownNotConnected() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,
				side.Top,
			),
			feature.New(
				feature.City,
				side.Bottom,
			),
			feature.New(
				feature.Field,
				side.LeftTopEdge|
					side.RightTopEdge|
					side.LeftBottomEdge|
					side.RightBottomEdge,
			),
		},
	}
}

/*
returns tiles.Tile having city edges on top and right. Not connected
*/
func TwoCityEdgesCornerNotConnected() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,
				side.Top,
			),
			feature.New(
				feature.City,
				side.Right,
			),
			feature.New(
				feature.Field,
				side.LeftTopEdge|
					side.LeftBottomEdge|
					side.BottomLeftEdge|
					side.BottomRightEdge,
			),
		},
	}
}

/*
returns tiles.Tile having city edges on top and down. Connected
*/
func TwoCityEdgesUpAndDownConnected() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,
				side.Top|
					side.Bottom,
			),
			feature.New(
				feature.Field,
				side.LeftTopEdge|
					side.LeftBottomEdge,
			),
			feature.New(
				feature.Field,
				side.RightTopEdge|
					side.RightBottomEdge,
			),
		},
	}
}

/*
returns tiles.Tile having city edges on top and down. Connected and shield
*/
func TwoCityEdgesUpAndDownConnectedShield() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,

				side.Top|
					side.Bottom,
				modifier.Shield,
			),
			feature.New(
				feature.Field,
				side.LeftTopEdge|
					side.LeftBottomEdge,
			),
			feature.New(
				feature.Field,
				side.RightTopEdge|
					side.RightBottomEdge,
			),
		},
	}
}

/*
returns tiles.Tile having city edges on top and right. Connected
*/
func TwoCityEdgesCornerConnected() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,
				side.Top|
					side.Right,
			),
			feature.New(
				feature.Field,
				side.LeftTopEdge|
					side.LeftBottomEdge|
					side.BottomLeftEdge|
					side.BottomRightEdge,
			),
		},
	}
}

/*
returns tiles.Tile having city edges on top and right. Connected and shield
*/
func TwoCityEdgesCornerConnectedShield() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,
				side.Top|
					side.Right,
				modifier.Shield,
			),
			feature.New(
				feature.Field,
				side.LeftTopEdge|
					side.LeftBottomEdge|
					side.BottomLeftEdge|
					side.BottomRightEdge,
			),
		},
	}
}

/*
returns tiles.Tile having city edges on top and right. Connected but also road from left to bottom
*/
func TwoCityEdgesCornerConnectedRoadTurn() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,
				side.Top|
					side.Right,
			),
			feature.New(
				feature.Road,
				side.Left|
					side.Bottom,
			),
			feature.New(
				feature.Field,
				side.LeftBottomEdge|
					side.BottomLeftEdge,
			),
			feature.New(
				feature.Field,
				side.LeftTopEdge|
					side.BottomRightEdge,
			),
		},
	}
}

/*
returns tiles.Tile having city edges on top and right. Connected, shield but also road from left to bottom
*/
func TwoCityEdgesCornerConnectedRoadTurnShield() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,
				side.Top|
					side.Right,
				modifier.Shield,
			),
			feature.New(
				feature.Road,
				side.Left|
					side.Bottom,
			),
			feature.New(
				feature.Field,
				side.LeftBottomEdge|
					side.BottomLeftEdge,
			),
			feature.New(
				feature.Field,
				side.LeftTopEdge|
					side.BottomRightEdge,
			),
		},
	}
}

/*
returns tiles.Tile having city edges on top, right and left. Connected
*/
func ThreeCityEdgesConnected() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,
				side.Top|
					side.Right|
					side.Left,
			),
			feature.New(
				feature.Field,
				side.BottomLeftEdge|
					side.BottomRightEdge,
			),
		},
	}
}

/*
returns tiles.Tile having city edges on top, right and left. Connected and shield
*/
func ThreeCityEdgesConnectedShield() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,

				side.Top|
					side.Right|
					side.Left,
				modifier.Shield,
			),
			feature.New(
				feature.Field,
				side.BottomLeftEdge|
					side.BottomRightEdge,
			),
		},
	}
}

/*
returns tiles.Tile having city edges on top, right and left. Connected and road at the bottom
*/
func ThreeCityEdgesConnectedRoad() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,
				side.Top|
					side.Right|
					side.Left,
			),
			feature.New(
				feature.Road,
				side.Bottom,
			),
			feature.New(
				feature.Field,
				side.BottomLeftEdge,
			),
			feature.New(
				feature.Field,
				side.BottomRightEdge,
			),
		},
	}
}

/*
returns tiles.Tile having city edges on top, right and left. Connected, shield and road at the bottom
*/
func ThreeCityEdgesConnectedRoadShield() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,
				side.Top|
					side.Right|
					side.Left,
				modifier.Shield,
			),
			feature.New(
				feature.Road,
				side.Bottom,
			),
			feature.New(
				feature.Field,
				side.BottomLeftEdge,
			),
			feature.New(
				feature.Field,
				side.BottomRightEdge,
			),
		},
	}
}

/*
returns tiles.Tile having 4 city edges. Connected
*/
func FourCityEdgesConnectedShield() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,
				side.Top|
					side.Right|
					side.Left|
					side.Bottom,
				modifier.Shield,
			),
		},
	}
}

/*
returns tiles.Tile consisting of only a single field. (Unused in game, useful in testing)
*/
func TestOnlyField() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.Field,
				side.Top|
					side.Right|
					side.Left|
					side.Bottom,
			),
		},
	}
}

/*
returns tiles.Tile consisting of only a single road. (Unused in game, useful in testing)
*/
func TestOnlyStraightRoads() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.Road,
				side.Right|
					side.Left,
			),
		},
	}
}

/*
returns tiles.Tile consisting of only a single monastery. (Unused in game, useful in testing)
*/
func TestOnlyMonastery() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.Monastery,
				side.NoSide,
			),
		},
	}
}
