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

func (manager CityManager) findNeighbouringCities(positions map[side.Side]elements.Position) map[side.Side]City {
	neighbouringCities := map[side.Side]City{}

	for s, pos := range positions {
		cityFound := false
		for _, c := range manager.cities {
			features, ok := c.GetFeaturesFromTile(pos)
			if ok {
				for _, f := range features {
					if f.Feature.Sides&s == s {
						neighbouringCities[s] = c
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

	return neighbouringCities
}

func (manager CityManager) findCitiesToJoin(foundCities map[side.Side]City, tile elements.PlacedTile) map[elements.PlacedFeature][]City {
	citiesToJoin := map[elements.PlacedFeature][]City{}
	cityFeatures := tile.GetCityFeatures()
	for _, cityFeature := range cityFeatures {
		var cToJoin = []City{}
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
	foundCities := manager.findNeighbouringCities(positions)
	if len(foundCities) > 0 {
		citiesToJoin := manager.findCitiesToJoin(foundCities, tile)
		for f, cToJoin := range citiesToJoin {
			if len(cToJoin) == 0 {
				// I guess panic
				break
			}
			firstCity := cToJoin[0]
			cToJoin = cToJoin[1:]
			if len(cToJoin) > 0 {
				for _, c := range cToJoin {
					firstCity.JoinCities(c)
				}
			}
			toAdd := []elements.PlacedFeature{f}
			firstCity.AddTile(tile.Position, toAdd, tile.TileWithMeeple.HasShield)
		}
	} else {
		for _, f := range tile.GetCityFeatures() {
			toAppend := []elements.PlacedFeature{f}
			manager.cities = append(manager.cities, NewCity(tile.Position, toAppend, tile.HasShield))
		}
	}

}
