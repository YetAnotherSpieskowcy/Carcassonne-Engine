package city

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
)

// Represents cities on board
type City struct {
	completed bool
	cities    map[elements.Position][]feature.Feature
}

func New(pos elements.Position, cityFeature []feature.Feature) City {
	return City{
		completed: false,
		cities: map[elements.Position][]feature.Feature{
			pos: cityFeature,
		},
	}
}

func (city City) GetCompleted() bool {
	return city.completed
}

func (city City) CheckCompleted() bool {
	// TODO: check if city is closed
	return city.completed
}

func (city City) GetFeaturesFromTile(pos elements.Position) ([]feature.Feature, bool) {
	cities, ok := city.cities[pos]
	return cities, ok
}

func (city City) AddTile(pos elements.Position, cityFeatures []feature.Feature) {
	city.cities[pos] = cityFeatures
	city.CheckCompleted()
}

// Merges two cities when they are connetced.
// Other city must be deleted after to avoid problems
func (city City) JoinCities(other City) {
	for key, otherFeature := range other.cities {
		feature, ok := city.GetFeaturesFromTile(key)
		if ok {
			feature := append(feature, otherFeature...)
			city.cities[key] = feature
		} else {
			city.cities[key] = otherFeature
		}
	}
	city.CheckCompleted()
}
