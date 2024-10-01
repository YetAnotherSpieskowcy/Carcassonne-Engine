package two_player_game_test

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	gameMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/test"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
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
	minitileSet := tilesets.OrderedMiniTileSet1()

	deckStack := stack.NewOrdered(minitileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: minitileSet.StartingTile}
	game, err := gameMod.NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

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
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.City},
		TurnNumber:   1,
	}.Run()

	// remoed meeple
	test.VerifyMeepleExistence{
		TestingT:    t,
		Game:        game,
		Position:    pos,
		Side:        side.Bottom,
		FeatureType: feature.City,
		MeepleExist: false,
		TurnNumber:  1,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{4, 0},
		PlayerMeeples: []uint8{7, 7},
		TurnNumber:    1,
	}.Run()
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
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Left, FeatureType: feature.Road},
		TurnNumber:   2,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:    t,
		Game:        game,
		Position:    pos,
		Side:        side.Left,
		FeatureType: feature.Road,
		MeepleExist: true,
		TurnNumber:  2,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{4, 0},
		PlayerMeeples: []uint8{7, 6},
		TurnNumber:    2,
	}.Run()
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

	// try illegal turn first (put meeple on a road)
	test.MakeWrongTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.Road},
		TurnNumber:   3,
	}.Run()

	// normal correct turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.TopLeftEdge, FeatureType: feature.Field},
		TurnNumber:   3,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:    t,
		Game:        game,
		Position:    pos,
		Side:        side.TopLeftEdge,
		FeatureType: feature.Field,
		MeepleExist: true,
		TurnNumber:  3,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{4, 0},
		PlayerMeeples: []uint8{6, 6},
		TurnNumber:    3,
	}.Run()
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

	// try illegal turn first (put meeple on a road)
	test.MakeWrongTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Right, FeatureType: feature.Road},
		TurnNumber:   4,
	}.Run()

	// normal correct turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road},
		TurnNumber:   4,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:    t,
		Game:        game,
		Position:    pos,
		Side:        side.Bottom,
		FeatureType: feature.Road,
		MeepleExist: true,
		TurnNumber:  4,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{4, 0},
		PlayerMeeples: []uint8{6, 5},
		TurnNumber:    4,
	}.Run()
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

	// try illegal turn first (put meeple on a road)
	test.MakeWrongTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.Road},
		TurnNumber:   5,
	}.Run()

	// normal correct turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.NoSide, FeatureType: feature.Monastery},
		TurnNumber:   5,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:    t,
		Game:        game,
		Position:    position.New(-1, 0),
		Side:        side.Bottom,
		FeatureType: feature.Road,
		MeepleExist: false,
		TurnNumber:  5,
	}.Run()
	test.VerifyMeepleExistence{
		TestingT:    t,
		Game:        game,
		Position:    pos,
		Side:        side.NoSide,
		FeatureType: feature.Monastery,
		MeepleExist: true,
		TurnNumber:  5,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{4, 2},
		PlayerMeeples: []uint8{5, 6},
		TurnNumber:    5,
	}.Run()

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
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Right, FeatureType: feature.City},
		TurnNumber:   6,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:    t,
		Game:        game,
		Position:    pos,
		Side:        side.Right,
		FeatureType: feature.City,
		MeepleExist: true,
		TurnNumber:  6,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{4, 2},
		PlayerMeeples: []uint8{5, 5},
		TurnNumber:    6,
	}.Run()
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

	// try illegal turn first (put meeple on a city)
	test.MakeWrongTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Left, FeatureType: feature.City},
		TurnNumber:   7,
	}.Run()

	// normal correct turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Right, FeatureType: feature.City},
		TurnNumber:   7,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:    t,
		Game:        game,
		Position:    position.New(1, -1),
		Side:        side.Right,
		FeatureType: feature.City,
		MeepleExist: false,
		TurnNumber:  7,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:    t,
		Game:        game,
		Position:    pos,
		Side:        side.Right,
		FeatureType: feature.City,
		MeepleExist: true,
		TurnNumber:  7,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{4, 6},
		PlayerMeeples: []uint8{4, 6},
		TurnNumber:    7,
	}.Run()
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
	pos := position.New(2, 1)

	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Road},
		TurnNumber:   8,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:    t,
		Game:        game,
		Position:    pos,
		Side:        side.Bottom,
		FeatureType: feature.Road,
		MeepleExist: true,
		TurnNumber:  8,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{4, 6},
		PlayerMeeples: []uint8{4, 5},
		TurnNumber:    8,
	}.Run()
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

	// try illegal turn first (put meeple on a road)
	test.MakeWrongTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Right, FeatureType: feature.Road},
		TurnNumber:   9,
	}.Run()

	// normal correct turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.TopRightEdge, FeatureType: feature.Field},
		TurnNumber:   9,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:    t,
		Game:        game,
		Position:    position.New(1, 1),
		Side:        side.Left,
		FeatureType: feature.Road,
		MeepleExist: false,
		TurnNumber:  9,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:    t,
		Game:        game,
		Position:    pos,
		Side:        side.TopRightEdge,
		FeatureType: feature.Field,
		MeepleExist: true,
		TurnNumber:  9,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{4, 12},
		PlayerMeeples: []uint8{3, 6},
		TurnNumber:    9,
	}.Run()
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
|	.|.    .  .  .
|-1	.5.   >6<>7<>A<M
| 	...    .  .  .
*/
func checkTenthTurn(game *gameMod.Game, t *testing.T) {
	pos := position.New(3, -1)

	// normal correct turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Right, FeatureType: feature.City},
		TurnNumber:   10,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:    t,
		Game:        game,
		Position:    position.New(2, -1),
		Side:        side.Right,
		FeatureType: feature.City,
		MeepleExist: false,
		TurnNumber:  10,
	}.Run()

	test.VerifyMeepleExistence{
		TestingT:    t,
		Game:        game,
		Position:    pos,
		Side:        side.Right,
		FeatureType: feature.City,
		MeepleExist: true,
		TurnNumber:  10,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{8, 12},
		PlayerMeeples: []uint8{4, 5},
		TurnNumber:    10,
	}.Run()
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

	// try illegal turn first (put meeple on a field)
	test.MakeWrongTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Field},
		TurnNumber:   11,
	}.Run()

	// normal correct turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.NoneMeeple(),
		TurnNumber:   11,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{8, 12},
		PlayerMeeples: []uint8{4, 5},
		TurnNumber:    11,
	}.Run()
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

	// try illegal turn first (put meeple on a field)
	test.MakeWrongTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Bottom, FeatureType: feature.Field},
		TurnNumber:   12,
	}.Run()

	// normal correct turn
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		TilePosition: pos,
		MeepleParams: test.NoneMeeple(),
		TurnNumber:   12,
	}.Run()

	test.CheckMeeplesAndScore{
		Game:          game,
		TestingT:      t,
		PlayerScores:  []uint32{8, 12},
		PlayerMeeples: []uint8{4, 5},
		TurnNumber:    12,
	}.Run()
}

/* player 1 scores additional 12 points:
	- 3 points for monastery
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
	score := uint32(8 + 12)
	if scores.ReceivedPoints[1] != score {
		t.Fatalf("Player 1 final score incorrect. Expected %d, got: %d", score, scores.ReceivedPoints[1])
	}

	score = 16
	if scores.ReceivedPoints[2] != 16 {
		t.Fatalf("Player 2 final score incorrect. Expected %d, got: %d", score, scores.ReceivedPoints[2])
	}
}
