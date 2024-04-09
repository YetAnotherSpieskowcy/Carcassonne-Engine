package tiles

import (
	buildings "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/buildings"
	featureMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	sideMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

/*
Immutable object
*/
type Tile struct {
	Features  []featureMod.Feature
	HasShield bool
	Building  buildings.Bulding
}

func (tile *Tile) Cities() []featureMod.Feature {
	var cities []featureMod.Feature
	for _, feature := range tile.Features {
		if feature.FeatureType == featureMod.City {
			cities = append(cities, feature)
		}
	}
	return cities
}

func (tile *Tile) Roads() []featureMod.Feature {
	var roads []featureMod.Feature
	for _, feature := range tile.Features {
		if feature.FeatureType == featureMod.Road {
			roads = append(roads, feature)
		}
	}
	return roads
}

func (tile *Tile) Fields() []featureMod.Feature {
	var fields []featureMod.Feature
	for _, feature := range tile.Features {
		if feature.FeatureType == featureMod.Field {
			fields = append(fields, feature)
		}
	}
	return fields
}

func (tile Tile) Rotate(rotations uint) Tile {

	var newFeatures []featureMod.Feature

	for _, feature := range tile.Features {
		var newSides []sideMod.Side
		for _, side := range feature.Sides {
			newSides = append(newSides, side.Rotate(rotations))
		}

		newFeatures = append(newFeatures, featureMod.Feature{FeatureType: feature.FeatureType, Sides: newSides})
	}

	tile.Features = newFeatures
	return tile
}
