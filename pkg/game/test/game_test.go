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

	test.MakeTurn{
		Game:         game,
		T:            t,
		TilePosition: position.New(0, 1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Left, FeatureType: feature.Road},
	}.Run()

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

	// Treat illegal move as correct (create error)
	captureFail := test.CaptureFail{}
	test.MakeTurnValidCheck{
		Game:         game,
		T:            &captureFail,
		TilePosition: position.New(0, 1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road},
		CorrectMove:  true,
		TurnNumber:   1,
	}.Run()
	if !captureFail.FailCaught() {
		t.Fatalf("Did not catch fail")
	}

	// // Treat legal move as incorrect (create error)
	captureFail = test.CaptureFail{}
	test.MakeTurnValidCheck{
		Game:         game,
		T:            &captureFail,
		TilePosition: position.New(1, 0),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road},
		CorrectMove:  false,
		TurnNumber:   1,
	}.Run()
	if !captureFail.FailCaught() {
		t.Fatalf("Did not catch fail")
	}
}

func TestMakeTurnValidCheckCatchFail(t *testing.T) {
	// create game
	minitileSet := tilesets.OrderedMiniTileSet2()
	deckStack := stack.NewOrdered(minitileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: minitileSet.StartingTile}
	game, err := game.NewFromDeck(deck, nil, 4)
	if err != nil {
		t.Fatal(err.Error())
	}

	// do any wrong move, and catch it

	test.MakeTurnValidCheck{
		Game:         game,
		T:            t,
		TilePosition: position.New(0, 1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road},
		CorrectMove:  false,
		TurnNumber:   1,
	}.Run()

	// do any correct move
	test.MakeTurnValidCheck{
		Game:         game,
		T:            t,
		TilePosition: position.New(1, 0),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road},
		CorrectMove:  true,
		TurnNumber:   1,
	}.Run()

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

	test.CheckMeeplesAndScore{
		Game:          game,
		T:             t,
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

	captureFail := test.CaptureFail{}
	test.CheckMeeplesAndScore{
		Game:          game,
		T:             &captureFail,
		PlayerScores:  []uint32{2, 2}, // should be {0,0}, so error
		PlayerMeeples: []uint8{7, 7},
		TurnNumber:    1,
	}.Run()
	if !captureFail.FailCaught() {
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

	captureFail := test.CaptureFail{}
	test.CheckMeeplesAndScore{
		Game:          game,
		T:             &captureFail,
		PlayerScores:  []uint32{0, 0},
		PlayerMeeples: []uint8{6, 6}, // should be {7,7}, so error
		TurnNumber:    1,
	}.Run()
	if !captureFail.FailCaught() {
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
	test.MakeTurnValidCheck{
		Game:         game,
		T:            t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road},
		CorrectMove:  true,
		TurnNumber:   1,
	}.Run()

	// verify that meeple truly exists
	test.VerifyMeepleExistence{
		T:           t,
		Game:        game,
		Pos:         pos,
		S:           side.Bottom,
		FeatureType: feature.Road,
		MeepleExist: true,
		TurnNumber:  1,
	}.Run()

	// verify that meeple does not exists
	test.VerifyMeepleExistence{
		T:           t,
		Game:        game,
		Pos:         pos,
		S:           side.BottomLeftEdge,
		FeatureType: feature.Field,
		MeepleExist: false,
		TurnNumber:  1,
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
	test.MakeTurnValidCheck{
		Game:         game,
		T:            t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road},
		CorrectMove:  true,
		TurnNumber:   1,
	}.Run()

	// Wrongly assume, that there is no meeeple (there is meeple)
	captureFail := test.CaptureFail{}
	test.VerifyMeepleExistence{
		T:           &captureFail,
		Game:        game,
		Pos:         pos,
		S:           side.Bottom,
		FeatureType: feature.Road,
		MeepleExist: false,
		TurnNumber:  1,
	}.Run()
	if !captureFail.FailCaught() {
		t.Fatalf("Did not catch fail")
	}

	// Wrongly assume that there is meeple
	captureFail = test.CaptureFail{}
	test.VerifyMeepleExistence{
		T:           &captureFail,
		Game:        game,
		Pos:         pos,
		S:           side.BottomLeftEdge,
		FeatureType: feature.Field,
		MeepleExist: true,
		TurnNumber:  1,
	}.Run()
	if !captureFail.FailCaught() {
		t.Fatalf("Did not catch fail")
	}
}
