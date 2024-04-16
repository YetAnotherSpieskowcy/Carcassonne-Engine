package city

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	sideMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
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

func (city *City) CheckCompleted() bool {
	city.completed = true
	for key, features := range city.cities {
		for _, feature := range features {
			for _, side := range feature.Sides {
				switch side {
				case sideMod.Top:
					_, ok := city.GetFeaturesFromTile(elements.NewPosition(key.X(), key.Y()+1))
					if !ok {
						city.completed = false
						continue
					}
				case sideMod.Bottom:
					_, ok := city.GetFeaturesFromTile(elements.NewPosition(key.X(), key.Y()-1))
					if !ok {
						city.completed = false
						continue
					}
				case sideMod.Right:
					_, ok := city.GetFeaturesFromTile(elements.NewPosition(key.X()+1, key.Y()))
					if !ok {
						city.completed = false
						continue
					}
				case sideMod.Left:
					_, ok := city.GetFeaturesFromTile(elements.NewPosition(key.X()-1, key.Y()))
					if !ok {
						city.completed = false
						continue
					}
				}
			}
		}
	}
	return city.completed
}

func (city City) GetFeaturesFromTile(pos elements.Position) ([]feature.Feature, bool) {
	cities, ok := city.cities[pos]
	return cities, ok
}

func (city *City) AddTile(pos elements.Position, cityFeatures []feature.Feature) {
	city.cities[pos] = cityFeatures
	city.CheckCompleted()
}

// Merges two cities when they are connetced.
// Other city must be deleted after to avoid problems
func (city *City) JoinCities(other City) {
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
