package game

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

/*
  All tests contains the board visualization. Each tile is represented by 5x5 ascii characters
*/

func makeTurn(game *Game, t *testing.T, tilePosition position.Position, rotations uint, meeple elements.MeepleType, featureSide side.Side, featureType feature.Type) {
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}

	var player = game.CurrentPlayer()

	ptile := elements.ToPlacedTile(tile.Rotate(rotations))
	ptile.Position = tilePosition
	if meeple != elements.NoneMeeple {
		ptile.GetPlacedFeatureAtSide(featureSide, featureType).Meeple = elements.Meeple{
			Type:     meeple,
			PlayerID: player.ID(),
		}
	}

	err = game.PlayTurn(ptile)

	if err != nil {
		t.Fatal(err.Error())
	}
}

/*
|				  0	   1

|		        .| |......
|		        .\ /......
|0		        --0----1-!
|	            ..........
|	            ..........
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
	game, err := NewFromTileSet(tileset, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	makeTurn(game, t, position.New(1, 0), 0, elements.NormalMeeple, side.Right, feature.Road)

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
|				  0	   1

|		        .| |......
|		        .\ /......
|0		        --0----1-!
|	            ..........
|	            ..........
|		        ..........
|		        ..........
|-1		        --2-@--3--
|	            ..........
|	            ..........
|		        .....
|		        .....
|-2		        --4-@
|	            .....
|	            .....
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
	game, err := NewFromTileSet(tileset, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	makeTurn(game, t, position.New(1, 0), 0, elements.NormalMeeple, side.Right, feature.Road)
	makeTurn(game, t, position.New(0, -1), 0, elements.NormalMeeple, side.Right, feature.Road)
	makeTurn(game, t, position.New(1, -1), 0, elements.NoneMeeple, side.NoSide, feature.NoneType)
	makeTurn(game, t, position.New(0, -2), 0, elements.NormalMeeple, side.Right, feature.Road)

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
|				  0	   1

|		        .| |......
|		        .\ /..[!].
|0		        --0---[1].
|	            ......[ ].
|	            ..........
*/
func TestFinalScoreMonastery(t *testing.T) {
	// ------ create tileset --------
	var tiles []tiles.Tile
	var err error
	tiles = append(tiles, tiletemplates.MonasteryWithSingleRoad())

	tileset := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}

	// ------ create game --------
	game, err := NewFromTileSet(tileset, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	makeTurn(game, t, position.New(1, 0), 1, elements.NormalMeeple, side.NoSide, feature.Monastery)

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
|				  0	   1

|		        .| |......
|		        .\ /..[!].
|0		        --0---[1].
|	            ......[ ].
|	            ..........
|		        ..........
|		        .[@]..[!].
|-1		        .[2]..[3].
|	            .[ ]..[ ].
|	            ..........
*/
func TestFinalScoreMonasteries(t *testing.T) {
	// ------ create tileset --------
	var tiles []tiles.Tile
	var err error
	tiles = append(tiles, tiletemplates.MonasteryWithSingleRoad())
	tiles = append(tiles, tiletemplates.MonasteryWithoutRoads())
	tiles = append(tiles, tiletemplates.MonasteryWithoutRoads())

	tileset := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}

	// ------ create game --------
	game, err := NewFromTileSet(tileset, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	makeTurn(game, t, position.New(1, 0), 1, elements.NormalMeeple, side.NoSide, feature.Monastery)
	makeTurn(game, t, position.New(0, -1), 0, elements.NormalMeeple, side.NoSide, feature.Monastery)
	makeTurn(game, t, position.New(1, -1), 0, elements.NormalMeeple, side.NoSide, feature.Monastery)

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
|				  0	   1

|	            .| |.
|	            ./!\.
|1		        | 1 |
|		        .\ /.
|		        .| |.
|		        .| |.
|		        .\ /.
|0		        --0--
|	            .....
|	            .....
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
	game, err := NewFromTileSet(tileset, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	makeTurn(game, t, position.New(0, 1), 0, elements.NormalMeeple, side.Top, feature.City)

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
|				  0	   1

|	       .|!|..| |..|@|.
|	       .\ /../!\..\ /.
|1		   ..3..| 1 |..2..
|		   ......\ /......
|		   ......| |......
|		        .| |.
|		        .\ /.
|0		        --0--
|	            .....
|	            .....
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
	game, err := NewFromTileSet(tileset, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	makeTurn(game, t, position.New(0, 1), 0, elements.NormalMeeple, side.Top, feature.City)
	makeTurn(game, t, position.New(1, 1), 0, elements.NormalMeeple, side.Top, feature.City)
	makeTurn(game, t, position.New(-1, 1), 0, elements.NormalMeeple, side.Top, feature.City)

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
|				  0	   1

|	            ..!..
|	            .....
|1		        ..1..
|		        ./ \.
|		        .| |.
|		        .| |.
|		        .\ /.
|0		        --0--
|	            .....
|	            .....
*/
func TestFinalScoreField(t *testing.T) {
	// ------ create tileset --------
	var tiles []tiles.Tile
	var err error
	tiles = append(tiles, tiletemplates.SingleCityEdgeNoRoads())

	tileset := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}

	// ------ create game --------
	game, err := NewFromTileSet(tileset, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	makeTurn(game, t, position.New(0, 1), 2, elements.NormalMeeple, side.Top, feature.Field)

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
