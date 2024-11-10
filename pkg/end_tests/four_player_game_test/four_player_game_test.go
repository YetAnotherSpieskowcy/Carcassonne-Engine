package four_player_game_test

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	gameMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/test"
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
|               |   |...............
|				.\ /................
|1				..2....4----9---[A].
|				./ \...|.../ \......
|				|   |..|..|   |.....
|		   .....|   |..|..|   |
|		   ......\ /...|...\ /.
|0		   ..8----0----1....3..
|	       ..|.........|.../ \.
|	       ..|.........|..|   |
|	       ..|.........|..|   |
|	       ..|.........|...\ /
|-1		   ..B----6----5....7..
|	       ..|............./ \.
|	       ..|............|   |
|				.....
|				.....
|-2				--C--
|				.....
|				.....
*/

func Test4PlayerFullGame(t *testing.T) {
	t.Skip("four player games are not currently supported")

	// create game
	minitileSet := MiniTileSet()
	deckStack := stack.NewOrdered(minitileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: minitileSet.StartingTile}
	game, err := gameMod.NewFromDeck(deck, nil, 4)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkFirstTurn(game, t)    // T Cross road
	checkSecondTurn(game, t)   // Two city edges not connected
	checkThirdTurn(game, t)    // Two city edges not connected
	checkFourthTurn(game, t)   // Road turn
	checkFifthTurn(game, t)    // Road Turn
	checkSixthTurn(game, t)    // Straight Road
	checkSeventhTurn(game, t)  // Two city edges not connected
	checkEighthTurn(game, t)   // Road turn
	checkNinthTurn(game, t)    // Straight road with city edge
	checkTenthTurn(game, t)    // Monastery with road
	checkEleventhTurn(game, t) // T cross road
	checkTwelfthTurn(game, t)  // Straight road
	checkFinalResult(game, t)

}

func MiniTileSet() tilesets.TileSet { //nolint:gocyclo
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
|		        |   |..|..
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
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road},
		TurnNumber:   1,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:     t,
		Game:         game,
		Position:     pos,
		Side:         side.Bottom,
		FeatureType:  feature.Road,
		MeepleExists: true,
		TurnNumber:   1,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{0, 0, 0, 0},
		PlayerMeeples: []uint8{7 - 1, 7, 7, 7},
		TurnNumber:    1,
	}.Run()

	test.CheckMidGameScore{
		Game:         game,
		TestingT:     t,
		PlayerScores: []uint32{0 + 1, 0, 0, 0},
		TurnNumber:   1,
	}.Run()
}

/*
// player2 places Two city edges not connected, and a meeple on a closing city

|				  0	   1    2    3
|
|
|               |   |
|				.\ /.
|1				..2..
|				./ \.
|				|   |
|		        |   |..|..
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
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.City},
		TurnNumber:   2,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:     t,
		Game:         game,
		Position:     pos,
		Side:         side.Bottom,
		FeatureType:  feature.City,
		MeepleExists: false,
		TurnNumber:   2,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{0, 0 + 4, 0, 0},
		PlayerMeeples: []uint8{6, 7 - 1 + 1, 7, 7},
		TurnNumber:    2,
	}.Run()

	test.CheckMidGameScore{
		Game:         game,
		TestingT:     t,
		PlayerScores: []uint32{0 + 1, 4, 0, 0},
		TurnNumber:   2,
	}.Run()
}

/*
// Player3 places Two city edges not connected and a meeple on bottom city

|				  0	   1    2    3
|
|
|               |   |
|				.\ /.
|1				..2..
|				./ \.
|				|   |
|		        |   |..|..|   |
|		        .\ /...|...\ /.
|0		        --0----1....3..
|	            .......!.../ \.
|	            .......|..| # |
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
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.City},
		TurnNumber:   3,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:     t,
		Game:         game,
		Position:     pos,
		Side:         side.Bottom,
		FeatureType:  feature.City,
		MeepleExists: true,
		TurnNumber:   3,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{0, 4, 0, 0},
		PlayerMeeples: []uint8{6, 7, 7 - 1, 7},
		TurnNumber:    3,
	}.Run()

	test.CheckMidGameScore{
		Game:         game,
		TestingT:     t,
		PlayerScores: []uint32{0 + 1, 4, 0 + 1, 0},
		TurnNumber:   3,
	}.Run()

}

/*
// Player4 places Road turn with a meeple on a road

|				  0	   1    2    3
|
|
|               |   |.....
|				.\ /......
|1				..2....4-$
|				./ \...|..
|				|   |..|..
|		        |   |..|..|   |
|		        .\ /...|...\ /.
|0		        --0----1....3..
|	            .......!.../ \.
|	            .......|..| # |
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
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Right, FeatureType: feature.Road},
		TurnNumber:   4,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:     t,
		Game:         game,
		Position:     pos,
		Side:         side.Right,
		FeatureType:  feature.Road,
		MeepleExists: true,
		TurnNumber:   4,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{0, 4, 0, 0},
		PlayerMeeples: []uint8{6, 7, 6, 7 - 1},
		TurnNumber:    4,
	}.Run()

	test.CheckMidGameScore{
		Game:         game,
		TestingT:     t,
		PlayerScores: []uint32{0 + 1, 4, 0 + 1, 0 + 2},
		TurnNumber:   4,
	}.Run()
}

/*
// Player1 places Road turn with a farmer on the right bottom field

|				  0	   1    2    3
|
|
|               |   |.....
|				.\ /......
|1				..2....4-$
|				./ \...|..
|				|   |..|..
|		        |   |..|..|   |
|		        .\ /...|...\ /.
|0		        --0----1....3..
|	            .......!.../ \.
|	            .......|..| # |
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
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.RightBottomEdge, FeatureType: feature.Field},
		TurnNumber:   5,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:     t,
		Game:         game,
		Position:     pos,
		Side:         side.BottomRightEdge,
		FeatureType:  feature.Field,
		MeepleExists: true,
		TurnNumber:   5,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{0, 4, 0, 0},
		PlayerMeeples: []uint8{6 - 1, 7, 6, 6},
		TurnNumber:    5,
	}.Run()

	test.CheckMidGameScore{
		Game:         game,
		TestingT:     t,
		PlayerScores: []uint32{0 + 2, 4, 0 + 1, 0 + 2},
		TurnNumber:   5,
	}.Run()
}

/*
// Player2 places Straight Road with a meeple on a top field

|				  0	   1    2    3
|
|
|               |   |.....
|				.\ /......
|1				..2....4-$
|				./ \...|..
|				|   |..|..
|		        |   |..|..|   |
|		        .\ /...|...\ /.
|0		        --0----1....3..
|	            .......!.../ \.
|	            .......|..| # |
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
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Field},
		TurnNumber:   6,
		WrongTurn:    true,
	}.Run()

	// normal correct turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.Field},
		TurnNumber:   6,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:     t,
		Game:         game,
		Position:     pos,
		Side:         side.Top,
		FeatureType:  feature.Field,
		MeepleExists: true,
		TurnNumber:   6,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{0, 4, 0, 0},
		PlayerMeeples: []uint8{5, 7 - 1, 6, 6},
		TurnNumber:    6,
	}.Run()

	test.CheckMidGameScore{
		Game:         game,
		TestingT:     t,
		PlayerScores: []uint32{0 + 3, 4, 0 + 1, 0 + 2},
		TurnNumber:   6,
	}.Run()
}

/*
// Player3 places Two city edges not connected, finishing own city, and placing a meeple on a new one

|				  0	   1    2    3
|
|
|               |   |.....
|				.\ /......
|1				..2....4-$
|				./ \...|..
|				|   |..|..
|		        |   |..|..|   |
|		        .\ /...|...\ /.
|0		        --0----1....3..
|	            .......!.../ \.
|	            .......|..|   |
|	            ..@....|..|   |
|	            .......|...\ /
|-1		        --6----5....7..
|	            ........!../ \.
|	            ..........| # |
|
|
|-2
|
|
*/

func checkSeventhTurn(game *gameMod.Game, t *testing.T) {
	pos := position.New(2, -1)

	// try illegal turn first (put meeple on a city)
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.City},
		TurnNumber:   7,
		WrongTurn:    true,
	}.Run()

	// normal correct turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.City},
		TurnNumber:   7,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:     t,
		Game:         game,
		Position:     position.New(2, 0),
		Side:         side.Bottom,
		FeatureType:  feature.City,
		MeepleExists: false,
		TurnNumber:   7,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:     t,
		Game:         game,
		Position:     pos,
		Side:         side.Bottom,
		FeatureType:  feature.City,
		MeepleExists: true,
		TurnNumber:   7,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{0, 4, 0 + 4, 0},
		PlayerMeeples: []uint8{5, 6, 6 - 1 + 1, 6},
		TurnNumber:    7,
	}.Run()

	test.CheckMidGameScore{
		Game:         game,
		TestingT:     t,
		PlayerScores: []uint32{0 + 3 + 3, 4, 4 + 1, 0 + 2},
		TurnNumber:   7,
	}.Run()
}

/*
// Player4 places Road turn and a meeple on a road

|				  0	   1    2    3
|
|
|               |   |.....
|				.\ /......
|1				..2....4-$
|				./ \...|..
|				|   |..|..
|		   .....|   |..|..|   |
|		   ......\ /...|...\ /.
|0		   ..8----0----1....3..
|	       ..|.........!.../ \.
|	       ..$.........|..|   |
|	            ..@....|..|   |
|	            .......|...\ /
|-1		        --6----5....7..
|	            ........!../ \.
|	            ..........| # |
|
|
|-2
|
|
*/
func checkEighthTurn(game *gameMod.Game, t *testing.T) {
	pos := position.New(-1, 0)

	// try illegal turn first (put meeple on a field)
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.BottomRightEdge, FeatureType: feature.Field},
		TurnNumber:   8,
		WrongTurn:    true,
	}.Run()

	// normal correct turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road},
		TurnNumber:   8,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:     t,
		Game:         game,
		Position:     pos,
		Side:         side.Bottom,
		FeatureType:  feature.Road,
		MeepleExists: true,
		TurnNumber:   8,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{0, 4, 4, 0},
		PlayerMeeples: []uint8{5, 6, 6, 6 - 1},
		TurnNumber:    8,
	}.Run()

	test.CheckMidGameScore{
		Game:         game,
		TestingT:     t,
		PlayerScores: []uint32{0 + 3 + 3, 4, 4 + 1, 0 + 2 + 3},
		TurnNumber:   8,
	}.Run()
}

/*
// Player1 places Straight road with city edge and places a meeple on the upper field.
No one scores for the finished city

|				  0	   1    2    3
|
|
|               |   |.......!..
|				.\ /...........
|1				..2....4-$--9--
|				./ \...|.../ \.
|				|   |..|..|   |
|		   .....|   |..|..|   |
|		   ......\ /...|...\ /.
|0		   ..8----0----1....3..
|	       ..|.........!.../ \.
|	       ..$.........|..|   |
|	            ..@....|..|   |
|	            .......|...\ /
|-1		        --6----5....7..
|	            ........!../ \.
|	            ..........| # |
|
|
|-2
|
|
*/
func checkNinthTurn(game *gameMod.Game, t *testing.T) {
	pos := position.New(2, 1)

	// try illegal turn first (put meeple on a road)
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Left, FeatureType: feature.Road},
		TurnNumber:   9,
		WrongTurn:    true,
	}.Run()

	// normal correct turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.Field},
		TurnNumber:   9,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:     t,
		Game:         game,
		Position:     pos,
		Side:         side.Top,
		FeatureType:  feature.Field,
		MeepleExists: true,
		TurnNumber:   9,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{0, 4, 4, 0},
		PlayerMeeples: []uint8{5 - 1, 6, 6, 5},
		TurnNumber:    9,
	}.Run()

	test.CheckMidGameScore{
		Game:         game,
		TestingT:     t,
		PlayerScores: []uint32{0 + 3 + 3*3, 4, 4 + 1, 0 + 3 + 3},
		TurnNumber:   9,
	}.Run()
}

/*
// Player2 places Monastery with road, with a meeple on a monastery.
Player4 scores 4 points for finished road

|				  0	   1    2    3
|
|
|               |   |.......!.......
|				.\ /............[ ].
|1				..2....4----9---[A].
|				./ \...|.../ \..[@].
|				|   |..|..|   |.....
|		   .....|   |..|..|   |
|		   ......\ /...|...\ /.
|0		   ..8----0----1....3..
|	       ..|.........!.../ \.
|	       ..$.........|..|   |
|	            ..@....|..|   |
|	            .......|...\ /
|-1		        --6----5....7..
|	            ........!../ \.
|	            ..........| # |
|
|
|-2
|
|
*/
func checkTenthTurn(game *gameMod.Game, t *testing.T) {
	pos := position.New(3, 1)

	// try illegal turn first (put meeple on a field)
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Field},
		TurnNumber:   10,
		WrongTurn:    true,
	}.Run()

	// normal correct turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.NoSide, FeatureType: feature.Monastery},
		TurnNumber:   10,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:     t,
		Game:         game,
		Position:     position.New(1, 1),
		Side:         side.Right,
		FeatureType:  feature.Road,
		MeepleExists: false,
		TurnNumber:   10,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:     t,
		Game:         game,
		Position:     pos,
		Side:         side.NoSide,
		FeatureType:  feature.Monastery,
		MeepleExists: true,
		TurnNumber:   10,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{0, 4, 4, 0 + 4},
		PlayerMeeples: []uint8{4, 6 - 1, 6, 5 + 1},
		TurnNumber:    10,
	}.Run()

	test.CheckMidGameScore{
		Game:         game,
		TestingT:     t,
		PlayerScores: []uint32{0 + 3 + 3*3, 4 + 3, 4 + 1, 4 + 3},
		TurnNumber:   10,
	}.Run()
}

/*
// Player3 places T cross road, with a meeple on a bottom road
player1 and player4 score 4 points for their roads

|				  0	   1    2    3
|
|
|               |   |.......!.......
|				.\ /............[ ].
|1				..2....4----9---[A].
|				./ \...|.../ \..[@].
|				|   |..|..|   |.....
|		   .....|   |..|..|   |
|		   ......\ /...|...\ /.
|0		   ..8----0----1....3..
|	       ..|.........|.../ \.
|	       ..|.........|..|   |
|	       ..|....@....|..|   |
|	       ..|.........|...\ /
|-1		   ..B----6----5....7..
|	       ..|..........!../ \.
|	       ..#............| # |
|
|
|-2
|
|
*/
func checkEleventhTurn(game *gameMod.Game, t *testing.T) {
	pos := position.New(-1, -1)

	// try illegal turn first (put meeple on a field)
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.TopLeftEdge, FeatureType: feature.Field},
		TurnNumber:   11,
		WrongTurn:    true,
	}.Run()

	// normal correct turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road},
		TurnNumber:   11,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:     t,
		Game:         game,
		Position:     position.New(-1, 0),
		Side:         side.Bottom,
		FeatureType:  feature.Road,
		MeepleExists: false,
		TurnNumber:   11,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:     t,
		Game:         game,
		Position:     position.New(1, 0),
		Side:         side.Bottom,
		FeatureType:  feature.Road,
		MeepleExists: false,
		TurnNumber:   11,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:     t,
		Game:         game,
		Position:     pos,
		Side:         side.Bottom,
		FeatureType:  feature.Road,
		MeepleExists: true,
		TurnNumber:   11,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{0 + 4, 4, 4, 4 + 4},
		PlayerMeeples: []uint8{4 + 1, 5, 6 - 1, 6 + 1},
		TurnNumber:    11,
	}.Run()

	test.CheckMidGameScore{
		Game:         game,
		TestingT:     t,
		PlayerScores: []uint32{4 + 3*3, 4 + 3, 4 + 1 + 1, 8},
		TurnNumber:   11,
	}.Run()
}

/*
// Player4 places Straight road with a meeple on a road

|				  0	   1    2    3
|
|
|               |   |.......!.......
|				.\ /............[ ].
|1				..2....4----9---[A].
|				./ \...|.../ \..[@].
|				|   |..|..|   |.....
|		   .....|   |..|..|   |
|		   ......\ /...|...\ /.
|0		   ..8----0----1....3..
|	       ..|.........|.../ \.
|	       ..|.........|..|   |
|	       ..|....@....|..|   |
|	       ..|.........|...\ /
|-1		   ..B----6----5....7..
|	       ..|..........!../ \.
|	       ..#............| # |
|				.....
|				.....
|-2				--C-$
|				.....
|				.....
*/

func checkTwelfthTurn(game *gameMod.Game, t *testing.T) {
	pos := position.New(0, -2)

	// try illegal turn first (put meeple on a field)
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.Field},
		TurnNumber:   12,
		WrongTurn:    true,
	}.Run()

	// normal correct turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Right, FeatureType: feature.Road},
		TurnNumber:   12,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:     t,
		Game:         game,
		Position:     pos,
		Side:         side.Right,
		FeatureType:  feature.Road,
		MeepleExists: true,
		TurnNumber:   12,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{4, 4, 4, 8},
		PlayerMeeples: []uint8{5, 5, 5, 7 - 1},
		TurnNumber:    12,
	}.Run()

	test.CheckMidGameScore{
		Game:         game,
		TestingT:     t,
		PlayerScores: []uint32{4 + 3*3, 4 + 3, 4 + 1 + 1, 8 + 1},
		TurnNumber:   12,
	}.Run()
}

/*
player 1 scores additional 9 points:
  - 9 points for both farmers (the same farm)

player 2 scores additional 3 points:
  - 3 points for a monastery
  - 0 for farmer in the center

player 3 scores additional 2 points:
  - 1 point for an unfinished city
  - 1 point for an unfinished road

player 4 scores additional 1 poinTestingT:
  - 1 point for an unfinished road
*/
func checkFinalResult(game *gameMod.Game, t *testing.T) {
	var scores, err = game.Finalize()
	if err != nil {
		t.Fatal(err.Error())
	}
	var expectedScores = []uint32{
		4 + 9,
		4 + 3,
		4 + 2,
		8 + 1,
	}

	for i := range 4 {
		if scores.ReceivedPoints[elements.ID(i+1)] != expectedScores[i] {
			t.Fatalf("Player %d final score incorrect. Expected %d, goTestingT: %d", i+1, expectedScores[i], scores.ReceivedPoints[elements.ID(i+1)])
		}
	}
}
