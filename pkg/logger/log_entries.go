package logger

import (
	"encoding/json"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
)

type EventType string

const (
	StartEvent     EventType = "start"
	PlaceTileEvent EventType = "place"
	ScoreEvent     EventType = "score"
)

type Entry struct {
	Event   EventType `json:"event"`
	Content []byte    `json:"content"`
}

func NewEntry(event EventType, content []byte) Entry {
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

func ParseStartEntryContent(entryContent []byte) StartEntryContent {
	var content StartEntryContent
	err := json.Unmarshal(entryContent, &content)
	if err != nil {
		return StartEntryContent{}
	}
	return content
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

func ParsePlaceTileEntryContent(entryContent []byte) PlaceTileEntryContent {
	var content PlaceTileEntryContent
	err := json.Unmarshal(entryContent, &content)
	if err != nil {
		return PlaceTileEntryContent{}
	}
	return content
}

type ScoreEntryContent struct {
	Scores elements.ScoreReport `json:"scores"`
}

func NewScoreEntryContent(scores elements.ScoreReport) ScoreEntryContent {
	return ScoreEntryContent{Scores: scores}
}

func ParseScoreEntryContent(entryContent []byte) ScoreEntryContent {
	var content ScoreEntryContent
	err := json.Unmarshal(entryContent, &content)
	if err != nil {
		return ScoreEntryContent{}
	}
	return content
}
