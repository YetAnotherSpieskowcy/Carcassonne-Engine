package city

import (
	"slices"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

// Represents a manager responsible for organising cities
type Manager struct {
	cities []City
}

func NewCityManager() Manager {
	return Manager{
		cities: make([]City, 0),
	}
}

func (manager Manager) DeepClone() Manager {
	cities := make([]City, len(manager.cities))
	for i, city := range manager.cities {
		cities[i] = city.DeepClone()
	}
	manager.cities = cities
	return manager
}

// Returns a pointer to a City that has the given feature at the given position, and its index in the city manager
// Returns nil if no such city exists
func (manager Manager) GetCity(position position.Position, feature elements.PlacedFeature) (*City, int) {
	for cityIndex, city := range manager.cities {
		cityFeatures, exists := city.features[position]
		if exists {
			for _, cityFeature := range cityFeatures {
				if cityFeature.Sides == feature.Sides {
					return &manager.cities[cityIndex], cityIndex
				}
			}
		}
	}
	return nil, -1
}

// Finds cities surrounding position of a tile
// Returns a map of indexes of cities in
// manager.cities list with side of a tile as a key.
func (manager Manager) findCities(pos position.Position) map[side.Side]int {
	positions := map[side.Side]position.Position{
		side.Top:    position.New(pos.X(), pos.Y()+1),
		side.Right:  position.New(pos.X()+1, pos.Y()),
		side.Bottom: position.New(pos.X(), pos.Y()-1),
		side.Left:   position.New(pos.X()-1, pos.Y()),
	}

	foundCities := map[side.Side]int{}
	for s, pos := range positions {
		cityFound := false
		for idx, c := range manager.cities {
			features, ok := c.GetFeaturesFromTile(pos)
			if ok {
				for _, f := range features {
					if f.Feature.Sides.HasSide(s.Rotate(2)) {
						foundCities[s] = idx
						cityFound = true
						break
					}
				}

			}
			if cityFound {
				break
			}
		}
	}
	return foundCities
}

// Checks which cities surrounding positions must be joined after this move.
// Returns a map of lists of indexes of cities to join. As a key is used a
// feature that will close the city.
func (manager Manager) findCitiesToJoin(foundCities map[side.Side]int, sides side.Side) []int {
	var cityIndexesToJoin = []int{}
	for _, mask := range side.PrimarySides {
		if sides.HasSide(mask) {
			neighbourCityIndex, ok := foundCities[mask]
			if ok {
				alreadyChecked := false
				for _, cityIndex := range cityIndexesToJoin {
					if cityIndex == neighbourCityIndex {
						alreadyChecked = true
						break
					}
				}
				if !alreadyChecked {
					cityIndexesToJoin = append(cityIndexesToJoin, neighbourCityIndex)
				}
			}
		}
	}
	return cityIndexesToJoin
}

// Checks whether the tile can be placed at given position taking into account
// the meeple placed on the given feature and the meeples that are already placed on
// any city that the feature would join.
func (manager *Manager) CanBePlaced(tile elements.PlacedTile, feat elements.PlacedFeature) bool {
	foundCities := manager.findCities(tile.Position)

	if len(foundCities) == 0 {
		// no existing cities found in tile's neighbourhood - the tile either has
		// no City features or only has a completely new city
		return true
	}

	// this may return an empty list for features that have no cities to join with
	citiesToJoin := manager.findCitiesToJoin(foundCities, feat.Sides)

	// Check each of the existing cities (if any) found in feature's neighbourhood
	// We need to check, if they have *any* meeple placed
	for _, cityIndex := range citiesToJoin {
		// score report function checks the whole city for placed meeples
		// and reports any meeples that would be returned which we can use here
		scoreReport := manager.cities[cityIndex].GetScoreReport()
		if len(scoreReport.ReturnedMeeples) != 0 {
			return false
		}
	}

	return true
}

// Performs required operations to add a new city feature.
func (manager *Manager) UpdateCities(tile elements.PlacedTile) {
	foundCities := manager.findCities(tile.Position)

	if len(foundCities) > 0 {
		for _, cityFeature := range tile.GetFeaturesOfType(feature.City) {
			citiesToRemove := make([]int, 0)
			cityIndexesToJoin := manager.findCitiesToJoin(foundCities, cityFeature.Sides)
			// join cities
			if len(cityIndexesToJoin) == 0 {
				toAppend := []elements.PlacedFeature{cityFeature}
				manager.cities = append(manager.cities, NewCity(tile.Position, toAppend))
			} else {
				if len(cityIndexesToJoin) > 1 {
					for _, cityIndex := range cityIndexesToJoin[1:] {
						manager.cities[cityIndexesToJoin[0]].JoinCities(manager.cities[cityIndex])
						citiesToRemove = append(citiesToRemove, cityIndex)
					}
				}
				toAdd := []elements.PlacedFeature{cityFeature}
				manager.cities[cityIndexesToJoin[0]].AddTile(tile.Position, toAdd)
			}

			// remove cities that were merged into another city
			if len(citiesToRemove) > 0 {
				newCities := make([]City, 0)
				for index, city := range manager.cities {
					if !slices.Contains(citiesToRemove, index) {
						newCities = append(newCities, city)
					}
				}
				manager.cities = newCities
			}
		}
	} else {
		for _, f := range tile.GetFeaturesOfType(feature.City) {
			toAppend := []elements.PlacedFeature{f}
			manager.cities = append(manager.cities, NewCity(tile.Position, toAppend))
		}
	}
}

// Calculates ScoreReport. When forceScore = false calculates score only based on
// closed cities and sets city.scored to true. Otherwise calculates score based on
// every city in array and keeps closed ones.
func (manager *Manager) ScoreCities(forceScore bool) elements.ScoreReport {
	scoreReport := elements.NewScoreReport()
	newCities := make([]City, 0)
	for _, city := range manager.cities {
		if !city.scored {
			if forceScore {
				scoreReport.Join(city.GetScoreReport())
			} else if city.IsCompleted() {
				scoreReport.Join(city.GetScoreReport())
				city.SetScored(true)
			}
		}
		newCities = append(newCities, city)
	}
	manager.cities = newCities

	return scoreReport
}
