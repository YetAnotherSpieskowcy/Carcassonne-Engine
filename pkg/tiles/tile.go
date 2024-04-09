package tiles

import (
	buildings "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/buildings"
	connectionMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/connection"
	featureMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
)

/*
Immutable object
*/
type Tile struct {
	Features  []featureMod.Feature
	HasShield bool
	Building  buildings.Bulding
}

func (tile *Tile) Cities() [][]connectionMod.Side {
	var cities [][]connectionMod.Side
	for _, feature := range tile.Features {
		if feature.FeatureType == featureMod.City {
			cities = append(cities, feature.Connections)
		}
	}
	return cities
}

func (tile *Tile) CitiesAppendConnection(connections []connectionMod.Side) {
	tile.Features = append(tile.Features, featureMod.Feature{
		FeatureType: featureMod.City,
		Connections: connections,
	})

}

func (tile *Tile) Roads() [][]connectionMod.Side {
	var roads [][]connectionMod.Side
	for _, feature := range tile.Features {
		if feature.FeatureType == featureMod.Road {
			roads = append(roads, feature.Connections)
		}
	}
	return roads
}

func (tile *Tile) RoadsAppendConnection(connections []connectionMod.Side) {
	tile.Features = append(tile.Features, featureMod.Feature{
		FeatureType: featureMod.Road,
		Connections: connections,
	})
}

func (tile *Tile) Fields() [][]connectionMod.Side {
	var fields [][]connectionMod.Side
	for _, feature := range tile.Features {
		if feature.FeatureType == featureMod.Field {
			fields = append(fields, feature.Connections)
		}
	}
	return fields
}

func (tile *Tile) FieldsAppendConnection(connections []connectionMod.Side) {
	tile.Features = append(tile.Features, featureMod.Feature{
		FeatureType: featureMod.Field,
		Connections: connections,
	})
}

func (tile Tile) Rotate(rotations uint) Tile {

	var newFeatures []featureMod.Feature

	for _, feature := range tile.Features {
		var newConnections []connectionMod.Side
		for _, connection := range feature.Connections {
			newConnections = append(newConnections, connection.Rotate(rotations))
		}

		newFeatures = append(newFeatures, featureMod.Feature{FeatureType: feature.FeatureType, Connections: newConnections})
	}

	tile.Features = newFeatures
	return tile
}
