package requestperformancetests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/engine"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func BenchmarkCloneGameAtStart1(b *testing.B) {
	b.StopTimer()

	eng, err := engine.StartGameEngine(4, b.TempDir())
	if err != nil {
		b.Fatal(err.Error())
	}

	gameWithID, err := eng.GenerateOrderedGame(tilesets.EveryTileOnceTileSet())
	if err != nil {
		b.Fatal(err.Error())
	}

	for range b.N {
		b.StartTimer()
		_, err := eng.CloneGame(gameWithID.ID, 1)
		b.StopTimer()
		if err != nil {
			b.Fatalf("Clone fail")
		}
	}

	eng.Close()
}

func BenchmarkCloneGameAtStart100(b *testing.B) {
	b.StopTimer()

	eng, err := engine.StartGameEngine(4, b.TempDir())
	if err != nil {
		b.Fatal(err.Error())
	}

	gameWithID, err := eng.GenerateOrderedGame(tilesets.EveryTileOnceTileSet())
	if err != nil {
		b.Fatal(err.Error())
	}

	for range b.N {
		b.StartTimer()
		_, err := eng.CloneGame(gameWithID.ID, 100)
		b.StopTimer()
		if err != nil {
			b.Fatalf("Clone fail")
		}
	}

	eng.Close()
}

func BenchmarkCloneGameAtEarlyGame1(b *testing.B) {
	b.StopTimer()

	eng, serializedGameWithID := CreateEarlyGameEngine(b.TempDir())

	for range b.N {
		b.StartTimer()
		_, err := eng.CloneGame(serializedGameWithID.ID, 1)
		b.StopTimer()
		if err != nil {
			b.Fatalf("Clone fail")
		}
	}

	eng.Close()
}

func BenchmarkCloneGameAtEarlyGame100(b *testing.B) {
	b.StopTimer()

	eng, serializedGameWithID := CreateEarlyGameEngine(b.TempDir())

	for range b.N {
		b.StartTimer()
		_, err := eng.CloneGame(serializedGameWithID.ID, 100)
		b.StopTimer()
		if err != nil {
			b.Fatalf("Clone fail")
		}
	}

	eng.Close()
}

func BenchmarkCloneGameAtLateGame1(b *testing.B) {
	b.StopTimer()

	eng, serializedGameWithID := CreateLateGameEngine(b.TempDir())

	for range b.N {
		b.StartTimer()
		_, err := eng.CloneGame(serializedGameWithID.ID, 1)
		b.StopTimer()
		if err != nil {
			b.Fatalf("Clone fail")
		}
	}

	eng.Close()
}

func BenchmarkCloneGameAtLateGame100(b *testing.B) {
	b.StopTimer()

	eng, serializedGameWithID := CreateLateGameEngine(b.TempDir())

	for range b.N {
		b.StartTimer()
		_, err := eng.CloneGame(serializedGameWithID.ID, 100)
		b.StopTimer()
		if err != nil {
			b.Fatalf("Clone fail")
		}
	}

	eng.Close()
}
