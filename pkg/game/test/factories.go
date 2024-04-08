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
