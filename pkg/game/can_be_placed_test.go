package game

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func TestPlaceFieldDirectlyAdjacentToCity(t *testing.T) {
	/*
		the board setup is as follows:
		M
		S

		S - starting tile
		M - monastery with a single road, going left

		The top edge of the city directly neighbours the field (invalid placement)
	*/
	boardInterface := NewBoard(tilesets.StandardTileSet())
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad().Rotate(1)),
	}

	// set positions
	tiles[0].Position = position.New(0, 1)

	_, err := board.PlaceTile(tiles[0])
	if err == nil {
		t.Fatalf("expected error placing first tile")
	}
}

func TestPlaceTwoAdjacentFieldsWithMeeples(t *testing.T) {
	/*
		the board setup is as follows:
		S
		M
		M

		S - starting tile
		M - monastery with a single road, going left

		The meeples are placed on both monasteries' fields
	*/
	boardInterface := NewBoard(tilesets.StandardTileSet())
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad().Rotate(1)),
		elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad().Rotate(1)),
	}

	// add meeple to the fields
	tiles[0].GetPlacedFeatureAtSide(side.All, feature.Field).Meeple =
		elements.Meeple{PlayerID: 1, Type: elements.NormalMeeple}
	tiles[1].GetPlacedFeatureAtSide(side.All, feature.Field).Meeple =
		elements.Meeple{PlayerID: 1, Type: elements.NormalMeeple}

	// set positions
	tiles[0].Position = position.New(0, -1)
	tiles[1].Position = position.New(0, -2)

	_, err := board.PlaceTile(tiles[0])
	if err != nil {
		t.Fatalf("error placing first tile: %#v", err)
	}

	_, err = board.PlaceTile(tiles[1])
	if err == nil {
		t.Fatalf("expected error placing second tile")
	}
}

func TestConnectTwoFieldsWithMeeplesWithAThirdMeeple(t *testing.T) {
	/*
		the board setup is as follows:
		─SM
		M

		S - starting tile
		─ - road
		M - monastery with a single road, going left

		The tiles are placed in the following order:
		(S), ─, M(bottom-left), M(right)

		The meeples are placed on the top field of the ─ tile and on both monasteries' fields
	*/
	boardInterface := NewBoard(tilesets.StandardTileSet())
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		elements.ToPlacedTile(tiletemplates.StraightRoads()),
		elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad().Rotate(1)),
		elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad().Rotate(1)),
	}

	// add meeple to the fields
	tiles[0].GetPlacedFeatureAtSide(side.Top, feature.Field).Meeple =
		elements.Meeple{PlayerID: 1, Type: elements.NormalMeeple}
	tiles[1].GetPlacedFeatureAtSide(side.All, feature.Field).Meeple =
		elements.Meeple{PlayerID: 1, Type: elements.NormalMeeple}
	tiles[2].GetPlacedFeatureAtSide(side.All, feature.Field).Meeple =
		elements.Meeple{PlayerID: 1, Type: elements.NormalMeeple}

	// set positions
	tiles[0].Position = position.New(-1, 0)
	tiles[1].Position = position.New(-1, -1)
	tiles[2].Position = position.New(1, 0)

	_, err := board.PlaceTile(tiles[0])
	if err != nil {
		t.Fatalf("error placing first tile: %#v", err)
	}

	_, err = board.PlaceTile(tiles[1])
	if err != nil {
		t.Fatalf("error placing second tile: %#v", err)
	}

	_, err = board.PlaceTile(tiles[2])
	if err == nil {
		t.Fatalf("expected error placing third tile")
	}
}
