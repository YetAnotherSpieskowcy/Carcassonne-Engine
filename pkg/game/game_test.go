package game

import (
	"errors"
	"io"
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/logger"
	pb "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/proto"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

type TestLogger struct {
	callCount int
}

func (l *TestLogger) LogEvent(_ pb.Entry) error {
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
	original, err := NewFromTileSet(tileSet, originalLogger, 2)
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
		tiletemplates.SingleCityEdgeNoRoads().Rotate(2),
		tiletemplates.StraightRoads(),
	}
	deckStack := stack.NewOrdered(tileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileSet.StartingTile}
	game, err := NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	// correct move with tile 0
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile := elements.ToPlacedTile(tile)
	ptile.Position = position.New(0, 1)
	err = game.PlayTurn(ptile)
	if err != nil {
		t.Fatal(err.Error())
	}

	// incorrect move - try placing tile 0 when 1 should be placed
	tile = tileSet.Tiles[0]
	ptile = elements.ToPlacedTile(tile)
	ptile.Position = position.New(0, -1)
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
	ptile.Position = position.New(0, -1)
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
	game, err := NewFromTileSet(tilesets.StandardTileSet(), nil, 2)
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

func TestGameSerializedCurrentTileNotSetWhenStackOutOfBounds(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	tileSet.Tiles = []tiles.Tile{}

	game, err := NewFromTileSet(tileSet, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	serialized := game.Serialized()
	if len(serialized.CurrentTile.Features) != 0 {
		t.Fatalf("expected no current tile, got %v instead", serialized.CurrentTile)
	}
}

func TestGameSerializedCurrentTileNotSetForClonesWithSwappableTiles(t *testing.T) {
	tileSet := tilesets.StandardTileSet()

	game, err := NewFromTileSet(tileSet, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	serialized := game.Serialized()
	if len(serialized.CurrentTile.Features) == 0 {
		t.Fatal("got empty tile when expected valid one to be provided")
	}

	clone := game.DeepCloneWithSwappableTiles()
	serialized = clone.Serialized()
	if len(serialized.CurrentTile.Features) != 0 {
		t.Fatalf("expected no current tile, got %v instead", serialized.CurrentTile)
	}
}

func TestGameGetLegalMovesForIncludesMeepleTypesCurrentPlayerDoesHave(t *testing.T) {
	tile := tiletemplates.MonasteryWithSingleRoad()
	tileSet := tilesets.TileSet{
		// non-default starting tile - limits number of possible positions to one
		StartingTile: tiletemplates.ThreeCityEdgesConnected(),
		Tiles:        []tiles.Tile{tile},
	}

	game, err := NewFromTileSet(tileSet, nil, 2)
	if err != nil {
		t.Fatal(err)
	}

	basePlacement := game.GetTilePlacementsFor(tile)[0]
	expected := []elements.PlacedTile{basePlacement}
	for i := range basePlacement.Features {
		ptile := basePlacement.DeepClone()
		ptile.Features[i].Meeple = elements.Meeple{
			Type: elements.NormalMeeple, PlayerID: 1,
		}
		expected = append(expected, ptile)
	}
	actual := game.GetLegalMovesFor(basePlacement)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestGameGetLegalMovesForExcludesMeepleTypesCurrentPlayerDoesNotHave(t *testing.T) {
	tile := tiletemplates.MonasteryWithSingleRoad()
	tileSet := tilesets.TileSet{
		// non-default starting tile - limits number of possible positions to one
		StartingTile: tiletemplates.ThreeCityEdgesConnected(),
		Tiles:        []tiles.Tile{tile},
	}

	game, err := NewFromTileSet(tileSet, nil, 2)
	if err != nil {
		t.Fatal(err)
	}

	ptile := game.GetTilePlacementsFor(tile)[0]
	player := game.CurrentPlayer()
	player.SetMeepleCount(elements.NormalMeeple, 0)

	expected := []elements.PlacedTile{ptile}
	actual := game.GetLegalMovesFor(ptile)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestGameSwapCurrentTileReturnsErrorOnOriginalGame(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	tileSet.Tiles = []tiles.Tile{
		tiletemplates.SingleCityEdgeNoRoads(),
		tiletemplates.StraightRoads(),
	}
	deckStack := stack.NewOrdered(tileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileSet.StartingTile}

	game, err := NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = game.SwapCurrentTile(tileSet.Tiles[1])
	if err == nil {
		t.Fatal("expected error to occur")
	}
	if !errors.Is(err, ErrCannotSwapTiles) {
		t.Fatal(err)
	}

	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err)
	}

	if !tile.ExactEquals(tileSet.Tiles[0]) {
		t.Fatalf("expected %#v, got %#v instead", tileSet.Tiles[0], tile)
	}
}

func TestGameSwapCurrentTileSwapTileOnCloneWithSwappableTiles(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	tileSet.Tiles = []tiles.Tile{
		tiletemplates.SingleCityEdgeNoRoads(),
		tiletemplates.StraightRoads(),
	}
	deckStack := stack.NewOrdered(tileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileSet.StartingTile}

	game, err := NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}
	clone := game.DeepCloneWithSwappableTiles()

	err = clone.SwapCurrentTile(tileSet.Tiles[1])
	if err != nil {
		t.Fatal(err)
	}

	tile, err := clone.GetCurrentTile()
	if err != nil {
		t.Fatal(err)
	}

	if !tile.ExactEquals(tileSet.Tiles[1]) {
		t.Fatalf("expected %#v, got %#v instead", tileSet.Tiles[1], tile)
	}
}

func TestGamePlayTurnDoesNotMutateInput(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	tileSet.Tiles = []tiles.Tile{tiletemplates.SingleCityEdgeNoRoads().Rotate(2)}

	game, err := NewFromTileSet(tileSet, nil, 2)
	if err != nil {
		t.Fatal(err)
	}

	expected := elements.Meeple{
		Type:     elements.NormalMeeple,
		PlayerID: game.CurrentPlayer().ID(),
	}
	ptile := elements.ToPlacedTile(tileSet.Tiles[0])
	ptile.Position = position.New(0, 1)
	ptile.GetPlacedFeatureAtSide(side.Bottom, feature.City).Meeple = expected

	err = game.PlayTurn(ptile)
	if err != nil {
		t.Fatal(err)
	}

	actual := ptile.GetPlacedFeatureAtSide(side.Bottom, feature.City).Meeple
	if expected != actual {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestGameGetPlayerById(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	game, err := NewFromTileSet(tileSet, nil, 2)
	if err != nil {
		t.Fatal(err)
	}

	player1 := game.GetPlayerByID(elements.ID(1))
	player2 := game.GetPlayerByID(elements.ID(2))

	// ID
	if player1.ID() != elements.ID(1) {
		t.Fatalf("Player1 id not valid expected 1")
	}

	if player2.ID() != elements.ID(2) {
		t.Fatalf("Player1 id not valid expected 2")
	}

	// score
	if player1.Score() != 0 {
		t.Fatalf("Player1 score not 0")
	}

	if player2.Score() != 0 {
		t.Fatalf("Player2 score not 0")
	}
}

func TestGameGetPlayerByIdNotFound(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	game, err := NewFromTileSet(tileSet, nil, 2)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Found not existing player, panic didn't happen")
		}
	}()

	game.GetPlayerByID(elements.ID(10))

}

func TestGetBoard(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	deckStack := stack.NewOrdered(tileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileSet.StartingTile}

	game, err := NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	if game.GetBoard() == nil {
		t.Fatalf("Couldn't get board")
	}
}
