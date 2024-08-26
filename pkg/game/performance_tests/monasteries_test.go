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
func TestLongMonastery(t *testing.T) {
	tileCount := 5000
	monasteryTile := tiletemplates.MonasteryWithoutRoads()

	monasteryStart := time.Now()
	PlayNTileGame(tileCount, monasteryTile, true)
	monasteryEnd := time.Now()

	emptyStart := time.Now()
	PlayNTileGame(tileCount, monasteryTile, false)
	emptyEnd := time.Now()

	monasteryGameDuration := monasteryEnd.Sub(monasteryStart)
	emptyGameDuration := emptyEnd.Sub(emptyStart)

	monasteryTimeCost := monasteryGameDuration - emptyGameDuration
	fmt.Printf("monasterytime total cost: %s, avg per tile: %s", monasteryTimeCost.String(), (monasteryTimeCost / time.Duration(tileCount)).String())
}

/*
Play multiple games with few tiles.
Compare total time of those games with empty games, to substract
the creating game time.
*/
func TestManySmallMonasteries(t *testing.T) {

	tileCount := 10
	gameCount := 100_000
	monasteryTile := tiletemplates.MonasteryWithoutRoads()

	monasteryStart := time.Now()
	for range gameCount {
		PlayNTileGame(tileCount, monasteryTile, true)
	}
	monasteryEnd := time.Now()

	emptyStart := time.Now()
	for range gameCount {
		PlayNTileGame(tileCount, monasteryTile, false)
	}
	emptyEnd := time.Now()

	monasteryGameDuration := monasteryEnd.Sub(monasteryStart)
	emptyGameDuration := emptyEnd.Sub(emptyStart)

	roadTimeCost := (monasteryGameDuration - emptyGameDuration) / time.Duration(gameCount)
	fmt.Printf("monasterytime total cost: %s, avg per tile: %s", roadTimeCost.String(), (roadTimeCost / time.Duration(tileCount)).String())

}
