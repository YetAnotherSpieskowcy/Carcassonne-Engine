package connection

type Side int64

const (
	NONE Side = iota
	TOP
	RIGHT
	LEFT
	BOTTOM

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
		default:
			result = NONE
		}
		rotations--
	}
	return result
}

type Connection struct {
	A Side
	B Side
}

func (connection Connection) Rotate(rotations uint) Connection {
	var result Connection
	result.A = connection.A.Rotate(rotations)
	result.B = connection.B.Rotate(rotations)
	return result
}

func NewConnection(A Side, B Side) Connection {
	var result Connection
	result.A = A
	result.B = B
	return result
}

func (connection Connection) ToString() string {
	return connection.A.ToString() + " " + connection.B.ToString()
}
