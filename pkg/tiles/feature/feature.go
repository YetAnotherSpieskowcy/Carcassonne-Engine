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
	Sides      side.Side
	BinaryData uint8
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

func New(featureType Type, sides side.Side, modifierType ...modifier.Type) Feature {
	var binaryData uint8
	binaryData |= uint8(featureType)
	if len(modifierType) != 0 && modifierType[0] != modifier.NoneType {
		binaryData |= modifierMask
	}

	return Feature{Sides: sides, BinaryData: binaryData}
}

func NewWithMeeple(featureType Type, sides side.Side, ownerID uint8, modifierType ...modifier.Type) Feature { // todo elements.ID
	feature := New(featureType, sides, modifierType...)

	feature.BinaryData |= (ownerID << ownerStartBit)

	return feature
}

func (feature Feature) Type() Type {
	return Type(feature.BinaryData & featureTypeMask)
}

func (feature Feature) OwnerID() uint8 {
	return feature.BinaryData >> ownerStartBit
}

func (feature Feature) ModifierType() modifier.Type {
	// TODO: currently hardcoded to one modifier type. Change this when adding more
	if feature.HasModifier() {
		return modifier.Shield
	}
	return modifier.NoneType
}

func (feature Feature) HasModifier() bool {
	return (feature.BinaryData & modifierMask) != 0
}

func (feature Feature) HasMeeple() bool {
	return (feature.BinaryData & meepleMask) != 0
}

// structs exposed through the bindings do not implement `__eq__`
// so temporarily implement this as a workaround
// TODO: nuke this after bindings depend only on the binary tile representation
func (feature Feature) Equals(other Feature) bool {
	return feature == other
}
