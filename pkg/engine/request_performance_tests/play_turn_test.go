package requestperformancetests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/engine"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func BenchmarkPlayTurnTest(b *testing.B) {
	gameCount := 100
	b.StopTimer()

	eng, err := engine.StartGameEngine(4, b.TempDir())
	if err != nil {
		b.Fatal()
	}

	var games = []engine.SerializedGameWithID{}
	for seed := range gameCount {
		game, err := eng.GenerateSeededGame(tilesets.StandardTileSet(), int64(seed+1000))
		if err != nil {
			b.Fatal()
		}
		games = append(games, game)
	}

	// for each turn
	for range len(tilesets.StandardTileSet().Tiles) {

		// get moves
		legalMovesRequests := []*engine.GetLegalMovesRequest{}
		for _, game := range games {
			legalMovesRequests = append(legalMovesRequests,
				&engine.GetLegalMovesRequest{
					BaseGameID:  game.ID,
					TileToPlace: game.Game.CurrentTile,
				},
			)
		}
		legalMoves := eng.SendGetLegalMovesBatch(legalMovesRequests)
		for i, legalMove := range legalMoves {
			if legalMove.Err() != nil {
				b.Fatalf("%#v legal move failed. Reason: %#v", i, legalMove.Err().Error())
			}
		}

		// make move
		makeTurnRequests := []*engine.PlayTurnRequest{}
		for i := range gameCount {
			makeTurnRequests = append(makeTurnRequests,
				&engine.PlayTurnRequest{
					GameID: games[i].ID,
					Move:   legalMoves[i].Moves[0].Move,
				},
			)
		}
		b.StartTimer()
		playTurnResp := eng.SendPlayTurnBatch(makeTurnRequests)
		b.StopTimer()

		// update games
		for i := range gameCount {
			games[i].Game = playTurnResp[i].Game
		}
	}
}
