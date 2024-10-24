package game

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/test"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func validateScores(game *Game, expectedScores []uint32, t *testing.T) {
	report := game.GetMidGameScore()
	for i := range 2 {
		if report.ReceivedPoints[elements.ID(i+1)] != expectedScores[i] {
			t.Fatalf("Player %d mid game score incorrect. Expected %d, got: %d", i+1, expectedScores[i], report.ReceivedPoints[elements.ID(i+1)])
		}
	}
}

/*
|            -1   0    1
|               | @ |
|               .\ /.
|1              ..4..
|               ./ \.
|               |   |
|     ..!.......|   |.....
|     ...........\ /..[!].
|0    --3----2-@--0---[1].
|     ..|.............[ ].
|     ..|.................
*/
func TestScoringMidGame(t *testing.T) { // nolint: gocyclo
	// ------ create tileset --------
	var tiles []tiles.Tile
	var err error
	tiles = append(tiles, tiletemplates.MonasteryWithSingleRoad().Rotate(1))
	tiles = append(tiles, tiletemplates.StraightRoads())
	tiles = append(tiles, tiletemplates.TCrossRoad())
	tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownNotConnected())

	tileset := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}

	// ------ create game --------
	deckStack := stack.NewOrdered(tileset.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileset.StartingTile}

	game, err := NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	var expectedScores []uint32
	// first turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(1, 0),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.NoSide, FeatureType: feature.Monastery},
	}.Run()
	validateScores(game, []uint32{2, 0}, t)

	// second turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(-1, 0),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Right, FeatureType: feature.Road},
	}.Run()
	validateScores(game, []uint32{2, 3}, t)

	// third turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(-2, 0),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.Field},
	}.Run()
	validateScores(game, []uint32{2, 4}, t)

	// fourth turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(0, 1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.City},
	}.Run()
	validateScores(game, []uint32{6, 5}, t)

	// finalize
	expectedScores = []uint32{6, 5}
	scores, err := game.Finalize()
	if err != nil {
		t.Fatal(err.Error())
	}

	for i := range 2 {
		if scores.ReceivedPoints[elements.ID(i+1)] != expectedScores[i] {
			t.Fatalf("Player %d final score incorrect. Expected %d, got: %d", i+1, expectedScores[i], scores.ReceivedPoints[elements.ID(i+1)])
		}
	}
}

/*
board drawing:
|			-1	  0	   1    2    3
|
|
|               |   |.....
|               .\ /......
|0              --0--!-1..
|               .......|..
|               .......|..
|          ..|.........|.......
|          ..|.........|.......
|-1        ..5----3----2...4...
|          ..|.................
|          ..!.................
*/
func TestMidGameScorePreventCalculatingSameMeepleOnRoadsMultipleTimes(t *testing.T) {
	// ------ create tileset --------
	var err error
	tiles := []tiles.Tile{
		tiletemplates.RoadsTurn(),            // 1
		tiletemplates.RoadsTurn().Rotate(1),  // 2
		tiletemplates.StraightRoads(),        // 3
		tiletemplates.TestOnlyField(),        // 4
		tiletemplates.TCrossRoad().Rotate(3), // 5
	}

	tileset := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}

	// create turns
	tilePositions := []position.Position{
		position.New(1, 0),
		position.New(1, -1),
		position.New(0, -1),
		position.New(2, -1),
		position.New(-1, -1),
	}

	meepleParams := []test.MeepleParams{
		{MeepleType: elements.NormalMeeple, FeatureSide: side.Left, FeatureType: feature.Road},
		test.NoneMeeple(),
		test.NoneMeeple(),
		test.NoneMeeple(),
		{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road},
	}

	// ------ create game --------
	deckStack := stack.NewOrdered(tileset.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileset.StartingTile}

	game, err := NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedScores := []map[elements.ID]uint32{
		{1: 2, 2: 0},
		{1: 3, 2: 0},
		{1: 4, 2: 0},
		{1: 4, 2: 0},
		{1: 6, 2: 0},
	}

	// --------------- Placing tile ----------------------
	for i := range len(tiles) {
		// first turn
		test.MakeTurn{
			Game:         game,
			TestingT:     t,
			Position:     tilePositions[i],
			MeepleParams: meepleParams[i],
		}.Run()

		if err != nil {
			t.Fatalf("error placing tile number: %#v ", i+1)
		}

		midGameScore := game.GetMidGameScore()
		for playerID, points := range midGameScore.ReceivedPoints {
			if points != expectedScores[i][playerID] {
				t.Fatalf("Player %#v placing tile number: %#v failed. Received points:%#v,  expected %#v", playerID, i+1, points, expectedScores[i][playerID])
			}
		}
	}
}
