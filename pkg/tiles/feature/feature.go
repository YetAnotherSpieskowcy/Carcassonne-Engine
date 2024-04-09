package feature

import connection "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/connection"

type Type int64

const (
	None Type = iota
	Road
	City
	Field
)

type Feature struct {
	FeatureType Type
	Connections []connection.Side
}
