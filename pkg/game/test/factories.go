package test

import (
	. "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
)

func GetTestTile() Tile {
	return SingleCityEdgeNoRoads().Rotate(2)
}

func GetTestPlacedTile() PlacedTile {
	return PlacedTile{
		LegalMove: LegalMove{Tile: GetTestTile(), Pos: NewPosition(0, 1)},
		Meeple: Meeple{Side: Bottom},
	}
}

func GetTestPlacedTileWithMeeple(meeple Meeple) PlacedTile {
	return PlacedTile{
		LegalMove: LegalMove{Tile: GetTestTile(), Pos: NewPosition(0, 1)},
		Meeple: meeple,
	}
}

func GetTestScoreReport() ScoreReport {
	return ScoreReport{
		ReceivedPoints: map[int]uint32{0: 5},
		ReturnedMeeples: map[int]uint8{},
	}
}
