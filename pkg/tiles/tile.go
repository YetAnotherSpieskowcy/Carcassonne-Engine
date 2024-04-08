package tiles

import (
	"strconv"

	buildings "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/buildings"
	connectionMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/connection"
)

/*
Immutable object
*/
type Tile struct {
	Features  []Feature
	HasShield bool
	Building  buildings.Bulding
}

func (tile *Tile) Cities() []connectionMod.Connection {
	for _, feature := range tile.Features {
		if feature.FeatureType == CITIES {
			return feature.Connections
		}
	}
	return []connectionMod.Connection{}
}

func (tile *Tile) CitiesAppendConnection(connection connectionMod.Connection) {
	var found = false
	for _, feature := range tile.Features {
		if feature.FeatureType == CITIES {
			feature.Connections = append(feature.Connections, connection)
			found = true
		}
	}
	if !found {
		tile.Features = append(tile.Features, Feature{
			FeatureType: CITIES,
			Connections: []connectionMod.Connection{connection},
		})
	}
}

func (tile *Tile) Roads() []connectionMod.Connection {
	for _, feature := range tile.Features {
		if feature.FeatureType == ROADS {
			return feature.Connections
		}
	}
	return []connectionMod.Connection{}
}

func (tile *Tile) RoadsAppendConnection(connection connectionMod.Connection) {
	var found = false
	for _, feature := range tile.Features {
		if feature.FeatureType == ROADS {
			feature.Connections = append(feature.Connections, connection)
			found = true
		}
	}
	if !found {
		tile.Features = append(tile.Features, Feature{
			FeatureType: ROADS,
			Connections: []connectionMod.Connection{connection},
		})
	}
}

func (tile *Tile) Fields() []connectionMod.Connection {
	for _, feature := range tile.Features {
		if feature.FeatureType == FIELDS {
			return feature.Connections
		}
	}
	return []connectionMod.Connection{}
}

func (tile *Tile) FieldsAppendConnection(connection connectionMod.Connection) {
	var found = false
	for _, feature := range tile.Features {
		if feature.FeatureType == FIELDS {
			feature.Connections = append(feature.Connections, connection)
			found = true
		}
	}
	if !found {
		tile.Features = append(tile.Features, Feature{
			FeatureType: FIELDS,
			Connections: []connectionMod.Connection{connection},
		})
	}
}

func (tile Tile) Rotate(rotations uint) Tile {
	var t Tile
	// rotate cities
	for _, cityConnection := range tile.Cities() {
		t.CitiesAppendConnection(cityConnection.Rotate(rotations))
	}

	// rotate roads
	for _, road := range tile.Roads() {
		t.RoadsAppendConnection(road.Rotate(rotations))
	}

	// rotate fields
	for _, field := range tile.Fields() {
		t.FieldsAppendConnection(field.Rotate(rotations))
	}

	// copy parameters
	t.HasShield = tile.HasShield
	t.Building = tile.Building
	return t
}

func (tile *Tile) String() string {
	var result string
	result = ""
	result += "Cities\n"
	for _, cityConnection := range tile.Cities() {
		result += cityConnection.String() + "\n"
	}

	result += "Roads\n"
	for _, road := range tile.Roads() {
		result += road.String() + "\n"
	}

	result += "Fields\n"
	for _, field := range tile.Fields() {
		result += field.String() + "\n"
	}

	result += "Has shields: " + strconv.FormatBool(tile.HasShield) + "\n"
	result += "Building: " + tile.Building.String() + "\n"

	return result
}
