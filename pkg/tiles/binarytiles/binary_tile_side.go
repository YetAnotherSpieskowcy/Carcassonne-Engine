package binarytiles

import (
	"fmt"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

type BinaryTileSide uint16

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

// Returns whether or not the given side has otherSide
func (side BinaryTileSide) HasSide(otherSide BinaryTileSide) bool {
	return side&otherSide == otherSide // todo copy tests from side
}

// Returns whether or not the given side overlaps otherSide. The overlap does not need to be exact.
func (side BinaryTileSide) OverlapsSide(otherSide BinaryTileSide) bool {
	return side&otherSide != 0
}

// Converts corners to sides
// For example:
// - SideTopRightCorner -> SideTop|SideRight
// - SideTopRightCorner|SideTopLeftCorner -> SideTop|SideRight|SideLeft
func (side BinaryTileSide) CornersToSides() BinaryTileSide {
	side &= diagonalSideMask // to clear the center bit

	return ((side >> 4) | (side >> 3) | (side >> 7)) & orthogonalSideMask // don't ask why, this just works.
}

// Converts sides to corners
// For example:
// - SideTop -> SideTopRightCorner|SideTopLeftCorner
// - SideTop|SideLeft -> SideTopRightCorner|SideTopLeftCorner|SideBottomLeftCorner
func (side BinaryTileSide) SidesToCorners() BinaryTileSide {
	return ((side << 4) | (side << 3) | (side << 7)) & diagonalSideMask // don't ask why, this just works.
}

// BinaryTileSide equivalent of position.FromSide(side)
func (side BinaryTileSide) PositionFromSide() position.Position {
	primarySides := 0
	for _, otherSide := range OrthogonalSides {
		if side.OverlapsSide(otherSide) {
			primarySides++
		}
	}

	if primarySides == 0 {
		return position.New(0, 0)
	} else if primarySides == 1 {
		switch {
		case side.OverlapsSide(SideTop):
			return position.New(0, 1)
		case side.OverlapsSide(SideRight):
			return position.New(1, 0)
		case side.OverlapsSide(SideLeft):
			return position.New(-1, 0)
		case side.OverlapsSide(SideBottom):
			return position.New(0, -1)
		}
	}
	panic(fmt.Sprintf("PositionFromSide called with more than one primary side. 'side' = %08b", side))
}

func SideToBinaryTileSide(sideToConvert side.Side, orthogonal bool) BinaryTileSide {
	result := SideNone

	if orthogonal {
		if sideToConvert.OverlapsSide(side.Top) {
			result |= SideTop
		}
		if sideToConvert.OverlapsSide(side.Right) {
			result |= SideRight
		}
		if sideToConvert.OverlapsSide(side.Bottom) {
			result |= SideBottom
		}
		if sideToConvert.OverlapsSide(side.Left) {
			result |= SideLeft
		}

	} else {
		if sideToConvert.OverlapsSide(side.TopRightEdge | side.RightTopEdge) {
			result |= SideTopRightCorner
		}
		if sideToConvert.OverlapsSide(side.RightBottomEdge | side.BottomRightEdge) {
			result |= SideBottomRightCorner
		}
		if sideToConvert.OverlapsSide(side.BottomLeftEdge | side.LeftBottomEdge) {
			result |= SideBottomLeftCorner
		}
		if sideToConvert.OverlapsSide(side.LeftTopEdge | side.TopLeftEdge) {
			result |= SideTopLeftCorner
		}
	}

	return result
}

// Returns a corner neighbouring the given corner(s) from the given direction
// Returns SideNone if no such corner exist
//
// Examples:
// - CornerFromSide(SideTopRightCorner, SideRight) -> SideTopLeftCorner
// - CornerFromSide(SideTopRightCorner, SideTop) -> SideBottomRightCorner
// - CornerFromSide(SideTopRightCorner, SideLeft) -> SideNone
//
// Picture:
//
//	     ---
//	     | |
//	     N-N
//
//	|-N  C-C  N-|
//	| |  | |  | |
//	|-N  C-C  N-|
//
//	     N-N
//	     | |
//	     ---
//
// C - input corners
// N - neighbours
// (each corner has two directions where it returns a neighbour, and two where it returns SideNone)
func CornerFromSide(corner BinaryTileSide, direction BinaryTileSide) BinaryTileSide {
	sideCorners := direction.SidesToCorners()
	corner &= sideCorners

	corner = (corner >> 3) | (corner >> 1) | (corner << 3) | (corner << 1) //neighbouring corners

	return diagonalSideMask & (^sideCorners) & corner
}
