package engine

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
)

type PlayTurnResponse struct {
	BaseResponse
	Game game.SerializedGame
}
type PlayTurnRequest struct {
	GameID int
	Move   elements.PlacedTile
}

func (req *PlayTurnRequest) gameID() int {
	return req.GameID
}

func (req *PlayTurnRequest) execute(game *game.Game) Response {
	err := game.PlayTurn(req.Move)
	resp := &PlayTurnResponse{
		BaseResponse: BaseResponse{
			gameID: req.gameID(),
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
	BaseResponse
}
type GameTreeRequest struct {
	GameID int
}

func (req *GameTreeRequest) gameID() int {
	return req.GameID
}

func (req *GameTreeRequest) execute(game *game.Game) Response {
	_ = game
	panic("game tree request not implemented")
}
