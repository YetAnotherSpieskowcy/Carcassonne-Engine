package four_player_game_test

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/end_tests"
	gameMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

/*
 diagonal edges represent cities, dots fields, straight lines roads. The big vertical line on the left is to prevent comment formating
 Player meeples will be represented as !@#$ signs ( you know, writing number but with shift!) 1->!, 2->@ and so on

 Final board: (each tile is represented by 5x5 ascii signs, at the center is the turn number in hex :/)

|				  0	   1    2    3
|
|
|               .| |................
|				.\ /................
|1				..2....4----9---[A].
|				./ \...|.../ \......
|				.| |...|...| |......
|		   ......| |...|...| |.
|		   ......\ /...|...\ /.
|0		   ..8----0----1....3..
|	       ..|.........|.../ \.
|	       ..|.........|...| |.
|	       ..|.........|...| |.
|	       ..|.........|...\ /
|-1		   ..B----6----5....7..
|	       ..|............./ \.
|	       ..|.............| |.
|				.....
|				.....
|-2				--C--
|				.....
|				.....
*/

func Test4PlayerFullGame(t *testing.T) {
	// create game
	minitileSet := end_tests.MiniTileSetRoadsAndFields()
	deckStack := stack.NewSeeded(minitileSet.Tiles, 94058654839) // random seed :P
	deck := deck.Deck{Stack: &deckStack, StartingTile: minitileSet.StartingTile}
	game, err := gameMod.NewFromDeck(deck, nil, 4)
	if err != nil {
		t.Fatal(err.Error())
	}

	// for i, _ := range deck.GetTiles() {
	// 	fmt.Printf("%v:\n", i+1)
	// 	var tile, _ = deck.Next()
	// 	fmt.Printf("%v\n", tile)
	// }
	checkFirstTurn(game, t)    // T Cross road
	checkSecondTurn(game, t)   // Two city edges not connected
	checkThirdTurn(game, t)    // Two city edges not connected
	checkFourthTurn(game, t)   // Road turn
	checkFifthTurn(game, t)    // Road Turn
	checkSixthTurn(game, t)    // Straight Road
	checkSeventhTurn(game, t)  // Two city edges not connected
	checkEightthTurn(game, t)  // Road turn
	checkNinethTurn(game, t)   // Straight road with city edge
	checkTenthTurn(game, t)    // Monastery with road
	checkEleventhTurn(game, t) // T cross road
	checkTwelvethTurn(game, t) // Straight road
	checkFinalResult(game, t)

}

/*
// player1 places T Cross road with meeple on a bottom road

|				  0	   1    2    3
|
|
|
|
|1
|
|
|		        .| |...|..
|		        .\ /...|..
|0		        --0----1..
|	            .......!..
|	            .......|..
|
|
|-1
|
|
|
|
|-2
|
|
*/
func checkFirstTurn(game *gameMod.Game, t *testing.T) {

	pos := elements.NewPosition(1, 0)
	end_tests.MakeTurn(game, t, pos, 1, elements.NormalMeeple, side.Bottom, feature.Road)

	end_tests.VerifyMeepleExistence(t, game, pos, side.Bottom, feature.Road, true)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{0, 0, 0, 0}, []uint8{6, 7, 7, 7})
}

/*
// player2 places Two city edges not connected, and a meeple on a closing city

|				  0	   1    2    3
|
|
|               .| |.
|				.\ /.
|1				..2..
|				./ \.
|				.| |.
|		        .| |...|..
|		        .\ /...|..
|0		        --0----1..
|	            .......!..
|	            .......|..
|
|
|-1
|
|
|
|
|-2
|
|
*/
func checkSecondTurn(game *gameMod.Game, t *testing.T) {
	pos := elements.NewPosition(0, 1)
	end_tests.MakeTurn(game, t, pos, 0, elements.NormalMeeple, side.Bottom, feature.City)

	end_tests.VerifyMeepleExistence(t, game, pos, side.Bottom, feature.City, false)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{0, 4, 0, 0}, []uint8{6, 7, 7, 7})
}

/*
// Player3 places Two city edges not connected and a meeple on bottom city

|				  0	   1    2    3
|
|
|               .| |.
|				.\ /.
|1				..2..
|				./ \.
|				.| |.
|		        .| |...|...| |.
|		        .\ /...|...\ /.
|0		        --0----1....3..
|	            .......!.../ \.
|	            .......|...|#|.
|
|
|-1
|
|
|
|
|-2
|
|
*/
func checkThirdTurn(game *gameMod.Game, t *testing.T) {
	pos := elements.NewPosition(2, 0)
	end_tests.MakeTurn(game, t, pos, 0, elements.NormalMeeple, side.Bottom, feature.City)

	end_tests.VerifyMeepleExistence(t, game, pos, side.Bottom, feature.City, true)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{0, 4, 0, 0}, []uint8{6, 7, 6, 7})

}

/*
// Player4 places Road turn with a meeple on a road

|				  0	   1    2    3
|
|
|               .| |......
|				.\ /......
|1				..2....4-$
|				./ \...|..
|				.| |...|..
|		        .| |...|...| |.
|		        .\ /...|...\ /.
|0		        --0----1....3..
|	            .......!.../ \.
|	            .......|...|#|.
|
|
|-1
|
|
|
|
|-2
|
|
*/
func checkFourthTurn(game *gameMod.Game, t *testing.T) {
	pos := elements.NewPosition(1, 1)
	end_tests.MakeTurn(game, t, pos, 3, elements.NormalMeeple, side.Right, feature.Road)

	end_tests.VerifyMeepleExistence(t, game, pos, side.Right, feature.Road, true)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{0, 4, 0, 0}, []uint8{6, 7, 6, 6})
}

/*
// Player1 places Road turn with a farmer on the right bototn

|				  0	   1    2    3
|
|
|               .| |......
|				.\ /......
|1				..2....4-$
|				./ \...|..
|				.| |...|..
|		        .| |...|...| |.
|		        .\ /...|...\ /.
|0		        --0----1....3..
|	            .......!.../ \.
|	            .......|...|#|.
|	                 ..|..
|	                 ..|..
|-1		             --5..
|	                 ...!.
|	                 .....
|
|
|-2
|
|
*/
func checkFifthTurn(game *gameMod.Game, t *testing.T) {
	pos := elements.NewPosition(1, -1)
	end_tests.MakeTurn(game, t, pos, 1, elements.NormalMeeple, side.RightBottomEdge, feature.Field)

	end_tests.VerifyMeepleExistence(t, game, pos, side.BottomRightEdge, feature.Field, true)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{0, 4, 0, 0}, []uint8{5, 7, 6, 6})
}

/*
// Player2 places Straight Road with a meeple on a top field

|				  0	   1    2    3
|
|
|               .| |......
|				.\ /......
|1				..2....4-$
|				./ \...|..
|				.| |...|..
|		        .| |...|...| |.
|		        .\ /...|...\ /.
|0		        --0----1....3..
|	            .......!.../ \.
|	            .......|...|#|.
|	            ..@....|..
|	            .......|..
|-1		        --6----5..
|	            ........!.
|	            ..........
|
|
|-2
|
|
*/
func checkSixthTurn(game *gameMod.Game, t *testing.T) {
	pos := elements.NewPosition(0, -1)
	end_tests.MakeTurn(game, t, pos, 0, elements.NormalMeeple, side.Top, feature.Field)

	end_tests.VerifyMeepleExistence(t, game, pos, side.Top, feature.Field, true)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{0, 4, 0, 0}, []uint8{5, 6, 6, 6})
}

/*
// Player3 places Two city edges not connected, finishing own city, and placing a meeple on a new one

|				  0	   1    2    3
|
|
|               .| |......
|				.\ /......
|1				..2....4-$
|				./ \...|..
|				.| |...|..
|		        .| |...|...| |.
|		        .\ /...|...\ /.
|0		        --0----1....3..
|	            .......!.../ \.
|	            .......|...| |.
|	            ..@....|...| |.
|	            .......|...\ /
|-1		        --6----5....7..
|	            ........!../ \.
|	            ...........|#|.
|
|
|-2
|
|
*/

func checkSeventhTurn(game *gameMod.Game, t *testing.T) {
	pos := elements.NewPosition(2, -1)
	end_tests.MakeTurn(game, t, pos, 0, elements.NormalMeeple, side.Bottom, feature.City)

	end_tests.VerifyMeepleExistence(t, game, elements.NewPosition(2, 0), side.Bottom, feature.City, false) // removed meeple
	end_tests.VerifyMeepleExistence(t, game, pos, side.Bottom, feature.City, true)                         // new meeple
	end_tests.CheckMeeplesAndScore(game, t, []uint32{0, 4, 4, 0}, []uint8{5, 6, 6, 6})
}

/*
// Player4 places Road turn and a meeple on a road

|				  0	   1    2    3
|
|
|               .| |......
|				.\ /......
|1				..2....4-$
|				./ \...|..
|				.| |...|..
|		   ......| |...|...| |.
|		   ......\ /...|...\ /.
|0		   ..8----0----1....3..
|	       ..|.........!.../ \.
|	       ..$.........|...| |.
|	            ..@....|...| |.
|	            .......|...\ /
|-1		        --6----5....7..
|	            ........!../ \.
|	            ...........|#|.
|
|
|-2
|
|
*/
func checkEightthTurn(game *gameMod.Game, t *testing.T) {
	pos := elements.NewPosition(-1, 0)
	end_tests.MakeTurn(game, t, pos, 3, elements.NormalMeeple, side.Bottom, feature.Road)

	end_tests.VerifyMeepleExistence(t, game, pos, side.Bottom, feature.Road, true)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{0, 4, 4, 0}, []uint8{5, 6, 6, 5})
}

/*
// Player1 places Straight road with city edge and places a meeple on the upper field.
No one scores for the finished city

|				  0	   1    2    3
|
|
|               .| |........!..
|				.\ /...........
|1				..2....4-$--9--
|				./ \...|.../ \.
|				.| |...|...| |.
|		   ......| |...|...| |.
|		   ......\ /...|...\ /.
|0		   ..8----0----1....3..
|	       ..|.........!.../ \.
|	       ..$.........|...| |.
|	            ..@....|...| |.
|	            .......|...\ /
|-1		        --6----5....7..
|	            ........!../ \.
|	            ...........|#|.
|
|
|-2
|
|
*/
func checkNinethTurn(game *gameMod.Game, t *testing.T) {
	pos := elements.NewPosition(2, 1)
	end_tests.MakeTurn(game, t, pos, 2, elements.NormalMeeple, side.Top, feature.Field)

	end_tests.VerifyMeepleExistence(t, game, pos, side.Top, feature.Field, true)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{0, 4, 4, 0}, []uint8{4, 6, 6, 5})
}

/*
// Player2 places Monastery with road, with a meeple on a monastery.
Player4 scores 3 points for finished road

|				  0	   1    2    3
|
|
|               .| |........!.......
|				.\ /............[ ].
|1				..2....4----9---[A].
|				./ \...|.../ \..[@].
|				.| |...|...| |......
|		   ......| |...|...| |.
|		   ......\ /...|...\ /.
|0		   ..8----0----1....3..
|	       ..|.........!.../ \.
|	       ..$.........|...| |.
|	            ..@....|...| |.
|	            .......|...\ /
|-1		        --6----5....7..
|	            ........!../ \.
|	            ...........|#|.
|
|
|-2
|
|
*/
func checkTenthTurn(game *gameMod.Game, t *testing.T) {
	pos := elements.NewPosition(3, 1)
	end_tests.MakeTurn(game, t, pos, 1, elements.NormalMeeple, side.None, feature.Monastery)

	end_tests.VerifyMeepleExistence(t, game, elements.NewPosition(1, 1), side.Right, feature.Road, false) // removed meeple
	end_tests.VerifyMeepleExistence(t, game, pos, side.None, feature.Field, true)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{0, 4, 4, 3}, []uint8{4, 5, 6, 6})
}

/*
// Player3 places T cross road, with a meeple on a bottom road
player1 and player4 score 4 points for their roads

|				  0	   1    2    3
|
|
|               .| |........!.......
|				.\ /............[ ].
|1				..2....4----9---[A].
|				./ \...|.../ \..[@].
|				.| |...|...| |......
|		   ......| |...|...| |.
|		   ......\ /...|...\ /.
|0		   ..8----0----1....3..
|	       ..|.........|.../ \.
|	       ..|.........|...| |.
|	       ..|....@....|...| |.
|	       ..|.........|...\ /
|-1		   ..B----6----5....7..
|	       ..|..........!../ \.
|	       ..#.............|#|.
|
|
|-2
|
|
*/
func checkEleventhTurn(game *gameMod.Game, t *testing.T) {
	pos := elements.NewPosition(3, 1)
	end_tests.MakeTurn(game, t, pos, 3, elements.NormalMeeple, side.None, feature.Monastery)

	end_tests.VerifyMeepleExistence(t, game, elements.NewPosition(-1, 0), side.Bottom, feature.Road, false) // removed meeple
	end_tests.VerifyMeepleExistence(t, game, elements.NewPosition(0, 1), side.Bottom, feature.Road, false)  // removed meeple
	end_tests.VerifyMeepleExistence(t, game, pos, side.Bottom, feature.City, true)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{4, 4, 4, 7}, []uint8{5, 5, 5, 7})
}

/*
// Player4 places Straight road with a meeple on a road

|				  0	   1    2    3
|
|
|               .| |........!.......
|				.\ /............[ ].
|1				..2....4----9---[A].
|				./ \...|.../ \..[@].
|				.| |...|...| |......
|		   ......| |...|...| |.
|		   ......\ /...|...\ /.
|0		   ..8----0----1....3..
|	       ..|.........|.../ \.
|	       ..|.........|...| |.
|	       ..|....@....|...| |.
|	       ..|.........|...\ /
|-1		   ..B----6----5....7..
|	       ..|..........!../ \.
|	       ..#.............|#|.
|				.....
|				.....
|-2				--C-$
|				.....
|				.....
*/

func checkTwelvethTurn(game *gameMod.Game, t *testing.T) {
	pos := elements.NewPosition(0, -2)
	end_tests.MakeTurn(game, t, pos, 0, elements.NormalMeeple, side.Right, feature.Road)

	end_tests.VerifyMeepleExistence(t, game, pos, side.Right, feature.Road, true)
	end_tests.CheckMeeplesAndScore(game, t, []uint32{4, 4, 4, 7}, []uint8{5, 5, 5, 6})
}

/*
player 1 scores additional 15 points:
  - 9 points for upper farmer
  - 6 points for lower farmer

player 2 scores additional 2 points:
  - 2 points for a monastery
  - 0 for farmer in the center

player 3 scores additional 2 points:
  - 1 point for an unfished city
  - 1 point for an unfished road

player 4 scores additional 1 point:
  - 1 point for an unfished road
*/
func checkFinalResult(game *gameMod.Game, t *testing.T) {
	var scores, err = game.Finalize()
	if err != nil {
		t.Fatal(err.Error())
	}
	var expectedScores = []uint32{4 + 15, 4 + 2, 4 + 2, 7 + 1}

	for i := range 4 {
		if scores[i] != expectedScores[i] {
			t.Fatalf("Player %d final score incorrect. Expected %d, got: %d", i+1, expectedScores[i], scores[1])
		}
	}
}
