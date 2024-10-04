package requestperformancetests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/engine"
)

func SendCloneGameRequest(games *[]engine.SerializedGameWithID, eng *engine.GameEngine, b *testing.B) {
	var ids []int
	var err error
	b.StartTimer()
	for _, game := range *games {
		ids, err = eng.CloneGame(game.ID, 10)
		if err != nil {
			b.Fatalf("Clone game failed for game: %#v, reason: %#v", game.ID, err.Error())
		}
	}
	b.StopTimer()

	eng.DeleteGames(ids) // free memory of copied games
}

func BenchmarkCloneGameTest(b *testing.B) {
	PlayGame(100, b, SendCloneGameRequest)
}
