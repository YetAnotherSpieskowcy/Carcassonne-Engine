package player

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
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

// how am I supposed to name this sensibly...
func (player *player) GetEligibleMovesFrom(moves []elements.LegalMove) []elements.LegalMove {
	result := []elements.LegalMove{}
	for _, move := range moves {
		if player.IsEligibleFor(move) {
			result = append(result, move)
		}
	}
	return result
}

// how am I supposed to name this sensibly...
func (player *player) IsEligibleFor(move elements.LegalMove) bool {
	if move.Meeple.Feature.FeatureType == feature.None {
		return true
	}
	return player.MeepleCount(move.Meeple.Type) != 0
}

func (player *player) PlaceTile(
	board elements.Board, move elements.LegalMove,
) (elements.ScoreReport, error) {
	if !player.IsEligibleFor(move) {
		return elements.ScoreReport{}, elements.ErrNoMeepleAvailable
	}

	tile := elements.PlacedTile{LegalMove: move, Player: player}
	scoreReport, err := board.PlaceTile(tile)
	if err != nil {
		return scoreReport, err
	}

	if move.Meeple.Feature.FeatureType != feature.None {
		meepleCount := player.MeepleCount(move.Meeple.Type)
		player.SetMeepleCount(move.Meeple.Type, meepleCount-1)
	}

	return scoreReport, nil
}
