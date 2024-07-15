package engine

import (
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

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
	engine.Shutdown()
}
