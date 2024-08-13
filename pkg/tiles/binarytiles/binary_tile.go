package binarytiles

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	featureMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature/modifier"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

type BinaryTile int64

// interpreting BinaryTile's bits:
//        0000000000000000001_000000011_00_0011_0000010011_0001001100_1000001110
//         owner (playerID)    meeple   ^    ^     city       road      field
//                                      |    |
//   monastery and unconnected field bits    city shield
//
// counting from right to left (from least to most significant bit):
//  - first four bits of the field section are the corners, starting from top-right, clockwise
//  - first four bits of the roads, cities and shields sections are the sides, starting from top, clockwise
//  - remaining six bits of fields, roads and cities are the connection between pairs
//    of the previous four bits (first and second, second and third, etc.)
//  - first four bits of the meeple section are the sides (same as with shields)
//  - next four meeple bits are the corners (same as with fields)
//  - the last meeple bit is the center
//  - the owner bits is just one-hot-encoded player ID. (ID(1) = 00...001, ID(2) = 00...010, etc.)

const (
	featureBitSize  = 10
	modifierBitSize = 4
	meepleBitSize   = 9

	connectionBitOffset  = 4
	diagonalMeepleOffset = 4

	fieldStartBit = 0
	fieldEndBit   = fieldStartBit + featureBitSize

	roadStartBit = fieldEndBit
	roadEndBit   = roadStartBit + featureBitSize

	cityStartBit = roadEndBit
	cityEndBit   = cityStartBit + featureBitSize

	shieldStartBit = cityEndBit
	shieldEndBit   = shieldStartBit + modifierBitSize

	monasteryBit    = shieldEndBit
	monasteryEndBit = monasteryBit + 1

	unconnectedFieldBit    = monasteryEndBit
	unconnectedFieldEndBit = unconnectedFieldBit + 1

	meepleStartBit = unconnectedFieldEndBit
	meepleEndBit   = meepleStartBit + meepleBitSize

	playerStartBit = meepleEndBit
)

var orthogonalFeaturesBits = []side.Side{
	side.Top,
	side.Right,
	side.Bottom,
	side.Left,
}

var diagonalFeaturesBits = []side.Side{
	side.TopRightEdge | side.RightTopEdge,
	side.RightBottomEdge | side.BottomRightEdge,
	side.BottomLeftEdge | side.LeftBottomEdge,
	side.LeftTopEdge | side.TopLeftEdge,
}

var connectionMasks = []BinaryTile{
	0b0011,
	0b0110,
	0b1100,
	0b1001,
	0b0101,
	0b1010,
}

func fromPlacedFeatures(features []elements.PlacedFeature) BinaryTile {
	var binaryTile BinaryTile

	for _, feature := range features {
		var bitOffset int

		switch feature.FeatureType {
		case featureMod.Field:
			bitOffset = fieldStartBit

		case featureMod.City:
			bitOffset = cityStartBit

		case featureMod.Road:
			bitOffset = roadStartBit

		case featureMod.Monastery:
			binaryTile.setBit(monasteryBit)
			if feature.Meeple.Type != elements.NoneMeeple {
				binaryTile.setOwner(feature.Meeple.PlayerID)
				binaryTile.setBit(meepleEndBit - 1) // last meeple bit is meeple in the center
			}

		default:
			panic("unknown feature type")
		}

		if feature.FeatureType == featureMod.Road || feature.FeatureType == featureMod.City {
			binaryTile.addOrthogonalFeature(feature, bitOffset)

		} else if feature.FeatureType == featureMod.Field {
			if feature.Sides == side.NoSide {
				binaryTile.setBit(unconnectedFieldBit)
				if feature.Meeple.Type != elements.NoneMeeple {
					binaryTile.setOwner(feature.Meeple.PlayerID)
					binaryTile.setBit(meepleEndBit - 1) // last meeple bit is meeple in the center
				}

			} else {
				binaryTile.addDiagonalFeature(feature, bitOffset)
			}
		}
	}

	return binaryTile
}

func FromTile(tile tiles.Tile) BinaryTile {
	binaryTile := fromPlacedFeatures(elements.ToPlacedTile(tile).Features)
	return binaryTile
}

func FromPlacedTile(tile elements.PlacedTile) BinaryTile {
	binaryTile := fromPlacedFeatures(tile.Features)

	return binaryTile
}

// Sets all necessary bits in the binary tile for a diagonal feature (field)
func (binaryTile *BinaryTile) addDiagonalFeature(feature elements.PlacedFeature, bitOffset int) {
	var tmpBinaryTile BinaryTile

	for bitIndex, side := range diagonalFeaturesBits {
		if feature.Sides.OverlapsSide(side) {
			tmpBinaryTile.setBit(bitOffset + bitIndex)

			// todo add more meeple types when they are implemented
			if feature.Meeple.Type != elements.NoneMeeple {
				binaryTile.setOwner(feature.Meeple.PlayerID)
				tmpBinaryTile.setBit(meepleStartBit + bitIndex + diagonalMeepleOffset)
			}
		}
	}
	for bitIndex, bitMask := range connectionMasks {
		bitMask <<= bitOffset
		if tmpBinaryTile&bitMask == bitMask {
			tmpBinaryTile.setBit(bitOffset + bitIndex + connectionBitOffset)
		}
	}
	*binaryTile |= tmpBinaryTile
}

// Sets all necessary bits in the binary tile for an orthogonal feature (city, road). Also handles city shields
func (binaryTile *BinaryTile) addOrthogonalFeature(feature elements.PlacedFeature, bitOffset int) {
	var tmpBinaryTile BinaryTile

	for bitIndex, side := range orthogonalFeaturesBits {
		if feature.Sides.HasSide(side) {
			tmpBinaryTile.setBit(bitOffset + bitIndex)

			// todo add more feature modifiers when they are implemented
			if feature.ModifierType == modifier.Shield {
				tmpBinaryTile.setBit(shieldStartBit + bitIndex)
			}

			// todo add more meeple types when they are implemented
			if feature.Meeple.Type != elements.NoneMeeple {
				binaryTile.setOwner(feature.Meeple.PlayerID)
				tmpBinaryTile.setBit(meepleStartBit + bitIndex)
			}
		}
	}
	for bitIndex, bitMask := range connectionMasks {
		bitMask <<= bitOffset
		if tmpBinaryTile&bitMask == bitMask {
			tmpBinaryTile.setBit(bitOffset + bitIndex + connectionBitOffset)
		}
	}
	*binaryTile |= tmpBinaryTile
}

// Sets the appropriate owner bit in the binary tile, if the owner ID is not 0
func (binaryTile *BinaryTile) setOwner(ownerID elements.ID) {
	if ownerID != 0 {
		binaryTile.setBit(playerStartBit + int(ownerID) - 1)
		// -1 because we don't want an empty bit (always zero) for the "NonePlayer" owner
	}
}

// Sets the bit at the specified index to 1
func (binaryTile *BinaryTile) setBit(bitIndex int) {
	*binaryTile |= (1 << bitIndex)
}
