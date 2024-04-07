package tiles

import (
	"strconv"

	buildings "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/buildings"
)

type Tile struct {
	Cities
	Roads
	Fields
	HasShield bool
	Building  buildings.Bulding
}

func (tile Tile) Rotate(rotations uint) Tile {
	var t Tile
	//rotate cities
	for _, cityConnection := range tile.Cities.Cities {
		t.Cities.Cities = append(t.Cities.Cities, cityConnection.Rotate(rotations))
	}

	//rotate roads
	for _, road := range tile.Roads.Roads {
		t.Roads.Roads = append(t.Roads.Roads, road.Rotate(rotations))
	}

	//rotate fields
	for _, field := range tile.Fields.Fields {
		t.Fields.Fields = append(t.Fields.Fields, field.Rotate(rotations))
	}

	//copy parameters
	t.HasShield = tile.HasShield
	t.Building = tile.Building
	return t
}

func (tile *Tile) String() string {
	var result string
	result = ""
	result += "Cities\n"
	for _, cityConnection := range tile.Cities.Cities {
		result += cityConnection.String() + "\n"
	}

	result += "Roads\n"
	for _, road := range tile.Roads.Roads {
		result += road.String() + "\n"
	}

	result += "Fields\n"
	for _, field := range tile.Fields.Fields {
		result += field.String() + "\n"
	}

	result += "Has shields: " + strconv.FormatBool(tile.HasShield) + "\n"
	result += "Building: " + tile.Building.String() + "\n"

	return result
}
