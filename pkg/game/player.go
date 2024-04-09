package game

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
)

type player struct {
	id          uint8
	meepleCount uint8
	score       uint32
}

func NewPlayer(id uint8) elements.Player {
	return &player{
		id:          id,
		meepleCount: 7,
		score:       0,
	}
}

func (player player) ID() uint8 {
	return player.id
}

func (player player) MeepleCount() uint8 {
	return player.meepleCount
}

func (player *player) SetMeepleCount(value uint8) {
	player.meepleCount = value
}

func (player player) Score() uint32 {
	return player.score
}

func (player *player) SetScore(value uint32) {
	player.score = value
}

func (player *player) PlaceTile(
	board elements.Board, tile elements.PlacedTile,
) (elements.ScoreReport, error) {
	if player.meepleCount == 0 && tile.Meeple.Side != elements.None {
		return elements.ScoreReport{}, NoMeepleAvailable
	}
	scoreReport, err := board.PlaceTile(tile)
	if err != nil {
		return scoreReport, err
	}
	if tile.Meeple.Side != elements.None {
		player.meepleCount--
	}
	return scoreReport, nil
}
