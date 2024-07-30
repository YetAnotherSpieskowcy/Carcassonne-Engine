package field

import (
	"fmt"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/city"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	featureMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/utilities"
)

/*
Assumptions:
 - if there is only one field feature on a tile, it neighbours ALL cities on this tile
 - if the field feature doesn't have any sides (i.e. its sides==side.NoSide), it neighbours ALL cities on this tile
   (this rule is only applicable to tiles from the expansions, as there are no such tiles in the base game)
 - in other cases, fields neighbour all cities they share a common tile corner with

Observation:
 - fieldFeature.sides.FlipCorners() produces different sides ONLY when the field neighbours a city,
   because only in that case the field doesn't contain both edge sides around the corner
*/

type fieldKey struct {
	feature  elements.PlacedFeature
	position position.Position
}

// Represents a field on the board
type Field struct {
	features           map[fieldKey]struct{}   // this is a set, not a dictionary (the value is always struct{} and only the keys matter)
	neighbouringCities map[int]struct{}        // this is a set of cities' IDs
	meeples            map[elements.ID][]uint8 // returned meeples for each player ID (same as in elements.ScoreReport)
}

func New(feature elements.PlacedFeature, position position.Position) Field {
	features := map[fieldKey]struct{}{
		{feature: feature, position: position}: {},
	}
	return Field{features: features, neighbouringCities: map[int]struct{}{}, meeples: map[elements.ID][]uint8{}}
}

func (field Field) FeaturesCount() int {
	return len(field.features)
}

func (field Field) CitiesCount() int {
	return len(field.neighbouringCities)
}

// Expands this field to maximum possible size (like flood fill) and finds all neighbouring cities
func (field *Field) Expand(board elements.Board, cityManager city.Manager) {
	newFeatures := map[fieldKey]struct{}{}

	for len(field.features) != 0 {
		element, _, _ := utilities.GetAnyElementFromMap(field.features)

		// add neighbouring tiles to the set
		_, exists := newFeatures[element]
		if !exists {
			newFeatures[element] = struct{}{}
			for _, neighbour := range findNeighbours(element, board) {
				_, exists := newFeatures[neighbour]
				if !exists {
					field.features[neighbour] = struct{}{}
				}
			}

		}

		// add meeple if it exists
		meeple := element.feature.Meeple
		if meeple.Type != elements.NoneMeeple {
			_, exists := field.meeples[meeple.PlayerID]
			if !exists {
				field.meeples[meeple.PlayerID] = make([]uint8, elements.MeepleTypeCount)
			}
			field.meeples[meeple.PlayerID][meeple.Type]++
		}

		// find neighbouring city features
		var neighbouringCityFeatures []elements.PlacedFeature

		if element.feature.Sides == side.NoSide {
			// field neighbours all cities on this tile
			tile, _ := board.GetTileAt(element.position)
			neighbouringCityFeatures = tile.GetFeaturesOfType(featureMod.City)

		} else {
			cornerFlippedSide := element.feature.Sides.FlipCorners()

			if cornerFlippedSide != element.feature.Sides {
				tile, _ := board.GetTileAt(element.position)
				if len(tile.GetFeaturesOfType(featureMod.Field)) == 1 {
					// field neighbours all cities on this tile
					neighbouringCityFeatures = tile.GetFeaturesOfType(featureMod.City)

				} else {
					// field neighbours only the cities it shares a common corner with
					neighbouringCityFeatures = tile.GetPlacedFeaturesOverlappingSide(cornerFlippedSide, featureMod.City)
				}
			} // else { the field feature doesn't neighbour any cities }
		}

		for _, cityFeature := range neighbouringCityFeatures {
			city, cityID := cityManager.GetCity(element.position, cityFeature)
			if city == nil {
				panic(fmt.Sprintf("city manager did not find city: %#v at position %#v", cityFeature, element.position))
			}
			if city.IsCompleted() {
				field.neighbouringCities[cityID] = struct{}{}
			}
		}

		// remove the processed field feature from field.features set
		delete(field.features, element)
	}
	field.features = newFeatures
}

// Returns a slice of fieldKey elements containing all features neighbouring a given fieldKey (feature and position)
// The slice may contain duplicates in some cases (todo?)
func findNeighbours(field fieldKey, board elements.Board) []fieldKey {
	neighbours := []fieldKey{}

	for _, side := range side.EdgeSides {
		if field.feature.Sides.OverlapsSide(side) {
			neighbourPosition := field.position.Add(position.FromSide(side))

			tile, tileExists := board.GetTileAt(neighbourPosition)
			if tileExists {
				mirroredSide := side.Mirror()

				feature := tile.GetPlacedFeatureAtSide(mirroredSide, featureMod.Field)
				if feature == nil {
					panic("No matching field found on adjacent tile! The field is directly touching another feature (e.g. city or road). This should never happen")
				}
				neighbours = append(neighbours, fieldKey{feature: *feature, position: neighbourPosition})
			}
		}
	}
	return neighbours
}

// Returns score report for this field. Has to be called after field.Expand() (todo?)
func (field Field) GetScoreReport() elements.ScoreReport {
	points := uint32(len(field.neighbouringCities) * 3)

	scoreReport := elements.NewScoreReport()

	scoreReport.ReturnedMeeples = field.meeples

	scoredPlayers := elements.GetPlayersWithMostMeeples(field.meeples)
	for _, player := range scoredPlayers {
		scoreReport.ReceivedPoints[player] = points
	}

	return scoreReport
}
