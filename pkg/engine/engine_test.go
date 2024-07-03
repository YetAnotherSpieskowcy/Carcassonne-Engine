package engine

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

type testRequest struct {
	baseRequest
	execute func(req *testRequest, game *game.Game) Response
}

type testResponse struct {
	baseResponse
}

func (req *testRequest) Execute(game *game.Game) Response {
	return req.execute(req, game)
}

func TestGameEngineSendsRequestToWorker(t *testing.T) {
	engine, err := StartGameEngine(4, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}

	requestCount := 5
	requests := make([]Request, 0, requestCount)
	for range requestCount {
		g, err := engine.GenerateGame(tilesets.StandardTileSet())
		if err != nil {
			t.Fatal(err.Error())
		}

		req := &testRequest{
			baseRequest: baseRequest{gameID: g.ID},
			execute:     func(req *testRequest, game *game.Game) Response {
				return &testResponse{baseResponse{gameID: req.GameID()}}
			},
		}
		requests = append(requests, req)
	}
	responses := engine.SendBatch(requests)
	for i, resp := range responses {
		expected := requests[i].GameID()
		actual := resp.GameID()
		if actual != expected {
			t.Fatalf("expected %v game ID, got %v instead", expected, actual)
		}
	}
	engine.Shutdown()
}
