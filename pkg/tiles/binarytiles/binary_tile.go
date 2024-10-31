package binarytiles

import (
	"fmt"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	featureMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature/modifier"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

type BinaryTile uint64
type BinaryTileSide uint16

// interpreting BinaryTile's bits:
//      00000000_00000000_1_01_000000011_00_0011_0000010011_0001001100_1000001110
//       X pos    Y pos   ^ ^   meeple   ^    ^     city       road      field
//                       /  |            |    |
//                      / owner playerID |    |
//                     /                 |    |
//             is placed                 |    |
//                                       |    |
//    unconnected field and monastery bits    city shield
//
// counting from right to left (from least to most significant bit):
//  - first four bits of the field section are the corners, starting from top-right, clockwise
//  - first four bits of the roads, cities and shields sections are the sides, starting from top, clockwise
//  - remaining six bits of fields, roads and cities are the connection between pairs
//    of the previous four bits (first and second, second and third, etc.)
//  - first four bits of the meeple section are the sides (same as with shields)
//  - next four meeple bits are the corners (same as with fields)
//  - the last meeple bit is the center (used only by monastery and unconnected field)
//  - the owner bits is just one-hot-encoded player ID. (ID(1) = 00...001, ID(2) = 00...010, etc.)
//  - is placed bit is always 1 on all placed tiles, and 0 on the non-placed tiles
//  - position bits are 8-bit reptesentations of tile position

const (
	featureBitSize  = 10
	modifierBitSize = 4
	meepleBitSize   = 9
	positionBitSize = 8
	maxPlayers      = 2

	connectionBitOffset = 4
	diagonalSideOffset  = 4

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
	playerEndBit   = playerStartBit + maxPlayers

	isPlacedBit    = playerEndBit
	isPlacedEndBit = isPlacedBit + 1

	positionYStartBit = isPlacedEndBit
	positionYEndBit   = positionYStartBit + positionBitSize

	positionXStartBit = positionYEndBit
	positionXEndBit   = positionXStartBit + positionBitSize

	// bit masks
	orthogonalSideMask = 0b0_0000_1111
	diagonalSideMask   = 0b0_1111_0000

	regularFieldMask     = 0b00000000_00000000_0_00_000000000_00_0000_0000000000_0000000000_1111111111
	unconnectedFieldMask = 0b00000000_00000000_0_00_000000000_10_0000_0000000000_0000000000_0000000000
	anyFieldMask         = regularFieldMask | unconnectedFieldMask

	roadMask      = 0b00000000_00000000_0_00_000000000_00_0000_0000000000_1111111111_0000000000
	cityMask      = 0b00000000_00000000_0_00_000000000_00_0000_1111111111_0000000000_0000000000
	shieldMask    = 0b00000000_00000000_0_00_000000000_00_1111_0000000000_0000000000_0000000000
	monasteryMask = 0b00000000_00000000_0_00_000000000_01_0000_0000000000_0000000000_0000000000
	meepleMask    = 0b00000000_00000000_0_00_111111111_00_0000_0000000000_0000000000_0000000000
	ownerMask     = 0b00000000_00000000_0_11_000000000_00_0000_0000000000_0000000000_0000000000
)

const ( // binary tile sides (different from side.Side)
	SideNone   BinaryTileSide = 0b0_0000_0000
	SideCenter BinaryTileSide = 0b1_0000_0000

	SideAllDiagonal   BinaryTileSide = 0b0_1111_0000
	SideAllOrthogonal BinaryTileSide = 0b0_0000_1111

	SideTop    BinaryTileSide = 0b0_0000_0001
	SideRight  BinaryTileSide = 0b0_0000_0010
	SideBottom BinaryTileSide = 0b0_0000_0100
	SideLeft   BinaryTileSide = 0b0_0000_1000

	SideTopRightCorner    BinaryTileSide = 0b0_0001_0000
	SideBottomRightCorner BinaryTileSide = 0b0_0010_0000
	SideBottomLeftCorner  BinaryTileSide = 0b0_0100_0000
	SideTopLeftCorner     BinaryTileSide = 0b0_1000_0000
)

var OrthogonalSides = []BinaryTileSide{
	SideTop,
	SideRight,
	SideBottom,
	SideLeft,
}

var DiagonalSides = []BinaryTileSide{
	SideTopRightCorner,
	SideBottomRightCorner,
	SideBottomLeftCorner,
	SideTopLeftCorner,
}

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

	binaryTile.addPosition(tile.Position)

	if tile.Features != nil {
		// turns out not all PlacedTiles are placed
		binaryTile.setBit(isPlacedBit)
	}

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
				tmpBinaryTile.setBit(meepleStartBit + bitIndex + diagonalSideOffset)
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

// Sets the appropriate owner bit in the binary tile, if the owner ID is not 0. Panics if ownerID is greater than maxPlayers
func (binaryTile *BinaryTile) setOwner(ownerID elements.ID) {
	if ownerID != 0 {
		if ownerID > maxPlayers {
			panic(fmt.Sprintf("cannot use player ID = %#v in binary tile. Max number of players = %#v", ownerID, maxPlayers))
		}
		binaryTile.setBit(playerStartBit + int(ownerID) - 1)
		// -1 because we don't want an empty bit (always zero) for the "NonePlayer" owner
	}
}

// Sets the position X and position Y bits in the binary tile
func (binaryTile *BinaryTile) addPosition(position position.Position) {
	if position.X() > 127 || position.Y() > 127 || position.X() < -128 || position.Y() < -128 {
		panic(fmt.Sprintf("position %#v out of range for binary tile. Allowed range: [-128, 127]", position))
	}
	var tmpBinaryTile BinaryTile
	tmpBinaryTile |= BinaryTile(uint8(position.X()))
	tmpBinaryTile <<= positionBitSize
	tmpBinaryTile |= BinaryTile(uint8(position.Y()))
	tmpBinaryTile <<= positionYStartBit
	*binaryTile |= tmpBinaryTile
}

// Sets the bit at the specified index to 1
func (binaryTile *BinaryTile) setBit(bitIndex int) {
	*binaryTile |= (1 << bitIndex)
}

func (binaryTile BinaryTile) Position() position.Position {
	return position.New(
		int16(int8(binaryTile>>positionXStartBit)),
		int16(int8(binaryTile>>positionYStartBit)),
	)
}

func (binaryTile BinaryTile) HasRegularField() bool { // todo test
	return binaryTile&regularFieldMask != 0
}

func (binaryTile BinaryTile) HasUnconnectedField() bool { // todo test
	return binaryTile&unconnectedFieldMask != 0
}

func (binaryTile BinaryTile) HasAnyField() bool { // todo test
	return binaryTile&anyFieldMask != 0
}

func (binaryTile BinaryTile) HasRoad() bool { // todo test
	return binaryTile&roadMask != 0
}

func (binaryTile BinaryTile) HasCity() bool { // todo test
	return binaryTile&cityMask != 0
}

func (binaryTile BinaryTile) HasMonastery() bool {
	return binaryTile&monasteryMask != 0
}

func (binaryTile BinaryTile) HasMeepleAtSide(side BinaryTileSide) bool { // todo test
	return binaryTile&(BinaryTile(side)<<meepleStartBit) != 0
}

// Returns player ID of meeple at the given side and on the given feature, or elements.NonePlayer if no such meeple exists
func (binaryTile BinaryTile) GetMeepleIDAtSide(side BinaryTileSide, featureType featureMod.Type) elements.ID { // todo test
	ownerID := binaryTile & ownerMask
	if ownerID == 0 {
		return elements.NonePlayer
	} else {
		ownerID >>= playerStartBit
	}

	switch featureType {
	case featureMod.Monastery:
		if binaryTile.HasMonastery() && side == SideCenter {
			return elements.ID(ownerID)
		}

	case featureMod.Field:
		if side == SideCenter && binaryTile&unconnectedFieldMask != 0 {
			return elements.ID(ownerID)
		}
		// todo check normal fields

	case featureMod.City:

	case featureMod.Road:
	}

	return elements.NonePlayer
}

// Returns all sides of the feature(s) of the given type connected to the given side
// Returns SideNone if no feature of the given type is at the given side
//
// In most use cases the sides in "side" argument should belong to only one feature
// For example, in tile with features (Top|Right, Bottom|Left):
// - argument side=Top will return Top|Right
// - argument side=Top|Left will return Top|Right|Bottom|Left
func (binaryTile BinaryTile) GetConnectedSides(side BinaryTileSide, featureType featureMod.Type) BinaryTileSide {
	switch featureType {
	case featureMod.Field:
		binaryTile &= regularFieldMask // todo: the masking *might* be unnecessary and can probably be safely removed
		binaryTile >>= fieldStartBit
		side >>= diagonalSideOffset // field sides are at the first four bits in binaryTiles, despite being diagonal

	case featureMod.Road:
		binaryTile &= roadMask
		binaryTile >>= roadStartBit
		side &= orthogonalSideMask

	case featureMod.City:
		binaryTile &= cityMask
		binaryTile >>= cityStartBit
		side &= orthogonalSideMask

	default:
		panic(fmt.Sprintf("method not supported for features of type: %#v", featureType))
	}

	if BinaryTile(side)&binaryTile == 0 {
		return SideNone
	}

	for i, mask := range connectionMasks {
		sideMask := BinaryTileSide(mask)

		if sideMask&side != 0 { // if any of the mask's sides is present
			if BinaryTile(1<<(i+connectionBitOffset))&binaryTile != 0 { // if connection bit [i] is set
				side |= sideMask
			}
		}
	}

	if featureType == featureMod.Field {
		side <<= diagonalSideOffset // reversing the earlier shift for diagonal sides
	}

	return side
}

// Returns a slice of sides of every feature of the given type of this tile
//
// For example, if a type has two cities, one on top and right sides, and one on the left,
//
//	the return slice will be: {SideTop|SideRight, SideLeft}
func (binaryTile BinaryTile) GetFeaturesOfType(featureType featureMod.Type) []BinaryTileSide {
	var result []BinaryTileSide

	var sidesToCheck [4]BinaryTileSide
	checkedSides := SideNone

	if featureType == featureMod.Field {
		sidesToCheck = [4]BinaryTileSide{SideTopRightCorner, SideBottomRightCorner, SideBottomLeftCorner, SideTopLeftCorner}
	} else {
		sidesToCheck = [4]BinaryTileSide{SideTop, SideRight, SideBottom, SideLeft}
	}

	for _, side := range sidesToCheck {
		if side&checkedSides != 0 {
			continue
		}
		connectedSides := binaryTile.GetConnectedSides(side, featureType)
		if connectedSides != SideNone {
			checkedSides |= connectedSides
			result = append(result, connectedSides)
		}
	}

	return result
}

// Returns whether or not the given side has otherSide
func (side BinaryTileSide) HasSide(otherSide BinaryTileSide) bool {
	return side&otherSide == otherSide // todo copy tests from side
}

// Returns whether or not the given side overlaps otherSide. The overlap does not need to be exact.
func (side BinaryTileSide) OverlapsSide(otherSide BinaryTileSide) bool {
	return side&otherSide != 0
}

