package engine

import (
	"errors"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

type testResponse struct {
	BaseResponse
}
type testRequest struct {
	GameID      int
	executeFunc func(req *testRequest, game *game.Game) Response
}

func (req *testRequest) gameID() int {
	return req.GameID
}

func (req *testRequest) execute(game *game.Game) Response {
	if req.executeFunc != nil {
		return req.executeFunc(req, game)
	}
	return &testResponse{BaseResponse{gameID: req.gameID()}}
}

func TestGameEngineDoubleCloseDoesNotPanic(t *testing.T) {
	engine, err := StartGameEngine(4, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}
	engine.Close()
	engine.Close()
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

		req := &testRequest{GameID: g.ID}
		requests = append(requests, req)
	}

	responses := engine.sendBatch(requests)
	for i, resp := range responses {
		err := resp.Err()
		if err != nil {
			t.Fatal(err.Error())
		}
		expected := requests[i].gameID()
		actual := resp.GameID()
		if actual != expected {
			t.Fatalf("expected %v game ID, got %v instead", expected, actual)
		}
	}
	engine.Close()
}

func TestGameEngineSendBatchReturnsFailureWhenGameIDNotFound(t *testing.T) {
	engine, err := StartGameEngine(4, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}

	requests := make([]Request, 0, 2)
	g, err := engine.GenerateGame(tilesets.StandardTileSet())
	if err != nil {
		t.Fatal(err.Error())
	}

	successfulReq := &testRequest{GameID: g.ID}
	requests = append(requests, successfulReq)

	wrongID := g.ID + 2
	failingReq := &testRequest{GameID: wrongID}
	requests = append(requests, failingReq)
	responses := engine.sendBatch(requests)

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
	engine.Close()
}

func TestGameEngineSendBatchReturnsFailuresWhenCommunicatorClosed(t *testing.T) {
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

		req := &testRequest{GameID: g.ID}
		requests = append(requests, req)
	}
	engine.Close()

	responses := engine.sendBatch(requests)
	for i, resp := range responses {
		err := resp.Err()
		if err == nil {
			t.Fatal("expected error to occur")
		}
		if !errors.Is(err, ErrCommunicatorClosed) {
			t.Fatal(err.Error())
		}
		expected := requests[i].gameID()
		actual := resp.GameID()
		if actual != expected {
			t.Fatalf("expected %v game ID, got %v instead", expected, actual)
		}
	}
}
