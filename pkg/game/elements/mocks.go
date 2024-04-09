package elements

// TODO: replace all of these with an import once full tile representation is defined

type Side int64
type Tile struct {
	ID int
}

func (tile Tile) Rotate(_ uint) Tile {
	return Tile{}
}

const (
	None Side = iota
	Bottom
)

var (
	StartingTile = PlacedTile{}
	BaseTileSet  = []Tile{SingleCityEdgeNoRoads(), FourCityEdgesConnectedShield()}
)

func SingleCityEdgeNoRoads() Tile {
	return Tile{}
}

func FourCityEdgesConnectedShield() Tile {
	return Tile{}
}
