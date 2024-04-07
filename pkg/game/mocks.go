package game

import "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"

// TODO: replace all of these with an import once full tile representation is defined

type Meeple struct {
	player   Player
	side     any
}
type Tile struct {}
type LegalMove struct {
	Tile
	pos  Position
}
type PlacedTile struct {
	LegalMove
	Meeple    *Meeple
}

var (
	StartingTile = PlacedTile{}
	BaseTileSet  = []Tile{}
)

// TODO: replace this with an import once logger is defined
type Logger struct {}

func (logger *Logger) Start(deck *stack.Stack[Tile], playerCount int) error {
	panic("not implemented")
}

func (logger *Logger) PlaceTile(playerId int, tile PlacedTile) error {
	panic("not implemented")
}

func (logger *Logger) End(scores []uint32) error {
	panic("not implemented")
}
