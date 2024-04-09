package feature

import (
	"slices"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

type Type int64

const (
	None Type = iota
	Road
	City
	Field
)

type Feature struct {
	FeatureType Type
	Sides       []side.Side
}

func (feature Feature) Equals(other Feature) bool {
	if feature.FeatureType != other.FeatureType {
		return false
	}
	return slices.Equal(feature.Sides, other.Sides)
}
