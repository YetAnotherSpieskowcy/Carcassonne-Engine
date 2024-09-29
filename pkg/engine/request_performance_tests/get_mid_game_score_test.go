package requestperformancetests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/engine"
)

func SendGetMidGameScoreRequest(games *[]engine.SerializedGameWithID, eng *engine.GameEngine, b *testing.B) {
	requests := []*engine.GetMidGameScoreRequest{}
	for _, game := range *games {
		requests = append(requests,
			&engine.GetMidGameScoreRequest{
				BaseGameID: game.ID,
			},
		)
	}
	b.StartTimer()
	eng.SendGetMidGameScoreBatch(requests)
	b.StopTimer()
}

func BenchmarkGetMidGameScoreTest(b *testing.B) {
	PlayGame(10, b, SendGetMidGameScoreRequest)
}
