package requestperformancetests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/engine"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func BenchmarkPlayTurnAtStart(b *testing.B) {
	b.StopTimer()

	move := CreateMovesArray()[0]

	for range b.N {
		eng, err := engine.StartGameEngine(4, b.TempDir())
		if err != nil {
			b.Fatal(err.Error())
		}

		gameWithID, err := eng.GenerateOrderedGame(tilesets.EveryTileOnceTileSet())
		if err != nil {
			b.Fatal(err.Error())
		}
		requests := []*engine.PlayTurnRequest{
			{GameID: gameWithID.ID, Move: move},
		}

		b.StartTimer()
		resp := eng.SendPlayTurnBatch(requests)[0]
		b.StopTimer()
		if resp.Err() != nil {
			b.Fatalf("Request fail")
		}

		eng.Close()
	}

}

func BenchmarkPlayTurnAtEarlyGame(b *testing.B) {
	b.StopTimer()

	move := CreateMovesArray()[10]

	for range b.N {
		eng, serializedGameWithID := CreateEarlyGameEngine(b.TempDir())
		requests := []*engine.PlayTurnRequest{
			{GameID: serializedGameWithID.ID, Move: move},
		}

		b.StartTimer()
		resp := eng.SendPlayTurnBatch(requests)[0]
		b.StopTimer()
		if resp.Err() != nil {
			b.Fatalf("Request fail")
		}

		eng.Close()
	}
}

func BenchmarkPlayTurnAtLateGame(b *testing.B) {
	b.StopTimer()

	move := CreateMovesArray()[20]

	for range b.N {
		eng, serializedGameWithID := CreateLateGameEngine(b.TempDir())
		requests := []*engine.PlayTurnRequest{
			{GameID: serializedGameWithID.ID, Move: move},
		}

		b.StartTimer()
		resp := eng.SendPlayTurnBatch(requests)[0]
		b.StopTimer()
		if resp.Err() != nil {
			b.Fatalf("Request fail")
		}

		eng.Close()
	}
}
