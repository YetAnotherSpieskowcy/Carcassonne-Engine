package tiles

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/building"
	featureMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	sideMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

/*
Immutable object
*/
type Tile struct {
	Features  []featureMod.Feature
	HasShield bool
	Building  building.Building
}

func (tile Tile) Equals(other Tile) bool {
outer:
	for rotations := range uint(4) {
		rotated := other.Rotate(rotations)
		if tile.HasShield != rotated.HasShield {
			continue
		}
		if tile.Building != rotated.Building {
			continue
		}
		if len(tile.Features) != len(rotated.Features) {
			continue
		}
		for i, feature := range tile.Features {
			if !feature.Equals(rotated.Features[i]) {
				continue outer
			}
		}
		return true
	}
	return false
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
	rotations %= 4
	if rotations == 0 {
		return tile
	}

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

func (tile *Tile) GetFeatureAtSide(sideToCheck sideMod.Side) *featureMod.Feature {
	for _, feature := range tile.Features {
		for _, side := range feature.Sides {
			if side == sideToCheck {
				return &feature
			}
		}
	}
	return nil
}
