package feature

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature/modifier"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

type Type int64

const (
	NoneType Type = iota
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

// structs exposed through the bindings do not implement `__eq__`
// so temporarily implement this as a workaround
// TODO: nuke this after bindings depend only on the binary tile representation
func (feature Feature) Equals(other Feature) bool {
	return feature == other
}
