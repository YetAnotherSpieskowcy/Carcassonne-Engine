package city

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine-API/pkg/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine-API/pkg/tiles/side"
)

// Represents cities on board
type City struct {
	completed bool
	scored    bool
	features  map[elements.Position][]elements.PlacedFeature
	shields   map[elements.Position]bool
}

func NewCity(pos elements.Position, cityFeature []elements.PlacedFeature, hasShield bool) City {
	return City{
		completed: false,
		scored:    false,
		features: map[elements.Position][]elements.PlacedFeature{
			pos: cityFeature,
		},
		shields: map[elements.Position]bool{
			pos: hasShield,
		},
	}
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
				if sides&mask == mask {
					_, ok := city.GetFeaturesFromTile(pos.Add(elements.PositionFromSide(mask)))
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
	scoreReport := elements.NewScoreReport()
	var totalScore uint32
	// calculate total value of the city and get all meeples
	for pos, features := range city.features {
		for _, feature := range features {
			if feature.MeepleType != elements.NoneMeeple {
				if _, ok := scoreReport.ReturnedMeeples[feature.PlayerID]; ok {
					scoreReport.ReturnedMeeples[feature.PlayerID][feature.MeepleType]++
				} else {
					scoreReport.ReturnedMeeples[feature.PlayerID] = make([]uint8, elements.MeepleTypeCount)
					scoreReport.ReturnedMeeples[feature.PlayerID][feature.MeepleType] = 1
				}
			}
		}
		totalScore += 2
		shield := city.shields[pos]
		if shield {
			totalScore += 2
		}
	}
	if !city.completed {
		totalScore /= 2
	}

	// determine winning players
	var max uint8
	winningPlayers := []elements.ID{}
	for playerID, numMeeples := range scoreReport.ReturnedMeeples {
		for meepleType, meepleCount := range numMeeples {
			if meepleCount > 0 && meepleType != int(elements.NoneMeeple) {
				// TODO: add excluding meeples like builder, etc. when they are implemented
				if meepleCount > max {
					max = meepleCount
					winningPlayers = nil // remove all values that are in array since there is a player with more meeples
					winningPlayers = append(winningPlayers, playerID)
				} else if meepleCount == max {
					winningPlayers = append(winningPlayers, playerID)
				}
			}
		}
	}

	// award points
	for _, player := range winningPlayers {
		scoreReport.ReceivedPoints[player] = totalScore
	}

	return scoreReport
}

// Returns all features from a tile at a given position that are part of a city
// and whether such a tile is in the city.
func (city City) GetFeaturesFromTile(pos elements.Position) ([]elements.PlacedFeature, bool) {
	cities, ok := city.features[pos]
	return cities, ok
}

func (city *City) AddTile(pos elements.Position, cityFeatures []elements.PlacedFeature, hasShield bool) {
	city.features[pos] = cityFeatures
	city.shields[pos] = hasShield
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
			city.shields[pos] = other.shields[pos]
		}
	}
	city.checkCompleted()
}
