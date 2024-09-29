package requestperformancetests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/engine"
)

func SendCloneGameRequest(games *[]engine.SerializedGameWithID, eng *engine.GameEngine, b *testing.B) {
	b.StartTimer()
	for _, game := range *games {
		_, err := eng.CloneGame(game.ID, 10)
		if err != nil {
			b.Fatalf("Clone game failed for game: %#v, reason: %#v", game.ID, err.Error())
		}
	}
	b.StopTimer()
}

func BenchmarkCloneGameTest(b *testing.B) {
	PlayGame(10, b, SendCloneGameRequest)
}
