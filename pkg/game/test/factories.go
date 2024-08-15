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
	tile.Position = position.New(0, 1)
	return tile
}
func GetTestScoreReport() elements.ScoreReport {
	return elements.ScoreReport{
		ReceivedPoints:  map[elements.ID]uint32{0: 5},
		ReturnedMeeples: map[elements.ID][]elements.MeepleWithPosition{},
	}
}
