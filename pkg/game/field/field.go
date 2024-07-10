package field

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	featureMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

type fieldKey struct {
	feature  elements.PlacedFeature
	position elements.Position
}

// Represents a field on the board
type Field struct {
	features map[fieldKey]struct{} // this is a set, not a dictionary (the value is always struct{} and only the keys matter)
	// todo neighbouring cities
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
// todo should find cities as well, probably
func (field *Field) Expand(board elements.Board) {
	newFeatures := map[fieldKey]struct{}{}

	for len(field.features) != 0 {
		element, _, _ := GetAny(field.features)

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

/* todo neighbouring cities:
Assumptions:
 - if there is only one field feature on a tile, it neighbours ALL cities on this tile
 - if the field feature doesn't have any sides (i.e. its sides==side.None), it neighbours ALL cities on this tile
 - in other cases, fields neighbour all cities they share a common tile corner with
*/

// todo would be nice to move this to some "utilities" package maybe? todo
// Retrieves key and value from a map, and a bool indicating if the map is empty.
// In case of multi-element maps, it is not specified which element is returned.
func GetAny[K comparable, V any](m map[K]V) (K, V, bool) {
	for key, value := range m {
		return key, value, true
	}
	return *new(K), *new(V), false
}
