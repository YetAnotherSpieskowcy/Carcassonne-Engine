package city

import (
	"slices"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine-API/pkg/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine-API/pkg/tiles/feature/modifier"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine-API/pkg/tiles/side"
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

// Finds cities surrounding position of a tile
// Returns a map of indexes of cities in
// manager.cities list with side of a tile as a key.
func (manager Manager) findCities(pos elements.Position) map[side.Side]int {
	positions := map[side.Side]elements.Position{
		side.Top:    elements.NewPosition(pos.X(), pos.Y()+1),
		side.Right:  elements.NewPosition(pos.X()+1, pos.Y()),
		side.Bottom: elements.NewPosition(pos.X(), pos.Y()-1),
		side.Left:   elements.NewPosition(pos.X()-1, pos.Y()),
	}

	foundCities := map[side.Side]int{}
	for s, pos := range positions {
		cityFound := false
		for idx, c := range manager.cities {
			features, ok := c.GetFeaturesFromTile(pos)
			if ok {
				for _, f := range features {
					if f.Feature.Sides&s.Rotate(2) == s.Rotate(2) {
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
func (manager Manager) findCitiesToJoin(foundCities map[side.Side]int, tile elements.PlacedTile) map[elements.PlacedFeature][]int {
	citiesToJoin := map[elements.PlacedFeature][]int{}
	cityFeatures := tile.GetCityFeatures()
	for _, cityFeature := range cityFeatures {
		var cToJoin = []int{}
		sides := cityFeature.Feature.Sides
		mask := side.Top
		for range 4 {
			if sides&mask == mask {
				c, ok := foundCities[mask]
				if ok {
					cToJoin = append(cToJoin, c)
				}
			}
			mask = mask.Rotate(1)
		}
		if len(cToJoin) > 0 {
			citiesToJoin[cityFeature] = cToJoin
		}
	}
	return citiesToJoin
}

// Performs required operations to add a new city feature.
func (manager *Manager) UpdateCities(tile elements.PlacedTile) {
	foundCities := manager.findCities(tile.Position)

	if len(foundCities) > 0 {
		citiesToRemove := make([]int, 0)
		citiesToJoin := manager.findCitiesToJoin(foundCities, tile)
		// join cities
		for f, cToJoin := range citiesToJoin {
			if len(cToJoin) == 0 {
				panic("No cities to join.")
			}
			if len(cToJoin) > 1 {
				for _, cityIndex := range cToJoin[1:] {
					manager.cities[cToJoin[0]].JoinCities(manager.cities[cityIndex])
					citiesToRemove = append(citiesToRemove, cityIndex)
				}
			}
			toAdd := []elements.PlacedFeature{f}
			manager.cities[cToJoin[0]].AddTile(tile.Position, toAdd, f.ModifierType == modifier.Shield)
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
	} else {
		for _, f := range tile.GetCityFeatures() {
			toAppend := []elements.PlacedFeature{f}
			manager.cities = append(manager.cities, NewCity(tile.Position, toAppend, tile.HasShield))
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
