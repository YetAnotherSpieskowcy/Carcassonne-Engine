package game

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/field"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/test"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func TestFeaturesLengthOfFieldExpand(t *testing.T) {
	/*
		the board setup is as follows:
		─┼SM·
		··┌─┐
		··M··

		· - empty
		M - monastery with a single road
		S - starting tile
		┌, ─, ┐, ┼ - roads

		The meeple is placed on the field feature in the higher monastery tile (position: 1,0).
	*/

	boardInterface := NewBoard(tilesets.StandardTileSet())
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		test.GetTestCustomPlacedTile(tiletemplates.MonasteryWithSingleRoad().Rotate(1)),
		test.GetTestCustomPlacedTile(tiletemplates.XCrossRoad()),
		test.GetTestCustomPlacedTile(tiletemplates.StraightRoads()),
		test.GetTestCustomPlacedTile(tiletemplates.StraightRoads()),
		test.GetTestCustomPlacedTile(tiletemplates.RoadsTurn()),
		test.GetTestCustomPlacedTile(tiletemplates.RoadsTurn().Rotate(3)),
		test.GetTestCustomPlacedTile(tiletemplates.MonasteryWithSingleRoad().Rotate(2)),
	}

	// add meeple to the field
	tiles[0].GetPlacedFeatureAtSide(side.All, feature.Field).Meeple.PlayerID = 1
	tiles[0].GetPlacedFeatureAtSide(side.All, feature.Field).Meeple.MeepleType = elements.NormalMeeple

	// set positions
	tiles[0].Position = elements.NewPosition(1, 0)
	tiles[1].Position = elements.NewPosition(-1, 0)
	tiles[2].Position = elements.NewPosition(-2, 0)

	tiles[3].Position = elements.NewPosition(1, -1)
	tiles[4].Position = elements.NewPosition(2, -1)
	tiles[5].Position = elements.NewPosition(0, -1)

	tiles[6].Position = elements.NewPosition(0, -2)

	// place tiles
	for i, tile := range tiles {
		_, err := board.PlaceTile(tile)
		if err != nil {
			t.Fatalf("error placing tile number: %#v: %#v", i, err)
		}
	}

	// test field.Expand()
	field := field.NewField(*tiles[0].GetPlacedFeatureAtSide(side.All, feature.Field), tiles[0].Position)
	field.Expand(board)

	if len(field.Features()) != 12 {
		t.Fatalf("expected %#v, got %#v instead", 12, len(field.Features()))
	}
}
