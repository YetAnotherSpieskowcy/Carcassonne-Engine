package engine

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/binarytiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

type testResponse struct {
	BaseResponse
}
type testRequest struct {
	GameID        int
	RequiresWrite bool
	executeFunc   func(req *testRequest, game *game.Game) Response
}

func (req *testRequest) gameID() int {
	return req.GameID
}

func (req *testRequest) requiresWrite() bool {
	return req.RequiresWrite
}

func (req *testRequest) execute(game *game.Game) Response {
	if req.executeFunc != nil {
		return req.executeFunc(req, game)
	}
	return &testResponse{BaseResponse{gameID: req.gameID()}}
}

func TestFullGame(t *testing.T) {
	engine, err := StartGameEngine(4, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}
	tileSet := tilesets.StandardTileSet()

	gameWithID, err := engine.GenerateGame(tileSet)
	if err != nil {
		t.Fatal(err.Error())
	}
	game, gameID := gameWithID.Game, gameWithID.ID

	t.Logf("before loop: %s\n", time.Now())
	for i := range len(tileSet.Tiles) {
		t.Logf(
			"iteration %v start: %v\n", i, binarytiles.FromTile(game.CurrentTile),
		)
		legalMovesReq := &GetLegalMovesRequest{
			BaseGameID: gameID, TileToPlace: game.CurrentTile,
		}
		legalMovesResp := engine.SendGetLegalMovesBatch(
			[]*GetLegalMovesRequest{legalMovesReq},
		)[0]
		if legalMovesResp.Err() != nil {
			t.Fatal(legalMovesResp.Err().Error())
		}
		t.Logf("iteration %v got moves\n", i)

		move := legalMovesResp.Moves[0].Move
		t.Logf(
			"iteration %v selecting move: %v at position %v\n",
			i,
			binarytiles.FromPlacedTile(move),
			move.Position,
		)
		playTurnReq := &PlayTurnRequest{GameID: gameID, Move: move}
		playTurnResp := engine.SendPlayTurnBatch([]*PlayTurnRequest{playTurnReq})[0]
		if playTurnResp.Err() != nil {
			t.Fatal(playTurnResp.Err().Error())
		}
		t.Logf("iteration %v played turn\n", i)

		game = playTurnResp.Game
		gameID = playTurnResp.GameID()
		t.Logf("iteration %v end: %s\n", i, time.Now())

		if len(game.CurrentTile.Features) == 0 {
			// number of tiles in the tile set and number of tiles that you actually
			// get to place can differ, if a tile that's next in the stack happens to
			// not have any position to place available
			break
		}
	}

	if len(game.CurrentTile.Features) != 0 {
		t.Fatalf("expected current tile to be nil, got %#v instead", game.CurrentTile)
	}
}

func TestConcurrentReadRequests(t *testing.T) {
	engine, err := StartGameEngine(4, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}
	tileSet := tilesets.StandardTileSet()

	gameWithID, err := engine.GenerateGame(tileSet)
	if err != nil {
		t.Fatal(err.Error())
	}
	game, gameID := gameWithID.Game, gameWithID.ID

	legalMovesReq := &GetLegalMovesRequest{
		BaseGameID: gameID, TileToPlace: game.CurrentTile,
	}
	legalMovesResp := engine.SendGetLegalMovesBatch(
		[]*GetLegalMovesRequest{legalMovesReq},
	)[0]
	if legalMovesResp.Err() != nil {
		t.Fatal(legalMovesResp.Err().Error())
	}

	requests := make([]*GetRemainingTilesRequest, len(legalMovesResp.Moves))
	for i, moveWithState := range legalMovesResp.Moves {
		requests[i] = &GetRemainingTilesRequest{
			BaseGameID: gameID, StateToCheck: moveWithState.State,
		}
	}
	responses := engine.SendGetRemainingTilesBatch(requests)

	for _, resp := range responses {
		err := resp.Err()
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestGameEngineDoubleCloseDoesNotPanic(t *testing.T) {
	engine, err := StartGameEngine(4, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}
	engine.Close()
	engine.Close()
}

func TestGameEngineCloneGameReturnsIndependentGames(t *testing.T) {
	engine, err := StartGameEngine(1, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}

	buf := bytes.Buffer{}
	engine.appLogger.SetOutput(&buf)

	g, err := engine.GenerateGame(tilesets.StandardTileSet())
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = engine.CloneGame(g.ID, 3)
	if err != nil {
		t.Fatal(err.Error())
	}

	logs := buf.String()
	if len(logs) > 0 {
		t.Fatalf(
			"expected logs to be empty but they weren't. full logs below:\n%v",
			logs,
		)
	}

	engine.Close()
}

func TestGameEngineSubCloneGameReturnsGamesMarkedAsChildren(t *testing.T) {
	engine, err := StartGameEngine(1, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}

	buf := bytes.Buffer{}
	engine.appLogger.SetOutput(&buf)

	g, err := engine.GenerateGame(tilesets.StandardTileSet())
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = engine.SubCloneGame(g.ID, 15)
	if err != nil {
		t.Fatal(err.Error())
	}

	req := &PlayTurnRequest{GameID: g.ID, Move: g.Game.ValidTilePlacements[0]}
	engine.SendPlayTurnBatch([]*PlayTurnRequest{req})

	logs := buf.String()
	expected := fmt.Sprintf(childrenCleanupWarnMsg, g.ID)
	if !strings.Contains(logs, expected) {
		t.Fatalf(
			"expected logs to contain %#v but they did not. full logs below:\n%v",
			expected,
			logs,
		)
	}

	engine.Close()
}

func TestGameEngineDeleteGamesWarnsAboutRemovedChildren(t *testing.T) {
	engine, err := StartGameEngine(1, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}

	buf := bytes.Buffer{}
	engine.appLogger.SetOutput(&buf)

	g, err := engine.GenerateGame(tilesets.StandardTileSet())
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = engine.SubCloneGame(g.ID, 15)
	if err != nil {
		t.Fatal(err.Error())
	}

	engine.DeleteGames([]int{g.ID})

	logs := buf.String()
	expected := fmt.Sprintf(childrenCleanupWarnMsg, g.ID)
	if !strings.Contains(logs, expected) {
		t.Fatalf(
			"expected logs to contain %#v but they did not. full logs below:\n%v",
			expected,
			logs,
		)
	}

	engine.Close()
}

func TestGameEngineSendPlayTurnBatchDoesNotWarnAboutRemovedChildren(t *testing.T) {
	engine, err := StartGameEngine(1, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}

	buf := bytes.Buffer{}
	engine.appLogger.SetOutput(&buf)

	g, err := engine.GenerateGame(tilesets.StandardTileSet())
	if err != nil {
		t.Fatal(err.Error())
	}

	ids, err := engine.SubCloneGame(g.ID, 15)
	if err != nil {
		t.Fatal(err.Error())
	}

	engine.DeleteGames(ids)

	req := &PlayTurnRequest{GameID: g.ID, Move: g.Game.ValidTilePlacements[0]}
	engine.SendPlayTurnBatch([]*PlayTurnRequest{req})

	logs := buf.String()
	if len(logs) > 0 {
		t.Fatalf(
			"expected logs to be empty but they weren't. full logs below:\n%v",
			logs,
		)
	}

	engine.Close()
}

func TestGameEngineSendPlayTurnBatchRemovesFinishedGames(t *testing.T) {
	engine, err := StartGameEngine(1, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}

	buf := bytes.Buffer{}
	engine.appLogger.SetOutput(&buf)

	g, err := engine.GenerateGame(
		tilesets.TileSet{
			StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
			Tiles:        []tiles.Tile{},
		},
	)
	if err != nil {
		t.Fatal(err.Error())
	}

	ids, err := engine.SubCloneGame(g.ID, 3)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, gameID := range []int{ids[0], g.ID} {
		if _, exists := engine.games[gameID]; !exists {
			t.Fatal("expected game to exist before final round")
		}

		req := &PlayTurnRequest{GameID: gameID}
		resp := engine.SendPlayTurnBatch([]*PlayTurnRequest{req})[0]
		if !errors.Is(resp.err, stack.ErrStackOutOfBounds) {
			t.Fatal(err.Error())
		}

		if _, exists := engine.games[gameID]; exists {
			t.Fatal("expected game to not exist after final round")
		}
	}

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
