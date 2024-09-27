package requestperformancetests

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/engine"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func CreateEarlyGameEngine(logdir string) (*engine.GameEngine, engine.SerializedGameWithID) {
	return PlayTurns(logdir, 10)
}

func CreateLateGameEngine(logdir string) (*engine.GameEngine, engine.SerializedGameWithID) {

	return PlayTurns(logdir, 20)
}

// TODO ADD MEEPLES
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
	var moves = []elements.PlacedTile{}
	for index := range len(positions) {
		ptile := elements.ToPlacedTile(tiles.Tiles[index])
		ptile.Position = positions[index]
		moves = append(moves, ptile)
	}

	return moves
}
