package game

import "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"

// TODO: replace all of these with an import once full tile representation is defined

type Side int64
type Tile struct {}

func (tile Tile) Rotate(rotations uint) Tile {
	return Tile{}
}

const (
	None Side = iota
	Bottom
)

var (
	StartingTile = PlacedTile{}
	BaseTileSet  = []Tile{}
)

func SingleCityEdgeNoRoads() Tile {
	return Tile{}
}

func FourCityEdgesConnectedShield() Tile {
	return Tile{}
}

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
