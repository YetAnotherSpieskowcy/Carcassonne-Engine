package test

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

func GetTestTile() tiles.Tile {
	return tiletemplates.SingleCityEdgeNoRoads().Rotate(2)
}

func GetTestPlacedTile() elements.PlacedTile {
	tile := elements.ToPlacedTile(GetTestTile())
	tile.Position = elements.NewPosition(0, 1)
	return tile
}
func GetTestScoreReport() elements.ScoreReport {
	return elements.ScoreReport{
		ReceivedPoints:  map[uint8]uint32{0: 5},
		ReturnedMeeples: map[uint8][]uint8{},
	}
}

func GetTestTileOnlyField() elements.PlacedTile {
	return elements.PlacedTile{
		LegalMove: elements.LegalMove{
			TilePlacement: elements.TilePlacement{
				Tile: tiletemplates.OnlyField(),
				Pos:  elements.NewPosition(0, 0),
			},
			Meeple: elements.MeeplePlacement{},
		},
		Player: player.New(1),
	}
}

func GetTestTileMonasteryWithoutRoads() elements.PlacedTile {
	return elements.PlacedTile{
		LegalMove: elements.LegalMove{
			TilePlacement: elements.TilePlacement{
				Tile: tiletemplates.MonasteryWithoutRoads(),
				Pos:  elements.NewPosition(0, 0),
			},
			Meeple: elements.MeeplePlacement{},
		},
		Player: player.New(1),
	}
}
