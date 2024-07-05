package city

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
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

// Prepares map of positions surrounding position chosen for tile
func getNeighbouringPositions(pos elements.Position) map[side.Side]elements.Position {
	return map[side.Side]elements.Position{
		side.Top:    elements.NewPosition(pos.X(), pos.Y()+1),
		side.Right:  elements.NewPosition(pos.X()+1, pos.Y()),
		side.Bottom: elements.NewPosition(pos.X(), pos.Y()-1),
		side.Left:   elements.NewPosition(pos.X()-1, pos.Y()),
	}
}

// Finds cities surrounding position of a tile
// Returns a map of indexes of cities in
// manager.cities list with side of a tile as a key.
func (manager Manager) findCities(positions map[side.Side]elements.Position) map[side.Side]int {
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
// Returns a map of lists of indexes of cities to join. As a key is used a side of a feature.
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
		citiesToJoin[cityFeature] = cToJoin
	}
	return citiesToJoin
}

// Performs required operations to add a new city feature.
func (manager *Manager) UpdateCities(tile elements.PlacedTile) {
	positions := getNeighbouringPositions(tile.Position)
	foundCities := manager.findCities(positions)

	if len(foundCities) > 0 {
		citiesToJoin := manager.findCitiesToJoin(foundCities, tile)
		for f, cToJoin := range citiesToJoin {
			if len(cToJoin) == 0 {
				// I guess panic
				break
			}
			cToIter := cToJoin[1:]
			if len(cToJoin) > 0 {
				for _, c := range cToIter {
					manager.cities[cToJoin[0]].JoinCities(manager.cities[c])
				}
			}
			toAdd := []elements.PlacedFeature{f}
			manager.cities[cToJoin[0]].AddTile(tile.Position, toAdd, tile.TileWithMeeple.HasShield)
		}
	} else {
		for _, f := range tile.GetCityFeatures() {
			toAppend := []elements.PlacedFeature{f}
			manager.cities = append(manager.cities, NewCity(tile.Position, toAppend, tile.HasShield))
		}
	}
}

// Calculates ScoreReport. When forceScore = false calculates score only based on
// closed cities and removes them from array. Otherwise calculates score based on
// every city in array and keeps closed ones.
func (manager *Manager) ScoreCities(forceScore bool) elements.ScoreReport {
	scoreReport := elements.NewScoreReport()
	for idx, city := range manager.cities {
		if forceScore {
			scoreReport.JoinReport(city.GetScoreReport())
		} else if city.IsCompleted() {
			scoreReport.JoinReport(city.GetScoreReport())
			manager.cities = append(manager.cities[:idx], manager.cities[idx+1:]...)
		}
	}
	return scoreReport
}
