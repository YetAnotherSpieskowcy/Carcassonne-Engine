package tiles

import (
	"slices"

	featureMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
)

/*
Immutable object
*/
type Tile struct {
	Features []featureMod.Feature
}

func (tile Tile) Equals(other Tile) bool {
outer:
	for rotations := range uint(4) {
		rotated := other.Rotate(rotations)
		if len(tile.Features) != len(rotated.Features) {
			continue
		}
		for i, feature := range tile.Features {
			if feature != rotated.Features[i] {
				continue outer
			}
		}
		return true
	}
	return false
}

func (tile Tile) Cities() []featureMod.Feature {
	var cities []featureMod.Feature
	for _, feature := range tile.Features {
		if feature.FeatureType == featureMod.City {
			cities = append(cities, feature)
		}
	}
	return cities
}

func (tile Tile) Roads() []featureMod.Feature {
	var roads []featureMod.Feature
	for _, feature := range tile.Features {
		if feature.FeatureType == featureMod.Road {
			roads = append(roads, feature)
		}
	}
	return roads
}

func (tile Tile) Fields() []featureMod.Feature {
	var fields []featureMod.Feature
	for _, feature := range tile.Features {
		if feature.FeatureType == featureMod.Field {
			fields = append(fields, feature)
		}
	}
	return fields
}

func (tile Tile) Monastery() *featureMod.Feature {
	for _, feature := range tile.Features {
		if feature.FeatureType == featureMod.Monastery {
			return &feature
		}
	}
	return nil
}

func (tile Tile) Rotate(rotations uint) Tile {
	rotations %= 4
	if rotations == 0 {
		return tile
	}

	var newFeatures []featureMod.Feature

	for _, feature := range tile.Features {
		newFeatures = append(
			newFeatures,
			featureMod.Feature{
				FeatureType:  feature.FeatureType,
				ModifierType: feature.ModifierType,
				Sides:        feature.Sides.Rotate(rotations),
			},
		)
	}

	tile.Features = newFeatures
	return tile
}

// Returns all possible rotations of the input tile,
// while ensuring that no duplicates are included in the result.
func (tile Tile) GetTileRotations() []Tile {
	rotations := []Tile{tile}
outer:
	for range 3 {
		tile = tile.Rotate(1)
	inner:
		for _, t := range rotations {
			for _, feature := range tile.Features {
				if !slices.Contains(t.Features, feature) {
					break inner
				}
			}
			break outer
		}
		rotations = append(rotations, tile)
	}
	return rotations
}
