package game

import (
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/test"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/player"
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

func TestBoardScoreInclompleteMonastery(t *testing.T) {
	var report elements.ScoreReport
	var extendedTileSet = tilesets.StandardTileSet()
	for range 3 {
		extendedTileSet.Tiles = append(extendedTileSet.Tiles, tiletemplates.OnlyField())
	}
	var boardInterface interface{} = NewBoard(extendedTileSet)
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		test.GetTestTileMonasteryWithoutRoads(),
		test.GetTestTileOnlyField(),
		test.GetTestTileOnlyField(),
		test.GetTestTileOnlyField(),
	}

	// add meeple to the monastery
	tiles[0].Meeple.Side = side.Center
	tiles[0].Meeple.Type = 0

	// set positions
	tiles[0].Pos = elements.NewPosition(0, 0)
	tiles[1].Pos = elements.NewPosition(1, 0)
	tiles[2].Pos = elements.NewPosition(0, 1)
	tiles[3].Pos = elements.NewPosition(1, 1)

	// place tiles
	for i, tile := range tiles {
		_, err := board.PlaceTile(tile)
		if err != nil {
			t.Fatalf("error placing tile number: %#v: %#v", i, err)
		}

		report = board.ScoreMonasteries(tile, false)
		if !reflect.DeepEqual(report, elements.NewScoreReport()) {
			t.Fatalf("ScoreMonasteries failed on tile number: %#v. expected %#v, got %#v instead", i, elements.NewScoreReport(), report)
		}
	}

	// test forceScore
	report = board.ScoreMonasteries(tiles[0], true)
	var expected = elements.ScoreReport{
		ReceivedPoints: map[uint8]uint32{
			1: 4,
		},
		ReturnedMeeples: map[uint8][]uint8{
			1: {1},
		},
	}
	if !reflect.DeepEqual(report, expected) {
		t.Fatalf("ScoreMonasteries failed when forceScore=true. expected %#v,\ngot %#v instead", expected, report)
	}
}

func TestBoardCompleteTwoMonasteriesAtOnce(t *testing.T) {
	/*
		the board setup is as follows:
		FFFF
		FMMF
		FFFF

		F - field
		M - monastery

		left monastery is at (0,0)
		right monastery is placed as the last tile
	*/

	var report elements.ScoreReport
	var extendedTileSet = tilesets.StandardTileSet()
	for range 10 {
		extendedTileSet.Tiles = append(extendedTileSet.Tiles, tiletemplates.OnlyField())
	}
	var boardInterface interface{} = NewBoard(extendedTileSet)
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		test.GetTestTileMonasteryWithoutRoads(),
		test.GetTestTileOnlyField(),
		test.GetTestTileOnlyField(),
		test.GetTestTileOnlyField(),
		test.GetTestTileOnlyField(),
		test.GetTestTileOnlyField(),
		test.GetTestTileOnlyField(),
		test.GetTestTileOnlyField(),
		test.GetTestTileOnlyField(),
		test.GetTestTileOnlyField(),
		test.GetTestTileOnlyField(),
		test.GetTestTileMonasteryWithoutRoads(),
	}

	// add meeple to the monastery
	tiles[0].Meeple.Side = side.Center
	tiles[0].Meeple.Type = 0

	tiles[11].Meeple.Side = side.Center
	tiles[11].Meeple.Type = 0
	tiles[11].Player = player.New(2)

	// set positions
	tiles[0].Pos = elements.NewPosition(0, 0)

	tiles[1].Pos = elements.NewPosition(-1, -1)
	tiles[2].Pos = elements.NewPosition(0, -1)
	tiles[3].Pos = elements.NewPosition(1, -1)
	tiles[4].Pos = elements.NewPosition(2, -1)

	tiles[5].Pos = elements.NewPosition(-1, 0)
	tiles[6].Pos = elements.NewPosition(2, 0)

	tiles[7].Pos = elements.NewPosition(-1, 1)
	tiles[8].Pos = elements.NewPosition(0, 1)
	tiles[9].Pos = elements.NewPosition(1, 1)
	tiles[10].Pos = elements.NewPosition(2, 1)

	tiles[11].Pos = elements.NewPosition(1, 0)

	// place tiles
	for i, tile := range tiles[:len(tiles)-1] {
		_, err := board.PlaceTile(tile)
		if err != nil {
			t.Fatalf("error placing tile number: %#v: %#v", i, err)
		}

		report = board.ScoreMonasteries(tile, false)
		if !reflect.DeepEqual(report, elements.NewScoreReport()) {
			t.Fatalf("ScoreMonasteries failed on tile number: %#v. expected %#v, got %#v instead", i, elements.NewScoreReport(), report)
		}
	}

	// place the last tile
	_, err := board.PlaceTile(tiles[11])
	if err != nil {
		t.Fatalf("error placing tile number: %#v: %#v", 11, err)
	}
	report = board.ScoreMonasteries(tiles[11], false)
	var expected = elements.ScoreReport{
		ReceivedPoints: map[uint8]uint32{
			1: 9,
			2: 9,
		},
		ReturnedMeeples: map[uint8][]uint8{
			1: {1},
			2: {1},
		},
	}
	if !reflect.DeepEqual(report, expected) {
		t.Fatalf("ScoreMonasteries failed on tile number: %#v. expected %#v, got %#v instead", 11, expected, report)
	}
}
