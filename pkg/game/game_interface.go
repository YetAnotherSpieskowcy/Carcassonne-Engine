package game

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
)

// used only in tests!
func (game *Game) GetBoard() elements.Board {
	return game.board
}
