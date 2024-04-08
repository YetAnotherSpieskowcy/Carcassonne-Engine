package game

import (
	. "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
)

type player struct {
	id          uint8
	meepleCount uint8
	score       uint32
}

func NewPlayer(id uint8) Player {
	return &player{
		id:          id,
		meepleCount: 7,
		score:       0,
	}
}

func (player player) Id() uint8 {
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

func (player *player) PlaceTile(board Board, tile PlacedTile) (ScoreReport, error) {
	if player.meepleCount == 0 && tile.Meeple.Side != None {
		return ScoreReport{}, NoMeepleAvailable
	}
	scoreReport, err := board.PlaceTile(tile)
	if err != nil {
		return scoreReport, err
	}
	if tile.Meeple.Side != None {
		player.meepleCount -= 1
	}
	return scoreReport, nil
}
