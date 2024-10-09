package requestperformancetests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/engine"
)

func SendGetRemainingTilesRequest(games *[]engine.SerializedGameWithID, eng *engine.GameEngine, b *testing.B) {
	requests := []*engine.GetRemainingTilesRequest{}
	for _, game := range *games {
		requests = append(requests,
			&engine.GetRemainingTilesRequest{
				BaseGameID: game.ID,
			},
		)
	}
	b.StartTimer()
	eng.SendGetRemainingTilesBatch(requests)
	b.StopTimer()
}

func BenchmarkGetRemainingTilesTest(b *testing.B) {
	PlayGame(100, SendGetRemainingTilesRequest, b)
}
