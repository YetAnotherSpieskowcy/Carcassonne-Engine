package logger

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
)

type StartEntry struct {
	Event       string       `json:"event"`
	Deck        []tiles.Tile `json:"deck"`
	PlayerCount int          `json:"playerCount"`
}

func NewStartEntry(deck *stack.Stack[tiles.Tile], playerCount int) StartEntry {
	return StartEntry{"start", deck.GetRemaining(), playerCount}
}

type PlaceTileEntry struct {
	Event    string              `json:"event"`
	PlayerID int                 `json:"playerID"`
	Tile     elements.PlacedTile `json:"tile"`
}

func NewPlaceTileEntry(playerID int, tile elements.PlacedTile) PlaceTileEntry {
	return PlaceTileEntry{"place", playerID, tile}
}

type EndEntry struct {
	Event  string   `json:"event"`
	Scores []uint32 `json:"scores"`
}

func NewEndEntry(scores []uint32) EndEntry {
	return EndEntry{"end", scores}
}
