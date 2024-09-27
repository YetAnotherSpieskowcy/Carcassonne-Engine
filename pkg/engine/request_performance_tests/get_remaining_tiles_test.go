package requestperformancetests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/engine"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func BenchmarkGetRemainingTilesAtStart(b *testing.B) {
	b.StopTimer()

	eng, err := engine.StartGameEngine(4, b.TempDir())
	if err != nil {
		b.Fatal(err.Error())
	}

	gameWithID, err := eng.GenerateOrderedGame(tilesets.StandardTileSet())
	if err != nil {
		b.Fatal(err.Error())
	}

	requests := []*engine.GetRemainingTilesRequest{{BaseGameID: gameWithID.ID}}
	for range b.N {
		b.StartTimer()
		eng.SendGetRemainingTilesBatch(requests)
		b.StopTimer()
	}

	eng.Close()
}

func BenchmarkGetRemainingTilesAtEarlyGame(b *testing.B) {
	b.StopTimer()

	eng, serializedGameWithID := CreateEarlyGameEngnine(b.TempDir())

	requests := []*engine.GetRemainingTilesRequest{{BaseGameID: serializedGameWithID.ID}}
	for range b.N {
		b.StartTimer()
		eng.SendGetRemainingTilesBatch(requests)
		b.StopTimer()
	}

	eng.Close()
}

func BenchmarkGetRemainingTilesAtLateGame(b *testing.B) {
	b.StopTimer()

	eng, serializedGameWithID := CreateLateGameEngnine(b.TempDir())

	requests := []*engine.GetRemainingTilesRequest{{BaseGameID: serializedGameWithID.ID}}
	for range b.N {
		b.StartTimer()
		eng.SendGetRemainingTilesBatch(requests)
		b.StopTimer()
	}

	eng.Close()
}
