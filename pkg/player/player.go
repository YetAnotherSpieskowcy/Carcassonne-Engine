package player

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

type player struct {
	id uint8
	// indexed by meeple's enum value
	meepleCounts []uint8
	score        uint32
}

func New(id uint8) elements.Player {
	meepleCounts := make([]uint8, elements.MeepleTypeCount)
	meepleCounts[elements.NormalMeeple] = 7
	return &player{
		id:           id,
		meepleCounts: meepleCounts,
		score:        0,
	}
}

func (player player) ID() uint8 {
	return player.id
}

func (player player) MeepleCount(meepleType elements.MeepleType) uint8 {
	return player.meepleCounts[meepleType]
}

func (player *player) SetMeepleCount(meepleType elements.MeepleType, value uint8) {
	player.meepleCounts[meepleType] = value
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
	meepleCount := player.MeepleCount(tile.Meeple.Type)
	if meepleCount == 0 && tile.Meeple.Side != side.None {
		return elements.ScoreReport{}, elements.ErrNoMeepleAvailable
	}
	scoreReport, err := board.PlaceTile(tile)
	if err != nil {
		return scoreReport, err
	}
	if tile.Meeple.Side != side.None {
		player.SetMeepleCount(tile.Meeple.Type, meepleCount-1)
	}
	return scoreReport, nil
}
