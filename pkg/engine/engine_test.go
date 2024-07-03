package engine

import (
	"errors"
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
	if req.execute != nil {
		return req.execute(req, game)
	}
	return &testResponse{baseResponse{gameID: req.GameID()}}
}

func TestGameEngineSendBatchReceivesCorrectResponsesAfterWorkerRequests(t *testing.T) {
	engine, err := StartGameEngine(4, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}

	requestCount := 100
	requests := make([]Request, 0, requestCount)
	for range requestCount {
		g, err := engine.GenerateGame(tilesets.StandardTileSet())
		if err != nil {
			t.Fatal(err.Error())
		}

		req := &testRequest{baseRequest: baseRequest{gameID: g.ID}}
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

func TestGameEngineSendBatchFailsWhenGameIDNotFound(t *testing.T) {
	engine, err := StartGameEngine(4, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}

	requests := make([]Request, 0, 2)
	g, err := engine.GenerateGame(tilesets.StandardTileSet())
	if err != nil {
		t.Fatal(err.Error())
	}

	successfulReq := &testRequest{baseRequest: baseRequest{gameID: g.ID}}
	requests = append(requests, successfulReq)

	wrongID := g.ID + 2
	failingReq := &testRequest{baseRequest: baseRequest{gameID: wrongID}}
	requests = append(requests, failingReq)
	responses := engine.SendBatch(requests)

	// successful req
	err = responses[0].Err()
	if err != nil {
		t.Fatal(err.Error())
	}
	expected := g.ID
	actual := responses[0].GameID()
	if expected != actual {
		t.Fatalf("expected %v game ID, got %v instead", expected, actual)
	}

	// failing req
	err = responses[1].Err()
	if err == nil {
		t.Fatal("expected error to occur")
	}
	if !errors.Is(err, ErrGameNotFound) {
		t.Fatal(err.Error())
	}
	expected = wrongID
	actual = responses[1].GameID()
	if expected != actual {
		t.Fatalf("expected %v game ID, got %v instead", expected, actual)
	}
}
