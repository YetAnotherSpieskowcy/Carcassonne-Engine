package test

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/player"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

func GetTestTile() tiles.Tile {
	return tiletemplates.SingleCityEdgeNoRoads().Rotate(2)
}

func GetTestTilePlacement() elements.TilePlacement {
	return elements.TilePlacement{
		Tile: GetTestTile(),
		Pos:  elements.NewPosition(0, 1),
	}
}

func GetTestPlacedTile() elements.PlacedTile {
	return elements.PlacedTile{
		LegalMove: elements.LegalMove{
			TilePlacement: GetTestTilePlacement(),
			Meeple: elements.MeeplePlacement{
				Feature: feature.Feature{
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
		},
		Player: player.New(1),
	}
}

func GetTestPlacedTileWithMeeple(meeple elements.MeeplePlacement) elements.PlacedTile {
	return elements.PlacedTile{
		LegalMove: elements.LegalMove{
			TilePlacement: GetTestTilePlacement(),
			Meeple:        meeple,
		},
		Player: player.New(1),
	}
}

func GetTestScoreReport() elements.ScoreReport {
	return elements.ScoreReport{
		ReceivedPoints:  map[int]uint32{0: 5},
		ReturnedMeeples: map[int][]uint8{},
	}
}
