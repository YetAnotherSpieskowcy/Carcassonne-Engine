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

type Connection struct {
	a Side
	b Side
}
