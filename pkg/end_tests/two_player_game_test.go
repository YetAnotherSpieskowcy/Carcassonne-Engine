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
 diagonal edges represent cities, dots fields, straight lines roads. The big vertical line on the left is to prevent comment formating

 Final board: (each tile is repestned by 3x3 ascii signs, at the center is the turn number in hex :/)

| 	.|........|.
| 	.9--1--2..8.
| 	.|./ \ |..|.
| 	.| \ / |..|....
| 	.4--0--3..B--C-
| 	.|.............
|	.|.   .........
| 	.5.   >6<>7<>A<
| 	...   .........
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
	checkFinalResult(game, t)

}

func MakeTurn(game *gameMod.Game, t *testing.T, tilePosition elements.Position, rotations uint, meeple elements.MeepleType, featureSide side.Side, featureType feature.Type) {
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}

	var player = game.CurrentPlayer()

	ptile := elements.ToPlacedTile(tile.Rotate(rotations))
	ptile.Position = tilePosition
	if meeple != elements.NoneMeeple {
		ptile.GetPlacedFeatureAtSide(featureSide, featureType).Meeple = elements.Meeple{
			MeepleType: meeple,
			PlayerID:   player.ID(),
		}
	}

	err = game.PlayTurn(ptile)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func CheckMeeplesAndScore(game *gameMod.Game, t *testing.T, player1Score uint32, player1Meeples uint8, player2Score uint32, player2Meeples uint8) {
	var player1 = game.GetPlayerByID(1)
	var player2 = game.GetPlayerByID(2)

	// player 1

	// check meeples count of first player
	if player1.MeepleCount(elements.NormalMeeple) == player1Meeples {
		t.Fatalf("meeples count does not match. Expected: %d  Got: %d", player1Meeples, player1.MeepleCount(elements.NormalMeeple))
	}
	// check points
	if player1.Score() != player1Score {
		t.Fatalf("Player received wrong amount of points! Expected: %d  Got: %d ", player1Score, player1.Score())
	}

	// player 2

	// check meeples count of first player
	if player2.MeepleCount(elements.NormalMeeple) == player2Meeples {
		t.Fatalf("meeples count does not match. Expected: %d  Got: %d", player2Meeples, player2.MeepleCount(elements.NormalMeeple))
	}
	// check points
	if player2.Score() != player2Score {
		t.Fatalf("Player received wrong amount of points! Expected: %d  Got: %d ", player2Score, player2.Score())
	}
}

// straight road with city edge
// player 1 places meeple on city, and closes it
/*
|       0
|
| 	   ...
|1 	   -1-
| 	   / \
| 	   \ /
|0 	   -0-
| 	   ...
*/
func checkFirstTurn(game *gameMod.Game, t *testing.T) {

	MakeTurn(game, t, elements.NewPosition(0, 1), 2, elements.NormalMeeple, side.Bottom, feature.City)
	CheckMeeplesAndScore(game, t, 4, 7, 0, 7)
}

// road turn
// player 2 places meeple (M) on a road
/*
|       0  1
|
| 	   ......
|1 	   -1-M2.
| 	   / \ |.
| 	   \ /
|0 	   -0-
| 	   ...
*/
func checkSecondTurn(game *gameMod.Game, t *testing.T) {
	MakeTurn(game, t, elements.NewPosition(1, 1), 0, elements.NormalMeeple, side.Left, feature.Road)
	CheckMeeplesAndScore(game, t, 4, 7, 0, 6)
}

// road turn
// player 1 places meeple (m) on a field
/*
|       0  1
|
| 	   ......
|1 	   -1-M2.
| 	   / \ |.
| 	   \ /m|.
|0 	   -0--3.
| 	   ......
*/
func checkThirdTurn(game *gameMod.Game, t *testing.T) {
	MakeTurn(game, t, elements.NewPosition(1, 0), 0, elements.NoneMeeple, side.None, feature.None)
	CheckMeeplesAndScore(game, t, 4, 6, 0, 6)
}

// T cross road
// player2 places meeple on road going down
/*
|   -1  0  1
|
| 	   ......
|1 	   -1-M2.
| 	   / \ |.
| 	.| \ /m|.
|0 	.4--0--3.
| 	.M.......
*/
func checkFourthTurn(game *gameMod.Game, t *testing.T) {
	MakeTurn(game, t, elements.NewPosition(-1, 0), 0, elements.NormalMeeple, side.Bottom, feature.Road)
	CheckMeeplesAndScore(game, t, 4, 6, 0, 5)
}

// monastery with single road
// player1 places meeple on a monastery
// road from 4 to 5 is finished, so player2 scores 2 points
/*
|   -1  0  1
|
| 	   ......
|1 	   -1-M2.
| 	   / \ |.
| 	.| \ /m|.
|0 	.4--0--3.
| 	.|.......
|	.|.
|-1 .5.
| 	...
*/
func checkFifthTurn(game *gameMod.Game, t *testing.T) {
	MakeTurn(game, t, elements.NewPosition(-1, 0), 0, elements.NormalMeeple, side.None, feature.Monastery)
	CheckMeeplesAndScore(game, t, 4, 5, 2, 6)

}

// Two city edges not connected
// player 2 places meeple(M) on the right city
/*
|   -1  0  1
|
| 	   ......
|1 	   -1-M2.
| 	   / \ |.
| 	.| \ /m|.
|0 	.4--0--3.
| 	.|.......
|	.|.   ...
|-1 .5.   >6<M
| 	...   ...
*/
func checkSixthTurn(game *gameMod.Game, t *testing.T) {
	MakeTurn(game, t, elements.NewPosition(1, -1), 0, elements.NormalMeeple, side.Right, feature.City)
	CheckMeeplesAndScore(game, t, 4, 5, 2, 5)
}

// Two city edges not connected
// player 1 places meeple on the right city
// playey 2 scores points for finished city

/*
|   -1  0  1  2
|
| 	   ......
|1 	   -1-M2.
| 	   / \ |.
| 	.| \ /m|.
|0 	.4--0--3.
| 	.|.......
|	.|.   ......
|-1	.5.   >6<>7<m
| 	...   ......
*/
func checkSeventhTurn(game *gameMod.Game, t *testing.T) {
	MakeTurn(game, t, elements.NewPosition(2, -1), 0, elements.NormalMeeple, side.Right, feature.City)
	CheckMeeplesAndScore(game, t, 4, 4, 6, 6)
}

// straight road
// player 2 places meeple on a bottom road

/*
|   -1  0  1  2  3
|
| 	   .......|.
|1 	   -1-M2..8.
| 	   / \ |..M.
| 	.| \ /m|.
|0 	.4--0--3.
| 	.|.......
|	.|.   ......
|-1 .5.   >6<>7<m
| 	...   ......
*/
func checkEightthTurn(game *gameMod.Game, t *testing.T) {
	MakeTurn(game, t, elements.NewPosition(1, 1), 0, elements.NormalMeeple, side.Bottom, feature.Road)
	CheckMeeplesAndScore(game, t, 4, 4, 6, 5)
}

// T cross road
// road is finished. Player 2 scores 6 points for a road
// player 1 places meeple (m) on a field

/*
|   -1  0  1  2  3
|
| 	.|m.......|.
|1 	.9--1--2..8.
| 	.|./ \ |..M.
| 	.| \ /m|.
|0 	.4--0--3.
| 	.|.......
|	.|.   ......
|-1	.5.   >6<>7<m
| 	...   ......
*/
func checkNinethTurn(game *gameMod.Game, t *testing.T) {
	MakeTurn(game, t, elements.NewPosition(-1, 1), 0, elements.NormalMeeple, side.TopRightEdge, feature.Field)
	CheckMeeplesAndScore(game, t, 4, 3, 12, 6)
}

// Two city edges not connected
// player 2 places meeple (M) on the right city
// player 1 scores points for city

/*
|   -1  0  1  2  3
|
| 	.|m.......|.
|1 	.9--1--2..8.
| 	.|./ \ |..M.
| 	.| \ /m|.
|0 	.4--0--3.
| 	.|.......
|	.|.   .........
|-1	.5.   >6<>7<>A<M
| 	...   .........
*/
func checkTenthTurn(game *gameMod.Game, t *testing.T) {
	MakeTurn(game, t, elements.NewPosition(3, -1), 0, elements.NormalMeeple, side.Right, feature.City)
	CheckMeeplesAndScore(game, t, 8, 4, 12, 5)
}

// road turn

/*
|   -1  0  1  2  3
|
| 	.|m.......|.
|1 	.9--1--2..8.
| 	.|./ \ |..M.
| 	.| \ /m|..|.
|0 	.4--0--3..B-
| 	.|..........
|	.|.   .........
|-1	.5.   >6<>7<>A<M
| 	...   .........
*/
func checkEleventhTurn(game *gameMod.Game, t *testing.T) {
	MakeTurn(game, t, elements.NewPosition(2, 0), 0, elements.NoneMeeple, side.None, feature.None)
	CheckMeeplesAndScore(game, t, 8, 4, 12, 5)
}

// straight road
/*
|   -1  0  1  2  3
|
| 	.|m.......|.
|1 	.9--1--2..8.
| 	.|./ \ |..M.
| 	.| \ /m|..|....
|0 	.4--0--3..B--C-
| 	.|.............
|	.|.   .........
|-1	.5.   >6<>7<>A<M
| 	...   .........
*/
func checkTwelvethTurn(game *gameMod.Game, t *testing.T) {
	MakeTurn(game, t, elements.NewPosition(3, 0), 0, elements.NoneMeeple, side.None, feature.None)
	CheckMeeplesAndScore(game, t, 8, 4, 12, 5)
}

/* player 1 scores additional 17 points:
	- 2 points for monastery
	- 6 points for farmer in the center
	- 9 points for farmer outside
   player 2 scores additional 4 points:
	- 3 points for a road
	- 1 point for unfishied city
*/

func checkFinalResult(game *gameMod.Game, t *testing.T) {
	var scores, err = game.Finalize()
	if err != nil {
		t.Fatal(err.Error())
	}
	if scores[0] != 25 {
		t.Fatal("Player 1 final score incorrect. Expected 25, got: %d", scores[1])
	}

	if scores[1] != 16 {
		t.Fatal("Player 2 final score incorrect. Expected 16, got: %d", scores[1])
	}
}
