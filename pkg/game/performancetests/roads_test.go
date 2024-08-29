package performancetests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

/*
Play single long games with maaaaaany tiles.
Compare total time of the game with only field game, to substract
the field calculation (to certain degree).
*/
func BenchmarkSingleExtraLongRoad(b *testing.B) {
	b.StopTimer()
	roadTile := tiletemplates.TestOnlyStraightRoads()
	tileCount := 5000

	for range b.N {
		b.StartTimer()
		err := PlayNTileGame(tileCount, roadTile, true)
		b.StopTimer()
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
}

// similar test to above, but only measuring the creating process of game
func BenchmarkSingleExtraLongRoadSetup(b *testing.B) {
	b.StopTimer()
	roadTile := tiletemplates.TestOnlyStraightRoads()
	tileCount := 5000

	for range b.N {
		b.StartTimer()
		err := PlayNTileGame(tileCount, roadTile, false)
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
func BenchmarkManyShortRoads(b *testing.B) {

	b.StopTimer()
	roadTile := tiletemplates.TestOnlyStraightRoads()
	tileCount := 10

	for range b.N {
		b.StartTimer()
		err := PlayNTileGame(tileCount, roadTile, true)
		b.StopTimer()
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
}

// similar test to above, but only measuring the creating process of game
func BenchmarkManyShortRoadsSetup(b *testing.B) {

	b.StopTimer()
	roadTile := tiletemplates.TestOnlyStraightRoads()
	tileCount := 10

	for range b.N {
		b.StartTimer()
		err := PlayNTileGame(tileCount, roadTile, false)
		b.StopTimer()
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
}
