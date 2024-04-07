package connection

type Side int64

const (
	NONE Side = iota
	TOP
	RIGHT
	LEFT
	BOTTOM

	CENTER

	//for farmers
	TOP_LEFT_CORNER
	TOP_RIGHT_CORNER
	BOTTOM_LEFT_CORNER
	BOTTOM_RIGHT_CORNER

	//for fields

	/*Left side of top edge*/
	TOP_LEFT_EDGE
	/*Right side of top edge*/
	TOP_RIGHT_EDGE

	/*Top side of right edge*/
	RIGHT_TOP_EDGE
	/*Bottom side of right edge*/
	RIGHT_BOTTOM_EDGE

	/*Top side of left edge*/
	LEFT_TOP_EDGE
	/*Bottom side of left edge*/
	LEFT_BOTTOM_EDGE

	/*Left side of bottom edge*/
	BOTTOM_LEFT_EDGE
	/*Right side of bottom edge*/
	BOTTOM_RIGHT_EDGE
)

func (side Side) String() string {

	switch side {
	case TOP:
		return "TOP"
	case RIGHT:
		return "RIGHT"
	case LEFT:
		return "LEFT"
	case BOTTOM:
		return "BOTTOM"

	case TOP_LEFT_CORNER:
		return "TOP_LEFT_CORNER"
	case TOP_RIGHT_CORNER:
		return "TOP_RIGHT_CORNER"
	case BOTTOM_LEFT_CORNER:
		return "BOTTOM_LEFT_CORNER"
	case BOTTOM_RIGHT_CORNER:
		return "BOTTOM_RIGHT_CORNER"

	case CENTER:
		return "CENTER"

	case TOP_LEFT_EDGE:
		return "TOP_LEFT_EDGE"
	case TOP_RIGHT_EDGE:
		return "TOP_RIGHT_EDGE"
	case RIGHT_TOP_EDGE:
		return "RIGHT_TOP_EDGE"
	case RIGHT_BOTTOM_EDGE:
		return "RIGHT_BOTTOM_EDGE"

	case LEFT_TOP_EDGE:
		return "LEFT_TOP_EDGE"
	case LEFT_BOTTOM_EDGE:
		return "LEFT_BOTTOM_EDGE"
	case BOTTOM_LEFT_EDGE:
		return "BOTTOM_LEFT_EDGE"
	case BOTTOM_RIGHT_EDGE:
		return "BOTTOM_RIGHT_EDGE"

	case NONE:
		return "NONE"
	default:
		return "ERROR"
	}
}

/*
Rotates side clockwise
*/
func (side Side) Rotate(rotations uint) Side {
	//limit rotations
	rotations = rotations % 4
	var result = side
	for rotations > 0 {
		switch side {
		case TOP:
			result = RIGHT
		case RIGHT:
			result = BOTTOM
		case LEFT:
			result = TOP
		case BOTTOM:
			result = LEFT

		case TOP_LEFT_CORNER:
			result = TOP_RIGHT_CORNER
		case TOP_RIGHT_CORNER:
			result = BOTTOM_RIGHT_CORNER
		case BOTTOM_LEFT_CORNER:
			result = TOP_LEFT_CORNER
		case BOTTOM_RIGHT_CORNER:
			result = BOTTOM_LEFT_CORNER

		case TOP_LEFT_EDGE:
			result = RIGHT_TOP_EDGE
		case TOP_RIGHT_EDGE:
			result = RIGHT_BOTTOM_EDGE
		case RIGHT_TOP_EDGE:
			result = BOTTOM_RIGHT_EDGE
		case RIGHT_BOTTOM_EDGE:
			result = BOTTOM_LEFT_EDGE

		case LEFT_TOP_EDGE:
			result = TOP_RIGHT_EDGE
		case LEFT_BOTTOM_EDGE:
			result = TOP_LEFT_EDGE
		case BOTTOM_LEFT_EDGE:
			result = LEFT_TOP_EDGE
		case BOTTOM_RIGHT_EDGE:
			result = LEFT_BOTTOM_EDGE

		case CENTER:
			result = CENTER
		default:
			result = NONE
		}
		rotations--
	}
	return result
}

type Connection struct {
	Sides []Side
}

func (connection Connection) Rotate(rotations uint) Connection {
	var result Connection
	for _, side := range connection.Sides {
		result.Sides = append(result.Sides, side.Rotate(rotations))
	}

	return result
}

func (connection Connection) String() string {
	var result string
	for _, side := range connection.Sides {
		result += side.String() + " "
	}
	return result
}
