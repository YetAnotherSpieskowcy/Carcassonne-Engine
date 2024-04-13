package logger

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
)

type StartEntry struct {
	Event        string       `json:"event"`
	StartingTile tiles.Tile   `json:"startingTile"`
	Stack        []tiles.Tile `json:"stack"`
	PlayerCount  int          `json:"playerCount"`
}

func NewStartEntry(deck deck.Deck, playerCount int) StartEntry {
	return StartEntry{
		Event:        "start",
		StartingTile: deck.StartingTile,
		Stack:        deck.GetRemaining(),
		PlayerCount:  playerCount,
	}
}

type PlaceTileEntry struct {
	Event    string             `json:"event"`
	PlayerID uint8              `json:"playerID"`
	Move     elements.LegalMove `json:"move"`
}

func NewPlaceTileEntry(player elements.Player, move elements.LegalMove) PlaceTileEntry {
	return PlaceTileEntry{"place", player.ID(), move}
}

type EndEntry struct {
	Event  string   `json:"event"`
	Scores []uint32 `json:"scores"`
}

func NewEndEntry(scores []uint32) EndEntry {
	return EndEntry{"end", scores}
}
