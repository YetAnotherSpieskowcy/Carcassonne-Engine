//go:build test

package four_player_game_test

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	gameMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
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
	minitileSet := CreateTileSet()
	deckStack := stack.NewOrdered(minitileSet.Tiles)
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

func CreateTileSet() tilesets.TileSet { //nolint:gocyclo // shallow loops for adding tiles
	var tiles []tiles.Tile
	// mini simple set containing (12 tiles in total):
	// 1 monastery with road
	// 2 straight roads
	// 1 straight road with city
	// 3 road turns
	// 2 T crossroads
	// 3 city edges up and down not connected

	tiles = append(tiles, tiletemplates.TCrossRoad().Rotate(1)) // 1 turn
	tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownNotConnected())
	tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownNotConnected())
	tiles = append(tiles, tiletemplates.RoadsTurn().Rotate(3))
	tiles = append(tiles, tiletemplates.RoadsTurn().Rotate(1)) // 5 turn
	tiles = append(tiles, tiletemplates.StraightRoads())
	tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownNotConnected())
	tiles = append(tiles, tiletemplates.RoadsTurn().Rotate(3))
	tiles = append(tiles, tiletemplates.SingleCityEdgeStraightRoads().Rotate(2))
	tiles = append(tiles, tiletemplates.MonasteryWithSingleRoad().Rotate(1)) // 10 turn
	tiles = append(tiles, tiletemplates.TCrossRoad().Rotate(3))
	tiles = append(tiles, tiletemplates.StraightRoads())

	return tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}
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

	pos := position.New(1, 0)
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road}, true, 1)

	gameMod.VerifyMeepleExistence(t, game, pos, side.Bottom, feature.Road, true, 1)
	gameMod.CheckMeeplesAndScore(game, t, []uint32{0, 0, 0, 0}, []uint8{6, 7, 7, 7}, 1)
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
	pos := position.New(0, 1)
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.City}, true, 2)

	gameMod.VerifyMeepleExistence(t, game, pos, side.Bottom, feature.City, false, 2)
	gameMod.CheckMeeplesAndScore(game, t, []uint32{0, 4, 0, 0}, []uint8{6, 7, 7, 7}, 2)
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
	pos := position.New(2, 0)
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.City}, true, 3)

	gameMod.VerifyMeepleExistence(t, game, pos, side.Bottom, feature.City, true, 3)
	gameMod.CheckMeeplesAndScore(game, t, []uint32{0, 4, 0, 0}, []uint8{6, 7, 6, 7}, 3)

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
	pos := position.New(1, 1)
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Right, FeatureType: feature.Road}, true, 4)

	gameMod.VerifyMeepleExistence(t, game, pos, side.Right, feature.Road, true, 4)
	gameMod.CheckMeeplesAndScore(game, t, []uint32{0, 4, 0, 0}, []uint8{6, 7, 6, 6}, 4)
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
	pos := position.New(1, -1)
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.RightBottomEdge, FeatureType: feature.Field}, true, 5)

	gameMod.VerifyMeepleExistence(t, game, pos, side.BottomRightEdge, feature.Field, true, 5)
	gameMod.CheckMeeplesAndScore(game, t, []uint32{0, 4, 0, 0}, []uint8{5, 7, 6, 6}, 5)
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
	pos := position.New(0, -1)

	// try illegal turn first (put meeple on a field)
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Field}, false, 6)

	// normal correct turn
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.Field}, true, 6)

	gameMod.VerifyMeepleExistence(t, game, pos, side.Top, feature.Field, true, 6)
	gameMod.CheckMeeplesAndScore(game, t, []uint32{0, 4, 0, 0}, []uint8{5, 6, 6, 6}, 6)
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
	pos := position.New(2, -1)

	// try illegal turn first (put meeple on a city)
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.City}, false, 7)

	// normal correct turn
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.City}, true, 7)

	gameMod.VerifyMeepleExistence(t, game, position.New(2, 0), side.Bottom, feature.City, false, 7) // removed meeple
	gameMod.VerifyMeepleExistence(t, game, pos, side.Bottom, feature.City, true, 7)                 // new meeple
	gameMod.CheckMeeplesAndScore(game, t, []uint32{0, 4, 4, 0}, []uint8{5, 6, 6, 6}, 7)
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
	pos := position.New(-1, 0)

	// try illegal turn first (put meeple on a field)
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.BottomRightEdge, FeatureType: feature.Field}, false, 8)

	// normal correct turn
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road}, true, 8)

	gameMod.VerifyMeepleExistence(t, game, pos, side.Bottom, feature.Road, true, 8)
	gameMod.CheckMeeplesAndScore(game, t, []uint32{0, 4, 4, 0}, []uint8{5, 6, 6, 5}, 8)
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
	pos := position.New(2, 1)

	// try illegal turn first (put meeple on a road)
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Left, FeatureType: feature.Road}, false, 9)

	// normal correct turn
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.Field}, true, 9)

	gameMod.VerifyMeepleExistence(t, game, pos, side.Top, feature.Field, true, 9)
	gameMod.CheckMeeplesAndScore(game, t, []uint32{0, 4, 4, 0}, []uint8{4, 6, 6, 5}, 9)
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
	pos := position.New(3, 1)

	// try illegal turn first (put meeple on a field)
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Field}, false, 10)

	// normal correct turn
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.NoSide, FeatureType: feature.Monastery}, true, 10)

	gameMod.VerifyMeepleExistence(t, game, position.New(1, 1), side.Right, feature.Road, false, 10) // removed meeple
	gameMod.VerifyMeepleExistence(t, game, pos, side.NoSide, feature.Monastery, true, 10)
	gameMod.CheckMeeplesAndScore(game, t, []uint32{0, 4, 4, 4}, []uint8{4, 5, 6, 6}, 10)
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
	pos := position.New(-1, -1)

	// try illegal turn first (put meeple on a field)
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.TopRightEdge, FeatureType: feature.Field}, false, 11)

	// normal correct turn
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road}, true, 11)

	gameMod.VerifyMeepleExistence(t, game, position.New(-1, 0), side.Bottom, feature.Road, false, 11) // removed meeple
	gameMod.VerifyMeepleExistence(t, game, position.New(1, 0), side.Bottom, feature.Road, false, 11)  // removed meeple
	gameMod.VerifyMeepleExistence(t, game, pos, side.Bottom, feature.Road, true, 11)
	gameMod.CheckMeeplesAndScore(game, t, []uint32{4, 4, 4, 8}, []uint8{5, 5, 5, 7}, 11)
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
	pos := position.New(0, -2)

	// try illegal turn first (put meeple on a field)
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.Field}, false, 12)

	// normal correct turn
	gameMod.MakeTurnValidCheck(game, t, pos, gameMod.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Right, FeatureType: feature.Road}, true, 12)

	gameMod.VerifyMeepleExistence(t, game, pos, side.Right, feature.Road, true, 12)
	gameMod.CheckMeeplesAndScore(game, t, []uint32{4, 4, 4, 8}, []uint8{5, 5, 5, 6}, 12)
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
	var expectedScores = []uint32{4 + 15, 4 + 2, 4 + 2, 8 + 1}

	for i := range 4 {
		if scores.ReceivedPoints[elements.ID(i)] != expectedScores[i] {
			t.Fatalf("Player %d final score incorrect. Expected %d, got: %d", i+1, expectedScores[i], scores.ReceivedPoints[1])
		}
	}
}
