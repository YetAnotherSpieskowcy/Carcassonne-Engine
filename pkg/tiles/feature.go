package tiles

import connection "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/connection"

type FeatureType int64

const (
	NONE FeatureType = iota
	ROADS
	CITIES
	FIELDS
)

type Feature struct {
	FeatureType FeatureType
	Connections []connection.Connection
}
