package engine

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func PlayTurnWithIndex(engine *GameEngine, gameID int, ptile elements.PlacedTile) error {
	// clone once
	clones, err := engine.cloneGame(gameID, 1, true)
	if err != nil {
		println("ERROR CLONE")
		println(err.Error())
		return err
	}
	clone := clones[0]

	// play turn
	playTurn := []*PlayTurnRequest{
		{
			GameID: clone,
			Move:   ptile,
		},
	}
	resp := engine.SendPlayTurnBatch(playTurn)[0]
	if resp.err != nil {
		println("ERROR PLAY TURN")
		println(resp.err.Error())
		return resp.err
	}

	return nil
}

func PlayFewTurns(engine *GameEngine, game SerializedGameWithID, turnCount int, t *testing.T) SerializedGameWithID {
	for range turnCount {
		// get legal moves
		legalMovesReq := &GetLegalMovesRequest{
			BaseGameID: game.ID, TileToPlace: game.Game.CurrentTile,
		}
		legalMoves := engine.SendGetLegalMovesBatch(
			[]*GetLegalMovesRequest{legalMovesReq},
		)[0]
		if legalMoves.err != nil {
			t.Fatal(legalMoves.err.Error())
		}

		// play turn
		playTurn := []*PlayTurnRequest{
			{
				GameID: game.ID,
				Move:   legalMoves.Moves[0].Move,
			},
		}
		resp := engine.SendPlayTurnBatch(playTurn)[0]
		if resp.err != nil {
			t.Fatal(resp.err.Error())

		}

		// update game
		if resp.Err() != nil {
			t.Fatalf("play turn failed. Reason: %#v", resp.Err().Error())
		}
		game.Game = resp.Game
	}
	return game
}

func TestManyThread(t *testing.T) {
	ThreadCount := 10000

	engine, err := StartGameEngine(4, "")
	if err != nil {
		t.Fatal(err.Error())
	}
	seed := int64(0)

	// create base game
	serializedGameWithID, err := engine.GenerateSeededGame(tilesets.StandardTileSet(), seed)
	if err != nil {
		t.Fatal(err.Error())
	}

	serializedGameWithID = PlayFewTurns(engine, serializedGameWithID, 10, t)

	// get legal moves
	legalMovesReq := &GetLegalMovesRequest{
		BaseGameID: serializedGameWithID.ID, TileToPlace: serializedGameWithID.Game.CurrentTile,
	}
	legalMoves := engine.SendGetLegalMovesBatch(
		[]*GetLegalMovesRequest{legalMovesReq},
	)[0]
	if legalMoves.err != nil {
		t.Fatal(err.Error())
	}

	movesCount := len(legalMoves.Moves)

	// play all turns
	errs := make(chan error, ThreadCount)
	for i := range ThreadCount {
		go func() {
			err := PlayTurnWithIndex(engine, serializedGameWithID.ID, legalMoves.Moves[i%movesCount].Move)
			errs <- err
		}()
	}

	// wait for all N to finish
	for i := 0; i < ThreadCount; i++ {
		err := <-errs
		if err != nil {
			t.Fatal(err)
		}
	}

	engine.Close()
}
