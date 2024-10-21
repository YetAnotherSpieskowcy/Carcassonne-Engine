package tiles

import (
	"slices"

	featureMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	sideMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

/*
Immutable object
*/
type Tile struct {
	Features []featureMod.Feature
}

// checks if two tiles are the same, ignoring their orientation
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

// checks if two tiles are the same, including their orientation
func (tile Tile) ExactEquals(other Tile) bool {
	return slices.Equal(tile.Features, other.Features)
}

func (tile Tile) Cities() []featureMod.Feature {
	var cities []featureMod.Feature
	for _, feature := range tile.Features {
		if feature.Type() == featureMod.City {
			cities = append(cities, feature)
		}
	}
	return cities
}

func (tile Tile) Roads() []featureMod.Feature {
	var roads []featureMod.Feature
	for _, feature := range tile.Features {
		if feature.Type() == featureMod.Road {
			roads = append(roads, feature)
		}
	}
	return roads
}

func (tile Tile) Fields() []featureMod.Feature {
	var fields []featureMod.Feature
	for _, feature := range tile.Features {
		if feature.Type() == featureMod.Field {
			fields = append(fields, feature)
		}
	}
	return fields
}

func (tile Tile) Monastery() *featureMod.Feature {
	for i, feature := range tile.Features {
		if feature.Type() == featureMod.Monastery {
			return &tile.Features[i]
		}
	}
	return nil
}

/*
Rotate tile clockwise
*/
func (tile Tile) Rotate(rotations uint) Tile {
	rotations %= 4
	if rotations == 0 {
		return tile
	}

	var newFeatures []featureMod.Feature

	for _, feature := range tile.Features {
		newFeatures = append(
			newFeatures,
			featureMod.New(
				feature.Type(),
				feature.Sides.Rotate(rotations),
				feature.ModifierType(),
			),
		)
	}

	tile.Features = newFeatures
	return tile
}

/*
Return the feature of certain type on desired side
*/
func (tile *Tile) GetFeatureAtSide(sideToCheck sideMod.Side, featureType featureMod.Type) *featureMod.Feature {
	for _, feature := range tile.Features {
		if feature.Sides.HasSide(sideToCheck) && feature.Type() == featureType {
			return &feature
		}
	}
	return nil
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
