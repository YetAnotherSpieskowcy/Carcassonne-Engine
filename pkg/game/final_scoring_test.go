package game

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
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
  All tests contains the board visualization. Each tile is represented by 5x5 ascii characters
*/

/*
|                 0    1

|               |   |.....
|               .\ /......
|0              --0----1-!
|               ..........
|               ..........
*/
func TestFinalScoreRoad(t *testing.T) {
	// ------ create tileset --------
	var tiles []tiles.Tile
	var err error
	tiles = append(tiles, tiletemplates.StraightRoads())

	tileset := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}

	// ------ create game --------
	deckStack := stack.NewOrdered(tileset.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileset.StartingTile}

	game, err := NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(1, 0),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Right, FeatureType: feature.Road},
	}.Run()

	scores, err := game.Finalize()
	if err != nil {
		t.Fatal(err.Error())
	}
	var expectedScores = []uint32{2, 0}

	for i := range 2 {
		if scores.ReceivedPoints[elements.ID(i+1)] != expectedScores[i] {
			t.Fatalf("Player %d final score incorrect. Expected %d, got: %d", i+1, expectedScores[i], scores.ReceivedPoints[elements.ID(i+1)])
		}
	}
}

/*
|                 0    1

|               |   |.....
|               .\ /......
|0              --0----1-!
|               ..........
|               ..........
|               ..........
|               ..........
|-1             --2-@--3--
|               ..........
|               ..........
|               .....
|               .....
|-2             --4-@
|               .....
|               .....
*/
func TestFinalScoreMultipleRoads(t *testing.T) {
	// ------ create tileset --------
	var tiles []tiles.Tile
	var err error
	for range 4 {
		tiles = append(tiles, tiletemplates.StraightRoads())
	}

	tileset := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}

	// ------ create game --------
	deckStack := stack.NewOrdered(tileset.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileset.StartingTile}

	game, err := NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(1, 0),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Right, FeatureType: feature.Road},
	}.Run()
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(0, -1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Right, FeatureType: feature.Road},
	}.Run()
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(1, -1),
		MeepleParams: test.NoneMeeple(),
	}.Run()
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(0, -2),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Right, FeatureType: feature.Road},
	}.Run()

	scores, err := game.Finalize()
	if err != nil {
		t.Fatal(err.Error())
	}
	var expectedScores = []uint32{2, 3}

	for i := range 2 {
		if scores.ReceivedPoints[elements.ID(i+1)] != expectedScores[i] {
			t.Fatalf("Player %d final score incorrect. Expected %d, got: %d", i+1, expectedScores[i], scores.ReceivedPoints[elements.ID(i+1)])
		}
	}
}

/*
|                 0    1

|               |   |.....
|               .\ /..[!].
|0              --0---[1].
|               ......[ ].
|               ..........
*/
func TestFinalScoreMonastery(t *testing.T) {
	// ------ create tileset --------
	var tiles []tiles.Tile
	var err error
	tiles = append(tiles, tiletemplates.MonasteryWithSingleRoad().Rotate(1))

	tileset := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}

	// ------ create game --------
	deckStack := stack.NewOrdered(tileset.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileset.StartingTile}

	game, err := NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(1, 0),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.NoSide, FeatureType: feature.Monastery},
	}.Run()

	scores, err := game.Finalize()
	if err != nil {
		t.Fatal(err.Error())
	}
	var expectedScores = []uint32{2, 0}

	for i := range 2 {
		if scores.ReceivedPoints[elements.ID(i+1)] != expectedScores[i] {
			t.Fatalf("Player %d final score incorrect. Expected %d, got: %d", i+1, expectedScores[i], scores.ReceivedPoints[elements.ID(i+1)])
		}
	}
}

/*
|                 0    1

|               |   |.....
|               .\ /..[!].
|0              --0---[1].
|               ......[ ].
|               ..........
|               ..........
|               .[@]..[!].
|-1             .[2]..[3].
|               .[ ]..[ ].
|               ..........
*/
func TestFinalScoreMonasteries(t *testing.T) {
	// ------ create tileset --------
	var tiles []tiles.Tile
	var err error
	tiles = append(tiles, tiletemplates.MonasteryWithSingleRoad().Rotate(1))
	tiles = append(tiles, tiletemplates.MonasteryWithoutRoads())
	tiles = append(tiles, tiletemplates.MonasteryWithoutRoads())

	tileset := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}

	// ------ create game --------
	deckStack := stack.NewOrdered(tileset.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileset.StartingTile}

	game, err := NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(1, 0),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.NoSide, FeatureType: feature.Monastery},
	}.Run()
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(0, -1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.NoSide, FeatureType: feature.Monastery},
	}.Run()
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(1, -1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.NoSide, FeatureType: feature.Monastery},
	}.Run()

	scores, err := game.Finalize()
	if err != nil {
		t.Fatal(err.Error())
	}
	var expectedScores = []uint32{8, 4}

	for i := range 2 {
		if scores.ReceivedPoints[elements.ID(i+1)] != expectedScores[i] {
			t.Fatalf("Player %d final score incorrect. Expected %d, got: %d", i+1, expectedScores[i], scores.ReceivedPoints[elements.ID(i+1)])
		}
	}
}

/*
|                 0    1

|               .| |.
|               ./!\.
|1              | 1 |
|               |   |
|               |   |
|               |   |
|               .\ /.
|0              --0--
|               .....
|               .....
*/
func TestFinalScoreCity(t *testing.T) {
	// ------ create tileset --------
	var tiles []tiles.Tile
	var err error
	tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownConnected())

	tileset := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}

	// ------ create game --------
	deckStack := stack.NewOrdered(tileset.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileset.StartingTile}

	game, err := NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(0, 1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.City},
	}.Run()

	scores, err := game.Finalize()
	if err != nil {
		t.Fatal(err.Error())
	}
	var expectedScores = []uint32{2, 0}

	for i := range 2 {
		if scores.ReceivedPoints[elements.ID(i+1)] != expectedScores[i] {
			t.Fatalf("Player %d final score incorrect. Expected %d, got: %d", i+1, expectedScores[i], scores.ReceivedPoints[elements.ID(i+1)])
		}
	}
}

/*
|                 0    1

|          | ! ||   || @ |
|          .\ /.| ! |.\ /.
|1         ..3..| 1 |..2..
|          .....|   |......
|          .....|   |.....
|               |   |
|               .\ /.
|0              --0--
|               .....
|               .....
*/

func TestFinalScoreCities(t *testing.T) {
	// ------ create tileset --------
	var tiles []tiles.Tile
	var err error
	tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownConnected())
	tiles = append(tiles, tiletemplates.SingleCityEdgeNoRoads())
	tiles = append(tiles, tiletemplates.SingleCityEdgeNoRoads())

	tileset := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}

	// ------ create game --------
	deckStack := stack.NewOrdered(tileset.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileset.StartingTile}

	game, err := NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(0, 1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.City},
	}.Run()
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(1, 1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.City},
	}.Run()
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(-1, 1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.City},
	}.Run()

	scores, err := game.Finalize()
	if err != nil {
		t.Fatal(err.Error())
	}
	var expectedScores = []uint32{3, 1}

	for i := range 2 {
		if scores.ReceivedPoints[elements.ID(i+1)] != expectedScores[i] {
			t.Fatalf("Player %d final score incorrect. Expected %d, got: %d", i+1, expectedScores[i], scores.ReceivedPoints[elements.ID(i+1)])
		}
	}
}

/*
|                 0    1

|               ..!..
|               .....
|1              ..1..
|               ./ \.
|               |   |
|               |   |
|               .\ /.
|0              --0--
|               .....
|               .....
*/
func TestFinalScoreField(t *testing.T) {
	// ------ create tileset --------
	var tiles []tiles.Tile
	var err error
	tiles = append(tiles, tiletemplates.SingleCityEdgeNoRoads().Rotate(2))

	tileset := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}

	// ------ create game --------
	deckStack := stack.NewOrdered(tileset.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileset.StartingTile}

	game, err := NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(0, 1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.Field},
	}.Run()

	scores, err := game.Finalize()
	if err != nil {
		t.Fatal(err.Error())
	}
	var expectedScores = []uint32{3, 0}

	for i := range 2 {
		if scores.ReceivedPoints[elements.ID(i+1)] != expectedScores[i] {
			t.Fatalf("Player %d final score incorrect. Expected %d, got: %d", i+1, expectedScores[i], scores.ReceivedPoints[elements.ID(i+1)])
		}
	}
}

/*
|                 0    1

|               ..!.......
|               ......[ ].
|1              ..1...[5]-
|               ./ \..[ ].
|               |   |.....
|               |   |.....
|               .\ /..[ ].
|0              --0---[4].
|               ......[ ].
|               ..........
|               ..@..
|               .....
|-1             ..2..
|               ./ \.
|               |   |
|               |   |
|               .\ /.
|-2             ..3..
|               .....
|               .....

	Testing multiple meeples on same field
*/
func TestFinalScoreFields1(t *testing.T) {
	// ------ create tileset --------
	var tiles []tiles.Tile
	var err error
	tiles = append(tiles, tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	tiles = append(tiles, tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	tiles = append(tiles, tiletemplates.SingleCityEdgeNoRoads())
	tiles = append(tiles, tiletemplates.MonasteryWithSingleRoad().Rotate(1))
	tiles = append(tiles, tiletemplates.MonasteryWithSingleRoad().Rotate(3))

	tileset := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}

	// ------ create game --------
	deckStack := stack.NewOrdered(tileset.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileset.StartingTile}

	game, err := NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(0, 1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.Field},
	}.Run()
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(0, -1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.Field},
	}.Run()
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(0, -2),
		MeepleParams: test.NoneMeeple(),
	}.Run()
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(1, 0),
		MeepleParams: test.NoneMeeple(),
	}.Run()
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(1, 1),
		MeepleParams: test.NoneMeeple(),
	}.Run()

	scores, err := game.Finalize()
	if err != nil {
		t.Fatal(err.Error())
	}
	var expectedScores = []uint32{6, 6}

	for i := range 2 {
		if scores.ReceivedPoints[elements.ID(i+1)] != expectedScores[i] {
			t.Fatalf("Player %d final score incorrect. Expected %d, got: %d", i+1, expectedScores[i], scores.ReceivedPoints[elements.ID(i+1)])
		}
	}
}

/*
|                 0

|               ...@.
|               .....
|2              ..2..
|               ./ \.
|               |   |
|               |   |
|               .\ /.
|1              ..1.!
|               ./ \.
|               |   |
|               |   |
|               .\ /.
|0              --0--
|               .....
|               .....

	Testing multiple meeples on same field
*/
func TestFinalScoreFields2(t *testing.T) {
	// ------ create tileset --------
	var tiles []tiles.Tile
	var err error
	tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownNotConnected().Rotate(2))
	tiles = append(tiles, tiletemplates.SingleCityEdgeNoRoads().Rotate(2))

	tileset := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}

	// ------ create game --------
	deckStack := stack.NewOrdered(tileset.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileset.StartingTile}

	game, err := NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(0, 1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Right, FeatureType: feature.Field},
	}.Run()

	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(0, 2),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.Field},
	}.Run()

	scores, err := game.Finalize()
	if err != nil {
		t.Fatal(err.Error())
	}
	var expectedScores = []uint32{6, 3}

	for i := range 2 {
		if scores.ReceivedPoints[elements.ID(i+1)] != expectedScores[i] {
			t.Fatalf("Player %d final score incorrect. Expected %d, got: %d", i+1, expectedScores[i], scores.ReceivedPoints[elements.ID(i+1)])
		}
	}
}

/*
|      -2   -1    0    1    2

|     .........................
|     .[ ].................[ ].
|2    .[K]...6....5....7...[H].
|     .[ ]../ \../ \../ \..[ ].
|     .....|   ||   ||   |.....
|     ....-    ||   ||   |-....
|     .../    /..\ /.\     \...
|1    ..G    X....1....X    9..
|     ...\  /...........\  /...
|     ....--.............--....
|     ....--.............--....
|     .../  \....[ ]..../  \...
|0    ..E    4...[0]...2    8..
|     ...\  /....[ ]....\  /...
|     ....--.............--....
|     ....--......!......--....
|     .../  \.........../  \...
|-1   ..F    X....3....X    A..
|     ...\    \../ \../    /...
|     ....-    ||   ||    -....
|     .....|   ||   ||   |...@.
|     .[ ]..\ /..\ /..\ /..[ ].
|-2   .[J]...D....B....C...[I].
|     .[ ].................[ ].
|     .........................

	Field surrounded be cities! 8 cities in total.
	For nice loops starting tile is changed for monastery field.
	One meeple inside(!) and one meeple outside (@)
*/
func TestFinalScoreFields3(t *testing.T) { //nolint:gocyclo
	// ------ create tileset --------
	rotations := []uint{0, 1, 2, 3}
	var tiles []tiles.Tile
	var err error

	// inner single city edges
	for _, rotation := range rotations {
		tiles = append(tiles, tiletemplates.SingleCityEdgeNoRoads().Rotate(rotation))
	}

	// outer city edges
	for _, rotation := range rotations {
		for range 3 {
			tiles = append(tiles, tiletemplates.SingleCityEdgeNoRoads().Rotate(rotation+2))
		}
	}

	// inner corners cities
	for _, rotation := range rotations {
		tiles = append(tiles, tiletemplates.TwoCityEdgesCornerConnected().Rotate(rotation))
	}

	// outer corner monasteries
	for _, rotation := range rotations {
		tiles = append(tiles, tiletemplates.MonasteryWithoutRoads().Rotate(rotation))
	}

	tileset := tilesets.TileSet{
		StartingTile: tiletemplates.MonasteryWithoutRoads(),
		Tiles:        tiles,
	}

	// ------ create game --------
	deckStack := stack.NewOrdered(tileset.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileset.StartingTile}

	game, err := NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	// ------ make turns --------
	offsets := []position.Position{position.New(0, 0), position.New(-1, 0), position.New(1, 0)}

	// start with inner single city edges (starting from top) (tiles 1-4)
	for _, rotation := range rotations {
		if rotation == 2 {
			test.MakeTurn{
				Game:         game,
				TestingT:     t,
				Position:     position.New(0, 1).Rotate(rotation),
				MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.Field},
			}.Run()
		} else {
			test.MakeTurn{
				Game:         game,
				TestingT:     t,
				Position:     position.New(0, 1).Rotate(rotation),
				MeepleParams: test.NoneMeeple(),
			}.Run()
		}
	}

	// Now outer single city edges (starting from top) (tiles 5-G)
	for _, rotation := range rotations {
		for _, offset := range offsets {
			test.MakeTurn{
				Game:         game,
				TestingT:     t,
				Position:     position.New(0, 2).Rotate(rotation).Add(offset.Rotate(rotation)),
				MeepleParams: test.NoneMeeple(),
			}.Run()
		}
	}

	// Inner corners
	for _, rotation := range rotations {
		test.MakeTurn{
			Game:         game,
			TestingT:     t,
			Position:     position.New(1, 1).Rotate(rotation),
			MeepleParams: test.NoneMeeple(),
		}.Run()
	}

	// Outer corners
	for _, rotation := range rotations {
		if rotation == 1 {
			test.MakeTurn{
				Game:         game,
				TestingT:     t,
				Position:     position.New(2, 2).Rotate(rotation),
				MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Top, FeatureType: feature.Field},
			}.Run()

		} else {
			test.MakeTurn{
				Game:         game,
				TestingT:     t,
				Position:     position.New(2, 2).Rotate(rotation),
				MeepleParams: test.NoneMeeple(),
			}.Run()
		}
	}

	// ------ game.Finalize --------
	scores, err := game.Finalize()
	if err != nil {
		t.Fatal(err.Error())
	}
	var expectedScores = []uint32{24, 24}

	for i := range 2 {
		if scores.ReceivedPoints[elements.ID(i+1)] != expectedScores[i] {
			t.Fatalf("Player %d final score incorrect. Expected %d, got: %d", i+1, expectedScores[i], scores.ReceivedPoints[elements.ID(i+1)])
		}
	}
}

func TestFinalScoreMultipleRoads123(t *testing.T) {
	// ------ create tileset --------
	var tiles []tiles.Tile
	var err error
	tiles = append(tiles, tiletemplates.TCrossRoad())
	tiles = append(tiles, tiletemplates.RoadsTurn())
	tiles = append(tiles, tiletemplates.RoadsTurn().Rotate(1))
	tiles = append(tiles, tiletemplates.RoadsTurn().Rotate(2))

	tileset := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}

	// ------ create game --------
	deckStack := stack.NewOrdered(tileset.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileset.StartingTile}

	game, err := NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(0, -1),
		MeepleParams: test.MeepleParams{MeepleType: elements.NormalMeeple, FeatureSide: side.Right, FeatureType: feature.Road},
	}.Run()
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(1, -1),
		MeepleParams: test.NoneMeeple(),
	}.Run()
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(1, -2),
		MeepleParams: test.NoneMeeple(),
	}.Run()
	test.MakeTurn{
		Game:         game,
		TestingT:     t,
		Position:     position.New(0, -2),
		MeepleParams: test.NoneMeeple(),
	}.Run()

	scores, err := game.Finalize()
	if err != nil {
		t.Fatal(err.Error())
	}
	var expectedScores = []uint32{4, 0}

	for i := range 2 {
		if scores.ReceivedPoints[elements.ID(i+1)] != expectedScores[i] {
			t.Fatalf("Player %d final score incorrect. Expected %d, got: %d", i+1, expectedScores[i], scores.ReceivedPoints[elements.ID(i+1)])
		}
	}
}
