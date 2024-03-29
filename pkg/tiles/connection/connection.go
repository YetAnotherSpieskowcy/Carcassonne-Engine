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

	TOPLEFT
	TOPRIGHT
	BOTTOMLEFT
	BOTTOMRIGHT
)

func (side Side) ToString() string {

	switch side {
	case TOP:
		return "TOP"
	case RIGHT:
		return "RIGHT"
	case LEFT:
		return "LEFT"
	case BOTTOM:
		return "BOTTOM"

	case TOPLEFT:
		return "TOPLEFT"
	case TOPRIGHT:
		return "TOPRIGHT"
	case BOTTOMLEFT:
		return "BOTTOMLEFT"
	case BOTTOMRIGHT:
		return "BOTTOMRIGHT"
	case CENTER:
		return "CENTER"
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

		case TOPLEFT:
			result = TOPRIGHT
		case TOPRIGHT:
			result = BOTTOMRIGHT
		case BOTTOMLEFT:
			result = TOPLEFT
		case BOTTOMRIGHT:
			result = BOTTOMLEFT
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

func (connection Connection) ToString() string {
	var result string
	for _, side := range connection.Sides {
		result += side.ToString() + " "
	}
	return result
}
