package performancetests

import (
	"fmt"
	"testing"
	"time"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

/*
Play single long games with maaaaaany tiles.
Compare total time of the game with only field game, to substract
the field calculation (to certain degree).
*/
func TestSingleExtraLongRoad(t *testing.T) {
	tileCount := 5000
	roadTile := tiletemplates.StraightRoads()
	fieldTile := tiletemplates.TestOnlyField()

	roadStart := time.Now()
	err := PlayNTileGame(tileCount, roadTile, true)
	if err != nil {
		t.Fatalf(err.Error())
	}
	roadEnd := time.Now()

	fieldStart := time.Now()
	err = PlayNTileGame(tileCount, fieldTile, true)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fieldEnd := time.Now()

	roadGameDuration := roadEnd.Sub(roadStart)
	fieldGameDuration := fieldEnd.Sub(fieldStart)

	roadTimeCost := roadGameDuration - fieldGameDuration
	fmt.Printf("Roadtime total cost: %s, avg per tile: %s", roadTimeCost.String(), (roadTimeCost / time.Duration(tileCount)).String())
}

/*
Play multiple games with few tiles.
Compare total time of those games with empty games, to substract
the creating game time.
*/
func TestManyShortRoads(t *testing.T) {

	tileCount := 10
	gameCount := 100_000
	roadTile := tiletemplates.StraightRoads()

	roadStart := time.Now()
	for range gameCount {
		err := PlayNTileGame(tileCount, roadTile, true)
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	roadEnd := time.Now()

	emptyStart := time.Now()
	for range gameCount {
		err := PlayNTileGame(tileCount, roadTile, false)
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	emptyEnd := time.Now()

	roadGameDuration := roadEnd.Sub(roadStart)
	emptyGameDuration := emptyEnd.Sub(emptyStart)

	roadTimeCost := (roadGameDuration - emptyGameDuration) / time.Duration(gameCount)
	fmt.Printf("Roadtime total cost: %s, avg per tile: %s", roadTimeCost.String(), (roadTimeCost / time.Duration(tileCount)).String())
}
