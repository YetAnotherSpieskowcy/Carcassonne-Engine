package test

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/player"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
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
			Meeple:        elements.MeeplePlacement{Side: side.Bottom},
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
		ReceivedPoints:  map[uint8]uint32{0: 5},
		ReturnedMeeples: map[uint8][]uint8{},
	}
}

func GetTestTestPlacedTile(tileTemplate tiles.Tile, playerID uint8) elements.PlacedTile {
	return elements.PlacedTile{
		LegalMove: elements.LegalMove{
			TilePlacement: elements.TilePlacement{
				Tile: tileTemplate,
				Pos:  elements.NewPosition(0, 0),
			},
			Meeple: elements.MeeplePlacement{},
		},
		Player: player.New(playerID),
	}
}
