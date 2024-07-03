package city

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

// Represents cities on board
type CityManager struct {
	cities []City
}

func NewCityManager() CityManager {
	return CityManager{
		cities: make([]City, 0),
	}
}

func getNeighbouringPositions(pos elements.Position) map[side.Side]elements.Position {
	return map[side.Side]elements.Position{
		side.Top:    elements.NewPosition(pos.X(), pos.Y()+1),
		side.Right:  elements.NewPosition(pos.X()+1, pos.Y()),
		side.Bottom: elements.NewPosition(pos.X(), pos.Y()-1),
		side.Left:   elements.NewPosition(pos.X()-1, pos.Y()),
	}
}

func (manager CityManager) findCities(positions map[side.Side]elements.Position) map[side.Side]int {
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

func (manager CityManager) findCitiesToJoin(foundCities map[side.Side]int, tile elements.PlacedTile) map[elements.PlacedFeature][]int {
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

func (manager *CityManager) UpdateCities(tile elements.PlacedTile) {
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

func (manager *CityManager) ScoreCities(forceScore bool) elements.ScoreReport {
	scoreReport := elements.NewScoreReport()
	for idx, city := range manager.cities {
		if forceScore {
			scoreReport.JoinReport(city.GetScoreReport())
		} else if city.GetCompleted() {
			println("scoring")
			scoreReport.JoinReport(city.GetScoreReport())
			manager.cities = append(manager.cities[:idx], manager.cities[idx+1:]...)
		}
	}
	return scoreReport
}
