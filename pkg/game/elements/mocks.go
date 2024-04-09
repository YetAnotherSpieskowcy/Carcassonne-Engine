package elements

// TODO: replace all of these with an import once full tile representation is defined

type Side int64
type Tile struct {
	ID int
}

func (tile Tile) Rotate(_ uint) Tile {
	return Tile{ID: tile.ID}
}

const (
	None Side = iota
	Bottom
)

var (
	StartingTile = PlacedTile{LegalMove: LegalMove{Tile: Tile{ID: 69}}}
	BaseTileSet  = []Tile{SingleCityEdgeNoRoads(), FourCityEdgesConnectedShield()}
)

func SingleCityEdgeNoRoads() Tile {
	return Tile{ID: 1}
}

func FourCityEdgesConnectedShield() Tile {
	return Tile{ID: 2}
}

func GetStandardTiles() []Tile {
	tiles := []Tile{}
	repeatCount := 3
	for tileID := range 71 + repeatCount {
		tiles = append(tiles, Tile{ID: tileID / repeatCount})
	}
	return tiles[repeatCount:]
}
