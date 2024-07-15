package logger

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
)

type StartEntry struct {
	Event        string       `json:"event"`
	StartingTile tiles.Tile   `json:"startingTile"`
	Stack        []tiles.Tile `json:"stack"`
	PlayerCount  int          `json:"playerCount"`
}

func NewStartEntry(startingTile tiles.Tile, stack []tiles.Tile, playerCount int) StartEntry {
	return StartEntry{
		Event:        "start",
		StartingTile: startingTile,
		Stack:        stack,
		PlayerCount:  playerCount,
	}
}

type PlaceTileEntry struct {
	Event    string              `json:"event"`
	PlayerID elements.ID         `json:"playerID"`
	Move     elements.PlacedTile `json:"move"`
}

func NewPlaceTileEntry(player elements.ID, move elements.PlacedTile) PlaceTileEntry {
	return PlaceTileEntry{"place", player, move}
}

type EndEntry struct {
	Event  string               `json:"event"`
	Scores elements.ScoreReport `json:"scores"`
}

func NewEndEntry(scores elements.ScoreReport) EndEntry {
	return EndEntry{"end", scores}
}
