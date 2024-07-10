package end_tests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	gameMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

/*
 diagonal edges represent cities

 Final board: (each tile is repestned by 3x3 ascii signs, at the center is the turn number in hex :/)

 .|........|.
 .9--1--2..8.
 .|./ \ |..|.
 .| \ / |..|....
 .4--0--3..B--C-
 .|.............
 .|.   .........
 .5.   >6<>7<>A<
 ...   .........
*/

func TestFullGame(t *testing.T) {
	// create game
	minitileSet := MiniTileSetRoadsAndFields()
	deckStack := stack.NewSeeded(minitileSet.Tiles, 47328235) // random seed :P
	deck := deck.Deck{Stack: &deckStack, StartingTile: minitileSet.StartingTile}
	game, err := gameMod.NewFromDeck(deck, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	/*
		for i, _ := range deck.GetTiles() {
			fmt.Printf("%v:\n", i+1)
			var tile, _ = deck.Next()
			fmt.Printf("%v\n", tile)
		}
	*/

	checkFirstTurn(game, t)    // straight road with city edge
	checkSecondTurn(game, t)   // road turn
	checkThirdTurn(game, t)    // road turn
	checkFourthTurn(game, t)   // T cross road
	checkFifthTurn(game, t)    // monastery with single road
	checkSixthTurn(game, t)    // Two city edges not connected
	checkSeventhTurn(game, t)  // Two city edges not connected
	checkEightthTurn(game, t)  // straight road
	checkNinethTurn(game, t)   // T cross road
	checkTenthTurn(game, t)    // Two city edges not connected
	checkEleventhTurn(game, t) // road turn
	checkTwelvethTurn(game, t) // straight road
}

// straight road with city edge
// player 1 places meeple on city, and closes it
/*
...
-1-
/ \
\ /
-0-
...
*/
func checkFirstTurn(game *gameMod.Game, t *testing.T) {
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	var startMeepleCount = game.GetPlayerByID(1).MeepleCount(elements.NormalMeeple)
	ptile := elements.ToPlacedTile(tile.Rotate(2))
	ptile.Position = elements.NewPosition(0, 1)
	ptile.GetPlacedFeatureAtSide(side.Bottom, feature.City).Meeple = elements.Meeple{
		MeepleType: elements.NormalMeeple,
		PlayerID:   game.CurrentPlayer().ID(),
	}

	err = game.PlayTurn(ptile)
	if err != nil {
		t.Fatal(err.Error())
	}

	var player = game.GetPlayerByID(1)
	// check meeples count of first player
	if player.MeepleCount(elements.NormalMeeple) == startMeepleCount-1 {
		t.Fatalf("meeples count did not decrease after placing meeple")
	}
	// check points
	if player.Score() != 2 {
		t.Fatalf("Player received wrong amount of points! expected: got:")
	}
}

// road turn
/*
.|........|.
 .9--1--2..8.
 .|./ \ |..|.
 .| \ / |..|....
 .4--0--3..B--C-
 .|.............
 .|.   .........
 .5.   >6<>7<>A<
 ...   .........
*/
func checkSecondTurn(game *gameMod.Game, t *testing.T) {
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile := elements.ToPlacedTile(tile)
	ptile.Position = elements.NewPosition(0, 1)
}

// road turn

/*
.|........|.

	.9--1--2..8.
	.|./ \ |..|.
	.| \ / |..|....
	.4--0--3..B--C-
	.|.............
	.|.   .........
	.5.   >6<>7<>A<
	...   .........
*/
func checkThirdTurn(game *gameMod.Game, t *testing.T) {
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile := elements.ToPlacedTile(tile)
	ptile.Position = elements.NewPosition(0, 1)
}

// T cross road

/*
.|........|.

	.9--1--2..8.
	.|./ \ |..|.
	.| \ / |..|....
	.4--0--3..B--C-
	.|.............
	.|.   .........
	.5.   >6<>7<>A<
	...   .........
*/
func checkFourthTurn(game *gameMod.Game, t *testing.T) {
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile := elements.ToPlacedTile(tile)
	ptile.Position = elements.NewPosition(0, 1)
}

// monastery with single road

/*
.|........|.

	.9--1--2..8.
	.|./ \ |..|.
	.| \ / |..|....
	.4--0--3..B--C-
	.|.............
	.|.   .........
	.5.   >6<>7<>A<
	...   .........
*/
func checkFifthTurn(game *gameMod.Game, t *testing.T) {
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile := elements.ToPlacedTile(tile)
	ptile.Position = elements.NewPosition(0, 1)
}

// Two city edges not connected

/*
.|........|.

	.9--1--2..8.
	.|./ \ |..|.
	.| \ / |..|....
	.4--0--3..B--C-
	.|.............
	.|.   .........
	.5.   >6<>7<>A<
	...   .........
*/
func checkSixthTurn(game *gameMod.Game, t *testing.T) {
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile := elements.ToPlacedTile(tile)
	ptile.Position = elements.NewPosition(0, 1)
}

// Two city edges not connected

/*
.|........|.

	.9--1--2..8.
	.|./ \ |..|.
	.| \ / |..|....
	.4--0--3..B--C-
	.|.............
	.|.   .........
	.5.   >6<>7<>A<
	...   .........
*/
func checkSeventhTurn(game *gameMod.Game, t *testing.T) {
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile := elements.ToPlacedTile(tile)
	ptile.Position = elements.NewPosition(0, 1)
}

// straight road

/*
	.......|.
	-1--2..8.
	/ \ |..|.

.| \ / |.
.4--0--3.
.|.......
.|.   ......
.5.   >6<>7<
...   ......
*/
func checkEightthTurn(game *gameMod.Game, t *testing.T) {
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile := elements.ToPlacedTile(tile)
	ptile.Position = elements.NewPosition(0, 1)
}

// T cross road

/*
.|........|.
.9--1--2..8.
.|./ \ |..|.
.| \ / |.
.4--0--3.
.|.......
.|.   ......
.5.   >6<>7<
...   ......
*/
func checkNinethTurn(game *gameMod.Game, t *testing.T) {
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile := elements.ToPlacedTile(tile)
	ptile.Position = elements.NewPosition(0, 1)
}

// Two city edges not connected

/*
.|........|.
.9--1--2..8.
.|./ \ |..|.
.| \ / |.
.4--0--3.
.|.......
.|.   .........
.5.   >6<>7<>A<
...   .........
*/
func checkTenthTurn(game *gameMod.Game, t *testing.T) {
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile := elements.ToPlacedTile(tile)
	ptile.Position = elements.NewPosition(0, 1)
}

// road turn

/*
.|........|.
.9--1--2..8.
.|./ \ |..|.
.| \ / |..|.
.4--0--3..B-
.|..........
.|.   .........
.5.   >6<>7<>A<
...   .........
*/
func checkEleventhTurn(game *gameMod.Game, t *testing.T) {
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile := elements.ToPlacedTile(tile)
	ptile.Position = elements.NewPosition(0, 1)
}

// straight road
/*
 .|........|.
 .9--1--2..8.
 .|./ \ |..|.
 .| \ / |..|....
 .4--0--3..B--C-
 .|.............
 .|.   .........
 .5.   >6<>7<>A<
 ...   .........
*/
func checkTwelvethTurn(game *gameMod.Game, t *testing.T) {
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile := elements.ToPlacedTile(tile)
	ptile.Position = elements.NewPosition(0, 1)
}
