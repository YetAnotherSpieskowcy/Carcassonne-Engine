package engine

import (
	"errors"
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

// --- early failure (before worker receives the request) tests ---

func TestGameEngineSendPlayTurnBatchReturnsFailureWhenCommunicatorClosed(t *testing.T) {
	engine, err := StartGameEngine(1, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}
	engine.Close()

	requests := []*PlayTurnRequest{{GameID: 123}}
	resp := engine.SendPlayTurnBatch(requests)[0]
	if resp.Err() == nil {
		t.Fatal("expected error to occur")
	}
	if !errors.Is(resp.Err(), ErrCommunicatorClosed) {
		t.Fatal(resp.Err().Error())
	}
}

func TestGameEngineSendGetRemainingTilesBatchReturnsFailureWhenCommunicatorClosed(t *testing.T) {
	engine, err := StartGameEngine(1, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}
	engine.Close()

	requests := []*GetRemainingTilesRequest{{BaseGameID: 123}}
	resp := engine.SendGetRemainingTilesBatch(requests)[0]
	if resp.Err() == nil {
		t.Fatal("expected error to occur")
	}
	if !errors.Is(resp.Err(), ErrCommunicatorClosed) {
		t.Fatal(resp.Err().Error())
	}
}

func TestGameEngineSendGetLegalMovesBatchReturnsFailureWhenCommunicatorClosed(t *testing.T) {
	engine, err := StartGameEngine(1, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}
	engine.Close()

	requests := []*GetLegalMovesRequest{{BaseGameID: 123}}
	resp := engine.SendGetLegalMovesBatch(requests)[0]
	if resp.Err() == nil {
		t.Fatal("expected error to occur")
	}
	if !errors.Is(resp.Err(), ErrCommunicatorClosed) {
		t.Fatal(resp.Err().Error())
	}
}

// --- logic tests ---

func TestGameEngineSendPlayTurnBatchReceivesCorrectResponsesAfterWorkerRequests(t *testing.T) {
	engine, err := StartGameEngine(4, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}

	requestCount := 100
	requests := make([]*PlayTurnRequest, 0, requestCount)
	games := make([]*game.Game, 0, requestCount)
	for range requestCount {
		g, err := engine.GenerateGame(tilesets.StandardTileSet())
		if err != nil {
			t.Fatal(err.Error())
		}
		games = append(games, engine.games[g.ID])

		req := &PlayTurnRequest{GameID: g.ID, Move: g.Game.ValidTilePlacements[0]}
		requests = append(requests, req)
	}

	responses := engine.SendPlayTurnBatch(requests)
	for i, resp := range responses {
		err := resp.Err()
		if err != nil {
			t.Fatal(err.Error())
		}
		expectedID := requests[i].gameID()
		actualID := resp.GameID()
		if actualID != expectedID {
			t.Fatalf("expected %v game ID, got %v instead", expectedID, actualID)
		}
		expectedGame := games[i].Serialized()
		actualGame := resp.Game
		if !reflect.DeepEqual(actualGame, expectedGame) {
			t.Fatalf("expected %#v serialized game, got %#v instead", expectedGame, actualGame)
		}
	}
	engine.Close()
}

func TestGameEngineSendGetRemainingTilesBatchReturnsRemainingTiles(t *testing.T) {
	t1 := tiletemplates.MonasteryWithSingleRoad()
	t2 := tiletemplates.RoadsTurn()
	allTiles := []tiles.Tile{t1, t2, t1}
	total := float32(len(allTiles))
	expectedCounts := map[float32]tiles.Tile{2.0 / total: t1, 1.0 / total: t2}
	tileSet := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        allTiles,
	}

	engine, err := StartGameEngine(1, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}

	g, err := engine.GenerateGame(tileSet)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := &GetRemainingTilesRequest{BaseGameID: g.ID}
	resp := engine.SendGetRemainingTilesBatch([]*GetRemainingTilesRequest{request})[0]
	if resp.Err() != nil {
		t.Fatal(resp.Err().Error())
	}

	for _, tileProbability := range resp.TileProbabilities {
		found := false
		for expected, tile := range expectedCounts {
			if tileProbability.Tile.Equals(tile) {
				if tileProbability.Probability != expected {
					t.Fatalf(
						"expected %v probability, got %v instead",
						expected,
						tileProbability.Probability,
					)
				}
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("could not find a tile matching the tile probability %#v", tileProbability)
		}
	}

	engine.Close()
}

func TestGameEngineSendGetLegalMovesBatchReturnsNoDuplicates(t *testing.T) {
	tile := tiletemplates.MonasteryWithoutRoads()
	tileSet := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        []tiles.Tile{tile},
	}

	engine, err := StartGameEngine(1, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}

	g, err := engine.GenerateGame(tileSet)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := &GetLegalMovesRequest{BaseGameID: g.ID, TileToPlace: tile}
	resp := engine.SendGetLegalMovesBatch([]*GetLegalMovesRequest{request})[0]
	if resp.Err() != nil {
		t.Fatal(resp.Err().Error())
	}

	// Monastery with no roads is symmetrical both horizontally and vertically
	if len(resp.Moves) != 1 {
		t.Fatalf("expected %v as number of available moves, got %v instead", 1, resp.Moves)
	}
	if resp.Moves[0].Move.Position.X() != 0 {
		t.Fatalf("expected %v X, got %v instead", 0, resp.Moves[0].Move.Position.X())
	}
	if resp.Moves[0].Move.Position.Y() != -1 {
		t.Fatalf("expected %v Y, got %v instead", -1, resp.Moves[0].Move.Position.Y())
	}
	if len(resp.Moves[0].Move.Features) != len(tile.Features) {
		t.Fatalf("expected %#v, got %#v instead", tile.Features, resp.Moves[0].Move.Features)
	}
	for i := range tile.Features {
		if !tile.Features[i].Equals(resp.Moves[0].Move.Features[i].Feature) {
			t.Fatalf("expected %#v to equal %#v", tile.Features, resp.Moves[0].Move.Features)
		}
	}

	engine.Close()
}

func TestGameEngineSendGetLegalMovesBatchReturnsAllLegalRotations(t *testing.T) {
	tile := tiletemplates.MonasteryWithSingleRoad()
	tileSet := tilesets.TileSet{
		// non-default starting tile - limits number of possible positions to one
		StartingTile: tiletemplates.ThreeCityEdgesConnected(),
		Tiles:        []tiles.Tile{tile},
	}

	engine, err := StartGameEngine(1, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}

	g, err := engine.GenerateGame(tileSet)
	if err != nil {
		t.Fatal(err.Error())
	}

	request := &GetLegalMovesRequest{BaseGameID: g.ID, TileToPlace: tile}
	resp := engine.SendGetLegalMovesBatch([]*GetLegalMovesRequest{request})[0]
	if resp.Err() != nil {
		t.Fatal(resp.Err().Error())
	}

	// Monastery with single road can only be placed at (0, -1)
	// but in 3 different orientations (only field connected with road is invalid)
	// For each orientation, there are 4 valid meeple placements:
	// - no meeple
	// - meeple on monastery
	// - meeple on field
	// - meeple on road
	// For now, meeple placement is not handled so last number is 1 instead of 4 below
	expectedMoveCount := 1 * 3 * 1
	if len(resp.Moves) != expectedMoveCount {
		t.Fatalf("expected %v as number of available moves, got %v instead", expectedMoveCount, resp.Moves)
	}

	for _, moveState := range resp.Moves {
		if moveState.Move.Position.X() != 0 {
			t.Fatalf("expected %v X, got %#v instead", 0, moveState)
		}
		if moveState.Move.Position.Y() != -1 {
			t.Fatalf("expected %v Y, got %#v instead", -1, moveState)
		}
	}

	expectedTiles := []tiles.Tile{
		// road on the bottom
		tile,
		// road on the left
		tile.Rotate(1),
		// road on the right
		tile.Rotate(3),
	}
	for i, moveState := range resp.Moves {
		expected := expectedTiles[i]
		if len(moveState.Move.Features) != len(expected.Features) {
			t.Fatalf("expected %#v, got %#v instead", expected, moveState)
		}
		for i := range expected.Features {
			if !expected.Features[i].Equals(moveState.Move.Features[i].Feature) {
				t.Fatalf("expected %#v to equal %#v", expected, moveState)
			}
		}
	}

	engine.Close()
}
