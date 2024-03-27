package main

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

/*
Rotates side clockwise
*/
func (side Side) Rotate(rotations int) Side {
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
		case TOP:
			return RIGHT
		case RIGHT:
			return BOTTOM
		case LEFT:
			return TOP
		case BOTTOM:
			return LEFT

		case TOPLEFT:
			return TOPRIGHT
		case TOPRIGHT:
			return BOTTOMRIGHT
		case BOTTOMLEFT:
			return TOPLEFT
		case BOTTOMRIGHT:
			return BOTTOMLEFT
		default:
			return NONE
		}
	}
}

type Connection struct {
	a Side
	b Side
}

func (connection Connection) Rotate(rotations int) Connection {
	var result Connection
	result.a = result.a.Rotate(rotations)
	result.b = result.b.Rotate(rotations)
	return result
}
