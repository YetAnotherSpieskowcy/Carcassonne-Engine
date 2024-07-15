package field

import (
	"fmt"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/city"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	featureMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/utilities"
)

type fieldKey struct {
	feature  elements.PlacedFeature
	position elements.Position
}

// Represents a field on the board
type Field struct {
	features           map[fieldKey]struct{} // this is a set, not a dictionary (the value is always struct{} and only the keys matter)
	neighbouringCities map[int]struct{}      // this is a set of cities' IDs
}

func NewField(feature elements.PlacedFeature, position elements.Position) Field {
	features := map[fieldKey]struct{}{
		{feature: feature, position: position}: {},
	}
	return Field{features: features}
}

func (field *Field) Features() map[fieldKey]struct{} { // todo maybe delete this and change field.features to public?
	return field.features
}

// Expands this field to maximum possible size - like flood fill
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

		// find neighbouring city features
		var neighbouringCityFeatures []elements.PlacedFeature

		if element.feature.Sides == side.None {
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
				panic(fmt.Sprintf("city manager didn't found city: %#v at position %#v", cityFeature, element.position))
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
	sides := []side.Side{
		side.TopLeftEdge,
		side.TopRightEdge,
		side.RightTopEdge,
		side.RightBottomEdge,
		side.BottomRightEdge,
		side.BottomLeftEdge,
		side.LeftBottomEdge,
		side.LeftTopEdge,
	}

	neighbours := []fieldKey{}

	for _, side := range sides {
		if field.feature.Sides&side != 0 {
			neighbourPosition := elements.AddPositions(field.position, elements.PositionFromSide(side))

			tile, tileExists := board.GetTileAt(neighbourPosition)
			if tileExists {
				mirroredSide := side.Mirror()

				feature := tile.GetPlacedFeatureAtSide(mirroredSide, featureMod.Field)
				if feature == nil {
					panic("No matching field found on adjacent tile!") // when a field is directly touching another feature (e.g. a city). Should never happen
				}
				neighbours = append(neighbours, fieldKey{feature: *feature, position: neighbourPosition})
			}
		}
	}
	return neighbours
}

/*
Assumptions:
 - if there is only one field feature on a tile, it neighbours ALL cities on this tile
 - if the field feature doesn't have any sides (i.e. its sides==side.None), it neighbours ALL cities on this tile
 - in other cases, fields neighbour all cities they share a common tile corner with

Observation:
 - fieldFeature.sides.FlipCorners() produces different sides ONLY when the field neighbours a city,
   because only in that case the field doesn't contain both edge sides around the corner
*/
