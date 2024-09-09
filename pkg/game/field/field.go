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
	features           map[fieldKey]struct{}         // this is a set, not a dictionary (the value is always struct{} and only the keys matter)
	neighbouringCities map[int]struct{}              // this is a set of cities' IDs
	meeples            []elements.MeepleWithPosition // returned meeples for each player ID (same as in elements.ScoreReport)

	startingTile elements.PlacedTile
}

func New(feature elements.PlacedFeature, startingTile elements.PlacedTile) Field {
	features := map[fieldKey]struct{}{
		{feature: feature, position: startingTile.Position}: {},
	}
	return Field{
		features:           features,
		neighbouringCities: map[int]struct{}{},
		meeples:            []elements.MeepleWithPosition{},

		startingTile: startingTile,
	}
}

func (field Field) FeaturesCount() int {
	return len(field.features)
}

func (field Field) CitiesCount() int {
	return len(field.neighbouringCities)
}

// Expands this field to maximum possible size (like flood fill) and finds all neighbouring cities
// Has to be called on a field which starting tile has already been placed (mostly only at the end of the game)
func (field *Field) Expand(board elements.Board, cityManager city.Manager) {
	_, startingTileExists := board.GetTileAt(field.startingTile.Position)
	if !startingTileExists {
		panic("the field's starting tile does not exist on the board")
	}

	newFeatures := map[fieldKey]struct{}{}

	for len(field.features) != 0 {
		element, _, _ := utilities.GetAnyElementFromMap(field.features)

		// add neighbouring tiles to the set
		_, exists := newFeatures[element]
		if !exists {
			newFeatures[element] = struct{}{}
			for _, neighbour := range field.findNeighbours(element, board) {
				_, exists := newFeatures[neighbour]
				if !exists {
					field.features[neighbour] = struct{}{}
				}
			}
		}

		// add meeple if it exists
		meeple := element.feature.Meeple
		if meeple.Type != elements.NoneMeeple {
			field.meeples = append(field.meeples, elements.NewMeepleWithPosition(
				meeple,
				element.position,
			))
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

// Returns true if the field has less than maxMeepleCount meeples, and false otherwise
func (field Field) IsFieldValid(board elements.Board, maxMeepleCount int) bool {
	newFeatures := map[fieldKey]struct{}{}
	meeples := 0

	for len(field.features) != 0 {
		element, _, _ := utilities.GetAnyElementFromMap(field.features)

		// add neighbouring tiles to the set
		_, exists := newFeatures[element]
		if !exists {
			newFeatures[element] = struct{}{}
			for _, neighbour := range field.findNeighbours(element, board) {
				_, exists := newFeatures[neighbour]
				if !exists {
					field.features[neighbour] = struct{}{}
				}
			}
		}

		// add meeple if it exists
		meeple := element.feature.Meeple
		if meeple.Type != elements.NoneMeeple {
			meeples++
			if meeples > maxMeepleCount {
				return false
			}
		}

		// remove the processed field feature from field.features set
		delete(field.features, element)
	}
	return true
}

// Returns a slice of fieldKey elements containing all features neighbouring a given fieldKey (feature and position)
// The slice may contain duplicates in some cases (todo?)
func (field *Field) findNeighbours(fieldK fieldKey, board elements.Board) []fieldKey {
	neighbours := []fieldKey{}

	for _, side := range side.EdgeSides {
		if fieldK.feature.Sides.OverlapsSide(side) {
			neighbourPosition := fieldK.position.Add(position.FromSide(side))

			tile, tileExists := field.getTile(neighbourPosition, board)
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

// If a tile exists in board at the given position, returns board.GetTileAt(position)
// If the tile doesn't exist in board, but the position is the same as field.startingTile, returns (field.startingTile, true)
// Otherwise returns board.GetTileAt(position) (that is, empty tile and false)
func (field Field) getTile(position position.Position, board elements.Board) (elements.PlacedTile, bool) {
	tile, ok := board.GetTileAt(position)
	if ok {
		return tile, ok
	}
	if position == field.startingTile.Position {
		return field.startingTile, true
	}
	return tile, ok
}

// Returns score report for this field. Has to be called after field.Expand() (todo?)
func (field Field) GetScoreReport() elements.ScoreReport {
	points := uint32(len(field.neighbouringCities) * 3)

	return elements.CalculateScoreReportOnMeeples(int(points), field.meeples)
}
