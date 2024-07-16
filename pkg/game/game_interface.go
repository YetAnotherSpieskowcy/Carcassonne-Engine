//go:build test

package game

import "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"

func (game *Game) GetBoard() elements.Board {
	return game.board
}
