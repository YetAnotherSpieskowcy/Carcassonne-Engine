package field

import (
	"fmt"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/city"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/binarytiles"
	featureMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/utilities"
)

/*
Assumptions:
 - if there is only one field feature on a tile, it neighbours ALL cities on this tile
 - if the field feature doesn't have any sides (i.e. its sides==side.NoSide), it neighbours ALL cities on this tile
   (this rule is only applicable to tiles from the expansions, as there are no such tiles in the base game)
 - in other cases, fields neighbour all cities they share a common tile corner with
*/

// Represents a field on the board
type Field struct {
	features           map[binarytiles.BinaryTileFeature]struct{} // this is a set, not a dictionary (the value is always struct{} and only the keys matter)
	neighbouringCities map[int]struct{}                           // this is a set of cities' IDs
	meeples            []elements.MeepleWithPosition              // returned meeples for each player ID (same as in elements.ScoreReport)

	startingTile binarytiles.BinaryTile
}

func New(side binarytiles.BinaryTileSide, startingTile binarytiles.BinaryTile) Field {
	features := map[binarytiles.BinaryTileFeature]struct{}{
		{Side: startingTile.GetConnectedSides(side, featureMod.Field), Tile: startingTile}: {},
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
	_, startingTileExists := board.GetTileAt(field.startingTile.Position())
	if !startingTileExists {
		panic("the field's starting tile does not exist on the board")
	}

	newFeatures := map[binarytiles.BinaryTileFeature]struct{}{}

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
		meepleID := element.Tile.GetMeepleIDAtSide(element.Side, featureMod.Field)
		if meepleID != elements.NonePlayer {
			field.meeples = append(field.meeples, elements.NewMeepleWithPosition(
				elements.Meeple{Type: elements.NormalMeeple, PlayerID: meepleID},
				element.Tile.Position(),
			))
		}

		// find neighbouring city features
		neighbouringCityFeatures := findNeighbouringCityFeatures(element)

		for _, cityFeature := range neighbouringCityFeatures {
			city, cityID := cityManager.GetCity(element.Tile.Position(), cityFeature)
			if city == nil {
				panic(fmt.Sprintf("city manager did not find city: %#v at position %#v", cityFeature, element.Tile.Position()))
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

// Returns true if the field has lesser or equal number of meeples than maxMeepleCount, and false otherwise
//
// Useful for testing whether or not a meeple can be placed on a field feature
// (i.e. are there any other meeples on the expanded field)
// In such cases, maxMeepleCount should be set to 1 if the tested tile already has a meeple and 0 if it does not.
func (field Field) IsFieldValid(board elements.Board, maxMeepleCount int) bool {
	newFeatures := map[binarytiles.BinaryTileFeature]struct{}{}
	meeples := 0

	// copy the original field.features into features, to avoid modifying it
	features := map[binarytiles.BinaryTileFeature]struct{}{}
	for key := range field.features {
		features[key] = struct{}{}
	}

	for len(features) != 0 {
		element, _, _ := utilities.GetAnyElementFromMap(features)

		// add neighbouring tiles to the set
		_, exists := newFeatures[element]
		if !exists {
			newFeatures[element] = struct{}{}
			for _, neighbour := range field.findNeighbours(element, board) {
				_, exists := newFeatures[neighbour]
				if !exists {
					features[neighbour] = struct{}{}
				}
			}
		}

		// add meeple if it exists
		meepleID := element.Tile.GetMeepleIDAtSide(element.Side, featureMod.Field)
		if meepleID != elements.NonePlayer {
			meeples++
			if meeples > maxMeepleCount {
				return false
			}
		}

		// remove the processed field feature from features set
		delete(features, element)
	}
	return true
}

// Returns a slice of fieldKey elements containing all features neighbouring a given fieldKey (feature and position)
// The slice may contain duplicates in some cases (todo?)
func (field *Field) findNeighbours(fieldK binarytiles.BinaryTileFeature, board elements.Board) []binarytiles.BinaryTileFeature {
	neighbours := []binarytiles.BinaryTileFeature{}

	neighbouringSides := fieldK.Side.CornersToSides()
	neighbouringSides &= ^fieldK.Tile.GetFeatureSides(featureMod.City) // clear sides that also have cities

	for _, side := range binarytiles.OrthogonalSides {
		if neighbouringSides.OverlapsSide(side) {
			neighbourTile, ok := field.getTile(side.PositionFromSide().Add(fieldK.Tile.Position()), board)
			if ok {
				neighbours = append(neighbours, binarytiles.BinaryTileFeature{
					Tile: neighbourTile,
					Side: neighbourTile.GetConnectedSides(binarytiles.CornerFromSide(fieldK.Side&side.SidesToCorners(), side), featureMod.Field),
				})
			}
		}
	}

	return neighbours
}

// If a tile exists in board at the given position, returns board.GetTileAt(position)
// If the tile doesn't exist in board, but the position is the same as field.startingTile, returns (field.startingTile, true)
// Otherwise returns board.GetTileAt(position) (that is, empty tile and false)
func (field Field) getTile(position position.Position, board elements.Board) (binarytiles.BinaryTile, bool) {
	tile, ok := board.GetTileAt(position)
	binarytile := binarytiles.FromPlacedTile(tile) // todo binarytiles rewrite

	if ok {
		return binarytile, ok
	}
	if position == field.startingTile.Position() {
		return field.startingTile, true
	}
	return binarytile, ok
}

// Returns a slice of neighbouring city features* on the same tile
// *feature = all sides that a city feature occupies
func findNeighbouringCityFeatures(fieldKey binarytiles.BinaryTileFeature) []binarytiles.BinaryTileSide {
	side := fieldKey.Side
	tile := fieldKey.Tile

	cityFeatures := tile.GetFeaturesOfType(featureMod.City)

	if side == binarytiles.SideNone {
		// unconnected field - neighbours all cities on this tile
		return cityFeatures

	} else {
		fieldFeatures := tile.GetFeaturesOfType(featureMod.Field)

		if len(fieldFeatures) == 1 { // only one field - neighbours all cities on this tile
			return cityFeatures

		} else { // field neighbours only the cities it shares a common corner with
			sidesNeighbouringFieldCorners := side.CornersToSides()

			// iterate over cityFeatures and leave only the elements that overlap sidesNeighbouringFieldCorners (delete the rest)
			// solution adapted from: https://stackoverflow.com/a/20551116
			i := 0
			for _, citySides := range cityFeatures {
				if citySides.OverlapsSide(sidesNeighbouringFieldCorners) {
					cityFeatures[i] = citySides
					i++
				}
			}
			cityFeatures = cityFeatures[:i]

			return cityFeatures
		}
	}
}

// Returns score report for this field. Has to be called after field.Expand() (todo?)
func (field Field) GetScoreReport() elements.ScoreReport {
	points := uint32(len(field.neighbouringCities) * 3)

	return elements.CalculateScoreReportOnMeeples(int(points), field.meeples)
}
