package logger

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
)

type StartEntry struct {
	Event       string          `json:"event"`
	Deck        []elements.Tile `json:"deck"`
	PlayerCount int             `json:"playerCount"`
}

// TODO: replace with real tile implementation
func NewStartEntry(deck *stack.Stack[elements.Tile], playerCount int) StartEntry {
	return StartEntry{"start", deck.GetRemaining(), playerCount}
}

type PlaceTileEntry struct {
	Event    string              `json:"event"`
	PlayerId int                 `json:"playerId"`
	Tile     elements.PlacedTile `json:"tile"`
}

func NewPlaceTileEntry(playerId int, tile elements.PlacedTile) PlaceTileEntry {
	return PlaceTileEntry{"place", playerId, tile}
}

type EndEntry struct {
	Event  string   `json:"event"`
	Scores []uint32 `json:"scores"`
}

func NewEndEntry(scores []uint32) EndEntry {
	return EndEntry{"end", scores}
}
