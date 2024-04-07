package game

import "errors"


type InvalidMove struct {
	msg string
}

func (err *InvalidMove) Error() string {
	return err.msg
}

var (
	InvalidPosition   = &InvalidMove{"The tile cannot be placed at given position."}
	NoMeepleAvailable = &InvalidMove{"The player does not have any meeples available."}
	WrongTile         = &InvalidMove{"The played tile is not the one that was drawn."}
	GameIsNotFinished = errors.New("The game is not finished yet.")
)
