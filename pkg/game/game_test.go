package game

import (
	"errors"
	"io"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/logger"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

type TestLogger struct {
	callCount int
}

func (l *TestLogger) LogEvent(_ logger.EventType, _ interface{}) error {
	l.callCount++
	return nil
}

func (l *TestLogger) AsWriter() io.Writer {
	return io.Discard
}

func (l *TestLogger) CopyTo(_ logger.Logger) error {
	return logger.ErrCopyToNotImplemented
}

func TestDeepClone(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	tileSet.Tiles = []tiles.Tile{tiletemplates.SingleCityEdgeNoRoads().Rotate(2)}

	originalLogger := &TestLogger{}
	original, err := NewFromTileSet(tileSet, originalLogger)
	if err != nil {
		t.Fatal(err.Error())
	}
	originalLogger.callCount = 0

	clone := original.DeepClone()
	actualTile, err := clone.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}

	ptile := elements.ToPlacedTile(actualTile)
	ptile.Position = position.New(0, -1)
	err = clone.PlayTurn(ptile)
	if err != nil {
		t.Fatal(err.Error())
	}

	if originalLogger.callCount != 0 {
		t.Fatal("original game's logger was not expected to be called by the clone")
	}

	expectedTile, err := original.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	if !expectedTile.Equals(actualTile) {
		t.Fatalf(
			"expected clone's current tile (%#v) to be identical to original (%#v)",
			actualTile,
			expectedTile,
		)
	}

	if original.deck.Stack == clone.deck.Stack {
		// comparison by pointers
		t.Fatalf(
			"Original deck's stack (%#v) and clone deck's stack (%#v) are equal",
			original.deck,
			clone.deck,
		)
	}

	if original.board == clone.board {
		// comparison by pointers
		t.Fatalf(
			"Original board (%#v) and clone board (%#v) are equal",
			original.board,
			clone.board,
		)
	}

	for i, clonePlayer := range clone.players {
		// comparison by pointers
		if original.players[i] == clonePlayer {
			t.Fatalf(
				"Original player (%#v) and clone player (%#v) are equal",
				original.players[i],
				clonePlayer,
			)
		}
	}

	originalID := original.CurrentPlayer().ID()
	cloneID := clone.CurrentPlayer().ID()
	if originalID == cloneID {
		t.Fatalf(
			"original's current player ID (%v) == clone's current player ID *after* move (%v)",
			originalID,
			cloneID,
		)
	}
}

func TestFullGame(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	tileSet.Tiles = []tiles.Tile{
		tiletemplates.SingleCityEdgeNoRoads(),
		tiletemplates.StraightRoads(),
	}
	deckStack := stack.NewOrdered(tileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileSet.StartingTile}
	game, err := NewFromDeck(deck, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	// correct move with tile 0
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile := elements.ToPlacedTile(tile)
	ptile.Position = position.New(0, -1)
	err = game.PlayTurn(ptile)
	if err != nil {
		t.Fatal(err.Error())
	}

	// incorrect move - try placing tile 0 when 1 should be placed
	tile = tileSet.Tiles[0]
	ptile = elements.ToPlacedTile(tile)
	ptile.Position = position.New(0, 1)
	err = game.PlayTurn(ptile)
	if err == nil {
		t.Fatal("expected error to occur")
	}
	if !errors.Is(err, elements.ErrWrongTile) {
		t.Fatal(err.Error())
	}

	// correct move with tile 1
	tile, err = game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile = elements.ToPlacedTile(tile)
	ptile.Position = position.New(0, 1)
	err = game.PlayTurn(ptile)
	if err != nil {
		t.Fatal(err.Error())
	}

	ptile = elements.ToPlacedTile(tileSet.Tiles[1])
	ptile.Position = position.New(0, 0)
	// check if out of bounds state is detected
	err = game.PlayTurn(ptile)
	if err == nil {
		t.Fatal("expected error to occur")
	}
	if !errors.Is(err, stack.ErrStackOutOfBounds) {
		t.Fatal(err.Error())
	}

	actualScores, err := game.Finalize()
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedScores := elements.NewScoreReport()
	expectedScores.ReceivedPoints[elements.ID(1)] = 0
	expectedScores.ReceivedPoints[elements.ID(2)] = 0
	for playerID, actual := range actualScores.ReceivedPoints {
		expected := expectedScores.ReceivedPoints[playerID]
		if actual != expected {
			t.Fatalf("expected %v, got %v for player %v instead", expected, actual, playerID)
		}
	}
}

func TestGameFinalizeErrorsBeforeGameIsFinished(t *testing.T) {
	game, err := NewFromTileSet(tilesets.StandardTileSet(), nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	// try finalizing before the game is finished
	_, err = game.Finalize()
	if err == nil {
		t.Fatal("expected error to occur")
	}
	if !errors.Is(err, elements.ErrGameIsNotFinished) {
		t.Fatal(err.Error())
	}
}

func TestGameSerializedCurrentTileNilWhenStackOutOfBounds(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	tileSet.Tiles = []tiles.Tile{}

	game, err := NewFromTileSet(tileSet, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	serialized := game.Serialized()
	if serialized.CurrentTile != nil {
		t.Fatalf("expected nil, got %v instead", serialized.CurrentTile)
	}
}
