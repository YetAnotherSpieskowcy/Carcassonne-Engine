package performancetests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

/*
Play single long games with maaaaaany tiles.
Compare total time of the game with empty game, to substract
the time to create a game.
*/
func BenchmarkLongMonastery(b *testing.B) {
	b.StopTimer()
	tileCount := 5000
	monasteryTile := tiletemplates.TestOnlyMonastery()

	for range b.N {
		err := PlayNTileGame(tileCount, monasteryTile, b)
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
}

/*
Play multiple games with few tiles.
Compare total time of those games with empty games, to substract
the creating game time.
*/
func BenchmarkManySmallMonasteries(b *testing.B) {
	b.StopTimer()
	tileCount := 10
	monasteryTile := tiletemplates.TestOnlyMonastery()

	for range b.N {
		err := PlayNTileGame(tileCount, monasteryTile, b)
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
}
