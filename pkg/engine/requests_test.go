package engine

import (
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
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

func TestGameEngineSendGetRemainingTilesBatchReturnsRemainingTiles(t *testing.T) {
	t1 := tiletemplates.MonasteryWithSingleRoad()
	t2 := tiletemplates.RoadsTurn()
	allTiles := []tiles.Tile{t1, t2, t1}
	expectedCounts := map[uint16]tiles.Tile{2: t1, 1: t2}
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

	if int(resp.Denominator) != len(allTiles) {
		t.Fatalf("expected %v denominator, got %v instead", len(allTiles), resp.Denominator)
	}
	for _, moveNumerator := range resp.MoveNumerators {
		found := false
		for expected, tile := range expectedCounts {
			if moveNumerator.Tile.Equals(tile) {
				if moveNumerator.Numerator != expected {
					t.Fatalf(
						"expected %v numerator, got %v instead",
						expected,
						moveNumerator.Numerator,
					)
				}
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("could not find a tile matching the move numerator %#v", moveNumerator)
		}
	}

	engine.Shutdown()
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

	engine.Shutdown()
}

func TestGameEngineSendGetLegalMovesBatchReturnsAllLegalRotations(t *testing.T) {
	tile := tiletemplates.MonasteryWithSingleRoad()
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

	// Monastery with single road can only be placed at (0, -1)
	// but in 3 different orientations (only field connected with road is invalid)
	if len(resp.Moves) != 3 {
		t.Fatalf("expected %v as number of available moves, got %v instead", 3, resp.Moves)
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

	engine.Shutdown()
}
