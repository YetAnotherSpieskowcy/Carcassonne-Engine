package game

type Player struct {
	id          uint8
	meepleCount uint8
	score       uint32
}

func NewPlayer(id uint8) *Player {
	return &Player{
		id:          id,
		meepleCount: 7,
		score:       0,
	}
}

func (player Player) Id() uint8 {
	return player.id
}

func (player Player) MeepleCount() uint8 {
	return player.meepleCount
}

func (player Player) Score() uint32 {
	return player.score
}

// XXX: `PlacedTile` may just become `Tile` if the meeple field does not get split out:
// see https://github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pull/9#discussion_r1554723567
func (player *Player) PlaceTile(board *Board, tile PlacedTile) (ScoreReport, error) {
	if player.meepleCount == 0 && tile.Meeple == nil {
		return ScoreReport{}, NoMeepleAvailable
	}
	scoreReport, err := board.PlaceTile(tile)
	if err != nil {
		return scoreReport, err
	}
	if tile.Meeple != nil {
		player.meepleCount -= 1
	}
	return scoreReport, nil
}
