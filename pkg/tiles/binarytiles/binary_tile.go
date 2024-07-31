package binarytiles

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	featureMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature/modifier"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

type BinaryTile int64

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

var connectionMasks = []int64{
	0b0011,
	0b0110,
	0b1100,
	0b1001,
	0b0101,
	0b1010,
}

func fromPlacedFeatures(features []elements.PlacedFeature) BinaryTile {
	var bits int64
	var owner elements.ID

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
			bits = setBit(bits, monasteryBit)
			if feature.Meeple.Type != elements.NoneMeeple {
				owner = feature.Meeple.PlayerID
				bits = setBit(bits, meepleEndBit-1) // last bit is meeple in the center
			}

		default:
			panic("unknown feature type")
		}

		if feature.FeatureType == featureMod.Road || feature.FeatureType == featureMod.City {
			var tmpBits int64
			for bitIndex, side := range orthogonalFeaturesBits {
				if feature.Sides.HasSide(side) {
					tmpBits = setBit(tmpBits, bitOffset+bitIndex)

					// todo add more feature modifiers when they are implemented
					if feature.ModifierType == modifier.Shield {
						tmpBits = setBit(tmpBits, shieldStartBit+bitIndex)
					}

					// todo add more meeple types when they are implemented
					if feature.Meeple.Type != elements.NoneMeeple {
						owner = feature.Meeple.PlayerID
						tmpBits = setBit(tmpBits, meepleStartBit+bitIndex)
					}
				}
			}
			for bitIndex, bitMask := range connectionMasks {
				bitMask = bitMask << int64(bitOffset)
				if tmpBits&bitMask == bitMask {
					tmpBits = setBit(tmpBits, bitOffset+bitIndex+connectionBitOffset)
				}
			}
			bits = bits | tmpBits

		} else if feature.FeatureType == featureMod.Field {
			if feature.Sides == side.NoSide {
				bits = setBit(bits, unconnectedFieldBit)
				if feature.Meeple.Type != elements.NoneMeeple {
					owner = feature.Meeple.PlayerID
					bits = setBit(bits, meepleEndBit-1) // last bit is meeple in the center
				}
			} else {
				var tmpBits int64
				for bitIndex, side := range diagonalFeaturesBits {
					if feature.Sides.OverlapsSide(side) {
						tmpBits = setBit(tmpBits, bitOffset+bitIndex)

						// todo add more meeple types when they are implemented
						if feature.Meeple.Type != elements.NoneMeeple {
							owner = feature.Meeple.PlayerID
							tmpBits = setBit(tmpBits, meepleStartBit+bitIndex+diagonalMeepleOffset)
						}
					}
				}
				for bitIndex, bitMask := range connectionMasks {
					bitMask = bitMask << int64(bitOffset)
					if tmpBits&bitMask == bitMask {
						tmpBits = setBit(tmpBits, bitOffset+bitIndex+connectionBitOffset)
					}
				}
				bits = bits | tmpBits
			}
		}
	}

	if owner != 0 {
		bits = setBit(bits, playerStartBit+int(owner)-1)
	}

	return BinaryTile(bits)
}

func FromTile(tile tiles.Tile) BinaryTile {
	binaryTile := fromPlacedFeatures(elements.ToPlacedTile(tile).Features)
	return binaryTile
}

func FromPlacedTile(tile elements.PlacedTile) BinaryTile {
	binaryTile := fromPlacedFeatures(tile.Features)

	return binaryTile
}

func setBit(number int64, bitIndex int) int64 {
	return number | (1 << bitIndex)
}
