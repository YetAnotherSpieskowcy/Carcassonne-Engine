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

	gameWithID, err := eng.GenerateOrderedGame(tilesets.EveryTileOnceTileSet())
	if err != nil {
		b.Fatal(err.Error())
	}

	requests := []*engine.GetLegalMovesRequest{{BaseGameID: gameWithID.ID, TileToPlace: gameWithID.Game.CurrentTile}}
	for range b.N {
		b.StartTimer()
		resp := eng.SendGetLegalMovesBatch(requests)[0]
		b.StopTimer()
		if resp.Err() != nil {
			b.Fatalf("Request fail")
		}
	}

	eng.Close()
}

func BenchmarkGetLegalMovesAtEarlyGame(b *testing.B) {
	b.StopTimer()

	eng, serializedGameWithID := CreateEarlyGameEngine(b.TempDir())

	requests := []*engine.GetLegalMovesRequest{{BaseGameID: serializedGameWithID.ID, TileToPlace: serializedGameWithID.Game.CurrentTile}}
	for range b.N {
		b.StartTimer()
		resp := eng.SendGetLegalMovesBatch(requests)[0]
		b.StopTimer()
		if resp.Err() != nil {
			b.Fatalf("Request fail")
		}
	}

	eng.Close()
}

func BenchmarkGetLegalMovesLateGame(b *testing.B) {
	b.StopTimer()

	eng, serializedGameWithID := CreateEarlyGameEngine(b.TempDir())

	requests := []*engine.GetLegalMovesRequest{{BaseGameID: serializedGameWithID.ID, TileToPlace: serializedGameWithID.Game.CurrentTile}}
	for range b.N {
		b.StartTimer()
		resp := eng.SendGetLegalMovesBatch(requests)[0]
		b.StopTimer()
		if resp.Err() != nil {
			b.Fatalf("Request fail")
		}
	}

	eng.Close()
}
