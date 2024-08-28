//go:build performanceTest

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
func BenchmarkLongField(b *testing.B) {
	b.StopTimer()
	tileCount := 5000
	fieldTile := tiletemplates.TestOnlyField()

	for range b.N {
		b.StartTimer()
		err := PlayNTileGame(tileCount, fieldTile, true)
		b.StopTimer()
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
}

// similar test to above, but only measuring the creating process of game
func BenchmarkLongFieldSetup(b *testing.B) {
	b.StopTimer()
	tileCount := 5000
	fieldTile := tiletemplates.TestOnlyField()

	for range b.N {
		b.StartTimer()
		err := PlayNTileGame(tileCount, fieldTile, false)
		b.StopTimer()
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
func BenchmarkManySmallFields(b *testing.B) {
	b.StopTimer()
	tileCount := 10
	fieldTile := tiletemplates.TestOnlyField()

	for range b.N {
		b.StartTimer()
		err := PlayNTileGame(tileCount, fieldTile, true)
		b.StopTimer()
		if err != nil {
			b.Fatalf(err.Error())
		}
	}

}

// similar test to above, but only measuring the creating process of game
func BenchmarkManySmallFieldsSetup(b *testing.B) {
	b.StopTimer()
	tileCount := 10
	fieldTile := tiletemplates.TestOnlyField()

	for range b.N {
		b.StartTimer()
		err := PlayNTileGame(tileCount, fieldTile, false)
		b.StopTimer()
		if err != nil {
			b.Fatalf(err.Error())
		}
	}

}
