package city

import (
	"maps"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature/modifier"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

// Represents cities on board
type City struct {
	completed bool
	scored    bool
	features  map[position.Position][]elements.PlacedFeature
	shields   uint8
}

func NewCity(pos position.Position, cityFeatures []elements.PlacedFeature) City {
	var shields = uint8(0)
	for _, feat := range cityFeatures {
		if feat.ModifierType == modifier.Shield {
			shields++
		}
	}
	return City{
		completed: false,
		scored:    false,
		features: map[position.Position][]elements.PlacedFeature{
			pos: cityFeatures,
		},
		shields: shields,
	}
}

func (city City) DeepClone() City {
	city.features = maps.Clone(city.features)
	// shields number is already copied for being uint8
	return city
}

func (city City) IsCompleted() bool {
	return city.completed
}

// Checks if city is closed and sets city.completed.
func (city *City) checkCompleted() bool {
	city.completed = true
	for pos, placedFeatures := range city.features {
		for _, placedFeature := range placedFeatures {
			sides := placedFeature.Feature.Sides
			mask := side.Top
			for range 4 {
				if sides.HasSide(mask) {
					_, ok := city.GetFeaturesFromTile(pos.Add(position.FromSide(mask)))
					if !ok {
						city.completed = false
						break
					}
				}
				mask = mask.Rotate(1)
			}
		}
	}
	return city.completed
}

func (city *City) SetScored(scored bool) {
	city.scored = scored
}

// Calculates score value of the city and
// determines players that should receive points.
func (city *City) GetScoreReport() elements.ScoreReport {
	var returnedMeeples = []elements.MeepleWithPosition{}
	var totalScore uint32
	// calculate total value of the city and get all meeples
	for pos, features := range city.features {
		for _, feature := range features {
			if feature.Meeple.Type != elements.NoneMeeple {
				returnedMeeples = append(returnedMeeples, elements.NewMeepleWithPosition(
					feature.Meeple,
					pos,
				))
			}
		}
		totalScore += 2
	}
	totalScore += uint32(city.shields) * 2

	if !city.completed {
		totalScore /= 2
	}

	return elements.CalculateScoreReportOnMeeples(int(totalScore), returnedMeeples)
}

// Returns all features from a tile at a given position that are part of a city
// and whether such a tile is in the city.
func (city City) GetFeaturesFromTile(pos position.Position) ([]elements.PlacedFeature, bool) {
	cities, ok := city.features[pos]
	return cities, ok
}

func (city *City) AddTile(pos position.Position, cityFeatures []elements.PlacedFeature) {
	hasShield := false
	for _, feat := range cityFeatures {
		if feat.ModifierType == modifier.Shield {
			hasShield = true
		}
	}
	_, tileInCity := city.GetFeaturesFromTile(pos)
	if !tileInCity {
		city.features[pos] = cityFeatures
		if hasShield {
			city.shields++
		}

	} else {
		city.features[pos] = append(city.features[pos], cityFeatures...)
		// If the tile being checked is already in the city, it must have the shield already counted,
		// because tiles with shields, has only one city feature, so it's impossible
		// that there are multiple city features on that tile.
	}
	city.checkCompleted()
}

// Merges two cities when they are connected.
// Other city must be deleted after to avoid problems
func (city *City) JoinCities(other City) {
	for pos, otherFeature := range other.features {
		feature, ok := city.GetFeaturesFromTile(pos)
		if ok {
			features := feature
			features = append(features, otherFeature...)
			city.features[pos] = features
		} else {
			city.features[pos] = otherFeature
		}
	}
	city.shields += other.shields
	city.checkCompleted()
}
