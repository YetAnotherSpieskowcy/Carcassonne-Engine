package farm_connection

/**
First direction is from center, then specyfing which side of edge
*/
type FarmSide int64

const (
	NONE FarmSide = iota
	TOP_LEFT
	TOP_RIGHT

	RIGHT_TOP
	RIGHT_BOTTOM

	LEFT_TOP
	LEFT_BOTTOM

	BOTTOM_LEFT
	BOTTOM_RIGHT

	CENTER
)

func (side FarmSide) ToString() string {

	switch side {
	case TOP_LEFT:
		return "TOP_LEFT"
	case TOP_RIGHT:
		return "TOP_RIGHT"
	case RIGHT_TOP:
		return "RIGHT_TOP"
	case RIGHT_BOTTOM:
		return "RIGHT_BOTTOM"

	case LEFT_TOP:
		return "LEFT_TOP"
	case LEFT_BOTTOM:
		return "LEFT_BOTTOM"
	case BOTTOM_LEFT:
		return "BOTTOM_LEFT"
	case BOTTOM_RIGHT:
		return "BOTTOM_RIGHT"
	case CENTER:
		return "CENTER"
	case NONE:
		return "NONE"
	default:
		return "ERROR"
	}
}

/*
Rotates Farmside clockwise
*/
func (side FarmSide) Rotate(rotations uint) FarmSide {
	//limit rotations
	rotations = rotations % 4
	var result = side
	for rotations > 0 {
		switch side {
		case TOP_LEFT:
			result = RIGHT_TOP
		case TOP_RIGHT:
			result = RIGHT_BOTTOM
		case RIGHT_TOP:
			result = BOTTOM_RIGHT
		case RIGHT_BOTTOM:
			result = BOTTOM_LEFT

		case LEFT_TOP:
			result = TOP_RIGHT
		case LEFT_BOTTOM:
			result = TOP_LEFT
		case BOTTOM_LEFT:
			result = LEFT_TOP
		case BOTTOM_RIGHT:
			result = LEFT_BOTTOM
		case CENTER:
			result = CENTER
		case NONE:
			result = NONE
		default:
			result = NONE
		}
		rotations--
	}
	return result
}

type FarmConnection struct {
	A FarmSide
	B FarmSide
}

func (connection FarmConnection) Rotate(rotations uint) FarmConnection {
	var result FarmConnection
	result.A = connection.A.Rotate(rotations)
	result.B = connection.B.Rotate(rotations)
	return result
}

func (connection FarmConnection) ToString() string {
	return connection.A.ToString() + " " + connection.B.ToString()
}

func NewFarmConnection(A FarmSide, B FarmSide) FarmConnection {
	var result FarmConnection
	result.A = A
	result.B = B
	return result
}
