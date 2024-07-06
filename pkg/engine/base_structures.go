package engine

import "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"

type workerInput struct {
	requestID    int
	outputBuffer chan workerOutput
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
	GameID() int
	Execute(*game.Game) Response
}

type baseRequest struct {
	gameID int
}

func (req *baseRequest) GameID() int {
	return req.gameID
}

type baseResponse struct {
	gameID int
	err    error
}

func (resp *baseResponse) GameID() int {
	return resp.gameID
}

func (resp *baseResponse) Err() error {
	return resp.err
}

// used when the request doesn't even reach the worker
type SyncResponse struct {
	baseResponse
}
