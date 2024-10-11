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

type CaptureFail struct {
	failureCaught bool
}

func (t *CaptureFail) Fatal(_ ...any) {
	t.failureCaught = true
}

func (t *CaptureFail) Fatalf(_ string, _ ...any) {
	t.failureCaught = true
}

func TestMakeTurnLegalMove(t *testing.T) {
	minitileSet := tilesets.StandardTileSet()
	deckStack := stack.NewOrdered(minitileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: minitileSet.StartingTile}
	game, err := game.NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(0, -1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.NoSide, FeatureType: feature.Monastery},
	}.Run()

	board := game.GetBoard()
	ptile, ok := board.GetTileAt(position.New(0, -1))
	if !ok {
		t.Fatal("expected to find a tile at (0, -1)")
	}

	tile := elements.ToTile(ptile)
	expected := minitileSet.Tiles[0]
	if !tile.Equals(expected) {
		t.Fatalf("expected %#v, got %#v instead", expected, tile)
	}

	feat := ptile.GetPlacedFeatureAtSide(side.NoSide, feature.Monastery)
	if feat.Meeple.Type != elements.NormalMeeple {
		t.Fatalf("expected normal meeple on road tile feature, got %#v instead", feat.Meeple.Type)
	}
}

func TestMakeTurnIllegalMove(t *testing.T) {
	minitileSet := tilesets.StandardTileSet()
	deckStack := stack.NewOrdered(minitileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: minitileSet.StartingTile}
	game, err := game.NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	captureFail := CaptureFail{}
	test.MakeTurn{
		Game:         game,
		TestingT:     &captureFail,
		Position:     position.New(0, 1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.NoSide, FeatureType: feature.Monastery},
		TurnNumber:   1,
		WrongTurn:    true,
	}.Run()
	if captureFail.failureCaught {
		t.Fatalf("Did not catch fail")
	}

}

func TestMakeWrongTurnWithLegalTurn(t *testing.T) {
	// create game
	minitileSet := tilesets.StandardTileSet()
	deckStack := stack.NewOrdered(minitileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: minitileSet.StartingTile}
	game, err := game.NewFromDeck(deck, nil, 4)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Treat correct move as illegal (create error)
	captureFail := CaptureFail{}
	test.MakeTurn{
		Game:         game,
		TestingT:     &captureFail,
		Position:     position.New(0, -1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.NoSide, FeatureType: feature.Monastery},
		TurnNumber:   1,
		WrongTurn:    true,
	}.Run()
	if !captureFail.failureCaught {
		t.Fatalf("Did not catch fail")
	}
}

func TestMakeTurnWithLegalTurnWichIsActuallyIncorect(t *testing.T) {
	// create game
	minitileSet := tilesets.StandardTileSet()
	deckStack := stack.NewOrdered(minitileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: minitileSet.StartingTile}
	game, err := game.NewFromDeck(deck, nil, 4)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Treat invcorrect move as legal (create error)
	captureFail := CaptureFail{}
	test.MakeTurn{
		Game:         game,
		TestingT:     &captureFail,
		Position:     position.New(1, 0),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.NoSide, FeatureType: feature.Monastery},
		TurnNumber:   1,
	}.Run()
	if !captureFail.failureCaught {
		t.Fatalf("Did not catch fail")
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

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{0, 0},
		PlayerMeeples: []uint8{7, 7},
		TurnNumber:    1,
	}.Run()
}

func TestCheckMeeplesAndScoreCatchFailScore(t *testing.T) {
	// create game
	minitileSet := tilesets.OrderedMiniTileSet2()
	deckStack := stack.NewOrdered(minitileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: minitileSet.StartingTile}
	game, err := game.NewFromDeck(deck, nil, 4)
	if err != nil {
		t.Fatal(err.Error())
	}

	captureFail := CaptureFail{}
	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      &captureFail,
		PlayerScores:  []uint32{2, 2}, // should be {0,0}, so error
		PlayerMeeples: []uint8{7, 7},
		TurnNumber:    1,
	}.Run()
	if !captureFail.failureCaught {
		t.Fatalf("Did not catch fail")
	}
}

func TestCheckMeeplesAndScoreCatchFailMeeples(t *testing.T) {
	// create game
	minitileSet := tilesets.OrderedMiniTileSet2()
	deckStack := stack.NewOrdered(minitileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: minitileSet.StartingTile}
	game, err := game.NewFromDeck(deck, nil, 4)
	if err != nil {
		t.Fatal(err.Error())
	}

	captureFail := CaptureFail{}
	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      &captureFail,
		PlayerScores:  []uint32{0, 0},
		PlayerMeeples: []uint8{6, 6}, // should be {7,7}, so error
		TurnNumber:    1,
	}.Run()
	if !captureFail.failureCaught {
		t.Fatalf("Did not catch fail")
	}
}

func TestVerifyMeepleExistenceCorrectCheck(t *testing.T) {
	// create game
	minitileSet := tilesets.OrderedMiniTileSet2()
	deckStack := stack.NewOrdered(minitileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: minitileSet.StartingTile}
	game, err := game.NewFromDeck(deck, nil, 4)
	if err != nil {
		t.Fatal(err.Error())
	}

	pos := position.New(1, 0)
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road},
		TurnNumber:   1,
	}.Run()

	// verify that meeple truly exists
	test.VerifyMeepleExistence{
		TestingT:     t,
		Game:         game,
		Position:     pos,
		Side:         side.Bottom,
		FeatureType:  feature.Road,
		MeepleExists: true,
		TurnNumber:   1,
	}.Run()

	// verify that meeple does not exist
	test.VerifyMeepleExistence{
		TestingT:     t,
		Game:         game,
		Position:     pos,
		Side:         side.BottomLeftEdge,
		FeatureType:  feature.Field,
		MeepleExists: false,
		TurnNumber:   1,
	}.Run()
}

func TestVerifyMeepleExistenceFailCapture(t *testing.T) {

	// create game
	minitileSet := tilesets.OrderedMiniTileSet2()
	deckStack := stack.NewOrdered(minitileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: minitileSet.StartingTile}
	game, err := game.NewFromDeck(deck, nil, 4)
	if err != nil {
		t.Fatal(err.Error())
	}

	pos := position.New(1, 0)
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road},
		TurnNumber:   1,
	}.Run()

	// Wrongly assume, that there is no meeeple (there is meeple)
	captureFail := CaptureFail{}
	test.VerifyMeepleExistence{
		TestingT:     &captureFail,
		Game:         game,
		Position:     pos,
		Side:         side.Bottom,
		FeatureType:  feature.Road,
		MeepleExists: false,
		TurnNumber:   1,
	}.Run()
	if !captureFail.failureCaught {
		t.Fatalf("Did not catch fail")
	}

	// Wrongly assume that there is meeple
	captureFail = CaptureFail{}
	test.VerifyMeepleExistence{
		TestingT:     &captureFail,
		Game:         game,
		Position:     pos,
		Side:         side.BottomLeftEdge,
		FeatureType:  feature.Field,
		MeepleExists: true,
		TurnNumber:   1,
	}.Run()
	if !captureFail.failureCaught {
		t.Fatalf("Did not catch fail")
	}
}
