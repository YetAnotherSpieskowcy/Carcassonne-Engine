package game

import (
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

/*
Test creating simple loop with starting tile
using straight roads and road turns, also tests two scoring players
Roads:

	2 -	0 -	5
	|		|
	3 -	1 -	4

Meeples are placed on tiles 1 and 2
*/
func TestBoardScoreRoadLoop(t *testing.T) {
	var report elements.ScoreReport
	var boardInterface interface{} = NewBoard(tilesets.StandardTileSet())
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		elements.ToPlacedTile(tiletemplates.StraightRoads()),
		elements.ToPlacedTile(tiletemplates.RoadsTurn().Rotate(3)),
		elements.ToPlacedTile(tiletemplates.RoadsTurn().Rotate(2)),

		elements.ToPlacedTile(tiletemplates.RoadsTurn().Rotate(1)),
		elements.ToPlacedTile(tiletemplates.RoadsTurn()),
	}

	// add meeple to first road
	tiles[0].GetPlacedFeatureAtSide(side.Right, feature.Road).Meeple.PlayerID = 1
	tiles[0].GetPlacedFeatureAtSide(side.Right, feature.Road).Meeple.Type = elements.NormalMeeple
	// add meeple to second road
	tiles[1].GetPlacedFeatureAtSide(side.Right, feature.Road).Meeple.PlayerID = 2
	tiles[1].GetPlacedFeatureAtSide(side.Right, feature.Road).Meeple.Type = elements.NormalMeeple

	// set positions
	tiles[0].Position = position.New(0, -1)
	tiles[1].Position = position.New(-1, 0)
	tiles[2].Position = position.New(-1, -1)
	tiles[3].Position = position.New(1, -1)
	tiles[4].Position = position.New(1, 0)

	expectedScores := []uint32{0, 0, 0, 0, 6}
	expectedMeeples := [][][]elements.MeepleWithPosition{
		{[]elements.MeepleWithPosition(nil), []elements.MeepleWithPosition(nil)},
		{[]elements.MeepleWithPosition(nil), []elements.MeepleWithPosition(nil)},
		{[]elements.MeepleWithPosition(nil), []elements.MeepleWithPosition(nil)},
		{[]elements.MeepleWithPosition(nil), []elements.MeepleWithPosition(nil)},
		{
			[]elements.MeepleWithPosition{elements.NewMeepleWithPosition(
				elements.Meeple{Type: elements.NormalMeeple, PlayerID: elements.ID(1)},
				position.New(0, -1),
			)},

			[]elements.MeepleWithPosition{elements.NewMeepleWithPosition(
				elements.Meeple{Type: elements.NormalMeeple, PlayerID: elements.ID(2)},
				position.New(-1, 0),
			)},
		},
	}

	// --------------- Placing tile ----------------------

	for i := range 5 {
		err := board.addTileToBoard(tiles[i])
		if err != nil {
			t.Fatalf("error placing tile number: %#v ", i)
		}

		binaryTile := elements.BinaryTileFromPlacedTile(tiles[i]) // todo binarytiles rewrite

		report = board.scoreRoads(binaryTile, false)
		for _, playerID := range []elements.ID{1, 2} {
			if report.ReceivedPoints[playerID] != expectedScores[i] {
				t.Fatalf("placing tile number: %#v failed. expected %+v for player %v, got %+v instead", i, expectedScores[i], playerID, report.ReceivedPoints[playerID])
			}

			if !reflect.DeepEqual(report.ReturnedMeeples[playerID], expectedMeeples[i][playerID-1]) {
				t.Fatalf("placing tile number: %#v failed. expected %+v meeples for player %v, got %+v instead", i, expectedMeeples[i][playerID-1], playerID, report.ReturnedMeeples[playerID])
			}
		}
	}
}

/*
Test loop, but the final tile is a crossRoad
Roads:
  - 0 -
    1 -	2
    |	|
  - 4 -	3
*/
func TestBoardScoreRoadLoopCrossroad(t *testing.T) {
	var boardInterface interface{} = NewBoard(tilesets.StandardTileSet())
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		elements.ToPlacedTile(tiletemplates.RoadsTurn().Rotate(3)),
		elements.ToPlacedTile(tiletemplates.RoadsTurn().Rotate(0)),
		elements.ToPlacedTile(tiletemplates.RoadsTurn().Rotate(1)),

		elements.ToPlacedTile(tiletemplates.TCrossRoad().Rotate(2)),
	}

	// add meeple to last road
	tiles[3].GetPlacedFeatureAtSide(side.Top, feature.Road).Meeple.PlayerID = 1
	tiles[3].GetPlacedFeatureAtSide(side.Top, feature.Road).Meeple.Type = elements.NormalMeeple

	// set positions
	tiles[0].Position = position.New(0, -1)
	tiles[1].Position = position.New(1, -1)
	tiles[2].Position = position.New(1, -2)
	tiles[3].Position = position.New(0, -2)

	expectedScores := []uint32{0, 0, 0, 4}
	expectedMeeples := [][]elements.MeepleWithPosition{
		[]elements.MeepleWithPosition(nil),
		[]elements.MeepleWithPosition(nil),
		[]elements.MeepleWithPosition(nil),
		[]elements.MeepleWithPosition{elements.NewMeepleWithPosition(
			elements.Meeple{Type: elements.NormalMeeple, PlayerID: elements.ID(1)},
			position.New(0, -2),
		)}}

	// --------------- Placing tile ----------------------

	for i := range 4 {
		println(i, ":")
		report, err := board.PlaceTile(tiles[i])
		if err != nil {
			t.Fatalf("error placing tile number: %#v ", i)
		}

		if report.ReceivedPoints[1] != expectedScores[i] {
			t.Fatalf("placing tile number: %#v failed. expected %+v for player %v, got %+v instead", i, expectedScores[i], 1, report.ReceivedPoints[1])
		}

		if !reflect.DeepEqual(report.ReturnedMeeples[1], expectedMeeples[i]) {
			t.Fatalf("placing tile number: %#v failed. expected %#v meeples for player %v, got %#v instead", i, expectedMeeples[i], 1, report.ReturnedMeeples[1])
		}

	}
}

/*
Test crossroad end to monastery and starting tile between

Roads:

	1 -	0 -	2
*/
func TestBoardScoreRoadCityMonastery(t *testing.T) {
	var report elements.ScoreReport
	var boardInterface interface{} = NewBoard(tilesets.StandardTileSet())
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad().Rotate(3)), // monastery to right
		elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad().Rotate(1)), // monastery to left
	}

	// add meeple to first road
	tiles[0].GetPlacedFeatureAtSide(side.Right, feature.Road).Meeple = elements.Meeple{PlayerID: 1, Type: elements.NormalMeeple}

	// set positions
	tiles[0].Position = position.New(-1, 0)
	tiles[1].Position = position.New(1, 0)

	expectedScores := []uint32{0, 3}
	expectedMeeples := [][]elements.MeepleWithPosition{
		[]elements.MeepleWithPosition(nil),
		[]elements.MeepleWithPosition{elements.NewMeepleWithPosition(
			elements.Meeple{Type: elements.NormalMeeple, PlayerID: elements.ID(1)},
			position.New(-1, 0),
		)}}

	// --------------- Placing tile ----------------------

	for i := range len(tiles) {
		err := board.addTileToBoard(tiles[i])

		if err != nil {
			t.Fatalf("error placing tile number: %#v ", i)
		}

		binaryTile := elements.BinaryTileFromPlacedTile(tiles[i]) // todo binarytiles rewrite

		report = board.scoreRoads(binaryTile, false)
		if report.ReceivedPoints[1] != expectedScores[i] {
			t.Fatalf("placing tile number: %#v failed. expected %+v, got %+v instead", i, expectedScores[i], report.ReceivedPoints[1])
		}

		if !reflect.DeepEqual(report.ReturnedMeeples[1], expectedMeeples[i]) {
			t.Fatalf("placing tile number: %#v failed. expected %+v meeples, got %+v instead", i, expectedMeeples[i], report.ReturnedMeeples[1])
		}
	}
}

/*
Roads:

	4 -	0 -	1
	|
	3 -	2
*/
func TestBoardScoreRoadMultipleMeeplesOnSameRoad(t *testing.T) {
	var report elements.ScoreReport
	var boardInterface interface{} = NewBoard(tilesets.StandardTileSet())
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad().Rotate(1)), // on the right
		elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad().Rotate(1)), // below
		elements.ToPlacedTile(tiletemplates.RoadsTurn().Rotate(2)),               // on the left bottom
		elements.ToPlacedTile(tiletemplates.RoadsTurn().Rotate(3)),               // on the left
	}

	// add meeples to monastery roads
	tiles[0].GetPlacedFeatureAtSide(side.Left, feature.Road).Meeple.PlayerID = 1
	tiles[0].GetPlacedFeatureAtSide(side.Left, feature.Road).Meeple.Type = elements.NormalMeeple
	tiles[1].GetPlacedFeatureAtSide(side.Left, feature.Road).Meeple.PlayerID = 1
	tiles[1].GetPlacedFeatureAtSide(side.Left, feature.Road).Meeple.Type = elements.NormalMeeple

	// set positions
	tiles[0].Position = position.New(1, 0)
	tiles[1].Position = position.New(0, -1)
	tiles[2].Position = position.New(-1, -1)
	tiles[3].Position = position.New(-1, 0)

	expectedScores := []uint32{0, 0, 0, 5}
	// expectedMeeples := [][]uint8{nil, nil, nil, {0, 2}}
	expectedMeeples := [][]elements.MeepleWithPosition{
		[]elements.MeepleWithPosition(nil),
		[]elements.MeepleWithPosition(nil),
		[]elements.MeepleWithPosition(nil),
		[]elements.MeepleWithPosition{
			elements.NewMeepleWithPosition(
				elements.Meeple{Type: elements.NormalMeeple, PlayerID: elements.ID(1)},
				position.New(1, 0)),
			elements.NewMeepleWithPosition(
				elements.Meeple{Type: elements.NormalMeeple, PlayerID: elements.ID(1)},
				position.New(0, -1)),
		}}

	// --------------- Placing tile ----------------------
	for i := range len(tiles) {
		err := board.addTileToBoard(tiles[i])

		if err != nil {
			t.Fatalf("error placing tile number: %#v ", i)
		}

		binaryTile := elements.BinaryTileFromPlacedTile(tiles[i]) // todo binarytiles rewrite

		report = board.scoreRoads(binaryTile, false)
		if report.ReceivedPoints[1] != expectedScores[i] {
			t.Fatalf("placing tile number: %#v failed. expected %+v, got %+v instead", i, expectedScores[i], report.ReceivedPoints[1])
		}

		if !reflect.DeepEqual(report.ReturnedMeeples[1], expectedMeeples[i]) {
			t.Fatalf("placing tile number: %#v failed. expected %+v meeples, got %+v instead", i, expectedMeeples[i], report.ReturnedMeeples[1])
		}
	}
}

/*
board drawing:
|			-1	  0	   1    2    3
|
|
|          .....|   |..|..
|          ......\ /...|..
|0         ..2----0----1..
|          ..|.........!..
|          ..@.........|..
|          ..|.........|..
|          ..|.........|..
|-1        ..5----4----3..
|          ..|............
|          ..!............
*/
func TestScoreRoadPreventCheckingWithNoSideAtTile5(t *testing.T) {
	var report elements.ScoreReport
	var boardInterface interface{} = NewBoard(tilesets.StandardTileSet())
	board := boardInterface.(*board)

	tiles := []elements.PlacedTile{
		elements.ToPlacedTile(tiletemplates.TCrossRoad().Rotate(1)), // 1
		elements.ToPlacedTile(tiletemplates.RoadsTurn().Rotate(3)),  // 2
		elements.ToPlacedTile(tiletemplates.RoadsTurn().Rotate(1)),  // 3
		elements.ToPlacedTile(tiletemplates.StraightRoads()),        // 4
		elements.ToPlacedTile(tiletemplates.TCrossRoad().Rotate(3)), // 5
	}

	// add meeples to roads
	tiles[0].GetPlacedFeatureAtSide(side.Bottom, feature.Road).Meeple.PlayerID = 1
	tiles[0].GetPlacedFeatureAtSide(side.Bottom, feature.Road).Meeple.Type = elements.NormalMeeple
	tiles[1].GetPlacedFeatureAtSide(side.Bottom, feature.Road).Meeple.PlayerID = 2
	tiles[1].GetPlacedFeatureAtSide(side.Bottom, feature.Road).Meeple.Type = elements.NormalMeeple
	tiles[4].GetPlacedFeatureAtSide(side.Bottom, feature.Road).Meeple.PlayerID = 1
	tiles[4].GetPlacedFeatureAtSide(side.Bottom, feature.Road).Meeple.Type = elements.NormalMeeple

	// set positions
	tiles[0].Position = position.New(1, 0)
	tiles[1].Position = position.New(-1, 0)
	tiles[2].Position = position.New(1, -1)
	tiles[3].Position = position.New(0, -1)
	tiles[4].Position = position.New(-1, -1)

	expectedScores := []map[elements.ID]uint32{
		{1: 0, 2: 0},
		{1: 0, 2: 0},
		{1: 0, 2: 0},
		{1: 0, 2: 0},
		{1: 4, 2: 4},
	}

	expectedMeeples := []map[elements.ID][]elements.MeepleWithPosition{
		{
			1: []elements.MeepleWithPosition(nil),
			2: []elements.MeepleWithPosition(nil),
		},
		{
			1: []elements.MeepleWithPosition(nil),
			2: []elements.MeepleWithPosition(nil),
		},
		{
			1: []elements.MeepleWithPosition(nil),
			2: []elements.MeepleWithPosition(nil),
		},
		{
			1: []elements.MeepleWithPosition(nil),
			2: []elements.MeepleWithPosition(nil),
		},
		{
			1: []elements.MeepleWithPosition{
				elements.NewMeepleWithPosition(
					elements.Meeple{Type: elements.NormalMeeple, PlayerID: elements.ID(1)},
					position.New(1, 0))},
			2: []elements.MeepleWithPosition{
				elements.NewMeepleWithPosition(
					elements.Meeple{Type: elements.NormalMeeple, PlayerID: elements.ID(2)},
					position.New(-1, 0))},
		},
	}

	// --------------- Placing tile ----------------------
	for i := range len(tiles) {
		err := board.addTileToBoard(tiles[i])

		if err != nil {
			t.Fatalf("error placing tile number: %#v ", i+1)
		}

		binaryTile := elements.BinaryTileFromPlacedTile(tiles[i]) // todo binarytiles rewrite

		report = board.scoreRoads(binaryTile, false)
		for playerID, points := range report.ReceivedPoints {
			if points != expectedScores[i][playerID] {
				t.Fatalf("Player %#v placing tile number: %#v failed. Received points:%#v,  expected %#v", playerID, i+1, points, expectedScores[i][playerID])
			}
		}

		for playerID, meeples := range report.ReturnedMeeples {
			if !reflect.DeepEqual(meeples, expectedMeeples[i][playerID]) {
				t.Fatalf("placing tile number: %#v failed. expected %#v meeples, got %#v instead", i+1, expectedMeeples[i][playerID], meeples)
			}
		}
	}
}
