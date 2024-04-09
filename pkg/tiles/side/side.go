package side

type Side int64

const (
	None Side = iota
	Top
	Right
	Left
	Bottom

	Center

	// for farmers
	TopLeftCorner
	TopRightCorner
	BottomLeftCorner
	BottomRightCorner

	// for fields

	/* Left side of top edge */
	TopLeftEdge
	/* Right side of top edge */
	TopRightEdge

	/* Top side of right edge */
	RightTopEdge
	/* Bottom side of right edge */
	RightBottomEdge

	/* Top side of left edge */
	LeftTopEdge
	/* Bottom side of left edge */
	LeftBottomEdge

	/* Left side of bottom edge */
	BottomLeftEdge
	/* Right side of bottom edge */
	BottomRightEdge
)

func (side Side) String() string { //nolint:gocyclo // splitting into multiple switches would be obscure

	switch side {
	case Top:
		return "TOP"
	case Right:
		return "RIGHT"
	case Left:
		return "LEFT"
	case Bottom:
		return "BOTTOM"

	case TopLeftCorner:
		return "TOP_LEFT_CORNER"
	case TopRightCorner:
		return "TOP_RIGHT_CORNER"
	case BottomLeftCorner:
		return "BOTTOM_LEFT_CORNER"
	case BottomRightCorner:
		return "BOTTOM_RIGHT_CORNER"

	case Center:
		return "CENTER"

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

	case TopLeftEdge:
		return "TOP_LEFT_EDGE"
	case TopRightEdge:
		return "TOP_RIGHT_EDGE"
	case RightTopEdge:
		return "RIGHT_TOP_EDGE"
	case RightBottomEdge:
		return "RIGHT_BOTTOM_EDGE"

	case LeftTopEdge:
		return "LEFT_TOP_EDGE"
	case LeftBottomEdge:
		return "LEFT_BOTTOM_EDGE"
	case BottomLeftEdge:
		return "BOTTOM_LEFT_EDGE"
	case BottomRightEdge:
		return "BOTTOM_RIGHT_EDGE"

	case None:
		return "NONE"
	default:
		return "ERROR"
	}
}

/*
Rotates side clockwise
*/
func (side Side) Rotate(rotations uint) Side { //nolint:gocyclo // splitting into multiple switches would be obscure
	// limit rotations
	rotations %= 4
	var result = side
	for rotations > 0 {
		switch side {
		case Top:
			result = Right
		case Right:
			result = Bottom
		case Left:
			result = Top
		case Bottom:
			result = Left

		case TopLeftCorner:
			result = TopRightCorner
		case TopRightCorner:
			result = BottomRightCorner
		case BottomLeftCorner:
			result = TopLeftCorner
		case BottomRightCorner:
			result = BottomLeftCorner

		case TopLeftEdge:
			result = RightTopEdge
		case TopRightEdge:
			result = RightBottomEdge
		case RightTopEdge:
			result = BottomRightEdge
		case RightBottomEdge:
			result = BottomLeftEdge

		case LeftTopEdge:
			result = TopRightEdge
		case LeftBottomEdge:
			result = TopLeftEdge
		case BottomLeftEdge:
			result = LeftTopEdge
		case BottomRightEdge:
			result = LeftBottomEdge

		case Center:
			result = Center
		default:
			result = None
		}
		rotations--
	}
	return result
}

func RotateSideArray(sides []Side, rotations uint) []Side {
	var rotatedSides []Side
	for _, side := range sides {
		rotatedSides = append(rotatedSides, side.Rotate(rotations))
	}
	return rotatedSides
}
