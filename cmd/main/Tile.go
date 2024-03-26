package main

type Tile struct {
	cities    []Connection
	roads     []Connection
	fields    []FarmConnection
	hasShield bool
	Bulding   Bulding
	meeple    Meeple
}
