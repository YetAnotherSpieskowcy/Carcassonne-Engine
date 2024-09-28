package requestperformancetests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/engine"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func BenchmarkPlayTurnAtStart(b *testing.B) {
	b.StopTimer()

	move := CreateMovesArray()[0]
	eng, err := engine.StartGameEngine(4, b.TempDir())
	if err != nil {
		b.Fatal(err.Error())
	}
	gameWithID, err := eng.GenerateOrderedGame(tilesets.EveryTileOnceTileSet())
	if err != nil {
		b.Fatal(err.Error())
	}

	for range b.N {
		// create clone game
		cloneGamesID, err := eng.CloneGame(gameWithID.ID, 1)
		if err != nil {
			b.Fatalf(err.Error())
		}
		cloneGameID := cloneGamesID[0]
		requests := []*engine.PlayTurnRequest{
			{GameID: cloneGameID, Move: move},
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
	eng, serializedGameWithID := CreateEarlyGameEngine(b.TempDir())

	for range b.N {
		// create clone game
		cloneGamesID, err := eng.CloneGame(serializedGameWithID.ID, 1)
		if err != nil {
			b.Fatalf(err.Error())
		}
		cloneGameID := cloneGamesID[0]
		requests := []*engine.PlayTurnRequest{
			{GameID: cloneGameID, Move: move},
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
	eng, serializedGameWithID := CreateLateGameEngine(b.TempDir())

	for range b.N {
		// create clone game
		cloneGamesID, err := eng.CloneGame(serializedGameWithID.ID, 1)
		if err != nil {
			b.Fatalf(err.Error())
		}
		cloneGameID := cloneGamesID[0]
		requests := []*engine.PlayTurnRequest{
			{GameID: cloneGameID, Move: move},
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
