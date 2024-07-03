package city

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

// Represents cities on board
type City struct {
	completed bool
	cities    map[elements.Position][]elements.PlacedFeature
	shields   map[elements.Position]bool
}

func NewCity(pos elements.Position, cityFeature []elements.PlacedFeature, hasShield bool) City {
	return City{
		completed: false,
		cities: map[elements.Position][]elements.PlacedFeature{
			pos: cityFeature,
		},
		shields: map[elements.Position]bool{
			pos: hasShield,
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

func (city *City) GetScoreReport() elements.ScoreReport {
	scoreReport := elements.ScoreReport{
		ReceivedPoints:  map[uint8]uint32{},
		ReturnedMeeples: map[uint8][]uint8{},
	}
	var totalScore uint32 = 0
	// calculate total value of the city and get all meeples
	for pos, features := range city.cities {
		for _, feature := range features {
			if feature.MeepleType != elements.NoneMeeple {
				if _, ok := scoreReport.ReturnedMeeples[uint8(feature.PlayerID)]; ok {
					scoreReport.ReturnedMeeples[uint8(feature.PlayerID)][feature.MeepleType] += 1
				} else {
					scoreReport.ReturnedMeeples[uint8(feature.PlayerID)] = make([]uint8, elements.MeepleTypeCount)
					scoreReport.ReturnedMeeples[uint8(feature.PlayerID)][feature.MeepleType] = 1
				}
			}
		}
		totalScore += 2
		shield := city.shields[pos]
		if shield {
			totalScore += 2
		}
	}

	// determine winning players
	var max uint8 = 0
	winningPlayers := []uint8{}
	for playerID, numMeeples := range scoreReport.ReturnedMeeples {
		for meepleType, ctr := range numMeeples {
			if ctr > 0 && meepleType != int(elements.NoneMeeple) {
				// TODO: add excluding meeples like builder, etc. when they are implemented
				if ctr > max {
					max = ctr
					winningPlayers = nil // remove all values that are in array since there is a player with more meeples
					winningPlayers = append(winningPlayers, playerID)
				} else if ctr == max {
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

func (city City) GetFeaturesFromTile(pos elements.Position) ([]elements.PlacedFeature, bool) {
	cities, ok := city.cities[pos]
	return cities, ok
}

func (city *City) AddTile(pos elements.Position, cityFeatures []elements.PlacedFeature, hasShield bool) {
	city.cities[pos] = cityFeatures
	city.shields[pos] = hasShield
	city.CheckCompleted()
}

// Merges two cities when they are connetced.
// Other city must be deleted after to avoid problems
func (city *City) JoinCities(other City) {
	for pos, otherFeature := range other.cities {
		feature, ok := city.GetFeaturesFromTile(pos)
		if ok {
			feature := append(feature, otherFeature...)
			city.cities[pos] = feature
		} else {
			city.cities[pos] = otherFeature
			city.shields[pos] = other.shields[pos]
		}
	}
	city.CheckCompleted()
}
