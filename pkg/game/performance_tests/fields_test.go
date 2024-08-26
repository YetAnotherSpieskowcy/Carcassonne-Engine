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
func TestLongField(t *testing.T) {
	tileCount := 5000
	fieldTile := tiletemplates.TestOnlyField()

	fieldStart := time.Now()
	PlayNTileGame(tileCount, fieldTile, true)
	fieldEnd := time.Now()

	emptyStart := time.Now()
	PlayNTileGame(tileCount, fieldTile, false)
	emptyEnd := time.Now()

	fieldGameDuration := fieldEnd.Sub(fieldStart)
	emptyGameDuration := emptyEnd.Sub(emptyStart)

	fieldTimeCost := fieldGameDuration - emptyGameDuration
	fmt.Printf("Fieldtime total cost: %s, avg per tile: %s", fieldTimeCost.String(), (fieldTimeCost / time.Duration(tileCount)).String())
}

/*
Play multiple games with few tiles.
Compare total time of those games with empty games, to substract
the creating game time.
*/
func TestManySmallFields(t *testing.T) {

	tileCount := 10
	gameCount := 100_000
	fieldTile := tiletemplates.TestOnlyField()

	fieldStart := time.Now()
	for range gameCount {
		PlayNTileGame(tileCount, fieldTile, true)
	}
	fieldEnd := time.Now()

	emptyStart := time.Now()
	for range gameCount {
		PlayNTileGame(tileCount, fieldTile, false)
	}
	emptyEnd := time.Now()

	fieldGameDuration := fieldEnd.Sub(fieldStart)
	emptyGameDuration := emptyEnd.Sub(emptyStart)

	fieldTimeCost := (fieldGameDuration - emptyGameDuration) / time.Duration(gameCount)
	fmt.Printf("Fieldtime total cost: %s, avg per tile: %s", fieldTimeCost.String(), (fieldTimeCost / time.Duration(tileCount)).String())

}
