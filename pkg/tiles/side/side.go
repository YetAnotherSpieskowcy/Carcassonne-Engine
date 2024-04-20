package side

type Side int16

// The order of sides in this enum matters for some bitwise operations in Rotate().
const (
	Top Side = 1 << iota
	Right
	Bottom
	Left

	// for fields

	/* Left side of top edge */
	TopLeftEdge

	/* Top side of right edge */
	RightTopEdge

	/* Right side of bottom edge */
	BottomRightEdge

	/* Bottom side of left edge */
	LeftBottomEdge

	/* Top side of left edge */
	LeftTopEdge

	/* Right side of top edge */
	TopRightEdge

	/* Bottom side of right edge */
	RightBottomEdge

	/* Left side of bottom edge */
	BottomLeftEdge

	Center

	None Side = 0
)

func (side Side) String() string { //nolint:gocyclo // splitting into multiple switches would be obscure
	sideNames := map[Side]string{
		Top:             "TOP",
		Right:           "RIGHT",
		Bottom:          "BOTTOM",
		Left:            "LEFT",
		TopLeftEdge:     "TOP_LEFT_EDGE",
		RightTopEdge:    "RIGHT_TOP_EDGE",
		BottomRightEdge: "BOTTOM_RIGHT_EDGE",
		LeftBottomEdge:  "LEFT_BOTTOM_EDGE",
		LeftTopEdge:     "LEFT_TOP_EDGE",
		TopRightEdge:    "TOP_RIGHT_EDGE",
		RightBottomEdge: "RIGHT_BOTTOM_EDGE",
		BottomLeftEdge:  "BOTTOM_LEFT_EDGE",
		Center:          "CENTER",
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
	for key, value := range sideNames {
		if side&key != 0 {
			output += value
		}
	}

	if output == "" {
		output = "NONE"
	}

	return output
}

/*
Rotates side clockwise
*/
func (side Side) Rotate(rotations uint) Side {
	/*
		Top             0b0000000000000001
		Right           0b0000000000000010
		Bottom          0b0000000000000100
		Left            0b0000000000001000

		TopLeftEdge     0b0000000000010000
		RightTopEdge    0b0000000000100000
		BottomRightEdge 0b0000000001000000
		LeftBottomEdge  0b0000000010000000

		LeftTopEdge     0b0000000100000000
		TopRightEdge    0b0000001000000000
		RightBottomEdge 0b0000010000000000
		BottomLeftEdge  0b0000100000000000

		Center          0b0001000000000000
		None            0b0000000000000000
	*/

	// limit rotations
	rotations %= 4
	if rotations == 0 {
		return side
	}

	// center doesn't rotate
	hasCenter := side&Center != 0
	result := side & ^Center

	for rotations > 0 {
		result <<= 1

		// loop the last 4 bits (rotated Left goes back into Top)
		if result&0b1_0000 != 0 {
			result &= ^(1 << 4) // clear the overflowed bit
			result |= (1 << 0)  // set the bit in correct position
		}

		// loop the next 4 bits (rotated LeftBottomEdge goes back into TopLeftEdge)
		if result&0b1_0000_0000 != 0 {
			result &= ^(1 << 8) // clear the overflowed bit
			result |= (1 << 4)  // set the bit in correct position
		}

		// loop the next 4 bits (rotated BottomLeftEdge goes back into LeftTopEdge)
		if result&0b1_0000_0000_0000 != 0 {
			result &= ^(1 << 12) // clear the overflowed bit
			result |= (1 << 8)   // set the bit in correct position
		}
		rotations--
	}

	if hasCenter {
		result |= Center
	}
	return result
}
