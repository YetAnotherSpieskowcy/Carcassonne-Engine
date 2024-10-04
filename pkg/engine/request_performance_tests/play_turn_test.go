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
	// create game clones
	clones := []IDWithCurrentTile{}
	for _, game := range *games {
		clone, err := eng.CloneGame(game.ID, 1)
		if err != nil {
			b.Fatal(err.Error())
		}
		clones = append(clones, IDWithCurrentTile{ID: clone[0], CurrentTile: game.Game.CurrentTile})
	}

	// get legal moves for clones
	legalMovesRequests := []*engine.GetLegalMovesRequest{}
	for _, clone := range clones {
		legalMovesRequests = append(legalMovesRequests,
			&engine.GetLegalMovesRequest{
				BaseGameID:  clone.ID,
				TileToPlace: clone.CurrentTile,
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
	for i := range len(clones) {
		makeTurnRequests = append(makeTurnRequests,
			&engine.PlayTurnRequest{
				GameID: clones[i].ID,
				Move:   legalMoves[i].Moves[0].Move,
			},
		)
	}

	b.StartTimer()
	eng.SendPlayTurnBatch(makeTurnRequests)
	b.StopTimer()
}
func BenchmarkPlayTurnTest(b *testing.B) {
	PlayGame(100, b, SendPlayTurnRequest)
}
