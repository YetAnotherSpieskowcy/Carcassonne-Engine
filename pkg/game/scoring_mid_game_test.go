package game

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

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
func TestScoringMidGame(t *testing.T) {
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

	game, err := NewFromDeck(deck, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	var expectedScores []uint32
	// first turn
	makeTurn(game, t, position.New(1, 0), MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.NoSide, FeatureType: feature.Monastery})
	report, err := game.GetMidGamePoints()
	if err != nil {
		t.Fatal(err.Error())
	}
	expectedScores = []uint32{2, 0}
	for i := range 2 {
		if report.ReceivedPoints[elements.ID(i+1)] != expectedScores[i] {
			t.Fatalf("Player %d mid game score incorrect. Expected %d, got: %d", i+1, expectedScores[i], report.ReceivedPoints[elements.ID(i+1)])
		}
	}

	// second turn
	makeTurn(game, t, position.New(-1, 0), MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Right, FeatureType: feature.Road})
	report, err = game.GetMidGamePoints()
	if err != nil {
		t.Fatal(err.Error())
	}
	expectedScores = []uint32{2, 3}
	for i := range 2 {
		if report.ReceivedPoints[elements.ID(i+1)] != expectedScores[i] {
			t.Fatalf("Player %d mid game score incorrect. Expected %d, got: %d", i+1, expectedScores[i], report.ReceivedPoints[elements.ID(i+1)])
		}
	}

	// third turn
	makeTurn(game, t, position.New(-2, 0), MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.Field})
	report, err = game.GetMidGamePoints()
	if err != nil {
		t.Fatal(err.Error())
	}
	expectedScores = []uint32{2, 4}
	for i := range 2 {
		if report.ReceivedPoints[elements.ID(i+1)] != expectedScores[i] {
			t.Fatalf("Player %d mid game score incorrect. Expected %d, got: %d", i+1, expectedScores[i], report.ReceivedPoints[elements.ID(i+1)])
		}
	}

	// fourth turn
	makeTurn(game, t, position.New(0, 1), MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.City})
	report, err = game.GetMidGamePoints()
	if err != nil {
		t.Fatal(err.Error())
	}
	expectedScores = []uint32{6, 5}
	for i := range 2 {
		if report.ReceivedPoints[elements.ID(i+1)] != expectedScores[i] {
			t.Fatalf("Player %d mid game score incorrect. Expected %d, got: %d", i+1, expectedScores[i], report.ReceivedPoints[elements.ID(i+1)])
		}
	}

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
