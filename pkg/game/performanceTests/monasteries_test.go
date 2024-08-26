//go:build performanceTest

package performancetests

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
	err := PlayNTileGame(tileCount, monasteryTile, true)
	if err != nil {
		t.Fatalf(err.Error())
	}
	monasteryEnd := time.Now()

	emptyStart := time.Now()
	err = PlayNTileGame(tileCount, monasteryTile, false)
	if err != nil {
		t.Fatalf(err.Error())
	}
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
		err := PlayNTileGame(tileCount, monasteryTile, true)
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	monasteryEnd := time.Now()

	emptyStart := time.Now()
	for range gameCount {
		err := PlayNTileGame(tileCount, monasteryTile, false)
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	emptyEnd := time.Now()

	monasteryGameDuration := monasteryEnd.Sub(monasteryStart)
	emptyGameDuration := emptyEnd.Sub(emptyStart)

	roadTimeCost := (monasteryGameDuration - emptyGameDuration) / time.Duration(gameCount)
	fmt.Printf("monasterytime total cost: %s, avg per tile: %s", roadTimeCost.String(), (roadTimeCost / time.Duration(tileCount)).String())

}
