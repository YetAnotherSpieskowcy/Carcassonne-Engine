package feature

import connection "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/connection"

type FeatureType int64 //nolint:revive // name type would be too ambigous

const (
	None FeatureType = iota
	Road
	City
	Field
)

type Feature struct {
	FeatureType FeatureType
	Connections []connection.Side
}
