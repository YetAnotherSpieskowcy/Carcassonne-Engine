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
	canWrite     bool
}

// Internal struct returned by the worker function through the received output buffer.
type workerOutput struct {
	requestID int
	resp      Response
}

type outputItemInfo struct {
	GameID        int
	RequestIndex  int
	AcquiredWrite bool
}

// The base interface of a response returned by the API.
type Response interface {
	GameID() int
	Err() error
}

// Responses implementing this interface may indicate to the sender
// that the child games of the game with `GameID()` should be GC-able.
type ResponseChildGamesRemovable interface {
	canRemoveChildGames() bool
}

// Responses implementing this interface may indicate to the sender
// that the the game with `GameID()` should be removed from the engine.
type ResponseGameRemovable interface {
	canRemoveGame() bool
}

// The base interface of a request returned by the API.
type Request interface {
	// gameID() is a private getter -> classes from Python
	// create a new request object using the `GameID` field directly
	gameID() int
	// indicates, if the request requires exclusive write access to the game
	requiresWrite() bool
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

func (resp *SyncResponse) canRemoveGame() bool {
	err := resp.Err()
	if err == nil {
		return false
	}
	_, panicOccured := err.(*ExecutionPanicError)
	return panicOccured
}
