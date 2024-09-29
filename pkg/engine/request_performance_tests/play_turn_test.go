package requestperformancetests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/engine"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
)

func SendPlayTurnRequest(games *[]engine.SerializedGameWithID, eng *engine.GameEngine, b *testing.B) {
	type IDWithCurrentTile struct {
		ID          int
		CurrentTile tiles.Tile
	}
	// create game copies
	copies := []IDWithCurrentTile{}
	for _, game := range *games {
		copy, err := eng.CloneGame(game.ID, 1)
		if err != nil {
			b.Fatal(err.Error())
		}
		copies = append(copies, IDWithCurrentTile{ID: copy[0], CurrentTile: game.Game.CurrentTile})
	}

	// get legal moves for copies
	legalMovesRequests := []*engine.GetLegalMovesRequest{}
	for _, copy := range copies {
		legalMovesRequests = append(legalMovesRequests,
			&engine.GetLegalMovesRequest{
				BaseGameID:  copy.ID,
				TileToPlace: copy.CurrentTile,
			},
		)
	}
	legalMoves := eng.SendGetLegalMovesBatch(legalMovesRequests)
	for i, legalMove := range legalMoves {
		if legalMove.Err() != nil {
			b.Fatalf("%#v legal move failed. Reason: %#v", i, legalMove.Err().Error())
		}
	}

	// create make move requests
	// make move
	makeTurnRequests := []*engine.PlayTurnRequest{}
	for i := range len(copies) {
		makeTurnRequests = append(makeTurnRequests,
			&engine.PlayTurnRequest{
				GameID: copies[i].ID,
				Move:   legalMoves[i].Moves[0].Move,
			},
		)
	}

	b.StartTimer()
	eng.SendPlayTurnBatch(makeTurnRequests)
	b.StopTimer()
}
func BenchmarkPlayTurnTest(b *testing.B) {
	PlayGame(10, b, SendPlayTurnRequest)
}
