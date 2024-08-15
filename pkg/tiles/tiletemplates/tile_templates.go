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
			{
				FeatureType: feature.Field,
				Sides: side.TopLeftEdge |
					side.TopRightEdge |

					side.RightTopEdge |
					side.RightBottomEdge |

					side.LeftTopEdge |
					side.LeftBottomEdge |

					side.BottomLeftEdge |
					side.BottomRightEdge,
			},
			{
				FeatureType: feature.Monastery,
			},
		},
	}
}

/*
returns tiles.Tile having monastery and road going bottom
*/
func MonasteryWithSingleRoad() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.Road,
				Sides:       side.Bottom,
			},
			{
				FeatureType: feature.Field,
				Sides: side.TopLeftEdge |
					side.TopRightEdge |

					side.RightTopEdge |
					side.RightBottomEdge |

					side.LeftTopEdge |
					side.LeftBottomEdge |

					side.BottomLeftEdge |
					side.BottomRightEdge,
			},
			{
				FeatureType: feature.Monastery,
			},
		},
	}
}

/*
returns tiles.Tile having road from left to right
*/
func StraightRoads() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.Road,
				Sides: side.Left |
					side.Right,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftBottomEdge |
					side.BottomLeftEdge |
					side.BottomRightEdge |
					side.RightBottomEdge,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftTopEdge |
					side.TopLeftEdge |
					side.TopRightEdge |
					side.RightTopEdge,
			},
		},
	}
}

/*
returns tiles.Tile having road from left to bottom
*/
func RoadsTurn() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.Road,
				Sides: side.Left |
					side.Bottom,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftBottomEdge |
					side.BottomLeftEdge,
			},

			{
				FeatureType: feature.Field,
				Sides: side.LeftTopEdge |
					side.TopLeftEdge |
					side.TopRightEdge |
					side.RightTopEdge |
					side.RightBottomEdge |
					side.BottomRightEdge,
			},
		},
	}
}

/*
returns tiles.Tile having road from left,bottom,right to center
*/
func TCrossRoad() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.Road,
				Sides:       side.Left,
			},
			{
				FeatureType: feature.Road,
				Sides:       side.Right,
			},
			{
				FeatureType: feature.Road,
				Sides:       side.Bottom,
			},

			{
				FeatureType: feature.Field,
				Sides: side.LeftBottomEdge |
					side.BottomLeftEdge,
			},
			{
				FeatureType: feature.Field,
				Sides: side.RightBottomEdge |
					side.BottomRightEdge,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftTopEdge |
					side.TopLeftEdge |
					side.TopRightEdge |
					side.RightTopEdge,
			},
		},
	}
}

func XCrossRoad() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.Road,
				Sides:       side.Left,
			},
			{
				FeatureType: feature.Road,
				Sides:       side.Bottom,
			},
			{
				FeatureType: feature.Road,
				Sides:       side.Right,
			},
			{
				FeatureType: feature.Road,
				Sides:       side.Top,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftBottomEdge |
					side.BottomLeftEdge,
			},
			{
				FeatureType: feature.Field,
				Sides: side.RightBottomEdge |
					side.BottomRightEdge,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftTopEdge |
					side.TopLeftEdge,
			},
			{
				FeatureType: feature.Field,
				Sides: side.TopRightEdge |
					side.RightTopEdge,
			},
		},
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
				Sides:       side.Top,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftTopEdge |
					side.RightTopEdge |
					side.RightBottomEdge |
					side.BottomRightEdge |
					side.LeftBottomEdge |
					side.BottomLeftEdge,
			},
		},
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
				Sides:       side.Top,
			},
			{
				FeatureType: feature.Road,
				Sides: side.Right |
					side.Left,
			},
			{
				FeatureType: feature.Field,
				Sides: side.RightBottomEdge |
					side.BottomRightEdge |
					side.LeftBottomEdge |
					side.BottomLeftEdge,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftTopEdge |
					side.RightTopEdge,
			},
		},
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
				Sides:       side.Top,
			},
			{
				FeatureType: feature.Road,
				Sides: side.Left |
					side.Bottom,
			},
			{
				FeatureType: feature.Field,
				Sides: side.RightBottomEdge |
					side.BottomRightEdge |
					side.LeftTopEdge |
					side.RightTopEdge,
			},
			{
				FeatureType: feature.Field,
				Sides: side.BottomLeftEdge |
					side.LeftBottomEdge,
			},
		},
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
				Sides:       side.Top,
			},
			{
				FeatureType: feature.Road,
				Sides: side.Right |
					side.Bottom,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftTopEdge |
					side.RightTopEdge |
					side.BottomLeftEdge |
					side.LeftBottomEdge,
			},
			{
				FeatureType: feature.Field,
				Sides: side.RightBottomEdge |
					side.BottomRightEdge,
			},
		},
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
				Sides:       side.Top,
			},
			{
				FeatureType: feature.Road,
				Sides:       side.Right,
			},
			{
				FeatureType: feature.Road,
				Sides:       side.Left,
			},
			{
				FeatureType: feature.Road,
				Sides:       side.Bottom,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftTopEdge |
					side.RightTopEdge,
			},
			{
				FeatureType: feature.Field,
				Sides: side.RightBottomEdge |
					side.BottomRightEdge,
			},
			{
				FeatureType: feature.Field,
				Sides: side.BottomLeftEdge |
					side.LeftBottomEdge,
			},
		},
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
				Sides:       side.Top,
			},
			{
				FeatureType: feature.City,
				Sides:       side.Bottom,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftTopEdge |
					side.RightTopEdge |
					side.LeftBottomEdge |
					side.RightBottomEdge,
			},
		},
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
				Sides:       side.Top,
			},
			{
				FeatureType: feature.City,
				Sides:       side.Right,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftTopEdge |
					side.LeftBottomEdge |
					side.BottomLeftEdge |
					side.BottomRightEdge,
			},
		},
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
				Sides: side.Top |
					side.Bottom,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftTopEdge |
					side.LeftBottomEdge,
			},
			{
				FeatureType: feature.Field,
				Sides: side.BottomLeftEdge |
					side.BottomRightEdge,
			},
		},
	}
}

/*
returns tiles.Tile having city edges on top and down. Connected and shield
*/
func TwoCityEdgesUpAndDownConnectedShield() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType:  feature.City,
				ModifierType: modifier.Shield,
				Sides: side.Top |
					side.Bottom,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftTopEdge |
					side.LeftBottomEdge,
			},
			{
				FeatureType: feature.Field,
				Sides: side.BottomLeftEdge |
					side.BottomRightEdge,
			},
		},
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
				Sides: side.Top |
					side.Right,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftTopEdge |
					side.LeftBottomEdge |
					side.BottomLeftEdge |
					side.BottomRightEdge,
			},
		},
	}
}

/*
returns tiles.Tile having city edges on top and right. Connected and shield
*/
func TwoCityEdgesCornerConnectedShield() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType:  feature.City,
				ModifierType: modifier.Shield,
				Sides: side.Top |
					side.Right,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftTopEdge |
					side.LeftBottomEdge |
					side.BottomLeftEdge |
					side.BottomRightEdge,
			},
		},
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
				Sides: side.Top |
					side.Right,
			},
			{
				FeatureType: feature.Road,
				Sides: side.Left |
					side.Bottom,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftBottomEdge |
					side.BottomLeftEdge,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftTopEdge |
					side.BottomRightEdge,
			},
		},
	}
}

/*
returns tiles.Tile having city edges on top and right. Connected, shield but also road from left to bottom
*/
func TwoCityEdgesCornerConnectedRoadTurnShield() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType:  feature.City,
				ModifierType: modifier.Shield,
				Sides: side.Top |
					side.Right,
			},
			{
				FeatureType: feature.Road,
				Sides: side.Left |
					side.Bottom,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftBottomEdge |
					side.BottomLeftEdge,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftTopEdge |
					side.BottomRightEdge,
			},
		},
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
				Sides: side.Top |
					side.Right |
					side.Left,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftBottomEdge |
					side.BottomLeftEdge,
			},
		},
	}
}

/*
returns tiles.Tile having city edges on top, right and left. Connected and shield
*/
func ThreeCityEdgesConnectedShield() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType:  feature.City,
				ModifierType: modifier.Shield,
				Sides: side.Top |
					side.Right |
					side.Left,
			},
			{
				FeatureType: feature.Field,
				Sides: side.LeftBottomEdge |
					side.BottomLeftEdge,
			},
		},
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
				Sides: side.Top |
					side.Right |
					side.Left,
			},
			{
				FeatureType: feature.Road,
				Sides:       side.Bottom,
			},
			{
				FeatureType: feature.Field,
				Sides:       side.LeftBottomEdge,
			},
			{
				FeatureType: feature.Field,
				Sides:       side.BottomLeftEdge,
			},
		},
	}
}

/*
returns tiles.Tile having city edges on top, right and left. Connected, shield and road at the bottom
*/
func ThreeCityEdgesConnectedRoadShield() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType:  feature.City,
				ModifierType: modifier.Shield,
				Sides: side.Top |
					side.Right |
					side.Left,
			},
			{
				FeatureType: feature.Road,
				Sides:       side.Bottom,
			},
			{
				FeatureType: feature.Field,
				Sides:       side.LeftBottomEdge,
			},
			{
				FeatureType: feature.Field,
				Sides:       side.BottomLeftEdge,
			},
		},
	}
}

/*
returns tiles.Tile having 4 city edges. Connected
*/
func FourCityEdgesConnectedShield() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType:  feature.City,
				ModifierType: modifier.Shield,
				Sides: side.Top |
					side.Right |
					side.Left |
					side.Bottom,
			},
		},
	}
}

/*
returns tiles.Tile consisting of only a single field. (Unused in game, useful in testing)
*/
func TestOnlyField() tiles.Tile {
	return tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.Field,
				Sides: side.Top |
					side.Right |
					side.Left |
					side.Bottom,
			},
		},
	}
}
