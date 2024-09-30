package test_test

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/test"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

// If test functions contain logic, they need to be tested as well :)

func TestMakeTurn(t *testing.T) {
	minitileSet := tilesets.OrderedMiniTileSet1()
	deckStack := stack.NewOrdered(minitileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: minitileSet.StartingTile}
	game, err := game.NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	test.MakeTurn(game, t, position.New(0, 1), test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Left, FeatureType: feature.Road})

	board := game.GetBoard()
	ptile, ok := board.GetTileAt(position.New(0, 1))
	if !ok {
		t.Fatal("expected to find a tile at (0, 1)")
	}

	tile := elements.ToTile(ptile)
	expected := minitileSet.Tiles[0]
	if !tile.Equals(expected) {
		t.Fatalf("expected %#v, got %#v instead", expected, tile)
	}

	feat := ptile.GetPlacedFeatureAtSide(side.Left, feature.Road)
	if feat.Meeple.Type != elements.NormalMeeple {
		t.Fatalf("expected normal meeple on road tile feature, got %#v instead", feat.Meeple.Type)
	}
}

func TestMakeTurnValidCheck(t *testing.T) {
	// create game
	minitileSet := tilesets.OrderedMiniTileSet2()
	deckStack := stack.NewOrdered(minitileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: minitileSet.StartingTile}
	game, err := game.NewFromDeck(deck, nil, 4)
	if err != nil {
		t.Fatal(err.Error())
	}

	test.MakeTurnValidCheck(game, t, position.New(0, 1), test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road}, false, 1) // do any wrong move, and catch it
	test.MakeTurnValidCheck(game, t, position.New(1, 0), test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road}, true, 1)  // do any correct move

	// check if meeple was placed
	ptile, exist := game.GetBoard().GetTileAt(position.New(1, 0))
	if !exist {
		t.Fatalf("Tile doesn't exist!")
	}
	pfeature := ptile.GetPlacedFeatureAtSide(side.Bottom, feature.Road)
	if pfeature.Meeple.PlayerID != elements.ID(1) &&
		pfeature.Meeple.Type != elements.NormalMeeple {
		t.Fatalf("Wrong meeple params!")
	}

}

func TestCheckMeeplesAndScore(t *testing.T) {
	// create game
	minitileSet := tilesets.OrderedMiniTileSet2()
	deckStack := stack.NewOrdered(minitileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: minitileSet.StartingTile}
	game, err := game.NewFromDeck(deck, nil, 4)
	if err != nil {
		t.Fatal(err.Error())
	}

	test.CheckMeeplesAndScore(game, t, []uint32{0, 0}, []uint8{7, 7}, 1)
}

func TestVerifyMeepleExistence(t *testing.T) {
	// create game
	minitileSet := tilesets.OrderedMiniTileSet2()
	deckStack := stack.NewOrdered(minitileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: minitileSet.StartingTile}
	game, err := game.NewFromDeck(deck, nil, 4)
	if err != nil {
		t.Fatal(err.Error())
	}

	pos := position.New(1, 0)
	test.MakeTurnValidCheck(game, t, pos, test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road}, true, 1) // do any correct move

	test.VerifyMeepleExistence(t, game, pos, side.Bottom, feature.Road, true, 1)           // verify that meeple truly exists
	test.VerifyMeepleExistence(t, game, pos, side.BottomLeftEdge, feature.Field, false, 1) // verify that meeple truly exists
}
