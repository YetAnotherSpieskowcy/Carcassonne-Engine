package test

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
)

func GetTestTile() elements.Tile {
	return elements.SingleCityEdgeNoRoads().Rotate(2)
}

func GetTestPlacedTile() elements.PlacedTile {
	return elements.PlacedTile{
		LegalMove: elements.LegalMove{
			Tile: GetTestTile(),
			Pos: elements.NewPosition(0, 1),
		},
		Meeple: elements.Meeple{Side: elements.Bottom},
	}
}

func GetTestPlacedTileWithMeeple(meeple elements.Meeple) elements.PlacedTile {
	return elements.PlacedTile{
		LegalMove: elements.LegalMove{
			Tile: GetTestTile(),
			Pos: elements.NewPosition(0, 1),
		},
		Meeple: meeple,
	}
}

func GetTestScoreReport() elements.ScoreReport {
	return elements.ScoreReport{
		ReceivedPoints: map[int]uint32{0: 5},
		ReturnedMeeples: map[int]uint8{},
	}
}
