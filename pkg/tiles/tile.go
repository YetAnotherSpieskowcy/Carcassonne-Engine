package tiles

type Tile struct {
	cities    []Connection //<- rozdzielić na osobne klasy
	roads     []Connection
	fields    []FarmConnection
	hasShield bool
	Bulding   Bulding

	//dać Building po prostu by skomponować

	//not sure how to include undefied/null?
	//meeple    Meeple
}

func (tile *Tile) Rotate(rotations int) Tile {
	var t Tile
	return t
}
