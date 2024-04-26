package feature

import (
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
	Sides       side.Side
}
