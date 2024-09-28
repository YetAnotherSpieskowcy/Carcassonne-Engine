package requestperformancetests

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/engine"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func CreateEarlyGameEngine(logdir string) (*engine.GameEngine, engine.SerializedGameWithID) {
	return PlayTurns(logdir, 10)
}

func CreateLateGameEngine(logdir string) (*engine.GameEngine, engine.SerializedGameWithID) {

	return PlayTurns(logdir, 20)
}

func PlayTurns(logdir string, turnsNumber int) (*engine.GameEngine, engine.SerializedGameWithID) {
	var moves = CreateMovesArray()
	eng, err := engine.StartGameEngine(4, logdir)
	if err != nil {
		panic("start game engine failed")
	}

	gameWithID, err := eng.GenerateOrderedGame(tilesets.EveryTileOnceTileSet())
	if err != nil {
		panic("GenerateOrderedGame failed")
	}

	var serializedGame game.SerializedGame
	for i := range turnsNumber {
		requests := []*engine.PlayTurnRequest{
			{GameID: gameWithID.ID, Move: moves[i]},
		}
		resp := eng.SendPlayTurnBatch(requests)[0]
		serializedGame = resp.Game

	}
	return eng, engine.SerializedGameWithID{ID: gameWithID.ID, Game: serializedGame}
}

// the same game as in pythonendtest alltilegame
func CreateMovesArray() []elements.PlacedTile {

	type MeepleParams struct {
		side        side.Side
		featureType feature.Type
	}

	var tiles = tilesets.EveryTileOnceTileSet()
	var positions = []position.Position{
		position.New(0, -1), // 1
		position.New(1, 0),  // 2
		position.New(-1, 0), // 3
		position.New(-2, 0), // 4
		position.New(-2, 1), // 5
		position.New(-2, 2), // 6
		position.New(1, -1), // 7
		position.New(-1, 2), // 8
		position.New(2, -1), // 9
		position.New(0, 2),  // A
		position.New(0, 3),  // B
		position.New(1, 2),  // C
		position.New(2, 2),  // D
		position.New(3, 2),  // E
		position.New(-1, 1), // F
		position.New(3, 1),  // G
		position.New(-3, 0), // H
		position.New(1, 3),  // I
		position.New(3, 3),  // J
		position.New(2, 3),  // K
		position.New(1, 1),  // L
		position.New(2, 1),  // M
		position.New(0, 4),  // N
		position.New(0, 1),  // O
	}
	var meeples = map[int]MeepleParams{
		1:  {side.NoSide, feature.Monastery},
		2:  {side.NoSide, feature.Monastery},
		3:  {side.Left, feature.Road},
		4:  {side.Bottom, feature.Field},
		5:  {side.Top, feature.Road},
		6:  {side.Right, feature.Road},
		7:  {side.Right, feature.City},
		8:  {side.Bottom, feature.City},
		9:  {side.TopRightEdge, feature.Field},
		10: {side.Bottom, feature.City},          // A
		11: {side.BottomLeftEdge, feature.Field}, // B
		// C no meeple
		13: {side.Bottom, feature.City},  // D
		14: {side.Bottom, feature.City},  // E
		15: {side.Right, feature.City},   // F
		16: {side.Bottom, feature.Field}, // G
		17: {side.Bottom, feature.City},  // H
		18: {side.Top, feature.Road},     // I
		19: {side.Right, feature.Road},   // J
		// K no meeple
		// L no meeple
		// M no meeple
		// N no meeple
		// O no meeple
	}
	var moves = []elements.PlacedTile{}
	for index := range len(positions) {
		ptile := elements.ToPlacedTile(tiles.Tiles[index])
		ptile.Position = positions[index]
		meeple, exists := meeples[index+1]
		if exists {
			ptile.GetPlacedFeatureAtSide(meeple.side, meeple.featureType).Meeple = elements.Meeple{
				Type:     elements.NoneMeeple,
				PlayerID: elements.ID((index % 2) + 1),
			}
		}
		moves = append(moves, ptile)
	}

	return moves
}
