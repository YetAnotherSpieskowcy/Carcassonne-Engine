package game

import (
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/test"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
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
	ptile := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	ptile.Position = elements.NewPosition(0, 1)
	_, err := board.PlaceTile(ptile)
	if err != nil {
		t.Fatal(err.Error())
	}

	expected := []elements.PlacedTile{}
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

	actual, ok := board.GetTileAt(expected.Position)
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
	expected[3].Position = elements.NewPosition(0, -1)

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
Test creating simple loop with starting tile
using straight roads and road turns, also tests two scoring players
*/
func TestBoardScoreRoadLoop(t *testing.T) {
	var report elements.ScoreReport
	var boardInterface interface{} = NewBoard(tilesets.StandardTileSet())
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		test.GetTestCustomPlacedTile(tiletemplates.StraightRoads(), 1),
		test.GetTestCustomPlacedTile(tiletemplates.RoadsTurn(), 2),
		test.GetTestCustomPlacedTile(tiletemplates.RoadsTurn(), 1),

		test.GetTestCustomPlacedTile(tiletemplates.RoadsTurn(), 1),
		test.GetTestCustomPlacedTile(tiletemplates.RoadsTurn(), 1),
	}

	// add meeple to first road
	tiles[0].Meeple.Side = side.Right
	tiles[0].Meeple.Type = 0

	tiles[1].Meeple.Side = side.Right
	tiles[1].Meeple.Type = 0

	// set positions
	tiles[0].Pos = elements.NewPosition(0, -1)
	tiles[1].Pos = elements.NewPosition(-1, 0)
	tiles[2].Pos = elements.NewPosition(-1, -1)
	tiles[3].Pos = elements.NewPosition(1, -1)
	tiles[4].Pos = elements.NewPosition(1, 0)

	// rotate tiles
	tiles[1].TilePlacement.Tile = tiles[1].TilePlacement.Tile.Rotate(3)
	tiles[2].TilePlacement.Tile = tiles[2].TilePlacement.Tile.Rotate(2)
	tiles[3].TilePlacement.Tile = tiles[3].TilePlacement.Tile.Rotate(1)

	expectedScores := []uint32{0, 0, 0, 0, 6}
	expectedMeeples := [][]uint8{nil, nil, nil, nil, {1}}

	// --------------- Placing tile ----------------------

	for i := range 5 {

		_, err := board.PlaceTile(tiles[i])
		if err != nil {
			t.Fatalf("error placing tile number: %#v ", i)
		}
		report = board.ScoreRoads(tiles[i])
		for _, playerID := range []uint8{1, 2} {
			if report.ReceivedPoints[playerID] != expectedScores[i] {
				t.Fatalf("placing tile number: %#v failed. expected %+v for player %v, got %+v instead", i, expectedScores[i], playerID, report.ReceivedPoints[playerID])
			}

			if !reflect.DeepEqual(report.ReturnedMeeples[playerID], expectedMeeples[i]) {
				t.Fatalf("placing tile number: %#v failed. expected %+v meeples for player %v, got %+v instead", i, expectedMeeples[i], playerID, report.ReturnedMeeples[playerID])
			}
		}
	}
}

/*
Test crossroad end to monastery and starting tile between
*/
func TestBoardScoreRoadCityMonastery(t *testing.T) {
	var report elements.ScoreReport
	var boardInterface interface{} = NewBoard(tilesets.StandardTileSet())
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		test.GetTestCustomPlacedTile(tiletemplates.MonasteryWithSingleRoad(), 1),
		test.GetTestCustomPlacedTile(tiletemplates.MonasteryWithSingleRoad(), 1),
	}

	// rotate tiles
	tiles[0].TilePlacement.Tile = tiles[0].TilePlacement.Tile.Rotate(3)
	tiles[1].TilePlacement.Tile = tiles[1].TilePlacement.Tile.Rotate(1)

	// add meeple to first road
	tiles[0].Meeple.Side = side.Right
	tiles[0].Meeple.Type = 0

	// set positions
	tiles[0].Pos = elements.NewPosition(-1, 0)
	tiles[1].Pos = elements.NewPosition(1, 0)

	expectedScores := []uint32{0, 3}
	expectedMeeples := [][]uint8{nil, {1}}

	// --------------- Placing tile ----------------------

	for i := range len(tiles) {

		_, err := board.PlaceTile(tiles[i])

		if err != nil {
			t.Fatalf("error placing tile number: %#v ", i)
		}

		report = board.ScoreRoads(tiles[i])
		if report.ReceivedPoints[1] != expectedScores[i] {
			t.Fatalf("placing tile number: %#v failed. expected %+v, got %+v instead", i, expectedScores[i], report.ReceivedPoints[1])
		}

		if !reflect.DeepEqual(report.ReturnedMeeples[1], expectedMeeples[i]) {
			t.Fatalf("placing tile number: %#v failed. expected %+v meeples, got %+v instead", i, expectedMeeples[i], report.ReturnedMeeples[1])
		}
	}
}

func TestBoardScoreRoadMultipleMeeplesOnSameRoad(t *testing.T) {
	var report elements.ScoreReport
	var boardInterface interface{} = NewBoard(tilesets.StandardTileSet())
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		test.GetTestCustomPlacedTile(tiletemplates.MonasteryWithSingleRoad(), 1), // on the right
		test.GetTestCustomPlacedTile(tiletemplates.MonasteryWithSingleRoad(), 1), // below
		test.GetTestCustomPlacedTile(tiletemplates.RoadsTurn(), 1),               // on the left bottom
		test.GetTestCustomPlacedTile(tiletemplates.RoadsTurn(), 1),               // on the left
	}

	// rotate tiles
	tiles[0].TilePlacement.Tile = tiles[0].TilePlacement.Tile.Rotate(1)
	tiles[1].TilePlacement.Tile = tiles[1].TilePlacement.Tile.Rotate(1)
	tiles[2].TilePlacement.Tile = tiles[2].TilePlacement.Tile.Rotate(2)
	tiles[3].TilePlacement.Tile = tiles[3].TilePlacement.Tile.Rotate(3)

	// add meeples to monastery roads
	tiles[0].Meeple.Side = side.Left
	tiles[0].Meeple.Type = 0
	tiles[1].Meeple.Side = side.Left
	tiles[1].Meeple.Type = 0

	// set positions
	tiles[0].Pos = elements.NewPosition(1, 0)
	tiles[1].Pos = elements.NewPosition(0, -1)
	tiles[2].Pos = elements.NewPosition(-1, -1)
	tiles[3].Pos = elements.NewPosition(-1, 0)

	expectedScores := []uint32{0, 0, 0, 5}
	expectedMeeples := [][]uint8{nil, nil, nil, {2}}

	// --------------- Placing tile ----------------------
	for i := range len(tiles) {
		_, err := board.PlaceTile(tiles[i])

		if err != nil {
			t.Fatalf("error placing tile number: %#v ", i)
		}

		report = board.ScoreRoads(tiles[i])
		if report.ReceivedPoints[1] != expectedScores[i] {
			t.Fatalf("placing tile number: %#v failed. expected %+v, got %+v instead", i, expectedScores[i], report.ReceivedPoints[1])
		}

		if !reflect.DeepEqual(report.ReturnedMeeples[1], expectedMeeples[i]) {
			t.Fatalf("placing tile number: %#v failed. expected %+v meeples, got %+v instead", i, expectedMeeples[i], report.ReturnedMeeples[1])
		}
	}
}
