package main

type Tile struct {
	cities    []Connection
	roads     []Connection
	fields    []FarmConnection
	hasShield bool
	Bulding   Bulding

	//not sure how to include undefied/null?
	//meeple    Meeple
}

func (tile *Tile) Rotate(rotations int) Tile {
	var t Tile
	return t
}
