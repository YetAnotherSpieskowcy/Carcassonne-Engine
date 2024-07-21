package engine

import (
	"sync"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
)

// Internal struct passed to the worker function through the input buffer
// with the game pointer, the request, and ways to communicate with the requestor.
type workerInput struct {
	requestID    int
	outputBuffer chan workerOutput
	waitGroup    *sync.WaitGroup
	game         *game.Game
	request      Request
}

// Internal struct returned by the worker function through the received output buffer.
type workerOutput struct {
	requestID int
	resp      Response
}

// The base interface of a response returned by the API.
type Response interface {
	GameID() int
	Err() error
}

// The base interface of a request returned by the API.
type Request interface {
	// gameID() is a private getter -> classes from Python
	// create a new request object using the `GameID` field directly
	gameID() int
	// method that will be executed by the worker
	execute(*game.Game) Response
}

// Concrete type implementing the `Response` interface
// that can be conveniently embedded into every response type.
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
