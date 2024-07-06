package engine

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
)

type PlayTurnResponse struct {
	baseResponse
	Game game.SerializedGame
}
type PlayTurnRequest struct {
	baseRequest
	Move elements.PlacedTile
}

func (req *PlayTurnRequest) Execute(game *game.Game) Response {
	err := game.PlayTurn(req.Move)
	resp := &PlayTurnResponse{
		baseResponse: baseResponse{
			gameID: req.GameID(),
			err:    err,
		},
	}
	if err != nil {
		return resp
	}

	resp.Game = game.Serialized()
	return resp
}

// TODO: implement game tree request based on agent's needs
type GameTreeResponse struct {
	baseResponse
}
type GameTreeRequest struct {
	baseRequest
}

func (req *GameTreeRequest) Execute(game *game.Game) Response {
	_ = game
	panic("game tree request not implemented")
}
