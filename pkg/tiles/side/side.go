package side

type Side uint8

const (
	/* Left side of top edge */
	TopLeftEdge Side = 0b1000_0000

	/* Right side of top edge */
	TopRightEdge Side = 0b0100_0000

	/* Top side of right edge */
	RightTopEdge Side = 0b0010_0000

	/* Bottom side of right edge */
	RightBottomEdge Side = 0b0001_0000

	/* Right side of bottom edge */
	BottomRightEdge Side = 0b0000_1000

	/* Left side of bottom edge */
	BottomLeftEdge Side = 0b0000_0100

	/* Bottom side of left edge */
	LeftBottomEdge Side = 0b0000_0010

	/* Top side of left edge */
	LeftTopEdge Side = 0b0000_0001

	Top    Side = 0b1100_0000
	Right  Side = 0b0011_0000
	Bottom Side = 0b0000_1100
	Left   Side = 0b0000_0011

	NoSide Side = 0b0000_0000

	All Side = 0b1111_1111
)

var PrimarySides = []Side{
	Top,
	Right,
	Bottom,
	Left,
}

var EdgeSides = []Side{
	TopLeftEdge,
	TopRightEdge,
	RightTopEdge,
	RightBottomEdge,
	BottomRightEdge,
	BottomLeftEdge,
	LeftBottomEdge,
	LeftTopEdge,
}

func (side Side) String() string {
	type sideNamesStruct struct {
		primary     Side
		primaryName string
		secondary   map[Side]string
	}

	sideNames := []sideNamesStruct{
		{
			primary:     Top,
			primaryName: "TOP",
			secondary: map[Side]string{
				TopLeftEdge:  "TOP_LEFT_EDGE",
				TopRightEdge: "TOP_RIGHT_EDGE",
			},
		},
		{
			primary:     Right,
			primaryName: "RIGHT",
			secondary: map[Side]string{
				RightTopEdge:    "RIGHT_TOP_EDGE",
				RightBottomEdge: "RIGHT_BOTTOM_EDGE",
			},
		},
		{
			primary:     Bottom,
			primaryName: "BOTTOM",
			secondary: map[Side]string{
				BottomRightEdge: "BOTTOM_RIGHT_EDGE",
				BottomLeftEdge:  "BOTTOM_LEFT_EDGE",
			},
		},
		{
			primary:     Left,
			primaryName: "LEFT",
			secondary: map[Side]string{
				LeftBottomEdge: "LEFT_BOTTOM_EDGE",
				LeftTopEdge:    "LEFT_TOP_EDGE",
			},
		},
	}
	/*
		First direction indicates the main edge of square, the second tells which side of the edge.
		Example:"
		TopLeftEdge
		 <______
				|
				|
				|
			tile center
	*/
	output := ""

	for _, names := range sideNames {
		if side.HasSide(names.primary) {
			output += names.primaryName
		} else {
			for key, value := range names.secondary {
				if side.HasSide(key) {
					output += value
				}
			}
		}
	}

	if output == "" {
		output = "NO_SIDE"
	}

	return output
}

// Returns whether or not the given side has otherSide
// For example:
// - (Right|Top).HasSide(Right) == true
// - (Right|Top).HasSide(TopRightEdge) == true
// - (Right|Top).HasSide(Left) == false
// - (AnySide).HasSide(NoSide) == always True !
func (side Side) HasSide(otherSide Side) bool {
	return side&otherSide == otherSide
}

// Returns whether or not the given side overlaps otherSide. The overlap does not need to be exact.
// For example:
// - (Right|Top).OverlapsSide(Right|Bottom) == true
// - (Right|Top).OverlapsSide(TopRightEdge|Bottom|Left) == true
// - (Right|Top).OverlapsSide(Left|BottomLeftEdge) == false
func (side Side) OverlapsSide(otherSide Side) bool {
	return side&otherSide != 0
}

// Rotates side clockwise
func (side Side) Rotate(rotations uint) Side {
	/*
		TopLeftEdge     0b10000000
		TopRightEdge    0b01000000

		RightTopEdge    0b00100000
		RightBottomEdge 0b00010000

		BottomRightEdge 0b00001000
		BottomLeftEdge  0b00000100

		LeftBottomEdge  0b00000010
		LeftTopEdge     0b00000001

		Top             0b11000000
		Right           0b00110000
		Bottom          0b00001100
		Left            0b00000011

		NoSide          0b00000000
	*/

	rotations %= 4
	if rotations == 0 {
		return side
	}

	var shift = rotations * 2
	return (side >> shift) | (side << (8 - shift)) // circular bitshift (bitwise rotate) side to the right by {2*rotations} bits
}

/*
Mirrors the side:
TopLeftEdge     <->  BottomLeftEdge
TopRightEdge    <->  BottomRightEdge

RightTopEdge    <->  LeftTopEdge
RightBottomEdge <->  LeftBottomEdge
*/
func (side Side) Mirror() Side {
	return side.Rotate(2).FlipSides()
}

/*
Flips each part of the side relative to the side's center:
TopLeftEdge     <->  TopRightEdge
RightTopEdge    <->  RightBottomEdge

BottomLeftEdge  <->  BottomRightEdge
LeftTopEdge     <->  LeftBottomEdge
*/
func (side Side) FlipSides() Side {
	// swap bits in each pair (abcdefgh -> badcfehg)
	return ((side & 0b10101010) >> 1) | ((side & 0b01010101) << 1)
}

/*
Flips each part of the side relative to the adjacent corner:
TopLeftEdge     <->  LeftTopEdge
TopRightEdge    <->  RightTopEdge

BottomLeftEdge  <->  LeftBottomEdge
BottomRightEdge <->  RightBottomEdge
*/
func (side Side) FlipCorners() Side {
	// swap bits in each pair, but offset by one (abcdefgh -> hcbedgfa)
	return ((side & 0b01010100) >> 1) | ((side & 0b00101010) << 1) | ((side & 0b10000000) >> 7) | ((side & 0b00000001) << 7)
}

/*
Returns other connected side on the same tile.
It allows getting other side of the road feature.
direction must indicate only one cardinal direction!
*/
func (side Side) GetConnectedOtherCardinalDirection(direction Side) Side {
	for _, cardinal := range PrimarySides {
		if side.HasSide(cardinal) && cardinal != direction {
			return cardinal
		}
	}
	return NoSide
}

/*
Return nth existing direction indicated by Side.
For example Side indicates Top,Right,Bottom at once.
First cardinal direction would be Top, second Right, third Bottom.
If nth direction doesn't exist, NoSide is returned.
*/
func (side Side) GetNthCardinalDirection(n uint8) Side {
	found := uint8(0)
	for _, cardinal := range PrimarySides {
		if side.HasSide(cardinal) {
			found++
		}
		if found > n {
			return cardinal
		}
	}
	return NoSide
}

func (side Side) GetCardinalDirectionsLength() int {
	found := int(0)
	for _, cardinal := range PrimarySides {
		if side.HasSide(cardinal) {
			found++
		}
	}
	return found
}
