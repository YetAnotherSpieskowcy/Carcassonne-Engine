package feature

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature/modifier"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

type Type uint8

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
	binaryData   uint8
}

// interpreting Feature's binaryData:
//      XX_00_0_000
//      ^  ^  ^  ^
// ignored |  | featureType
//         |  |
//   ownerID  hasModifier

const (
	featureTypeMask = 0b0000_0111
	modifierMask    = 0b0000_1000
	meepleMask      = 0b0011_0000
	ownerStartBit   = 4
)

func New(sides side.Side, featureType Type, hasModifier bool) Feature {
	var binaryData uint8
	binaryData |= uint8(featureType)
	if hasModifier {
		binaryData |= modifierMask
	}

	return Feature{Sides: sides, binaryData: binaryData}
}

func NewWithMeeple(sides side.Side, featureType Type, hasModifier bool, ownerID uint8) Feature { // todo elements.ID
	feature := New(sides, featureType, hasModifier)

	feature.binaryData |= (ownerID << ownerStartBit)

	return feature
}

func (feature Feature) Type() Type {
	return Type(feature.binaryData & featureTypeMask)
}

func (feature Feature) OwnerID() uint8 {
	return feature.binaryData >> ownerStartBit
}

func (feature Feature) HasModifier() bool {
	return (feature.binaryData & modifierMask) != 0
}

func (feature Feature) HasMeeple() bool {
	return (feature.binaryData & meepleMask) != 0
}

// structs exposed through the bindings do not implement `__eq__`
// so temporarily implement this as a workaround
// TODO: nuke this after bindings depend only on the binary tile representation
func (feature Feature) Equals(other Feature) bool {
	return feature == other
}
