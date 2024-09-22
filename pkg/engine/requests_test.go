package engine

import (
	"errors"
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
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

func TestGameEngineSendGetRemainingTilesBatchReturnsFailureWhenInvalidGameStateIsPassed(t *testing.T) {
	engine, err := StartGameEngine(1, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}

	game, err := engine.GenerateGame(tilesets.StandardTileSet())
	if err != nil {
		t.Fatal(err)
	}
	ptile := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	ptile.Position = position.New(0, 1)

	requests := []*GetRemainingTilesRequest{
		{
			BaseGameID: game.ID,
			// constructing these manually is not an expected use of the API
			// but we want to simulate someone somehow passing an invalid game state
			StateToCheck: &GameState{simulatedMoves: []elements.PlacedTile{ptile}},
		},
	}
	resp := engine.SendGetRemainingTilesBatch(requests)[0]
	if resp.Err() == nil {
		t.Fatal("expected error to occur")
	}
	if !errors.Is(resp.Err(), elements.ErrInvalidPosition) {
		t.Fatal(resp.Err())
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

func TestGameEngineSendGetLegalMovesBatchReturnsFailureWhenInvalidGameStateIsPassed(t *testing.T) {
	engine, err := StartGameEngine(1, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}

	game, err := engine.GenerateGame(tilesets.StandardTileSet())
	if err != nil {
		t.Fatal(err)
	}
	ptile := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	ptile.Position = position.New(0, 1)

	requests := []*GetLegalMovesRequest{
		{
			BaseGameID: game.ID,
			// constructing these manually is not an expected use of the API
			// but we want to simulate someone somehow passing an invalid game state
			StateToCheck: &GameState{simulatedMoves: []elements.PlacedTile{ptile}},
		},
	}
	resp := engine.SendGetLegalMovesBatch(requests)[0]
	if resp.Err() == nil {
		t.Fatal("expected error to occur")
	}
	if !errors.Is(resp.Err(), elements.ErrInvalidPosition) {
		t.Fatal(resp.Err())
	}
}

func TestGameEngineSendGetMidGameScoreBatchReturnsFailureWhenCommunicatorClosed(t *testing.T) {
	engine, err := StartGameEngine(1, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}
	engine.Close()

	requests := []*GetMidGameScoreRequest{{BaseGameID: 123}}
	resp := engine.SendGetMidGameScoreBatch(requests)[0]
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
	// and there's just one valid position (below starting tile, i.e. (0, -1))
	// so there's three legal moves: without meeple and with meeple on each of
	// the 2 features

	// without meeple
	ptile1 := elements.ToPlacedTile(tile)
	ptile1.Position = position.New(0, -1)
	// with meeple on the field
	ptile2 := elements.ToPlacedTile(tile)
	ptile2.Position = position.New(0, -1)
	ptile2.Features[0].Meeple = elements.Meeple{
		Type:     elements.NormalMeeple,
		PlayerID: elements.ID(1),
	}
	// with meeple on the monastery
	ptile3 := elements.ToPlacedTile(tile)
	ptile3.Position = position.New(0, -1)
	ptile3.Features[1].Meeple = elements.Meeple{
		Type:     elements.NormalMeeple,
		PlayerID: elements.ID(1),
	}
	expectedMoves := []elements.PlacedTile{ptile1, ptile2, ptile3}

	actualMoves := make([]elements.PlacedTile, len(resp.Moves))
	for i, moveWithState := range resp.Moves {
		actualMoves[i] = moveWithState.Move
	}
	if !reflect.DeepEqual(expectedMoves, actualMoves) {
		t.Fatalf("expected following moves: %#v, got %#v instead", expectedMoves, actualMoves)
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
	expectedMoveCount := 1 * 3 * 4
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
	moveIndex := 0
	for _, expected := range expectedTiles {
		for range 4 {
			moveState := resp.Moves[moveIndex]
			for i := range expected.Features {
				if len(moveState.Move.Features) != len(expected.Features) {
					t.Fatalf("expected %#v, got %#v instead", moveState, expected)
				}
				if !expected.Features[i].Equals(moveState.Move.Features[i].Feature) {
					t.Fatalf("expected %#v to equal %#v", moveState, expected)
				}
			}
			moveIndex++
		}
	}

	engine.Close()
}

func TestGameEngineSendGetMidGameScoreBatchAtGameStartReturnsZeroScores(t *testing.T) {

	engine, err := StartGameEngine(4, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}
	tileSet := tilesets.StandardTileSet()

	gameWithID, err := engine.GenerateGame(tileSet)
	if err != nil {
		t.Fatal(err.Error())
	}
	_, gameID := gameWithID.Game, gameWithID.ID

	midGameScoreReq := &GetMidGameScoreRequest{
		BaseGameID: gameID,
	}

	midGameScoreResp := engine.SendGetMidGameScoreBatch(
		[]*GetMidGameScoreRequest{midGameScoreReq},
	)[0]

	if midGameScoreResp.Err() != nil {
		t.Fatal(midGameScoreResp.Err().Error())
	}

	if midGameScoreResp.Scores[0] != 0 {
		t.Fatal("Player 0 score not zero")
	}

	if midGameScoreResp.Scores[1] != 0 {
		t.Fatal("Player 0 score not zero")
	}
}

func TestGameEngineSendGetMidGameScoreBatchAtMidGameReturnsExpectedScores(t *testing.T) {
	// --------- setup  game -------------
	engine, err := StartGameEngine(4, t.TempDir())
	if err != nil {
		t.Fatal(err.Error())
	}
	tileSet := tilesets.TileSet{
		Tiles: []tiles.Tile{
			tiletemplates.FourCityEdgesConnectedShield(),
			tiletemplates.StraightRoads(),
			tiletemplates.StraightRoads(),
		},
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
	}

	gameWithID, err := engine.GenerateOrderedGame(tileSet)
	if err != nil {
		t.Fatal(err.Error())
	}
	gameState, gameID := gameWithID.Game, gameWithID.ID

	// --------- play two turns -------------
	ptile := elements.ToPlacedTile(gameState.CurrentTile)
	ptile.Position = position.New(0, 1)
	ptile.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple = elements.Meeple{PlayerID: elements.ID(1), Type: elements.NormalMeeple}
	playTurn1 := []*PlayTurnRequest{
		{
			GameID: gameID,
			Move:   ptile,
		},
	}
	resp := engine.SendPlayTurnBatch(playTurn1)[0]
	gameState, gameID = resp.Game, resp.GameID()

	ptile = elements.ToPlacedTile(gameState.CurrentTile)
	ptile.Position = position.New(1, 0)
	ptile.GetPlacedFeatureAtSide(side.Right, feature.Road).Meeple = elements.Meeple{PlayerID: elements.ID(2), Type: elements.NormalMeeple}
	playTurn2 := []*PlayTurnRequest{
		{
			GameID: gameID,
			Move:   ptile,
		},
	}
	resp = engine.SendPlayTurnBatch(playTurn2)[0]
	gameState, gameID = resp.Game, resp.GameID()

	// --------- check scores -------------
	midGameScoreReq := &GetMidGameScoreRequest{
		BaseGameID: gameID,
	}

	midGameScoreResp := engine.SendGetMidGameScoreBatch(
		[]*GetMidGameScoreRequest{midGameScoreReq},
	)[0]

	if midGameScoreResp.Err() != nil {
		t.Fatal(midGameScoreResp.Err().Error())
	}

	expected := uint32(3)
	if midGameScoreResp.Scores[1] != expected {
		t.Fatalf("Player 1 score: %#v, expected: %#v", midGameScoreResp.Scores[1], expected)
	}

	expected = uint32(2)
	if midGameScoreResp.Scores[2] != expected {
		t.Fatalf("Player 2 score: %#v, expected: %#v", midGameScoreResp.Scores[2], expected)
	}
}
