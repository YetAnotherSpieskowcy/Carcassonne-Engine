package field

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	featureMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

type fieldKey struct {
	feature  featureMod.Feature
	position elements.Position
}

// todo would be nice to move this to some "utilities" package maybe? todo
func GetAny[K comparable, V any](m map[K]V) (K, V, bool) {
	for key, value := range m {
		return key, value, true
	}
	return *new(K), *new(V), false
}

// Represents a field on the board
type Field struct {
	features map[fieldKey]struct{}
}

func New(feature featureMod.Feature, position elements.Position) Field {
	features := map[fieldKey]struct{}{
		{feature: feature, position: position}: {},
	}
	return Field{features: features}
}

// Expands this field to maximum possible size - like flood fill
func (field *Field) Expand(board elements.Board) { // todo not sure if board should be an argument..
	newFeatures := map[fieldKey]struct{}{}
	for len(field.features) != 0 {
		element, _, _ := GetAny(field.features)

		if _, ok := newFeatures[element]; !ok {
			newFeatures[element] = struct{}{}

			for _, neighbour := range findNeighbours(element, board) {
				if _, ok := newFeatures[neighbour]; !ok {
					field.features[neighbour] = struct{}{}
					newFeatures[neighbour] = struct{}{}
				}
			}
		}
		delete(field.features, element)
	}
}

func findNeighbours(field fieldKey, board elements.Board) []fieldKey { // todo not sure if board should be an argument..
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
			tile, ok := board.GetTileAt(neighbourPosition)
			if ok {
				mirroredSide := side.Mirror()
				for _, feature := range tile.Features {
					if feature.FeatureType == featureMod.Field && feature.Sides&mirroredSide != 0 {
						neighbours = append(neighbours, fieldKey{feature: feature.Feature, position: neighbourPosition})
						break
					}
				}
				panic("No matching field found on adjacent tile!")
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
