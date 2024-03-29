package farm_connection

/**
First direction is from center, then specyfing which side of edge
*/
type FarmSide int64

const (
	TOP_LEFT FarmSide = iota
	TOP_RIGHT

	RIGHT_TOP
	RIGHT_BOTTOM

	LEFT_TOP
	LEFT_BOTTOM

	BOTTOM_LEFT
	BOTTOM_RIGHT

	CENTER
)

/*
Rotates Farmside clockwise
*/
func (side FarmSide) Rotate(rotations int) FarmSide {
	//limit rotations
	rotations = rotations % 4
	//check if more rotations needed
	if rotations > 1 {
		return side.Rotate(rotations - 1)
		//check if doesn't need to rotate
	} else if rotations == 0 {
		return side
		//rotate once
	} else {
		switch side {
		case TOP_LEFT:
			return RIGHT_TOP
		case TOP_RIGHT:
			return RIGHT_BOTTOM
		case RIGHT_TOP:
			return BOTTOM_RIGHT
		case RIGHT_BOTTOM:
			return BOTTOM_LEFT

		case LEFT_TOP:
			return TOP_RIGHT
		case LEFT_BOTTOM:
			return TOP_LEFT
		case BOTTOM_LEFT:
			return LEFT_TOP
		case BOTTOM_RIGHT:
			return LEFT_BOTTOM
		case CENTER:
			return CENTER
		default:
			return CENTER
		}
	}
}

type FarmConnection struct {
	A FarmSide
	B FarmSide
}

func (connection FarmConnection) Rotate(rotations int) FarmConnection {
	var result FarmConnection
	result.A = result.A.Rotate(rotations)
	result.B = result.B.Rotate(rotations)
	return result
}

func NewFarmConnection(A FarmSide, B FarmSide) FarmConnection {
	var result FarmConnection
	result.A = A
	result.B = B
	return result
}
