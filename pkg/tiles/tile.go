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

	//dać Building po prostu by skomponować

	//not sure how to include undefied/null?
	//meeple    Meeple
}

func (tile *Tile) Rotate(rotations uint) Tile {
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

func (tile *Tile) ToString() string {
	var result string
	result = ""
	result += "Cities\n"
	for _, cityConnection := range tile.Cities.Cities {
		result += cityConnection.ToString() + "\n"
	}

	result += "Roads\n"
	for _, road := range tile.Roads.Roads {
		result += road.ToString() + "\n"
	}

	result += "Fields\n"
	for _, field := range tile.Fields.Fields {
		result += field.ToString() + "\n"
	}

	result += "Has shields: " + strconv.FormatBool(tile.HasShield) + "\n"
	result += "Building: " + tile.Building.ToString() + "\n"

	return result
}
