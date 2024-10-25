package game

import (
	"fmt"
	"reflect"
	"slices"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/test"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/binarytiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func TestBoardDeepClone(t *testing.T) {
	oldPos := position.New(0, 2)
	newPos := position.New(0, 3) // just one of new positions
	newTile := test.GetTestPlacedTile()
	newTile.Position = oldPos

	expectedMeeplePos := position.New(0, 1)
	expectedMeeple := elements.Meeple{
		PlayerID: 1,
		Type:     elements.NormalMeeple,
	}

	original := NewBoard(tilesets.StandardTileSet()).(*board)
	// add a tile with meeple to verify that it's still there on the original later
	ptile := elements.ToPlacedTile(tiletemplates.TwoCityEdgesUpAndDownConnected())
	ptile.Position = expectedMeeplePos
	ptile.GetPlacedFeatureAtSide(side.Bottom, feature.City).Meeple = expectedMeeple
	_, err := original.PlaceTile(ptile)
	if err != nil {
		t.Fatal(err)
	}

	clone := original.DeepClone().(*board)
	_, err = clone.PlaceTile(newTile)
	if err != nil {
		t.Fatal(err)
	}

	// --- placeablePositions check ---
	if !slices.Contains(original.placeablePositions, oldPos) {
		t.Fatalf("expected to find %#v in %#v", oldPos, original.placeablePositions)
	}
	// just to confirm that `oldPos` actually makes sense
	if slices.Contains(clone.placeablePositions, oldPos) {
		t.Fatalf("expected NOT to find %#v in %#v", oldPos, clone.placeablePositions)
	}

	if !slices.Contains(clone.placeablePositions, newPos) {
		t.Fatalf("expected to find %#v in %#v", newPos, clone.placeablePositions)
	}
	// just to confirm that `newPos` actually makes sense
	if slices.Contains(original.placeablePositions, newPos) {
		t.Fatalf("expected NOT to find %#v in %#v", newPos, original.placeablePositions)
	}

	// --- tiles check ---
	cmpFunc := func(v elements.PlacedTile) bool {
		return slices.Equal(v.Features, newTile.Features)
	}

	originalTiles := original.Tiles()
	if slices.ContainsFunc(originalTiles, cmpFunc) {
		t.Fatalf("expected NOT to find %#v in %#v", newTile, originalTiles)
	}

	cloneTiles := clone.Tiles()
	if !slices.ContainsFunc(cloneTiles, cmpFunc) {
		t.Fatalf("expected to find %#v in %#v", newTile, cloneTiles)
	}

	// check that meeple is still present on the original
	originalTile, ok := original.GetTileAt(expectedMeeplePos)
	if !ok {
		t.Fatalf("expected to find a tile at %#v", expectedMeeplePos)
	}
	actual := originalTile.GetPlacedFeatureAtSide(side.Bottom, feature.City)
	actualMeeple := actual.Meeple

	if expectedMeeple != actualMeeple {
		t.Fatalf("expected %#v, got %#v instead", expectedMeeple, actualMeeple)
	}
	// just to confirm that clone does not have the meeple
	clonedTile, ok := clone.GetTileAt(expectedMeeplePos)
	if !ok {
		t.Fatalf("expected to find a tile at %#v", expectedMeeplePos)
	}

	clonedMeeple := clonedTile.GetPlacedFeatureAtSide(side.Bottom, feature.City).Meeple
	if clonedMeeple.Type != elements.NoneMeeple {
		t.Fatalf("expected %#v to have no meeple", clonedTile)
	}
}

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
	ptile.Position = position.New(0, 1)
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

func TestBoardGetLegalMovesForDoesNotIncludeInvalidMeeplePlacements(t *testing.T) {
	// starting tile has a city on top, we want to expand it with an unclosed city
	// and then try finding legal moves for a tile with a city and some other feature.
	board := NewBoard(tilesets.StandardTileSet())
	ptile := elements.ToPlacedTile(tiletemplates.TwoCityEdgesUpAndDownConnected())
	ptile.Position = position.New(0, 1)
	ptile.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple = elements.Meeple{
		Type: elements.NormalMeeple, PlayerID: 1,
	}
	_, err := board.PlaceTile(ptile)
	if err != nil {
		t.Fatal(err.Error())
	}

	basePlacement := elements.ToPlacedTile(
		tiletemplates.SingleCityEdgeNoRoads().Rotate(2),
	)
	basePlacement.Position = position.New(0, 2)
	placementWithMeeple := basePlacement.DeepClone()
	placementWithMeeple.GetPlacedFeatureAtSide(
		side.Top, feature.Field,
	).Meeple = elements.Meeple{Type: elements.NormalMeeple}

	expected := []elements.PlacedTile{basePlacement, placementWithMeeple}
	actual := board.GetLegalMovesFor(basePlacement)

	if !reflect.DeepEqual(expected, actual) {
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

func TestBoardCanBePlacedReturnsFalseWhenMultipleFeaturesHaveMeeples(t *testing.T) {
	board := NewBoard(tilesets.StandardTileSet())
	ptile := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	ptile.Position = position.New(0, 1)
	ptile.Features[0].Meeple = elements.Meeple{Type: elements.NormalMeeple, PlayerID: 1}
	ptile.Features[1].Meeple = elements.Meeple{Type: elements.NormalMeeple, PlayerID: 1}

	expected := false
	actual := board.CanBePlaced(ptile)

	if expected != actual {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestBoardCanBePlacedReturnsFalseWhenPlacingAtInvalidPosition(t *testing.T) {
	board := NewBoard(tilesets.StandardTileSet())
	ptile := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	ptile.Position = position.New(0, 2)

	expected := false
	actual := board.CanBePlaced(ptile)

	if expected != actual {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestBoardFieldCanBePlacedReturnsFalseWhenExpandToFieldWithMeepleHappensOverAnotherField(t *testing.T) {
	board := NewBoard(tilesets.StandardTileSet()).(*board)

	// prepare board layout (graphical representation can be found in issue GH-86)
	tilesToPlace := []elements.PlacedTile{}
	ptile := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	ptile.Position = position.New(0, 1)
	ptile.GetPlacedFeatureAtSide(side.Top, feature.Field).Meeple = elements.Meeple{
		Type: elements.NormalMeeple, PlayerID: 1,
	}
	tilesToPlace = append(tilesToPlace, ptile)

	ptile = elements.ToPlacedTile(tiletemplates.StraightRoads().Rotate(1))
	ptile.Position = position.New(1, 1)
	tilesToPlace = append(tilesToPlace, ptile)

	ptile = elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad().Rotate(3))
	ptile.Position = position.New(-1, 0)
	tilesToPlace = append(tilesToPlace, ptile)

	for _, ptile := range tilesToPlace {
		if _, err := board.PlaceTile(ptile); err != nil {
			t.Fatal(err)
		}
	}

	ptile = elements.ToPlacedTile(tiletemplates.RoadsTurn().Rotate(1))
	ptile.Position = position.New(1, 0)
	feat := ptile.GetPlacedFeatureAtSide(side.Right, feature.Field)
	feat.Meeple = elements.Meeple{
		Type: elements.NormalMeeple, PlayerID: 2,
	}

	expected := false
	actual := board.fieldCanBePlaced(ptile, *feat)

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
	expected[3].Position = position.New(0, -1)

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

func TestIsPositionValidWhenPositionIsInvalid(t *testing.T) {
	boardInterface := NewBoard(tilesets.StandardTileSet())
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad().Rotate(1)), // field adjacent to city
		elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad()),           // road adjacent to city
		elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad().Rotate(2)), // road adjacent to field

		elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(1)), // city adjacent to road
		elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads()),           // field adjacent to road
		elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads()),           // city adjacent to field
	}

	// set positions
	tiles[0].Position = position.New(0, 1)
	tiles[1].Position = position.New(0, 1)
	tiles[2].Position = position.New(0, -1)

	tiles[3].Position = position.New(-1, 0)
	tiles[4].Position = position.New(-1, 0)
	tiles[5].Position = position.New(0, -1)

	// place tiles
	for i, tile := range tiles {
		valid := board.isPositionValid(tile)
		if valid == true {
			t.Fatalf("expected invalid position when placing tile number: %#v", i)
		}
	}
}

func TestIsPositionValidWhenPositionIsValid(t *testing.T) {
	boardInterface := NewBoard(tilesets.StandardTileSet())
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad().Rotate(1)),
		elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad()),
		elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2)),
	}

	// set positions
	tiles[0].Position = position.New(1, 0)
	tiles[1].Position = position.New(0, -1)
	tiles[2].Position = position.New(0, 1)

	// place tiles
	for i, tile := range tiles {
		valid := board.isPositionValid(tile)
		if valid == false {
			t.Fatalf("expected valid position when placing tile number: %#v", i)
		}
	}
}

func TestBoardScoreIncompleteMonastery(t *testing.T) {
	var report elements.ScoreReport
	var extendedTileSet = tilesets.StandardTileSet()
	for range 3 {
		extendedTileSet.Tiles = append(extendedTileSet.Tiles, tiletemplates.TestOnlyField())
	}
	boardInterface := NewBoard(extendedTileSet)
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		elements.ToPlacedTile(tiletemplates.MonasteryWithoutRoads()),
		elements.ToPlacedTile(tiletemplates.TestOnlyField()),
		elements.ToPlacedTile(tiletemplates.TestOnlyField()),
		elements.ToPlacedTile(tiletemplates.TestOnlyField()),
	}

	// add meeple to the monastery
	tiles[0].Monastery().Meeple.PlayerID = 1
	tiles[0].Monastery().Meeple.Type = elements.NormalMeeple

	// set positions
	tiles[0].Position = position.New(0, -1)
	tiles[1].Position = position.New(1, -1)
	tiles[2].Position = position.New(0, -2)
	tiles[3].Position = position.New(1, -2)

	// place tiles
	for i, tile := range tiles {
		binaryTile := binarytiles.FromPlacedTile(tile) // todo binarytiles rewrite
		err := board.addTileToBoard(tile)
		if err != nil {
			t.Fatalf("error placing tile number: %#v: %#v", i, err)
		}

		report = board.scoreMonasteries(binaryTile, false)
		if !reflect.DeepEqual(report, elements.NewScoreReport()) {
			t.Fatalf("scoreMonasteries() failed on tile number: %#v. expected %#v, got %#v instead", i, elements.NewScoreReport(), report)
		}
	}

	// test forceScore
	binaryTile := binarytiles.FromPlacedTile(tiles[0]) // todo binarytiles rewrite
	report = board.scoreMonasteries(binaryTile, true)

	expectedReport := elements.NewScoreReport()
	expectedReport.ReceivedPoints = map[elements.ID]uint32{
		1: 5,
	}
	expectedReport.ReturnedMeeples = map[elements.ID][]elements.MeepleWithPosition{
		1: {elements.NewMeepleWithPosition(
			elements.Meeple{Type: elements.NormalMeeple, PlayerID: elements.ID(1)},
			position.New(0, -1),
		)},
	}

	if !reflect.DeepEqual(report, expectedReport) {
		t.Fatalf("scoreMonasteries() failed when forceScore=true. expected:\n%#v,\ngot:\n%#v instead", expectedReport, report)
	}
}

func TestBoardCompleteTwoMonasteriesAtOnce(t *testing.T) {
	/*
		the board setup is as follows:
		 S
		FFFF
		FMMF
		FFFF

		F - field
		M - monastery
		S - starting tile

		left monastery is at (0,-2)
		right monastery is at (1,-2)
		right monastery is placed as the last tile
	*/

	var report elements.ScoreReport
	var extendedTileSet = tilesets.StandardTileSet()
	for range 10 {
		extendedTileSet.Tiles = append(extendedTileSet.Tiles, tiletemplates.TestOnlyField())
	}
	boardInterface := NewBoard(extendedTileSet)
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		elements.ToPlacedTile(tiletemplates.TestOnlyField()),
		elements.ToPlacedTile(tiletemplates.MonasteryWithoutRoads()),
		elements.ToPlacedTile(tiletemplates.TestOnlyField()),
		elements.ToPlacedTile(tiletemplates.TestOnlyField()),
		elements.ToPlacedTile(tiletemplates.TestOnlyField()),
		elements.ToPlacedTile(tiletemplates.TestOnlyField()),
		elements.ToPlacedTile(tiletemplates.TestOnlyField()),
		elements.ToPlacedTile(tiletemplates.TestOnlyField()),
		elements.ToPlacedTile(tiletemplates.TestOnlyField()),
		elements.ToPlacedTile(tiletemplates.TestOnlyField()),
		elements.ToPlacedTile(tiletemplates.TestOnlyField()),
		elements.ToPlacedTile(tiletemplates.MonasteryWithoutRoads()),
	}

	// add meeple to the monastery
	tiles[1].Monastery().Meeple.PlayerID = 1
	tiles[1].Monastery().Meeple.Type = elements.NormalMeeple

	tiles[11].Monastery().Meeple.PlayerID = 2
	tiles[11].Monastery().Meeple.Type = elements.NormalMeeple

	// set positions
	tiles[0].Position = position.New(0, -1)
	tiles[1].Position = position.New(0, -2)
	tiles[2].Position = position.New(0, -3)

	tiles[3].Position = position.New(-1, -1)
	tiles[4].Position = position.New(-1, -2)
	tiles[5].Position = position.New(-1, -3)

	tiles[6].Position = position.New(1, -1)

	tiles[7].Position = position.New(2, -1)
	tiles[8].Position = position.New(2, -2)
	tiles[9].Position = position.New(2, -3)

	tiles[10].Position = position.New(1, -3)

	tiles[11].Position = position.New(1, -2)

	// place tiles
	for i, tile := range tiles[:len(tiles)-1] {
		binaryTile := binarytiles.FromPlacedTile(tile) // todo binarytiles rewrite
		err := board.addTileToBoard(tile)
		if err != nil {
			t.Fatalf("error placing tile number: %#v: %#v", i, err)
		}

		report = board.scoreMonasteries(binaryTile, false)
		if !reflect.DeepEqual(report, elements.NewScoreReport()) {
			t.Fatalf("scoreMonasteries() failed on tile number: %#v. expected %#v, got %#v instead", i, elements.NewScoreReport(), report)
		}
	}

	// place the last tile
	err := board.addTileToBoard(tiles[11])
	if err != nil {
		t.Fatalf("error placing tile number: %#v: %#v", 11, err)
	}
	binaryTile := binarytiles.FromPlacedTile(tiles[11]) // todo binarytiles rewrite
	report = board.scoreMonasteries(binaryTile, false)
	expectedReport := elements.NewScoreReport()
	expectedReport.ReceivedPoints = map[elements.ID]uint32{
		1: 9,
		2: 9,
	}
	expectedReport.ReturnedMeeples = map[elements.ID][]elements.MeepleWithPosition{
		1: {elements.NewMeepleWithPosition(
			elements.Meeple{Type: elements.NormalMeeple, PlayerID: elements.ID(1)},
			position.New(0, -2),
		)},
		2: {elements.NewMeepleWithPosition(
			elements.Meeple{Type: elements.NormalMeeple, PlayerID: elements.ID(2)},
			position.New(1, -2),
		)},
	}
	if !reflect.DeepEqual(report, expectedReport) {
		t.Fatalf("scoreMonasteries() failed on tile number: %#v. expected:\n%#v,\ngot:\n%#v instead", 11, expectedReport, report)
	}
}

/*
Test if meeples are counter multiple times. They shouldn't be!

	board:

|            0  1  2  3  4  5  6  7 -> X
+-----------------------------------
|
|           ........................
|1          .1..2..3..4..5..6..7..D.
|           ........................
|           ...   ................
|0          -0-   -8M-B--9M-C--AM
|           ...   ...............
V

Y
*/
func TestScoreNotFinalMeeplesOnSameFeature(t *testing.T) {
	// define tileSlice
	tileSlice := []elements.PlacedTile{}
	for range 7 {
		tileSlice = append(tileSlice, elements.ToPlacedTile(tiletemplates.TestOnlyField()))
	}
	for range 5 {
		tileSlice = append(tileSlice, elements.ToPlacedTile(tiletemplates.StraightRoads()))
	}
	tileSlice = append(tileSlice, elements.ToPlacedTile(tiletemplates.TestOnlyField()))

	// set positions
	for i := range 7 {
		tileSlice[i].Position = position.New(int16(i), 1)
	}
	for i := range 3 {
		tileSlice[i+7].Position = position.New(int16(2*i+2), 0)
	}
	for i := range 2 {
		tileSlice[i+7+3].Position = position.New(int16(2*i+3), 0)
	}
	tileSlice[7+5].Position = position.New(7, 0)

	// add meeples
	for i := range 3 {
		tileSlice[i+7].GetPlacedFeatureAtSide(side.Left, feature.Road).Meeple = elements.Meeple{
			Type:     elements.NormalMeeple,
			PlayerID: elements.ID(1),
		}
	}

	// create board
	tileSet := tilesets.StandardTileSet()
	tileSet.StartingTile = tiletemplates.StraightRoads()
	tileSet.Tiles = []tiles.Tile{}
	for _, tile := range tileSlice {
		tileSet.Tiles = append(tileSet.Tiles, elements.ToTile(tile))
	}

	boardInterface := NewBoard(tileSet)
	board := boardInterface.(*board)

	// play all turns but one
	for i, tile := range tileSlice[:len(tileSlice)-1] {
		_, err := board.PlaceTile(tile)
		if err != nil {
			fmt.Printf("Tile: %#v\n", tile)
			t.Fatalf("error placing tile number: %#v: %#v", i+1, err)
		}
	}

	report := board.ScoreMeeples(false)
	expected := uint32(5)
	if report.ReceivedPoints[elements.ID(1)] != expected {
		if report.ReceivedPoints[elements.ID(1)] == expected*3 {
			t.Fatalf("Road was scored for each meeple!")
		} else {
			t.Fatalf("Wrong amount of points while scoring roads: %#v, expected: %#v", report.ReceivedPoints[elements.ID(1)], expected)
		}
	}
}

func TestScoreNotFinalMeeplesOnIncompleteCity(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	tileSet.Tiles = []tiles.Tile{tiletemplates.TwoCityEdgesUpAndDownConnected()}

	board := NewBoard(tileSet)

	ptile := elements.ToPlacedTile(tileSet.Tiles[0])
	ptile.Position = position.New(0, 1)
	ptile.GetPlacedFeatureAtSide(side.Bottom, feature.City).Meeple = elements.Meeple{
		Type:     elements.NormalMeeple,
		PlayerID: elements.ID(1),
	}
	_, err := board.PlaceTile(ptile)
	if err != nil {
		t.Fatal(err)
	}

	report := board.ScoreMeeples(false)
	actual := report.ReceivedPoints[elements.ID(1)]
	expected := uint32(2)

	if actual != expected {
		t.Fatalf("expected %v, got %v instead", expected, actual)
	}
}

func TestScoreNotFinalMeeplesOnFieldWithIncompleteCity(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	tileSet.Tiles = []tiles.Tile{tiletemplates.StraightRoads()}

	board := NewBoard(tileSet)

	ptile := elements.ToPlacedTile(tileSet.Tiles[0])
	ptile.Position = position.New(1, 0)
	ptile.GetPlacedFeatureAtSide(side.Top, feature.Field).Meeple = elements.Meeple{
		Type:     elements.NormalMeeple,
		PlayerID: elements.ID(1),
	}
	_, err := board.PlaceTile(ptile)
	if err != nil {
		t.Fatal(err)
	}

	report := board.ScoreMeeples(false)
	actual := report.ReceivedPoints[elements.ID(1)]
	expected := uint32(0)

	if actual != expected {
		t.Fatalf("expected %v, got %v instead", expected, actual)
	}
}

func TestScoreNotFinalMeeplesOnFieldWithCompleteCity(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	tileSet.Tiles = []tiles.Tile{
		tiletemplates.SingleCityEdgeNoRoads().Rotate(2),
		tiletemplates.StraightRoads(),
	}

	board := NewBoard(tileSet)

	ptile := elements.ToPlacedTile(tileSet.Tiles[0])
	ptile.Position = position.New(0, 1)
	_, err := board.PlaceTile(ptile)
	if err != nil {
		t.Fatal(err)
	}

	ptile = elements.ToPlacedTile(tileSet.Tiles[1])
	ptile.Position = position.New(1, 0)
	ptile.GetPlacedFeatureAtSide(side.Top, feature.Field).Meeple = elements.Meeple{
		Type:     elements.NormalMeeple,
		PlayerID: elements.ID(2),
	}
	_, err = board.PlaceTile(ptile)
	if err != nil {
		t.Fatal(err)
	}

	report := board.ScoreMeeples(false)
	actual := report.ReceivedPoints[elements.ID(2)]
	expected := uint32(3)

	if actual != expected {
		t.Fatalf("expected %v, got %v instead", expected, actual)
	}
}
