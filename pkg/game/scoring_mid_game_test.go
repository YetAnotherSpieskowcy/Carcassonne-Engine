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
		TilePosition: position.New(1, 0),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.NoSide, FeatureType: feature.Monastery},
	}.Run()
	validateScores(game, []uint32{2, 0}, t)

	// second turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: position.New(-1, 0),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Right, FeatureType: feature.Road},
	}.Run()
	validateScores(game, []uint32{2, 3}, t)

	// third turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: position.New(-2, 0),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.Field},
	}.Run()
	validateScores(game, []uint32{2, 4}, t)

	// fourth turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: position.New(0, 1),
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
