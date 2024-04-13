package game

import (
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/test"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func TestBoardTileCountReturnsOnlyPlacedTiles(t *testing.T) {
	// starting tile has a city on top, we want to close it with a single city tile
	// and then try finding legal moves of a tile filled with a city terrain
	board := NewBoard(tilesets.StandardTileSet())
	_, err := board.PlaceTile(test.GetTestPlacedTile())
	if err != nil {
		t.Fatal(err.Error())
	}

	expected := 2
	actual := board.TileCount()

	if expected != actual {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestBoardGetTilePlacementsForReturnsEmptySliceWhenCityCannotBePlaced(t *testing.T) {
	// starting tile has a city on top, we want to close it with a single city tile
	// and then try finding legal moves of a tile filled with a city terrain
	board := NewBoard(tilesets.StandardTileSet())
	_, err := board.PlaceTile(
		elements.PlacedTile{
			LegalMove: elements.LegalMove{
				TilePlacement: elements.TilePlacement{
					Tile: tiletemplates.SingleCityEdgeNoRoads().Rotate(2),
					Pos:  elements.NewPosition(0, 1),
				},
				Meeple: elements.MeeplePlacement{Side: side.None},
			},
		},
	)
	if err != nil {
		t.Fatal(err.Error())
	}

	expected := []elements.TilePlacement{}
	actual := board.GetTilePlacementsFor(tiletemplates.FourCityEdgesConnectedShield())

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestBoardTileHasValidPlacementReturnsTrueWhenValidPlacementExists(t *testing.T) {
	board := NewBoard(tilesets.StandardTileSet())

	expected := true
	actual := board.TileHasValidPlacement(tiletemplates.SingleCityEdgeNoRoads())

	if expected != actual {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestBoardCanBePlacedReturnsTrueWhenPlacedTileCanBePlaced(t *testing.T) {
	board := NewBoard(tilesets.StandardTileSet())

	expected := true
	actual := board.CanBePlaced(test.GetTestPlacedTile())

	if expected != actual {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestBoardPlaceTileErrorsWhenCapacityIsExceeded(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	tileSet.Tiles = []tiles.Tile{}
	board := NewBoard(tileSet)

	_, err := board.PlaceTile(test.GetTestPlacedTile())
	if err == nil {
		t.Fatal("expected capacity exceeded error to be returned")
	}
}

func TestBoardPlaceTileUpdatesBoardFields(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	tileSet.Tiles = []tiles.Tile{
		test.GetTestTile(), tiletemplates.FourCityEdgesConnectedShield(),
	}
	board := NewBoard(tileSet)
	expected := test.GetTestPlacedTile()

	_, err := board.PlaceTile(expected)
	if err != nil {
		t.Fatal(err.Error())
	}

	actual := board.Tiles()[1]
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}

	actual, ok := board.GetTileAt(expected.Pos)
	if !ok || !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, got %#v instead (ok = %#v)", expected, actual, ok)
	}
}

func TestBoardPlaceTilePlacesTwoTilesOfSameTypeProperly(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	tileSet.Tiles = []tiles.Tile{
		test.GetTestTile(),
		tiletemplates.FourCityEdgesConnectedShield(),
		test.GetTestTile(),
	}
	board := NewBoard(tileSet)
	startingPlacedTile := elements.NewStartingTile(tileSet)
	expected := []elements.PlacedTile{
		startingPlacedTile,
		test.GetTestPlacedTile(),
		{},
		test.GetTestPlacedTile(),
	}
	// place the test tile (single city edge) below starting tile
	// (connecting with the field)
	expected[3].Pos = elements.NewPosition(0, -1)

	_, err := board.PlaceTile(expected[1])
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = board.PlaceTile(expected[3])
	if err != nil {
		t.Fatal(err.Error())
	}

	actual := board.Tiles()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}
