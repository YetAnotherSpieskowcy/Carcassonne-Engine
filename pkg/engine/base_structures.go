package engine

import (
	"sync"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
)

type workerInput struct {
	requestID    int
	outputBuffer chan workerOutput
	waitGroup    *sync.WaitGroup
	game         *game.Game
	request      Request
}
type workerOutput struct {
	requestID int
	resp      Response
}

type Response interface {
	GameID() int
	Err() error
}
type Request interface {
	gameID() int
	execute(*game.Game) Response
}

type BaseResponse struct {
	gameID int
	err    error
}

func (resp *BaseResponse) GameID() int {
	return resp.gameID
}

func (resp *BaseResponse) Err() error {
	return resp.err
}

// used when the request doesn't even reach the worker
type SyncResponse struct {
	BaseResponse
}
