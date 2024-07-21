package two_player_game_test

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/end_tests"
	gameMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

/*
 diagonal edges represent cities, dots fields, straight lines roads. The big vertical line on the left is to prevent comment formating

 Final board: (each tile is represented by 3x3 ascii signs, at the center is the turn number in hex :/)

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

func Test2PlayerFullGame(t *testing.T) {
	// create game
	minitileSet := end_tests.MiniTileSetRoadsAndFields()
	deckStack := stack.NewSeeded(minitileSet.Tiles, 47328235) // random seed :P
	deck := deck.Deck{Stack: &deckStack, StartingTile: minitileSet.StartingTile}
	game, err := gameMod.NewFromDeck(deck, nil, 2)
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
	pos := position.New(0, 1)
	end_tests.MakeTurn(game, t, pos, 2, elements.NormalMeeple, side.Bottom, feature.City)

	end_tests.VerifyMeepleExistence(t, game, pos, side.Bottom, feature.City, false, 1) // removed meeple
	end_tests.CheckMeeplesAndScore(game, t, []uint32{4, 0}, []uint8{7, 7}, 1)
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
	pos := position.New(1, 1)
	end_tests.MakeTurn(game, t, pos, 0, elements.NormalMeeple, side.Left, feature.Road)

	end_tests.VerifyMeepleExistence(t, game, pos, side.Left, feature.Road, true, 2)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{4, 0}, []uint8{7, 6}, 2)
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
	pos := position.New(1, 0)
	end_tests.MakeTurn(game, t, pos, 0, elements.NoneMeeple, side.NoSide, feature.NoneType)

	end_tests.VerifyMeepleExistence(t, game, pos, side.TopLeftEdge, feature.Field, true, 3)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{4, 0}, []uint8{6, 6}, 3)
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
	pos := position.New(-1, 0)
	end_tests.MakeTurn(game, t, pos, 0, elements.NormalMeeple, side.Bottom, feature.Road)

	end_tests.VerifyMeepleExistence(t, game, pos, side.Bottom, feature.Road, true, 4)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{4, 0}, []uint8{6, 5}, 4)
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
	pos := position.New(-1, -1)
	end_tests.MakeTurn(game, t, pos, 0, elements.NormalMeeple, side.NoSide, feature.Monastery)

	end_tests.VerifyMeepleExistence(t, game, pos, side.NoSide, feature.Monastery, true, 5)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{4, 2}, []uint8{5, 6}, 5)

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
	pos := position.New(1, -1)
	end_tests.MakeTurn(game, t, pos, 0, elements.NormalMeeple, side.Right, feature.City)

	end_tests.VerifyMeepleExistence(t, game, pos, side.Right, feature.City, true, 6)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{4, 2}, []uint8{5, 5}, 6)
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
	pos := position.New(2, -1)
	end_tests.MakeTurn(game, t, pos, 0, elements.NormalMeeple, side.Right, feature.City)

	end_tests.VerifyMeepleExistence(t, game, position.New(1, -1), side.Right, feature.City, false, 7) // removed meeple
	end_tests.VerifyMeepleExistence(t, game, pos, side.Right, feature.City, true, 7)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{4, 6}, []uint8{4, 6}, 7)
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
	pos := position.New(1, 1)
	end_tests.MakeTurn(game, t, pos, 0, elements.NormalMeeple, side.Bottom, feature.Road)

	end_tests.VerifyMeepleExistence(t, game, pos, side.Bottom, feature.Road, true, 8)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{4, 6}, []uint8{4, 5}, 8)
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
	pos := position.New(-1, 1)
	end_tests.MakeTurn(game, t, pos, 0, elements.NormalMeeple, side.TopRightEdge, feature.Field)

	end_tests.VerifyMeepleExistence(t, game, position.New(1, 1), side.Left, feature.Road, false, 9) // removed meeple
	end_tests.VerifyMeepleExistence(t, game, pos, side.TopRightEdge, feature.Field, true, 9)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{4, 12}, []uint8{3, 6}, 9)
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
	pos := position.New(3, -1)
	end_tests.MakeTurn(game, t, pos, 0, elements.NormalMeeple, side.Right, feature.City)

	end_tests.VerifyMeepleExistence(t, game, position.New(2, 1), side.Right, feature.City, false, 10) // removed meeple
	end_tests.VerifyMeepleExistence(t, game, pos, side.Right, feature.City, true, 10)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{8, 12}, []uint8{4, 5}, 10)
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
	pos := position.New(2, 0)
	end_tests.MakeTurn(game, t, pos, 0, elements.NoneMeeple, side.NoSide, feature.NoneType)

	end_tests.CheckMeeplesAndScore(game, t, []uint32{8, 12}, []uint8{4, 5}, 11)
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
	pos := position.New(3, 0)
	end_tests.MakeTurn(game, t, pos, 0, elements.NoneMeeple, side.NoSide, feature.NoneType)

	end_tests.CheckMeeplesAndScore(game, t, []uint32{8, 12}, []uint8{4, 5}, 12)
}

/* player 1 scores additional 11 points:
	- 2 points for monastery
	- 3 points for farmer in the center
	- 6 points for farmer outside
   player 2 scores additional 4 points:
	- 3 points for a road
	- 1 point for unfishied city
*/

func checkFinalResult(game *gameMod.Game, t *testing.T) {
	var scores, err = game.Finalize()
	if err != nil {
		t.Fatal(err.Error())
	}
	if scores.ReceivedPoints[0] != 19 {
		t.Fatalf("Player 1 final score incorrect. Expected 19, got: %d", scores.ReceivedPoints[1])
	}

	if scores.ReceivedPoints[1] != 16 {
		t.Fatalf("Player 2 final score incorrect. Expected 16, got: %d", scores.ReceivedPoints[1])
	}
}
