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
		resp := eng.SendGetRemainingTilesBatch(requests)[0]
		b.StopTimer()
		if resp.Err() != nil {
			b.Fatalf("Request fail")
		}
	}

	eng.Close()
}

func BenchmarkGetRemainingTilesAtEarlyGame(b *testing.B) {
	b.StopTimer()

	eng, serializedGameWithID := CreateEarlyGameEngine(b.TempDir())

	requests := []*engine.GetRemainingTilesRequest{{BaseGameID: serializedGameWithID.ID}}
	for range b.N {
		b.StartTimer()
		resp := eng.SendGetRemainingTilesBatch(requests)[0]
		b.StopTimer()
		if resp.Err() != nil {
			b.Fatalf("Request fail")
		}
	}

	eng.Close()
}

func BenchmarkGetRemainingTilesAtLateGame(b *testing.B) {
	b.StopTimer()

	eng, serializedGameWithID := CreateLateGameEngine(b.TempDir())

	requests := []*engine.GetRemainingTilesRequest{{BaseGameID: serializedGameWithID.ID}}
	for range b.N {
		b.StartTimer()
		resp := eng.SendGetRemainingTilesBatch(requests)[0]
		b.StopTimer()
		if resp.Err() != nil {
			b.Fatalf("Request fail")
		}
	}

	eng.Close()
}
