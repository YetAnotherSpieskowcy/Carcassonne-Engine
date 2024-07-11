package test

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

func GetTestTile() tiles.Tile {
	return tiletemplates.SingleCityEdgeNoRoads().Rotate(2)
}

func GetTestPlacedTile() elements.PlacedTile {
	tile := elements.ToPlacedTile(GetTestTile())
	tile.Position = position.NewPosition(0, 1)
	return tile
}
func GetTestScoreReport() elements.ScoreReport {
	return elements.ScoreReport{
		ReceivedPoints:  map[elements.ID]uint32{0: 5},
		ReturnedMeeples: map[elements.ID][]uint8{},
	}
}

func GetTestCustomPlacedTile(tileTemplate tiles.Tile) elements.PlacedTile {
	var placedFeatures []elements.PlacedFeature

	// convert features to placedFeature
	for _, feature := range tileTemplate.Features {
		placedFeatures = append(placedFeatures, elements.PlacedFeature{
			Feature: feature,
			Meeple: elements.Meeple{
				MeepleType: elements.NoneMeeple,
				PlayerID:   elements.NonePlayer},
		})
	}

	return elements.PlacedTile{
		TileWithMeeple: elements.TileWithMeeple{
			Features:  placedFeatures,
			HasShield: false,
		},
		Position: position.NewPosition(0, 0),
	}
}
