package tiles

import (
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

func (tile *Tile) Cities() [][]connectionMod.Side {
	var cities [][]connectionMod.Side
	for _, feature := range tile.Features {
		if feature.FeatureType == CITY {
			cities = append(cities, feature.Connections)
		}
	}
	return cities
}

func (tile *Tile) CitiesAppendConnection(connections []connectionMod.Side) {
	tile.Features = append(tile.Features, Feature{
		FeatureType: CITY,
		Connections: connections,
	})

}

func (tile *Tile) Roads() [][]connectionMod.Side {
	var roads [][]connectionMod.Side
	for _, feature := range tile.Features {
		if feature.FeatureType == ROAD {
			roads = append(roads, feature.Connections)
		}
	}
	return roads
}

func (tile *Tile) RoadsAppendConnection(connections []connectionMod.Side) {
	tile.Features = append(tile.Features, Feature{
		FeatureType: ROAD,
		Connections: connections,
	})
}

func (tile *Tile) Fields() [][]connectionMod.Side {
	var fields [][]connectionMod.Side
	for _, feature := range tile.Features {
		if feature.FeatureType == FIELD {
			fields = append(fields, feature.Connections)
		}
	}
	return fields
}

func (tile *Tile) FieldsAppendConnection(connections []connectionMod.Side) {
	tile.Features = append(tile.Features, Feature{
		FeatureType: FIELD,
		Connections: connections,
	})
}

func (tile Tile) Rotate(rotations uint) Tile {
	var t Tile
	// rotate cities
	for _, cityConnection := range tile.Cities() {
		t.CitiesAppendConnection(connectionMod.RotateSideArray(cityConnection, rotations))
	}

	// rotate roads
	for _, road := range tile.Roads() {
		t.RoadsAppendConnection(connectionMod.RotateSideArray(road, rotations))
	}

	// rotate fields
	for _, field := range tile.Fields() {
		t.FieldsAppendConnection(connectionMod.RotateSideArray(field, rotations))
	}

	// copy parameters
	t.HasShield = tile.HasShield
	t.Building = tile.Building
	return t
}
