package player

import (
	"slices"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
)

type player struct {
	id elements.ID
	// indexed by meeple's enum value
	meepleCounts []uint8
	score        uint32
}

func New(id elements.ID) elements.Player {
	meepleCounts := make([]uint8, elements.MeepleTypeCount)
	meepleCounts[elements.NormalMeeple] = 7
	return &player{
		id:           id,
		meepleCounts: meepleCounts,
		score:        0,
	}
}

func (player player) DeepClone() elements.Player {
	player.meepleCounts = slices.Clone(player.meepleCounts)
	return &player
}

func (player player) ID() elements.ID {
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

// how am I supposed to name this sensibly...
func (player *player) GetEligibleMovesFrom(moves []elements.PlacedTile) []elements.PlacedTile {
	result := []elements.PlacedTile{}
	for _, move := range moves {
		if player.IsEligibleFor(move) {
			result = append(result, move)
		}
	}
	return result
}

// how am I supposed to name this sensibly...
func (player *player) IsEligibleFor(move elements.PlacedTile) bool {
	count := 0
	for _, feature := range move.Features {
		if feature.Meeple.Type != elements.NoneMeeple {
			if feature.Meeple.PlayerID != player.id {
				return false
			}
			if player.MeepleCount(feature.Meeple.Type) == 0 {
				return false
			}
			if count > 1 {
				return false
			}
			count++
		}
	}
	return true
}

func (player *player) PlaceTile(
	board elements.Board, move elements.PlacedTile,
) (elements.ScoreReport, error) {
	if !player.IsEligibleFor(move) {
		return elements.ScoreReport{}, elements.ErrNoMeepleAvailable
	}

	scoreReport, err := board.PlaceTile(move)
	if err != nil {
		return scoreReport, err
	}
	for _, feature := range move.Features {
		if feature.Meeple.Type != elements.NoneMeeple {
			meepleCount := player.MeepleCount(feature.Meeple.Type)
			player.SetMeepleCount(feature.Meeple.Type, meepleCount-1)
		}
	}
	return scoreReport, nil
}
