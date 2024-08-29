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
func BenchmarkLongCity(b *testing.B) {
	b.StopTimer()
	tileCount := 5000
	cityTile := tiletemplates.FourCityEdgesConnectedShield()

	for range b.N {
		b.StartTimer()
		err := PlayNTileGame(tileCount, cityTile, true)
		b.StopTimer()
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
}

// similar test to above, but only measuring the creating process of game
func BenchmarkLongCitySetup(b *testing.B) {
	b.StopTimer()
	tileCount := 5000
	cityTile := tiletemplates.FourCityEdgesConnectedShield()

	for range b.N {
		b.StartTimer()
		err := PlayNTileGame(tileCount, cityTile, false)
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
func BenchmarkManySmallCities(b *testing.B) {
	b.StopTimer()
	tileCount := 10
	cityTile := tiletemplates.FourCityEdgesConnectedShield()

	for range b.N {
		b.StartTimer()
		err := PlayNTileGame(tileCount, cityTile, true)
		b.StopTimer()
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
}

// similar test to above, but only measuring the creating process of game
func BenchmarkManySmallCitiesSetup(b *testing.B) {
	b.StopTimer()
	tileCount := 10
	cityTile := tiletemplates.FourCityEdgesConnectedShield()

	for range b.N {
		b.StartTimer()
		err := PlayNTileGame(tileCount, cityTile, false)
		b.StopTimer()
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
}
