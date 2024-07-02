package city

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

// Represents cities on board
type City struct {
	completed bool
	cities    map[elements.Position][]elements.PlacedFeature
}

func New(pos elements.Position, cityFeature []elements.PlacedFeature) City {
	return City{
		completed: false,
		cities: map[elements.Position][]elements.PlacedFeature{
			pos: cityFeature,
		},
	}
}

func (city City) GetCompleted() bool {
	return city.completed
}

func (city *City) CheckCompleted() bool {
	city.completed = true
	for key, placedFeatures := range city.cities {
		for _, placedFeature := range placedFeatures {
			sides := placedFeature.Feature.Sides
			mask := side.Top
			for range 4 {
				if sides&mask == mask {
					switch mask {
					case side.Top:
						_, ok := city.GetFeaturesFromTile(elements.NewPosition(key.X(), key.Y()+1))
						if !ok {
							city.completed = false
							continue
						}
					case side.Bottom:
						_, ok := city.GetFeaturesFromTile(elements.NewPosition(key.X(), key.Y()-1))
						if !ok {
							city.completed = false
							continue
						}
					case side.Right:
						_, ok := city.GetFeaturesFromTile(elements.NewPosition(key.X()+1, key.Y()))
						if !ok {
							city.completed = false
							continue
						}
					case side.Left:
						_, ok := city.GetFeaturesFromTile(elements.NewPosition(key.X()-1, key.Y()))
						if !ok {
							city.completed = false
							continue
						}
					}
				}
				mask = mask.Rotate(1)
			}
		}
	}
	return city.completed
}

func (city City) GetFeaturesFromTile(pos elements.Position) ([]elements.PlacedFeature, bool) {
	cities, ok := city.cities[pos]
	return cities, ok
}

func (city *City) AddTile(pos elements.Position, cityFeatures []elements.PlacedFeature) {
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
