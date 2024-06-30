package feature

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature/modifier"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

type Type int64

const (
	None Type = iota
	Road
	City
	Field
	Monastery
)

type Feature struct {
	FeatureType  Type
	ModifierType modifier.Type
	Sides        side.Side
}
