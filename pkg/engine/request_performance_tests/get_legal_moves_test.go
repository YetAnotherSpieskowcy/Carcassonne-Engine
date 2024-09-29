package requestperformancetests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/engine"
)

func SendGetLegalMovesRequest(games *[]engine.SerializedGameWithID, eng *engine.GameEngine, b *testing.B) {
	requests := []*engine.GetLegalMovesRequest{}
	for _, game := range *games {
		requests = append(requests,
			&engine.GetLegalMovesRequest{
				BaseGameID:  game.ID,
				TileToPlace: game.Game.CurrentTile,
			},
		)
	}
	b.StartTimer()
	eng.SendGetLegalMovesBatch(requests)
	b.StopTimer()
}

func BenchmarkGetLegalMovesTest(b *testing.B) {
	PlayGame(10, b, SendGetLegalMovesRequest)
}
