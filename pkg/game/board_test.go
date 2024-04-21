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

/*
*
Test creating simple loop with starting tile
using straight roads and road turns
*/
func TestBoardScoreRoad(t *testing.T) {
	var report elements.ScoreReport
	var boardInterface interface{} = NewBoard(tilesets.StandardTileSet())
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		test.GetTestRoadTurnPlacedTile(),
		test.GetTestRoadTurnPlacedTile(),
		test.GetTestStraightRoadPlacedTile(),
		test.GetTestRoadTurnPlacedTile(),
		test.GetTestRoadTurnPlacedTile(),
	}

	// add meeple to first road
	tiles[0].Meeple.Side = side.Right
	tiles[0].Meeple.Type = 0

	// set positions
	tiles[0].Pos = elements.NewPosition(-1, 0)
	tiles[1].Pos = elements.NewPosition(-1, -1)
	tiles[2].Pos = elements.NewPosition(0, -1)
	tiles[3].Pos = elements.NewPosition(1, -1)
	tiles[4].Pos = elements.NewPosition(1, 0)

	// rotate tiles
	tiles[0].TilePlacement.Tile = tiles[0].TilePlacement.Tile.Rotate(3)
	tiles[1].TilePlacement.Tile = tiles[1].TilePlacement.Tile.Rotate(2)
	tiles[3].TilePlacement.Tile = tiles[3].TilePlacement.Tile.Rotate(1)

	expectedScores := []uint32{0, 0, 0, 0, 6}
	expectedMeeples := [][]uint8{nil, nil, nil, nil, {1}}

	// --------------- Placing tile ----------------------

	for i := range 5 {

		_, err := board.PlaceTile(tiles[i])
		if err != nil {
			t.Fatalf("error placing tile number: %#v ", i)
		}
		report = board.ScoreRoadCompletion(tiles[i], tiles[i].Tile.Roads()[0])
		if report.ReceivedPoints[1] != expectedScores[i] {
			t.Fatalf("placing tile number: %#v failed. expected %#v, got %#v instead", i, expectedScores[i], report.ReceivedPoints[1])
		}

		if !reflect.DeepEqual(report.ReturnedMeeples[1], expectedMeeples[i]) {
			t.Fatalf("placing tile number: %#v failed. expected %#v meeples, got %#v instead", i, report.ReturnedMeeples[1], expectedMeeples[i])
		}
	}
}
