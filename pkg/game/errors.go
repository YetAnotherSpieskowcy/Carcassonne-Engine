package game

import "errors"

type InvalidMove struct {
	msg string
}

func (err *InvalidMove) Error() string {
	return err.msg
}

var (
	InvalidPosition      = &InvalidMove{"the tile cannot be placed at given position"}
	NoMeepleAvailable    = &InvalidMove{"the player does not have any meeples available"}
	WrongTile            = &InvalidMove{"the played tile is not the one that was drawn"}
	ErrGameIsNotFinished = errors.New("the game is not finished yet")
)
