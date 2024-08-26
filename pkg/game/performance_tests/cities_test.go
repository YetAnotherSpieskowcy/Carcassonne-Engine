package performance_tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

/*
Play single long games with maaaaaany tiles.
Compare total time of the game with empty game, to substract
the time to create a game.
*/
func TestLongCity(t *testing.T) {
	tileCount := 5000
	cityTile := tiletemplates.FourCityEdgesConnectedShield()

	cityStart := time.Now()
	PlayNTileGame(tileCount, cityTile, true)
	cityEnd := time.Now()

	emptyStart := time.Now()
	PlayNTileGame(tileCount, cityTile, false)
	emptyEnd := time.Now()

	cityGameDuration := cityEnd.Sub(cityStart)
	emptyGameDuration := emptyEnd.Sub(emptyStart)

	cityTimeCost := cityGameDuration - emptyGameDuration
	fmt.Printf("Citytime total cost: %s, avg per tile: %s", cityTimeCost.String(), (cityTimeCost / time.Duration(tileCount)).String())
}

/*
Play multiple games with few tiles.
Compare total time of those games with empty games, to substract
the creating game time.
*/
func TestManySmallCities(t *testing.T) {

	tileCount := 10
	gameCount := 100_000
	cityTile := tiletemplates.FourCityEdgesConnectedShield()

	cityStart := time.Now()
	for range gameCount {
		PlayNTileGame(tileCount, cityTile, true)
	}
	cityEnd := time.Now()

	emptyStart := time.Now()
	for range gameCount {
		PlayNTileGame(tileCount, cityTile, false)
	}
	emptyEnd := time.Now()

	cityGameDuration := cityEnd.Sub(cityStart)
	emptyGameDuration := emptyEnd.Sub(emptyStart)

	roadTimeCost := (cityGameDuration - emptyGameDuration) / time.Duration(gameCount)
	fmt.Printf("citytime total cost: %s, avg per tile: %s", roadTimeCost.String(), (roadTimeCost / time.Duration(tileCount)).String())

}
