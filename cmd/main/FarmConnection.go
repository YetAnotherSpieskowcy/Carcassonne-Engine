package main

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

type FarmConnection struct {
	a FarmSide
	b FarmSide
}
