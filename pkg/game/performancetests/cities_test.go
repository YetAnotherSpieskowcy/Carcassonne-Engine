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
		err := PlayNTileGame(tileCount, cityTile, b)
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
		err := PlayNTileGame(tileCount, cityTile, b)
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
}
