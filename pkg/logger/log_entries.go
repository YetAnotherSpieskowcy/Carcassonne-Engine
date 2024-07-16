package logger

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
)

type Entry struct {
	Event   string `json:"event"`
	Content []byte `json:"content"`
}

func NewEntry(event string, content []byte) Entry {
	return Entry{
		Event:   event,
		Content: content,
	}
}

type StartEntryContent struct {
	StartingTile tiles.Tile   `json:"startingTile"`
	Stack        []tiles.Tile `json:"stack"`
	PlayerCount  int          `json:"playerCount"`
}

func NewStartEntryContent(startingTile tiles.Tile, stack []tiles.Tile, playerCount int) StartEntryContent {
	return StartEntryContent{
		StartingTile: startingTile,
		Stack:        stack,
		PlayerCount:  playerCount,
	}
}

type PlaceTileEntryContent struct {
	PlayerID elements.ID         `json:"playerID"`
	Move     elements.PlacedTile `json:"move"`
}

func NewPlaceTileEntryContent(player elements.ID, move elements.PlacedTile) PlaceTileEntryContent {
	return PlaceTileEntryContent{
		PlayerID: player,
		Move:     move,
	}
}

type EndEntryContent struct {
	Scores elements.ScoreReport `json:"scores"`
}

func NewEndEntryContent(scores elements.ScoreReport) EndEntryContent {
	return EndEntryContent{Scores: scores}
}
