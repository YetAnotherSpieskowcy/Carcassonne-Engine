package requestperformancetests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/engine"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func BenchmarkGetLegalMovesAtStart(b *testing.B) {
	b.StopTimer()

	eng, err := engine.StartGameEngine(4, b.TempDir())
	if err != nil {
		b.Fatal(err.Error())
	}

	gameWithID, err := eng.GenerateOrderedGame(tilesets.StandardTileSet())
	if err != nil {
		b.Fatal(err.Error())
	}

	requests := []*engine.GetLegalMovesRequest{{BaseGameID: gameWithID.ID}}
	for range b.N {
		b.StartTimer()
		eng.SendGetLegalMovesBatch(requests)
		b.StopTimer()
	}

	eng.Close()
}

func BenchmarkGetLegalMovesAtEarlyGame(b *testing.B) {
	b.StopTimer()

	eng, serializedGameWithID := CreateEarlyGameEngnine(b.TempDir())

	requests := []*engine.GetLegalMovesRequest{{BaseGameID: serializedGameWithID.ID}}
	for range b.N {
		b.StartTimer()
		eng.SendGetLegalMovesBatch(requests)
		b.StopTimer()
	}

	eng.Close()
}

func BenchmarkGetLegalMovesLateGame(b *testing.B) {
	b.StopTimer()

	eng, serializedGameWithID := CreateEarlyGameEngnine(b.TempDir())

	requests := []*engine.GetLegalMovesRequest{{BaseGameID: serializedGameWithID.ID}}
	for range b.N {
		b.StartTimer()
		eng.SendGetLegalMovesBatch(requests)
		b.StopTimer()
	}

	eng.Close()
}
