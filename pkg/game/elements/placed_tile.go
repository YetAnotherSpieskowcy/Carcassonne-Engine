package elements

import (
	"slices"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

// https://wikicarpedia.com/car/Game_Figures
type MeepleType uint8

const (
	NoneMeeple MeepleType = iota
	NormalMeeple

	MeepleTypeCount int = iota
)

type Meeple struct {
	Type     MeepleType
	PlayerID ID
}

type PlacedFeature struct {
	feature.Feature
	Meeple Meeple
}

// Represents a legal move (tile placement and meeple placement) on the board
type PlacedTile struct {
	Features []PlacedFeature
	Position position.Position
}

func (placedTile PlacedTile) DeepClone() PlacedTile {
	placedTile.Features = slices.Clone(placedTile.Features)
	return placedTile
}

func (placedTile PlacedTile) Rotate(rotations uint) PlacedTile {
	_ = rotations
	panic("Rotate() not supported on PlacedTile")
}

func ToPlacedTile(tile tiles.Tile) PlacedTile {
	features := []PlacedFeature{}
	for _, feature := range tile.Features {
		features = append(features, PlacedFeature{feature, Meeple{NoneMeeple, NonePlayer}})
	}
	return PlacedTile{
		Features: features,
		Position: position.New(0, 0),
	}
}

func ToTile(tile PlacedTile) tiles.Tile {
	features := []feature.Feature{}
	for _, feature := range tile.Features {
		features = append(features, feature.Feature)
	}
	return tiles.Tile{
		Features: features,
	}
}

// Returns true if placedTile equals tile and false otherwise
// The comparison ignores meeples and position but includes orientation
// Features of tile *MUST* be in the same order as in placedTile for the tiles to be considered equal
// (e.g. tile with monastery and field != tile with field and monastery, even if their sides are the same)
func (placedTile PlacedTile) ExactEqualsTile(tile tiles.Tile) bool {
	if len(tile.Features) != len(placedTile.Features) {
		return false
	}
	for i := range tile.Features {
		if tile.Features[i] != placedTile.Features[i].Feature {
			return false
		}
	}
	return true
}

// Returns true if placedTile equals tile and false otherwise
// The comparison ignores meeples, position and orientation
// Features of tile *MUST* be in the same order as in placedTile for the tiles to be considered equal
// (e.g. tile with monastery and field != tile with field and monastery, even if their sides are the same)
func (placedTile PlacedTile) EqualsTile(tile tiles.Tile) bool {
	if len(tile.Features) != len(placedTile.Features) {
		return false
	}
outer:
	for rotations := range uint(4) {
		rotated := tile.Rotate(rotations)
		for i, placedFeature := range placedTile.Features {
			if placedFeature.Feature != rotated.Features[i] {
				continue outer
			}
		}
		return true
	}
	return false
}

// Returns a list of all features of the given type on this tile
func (placedTile PlacedTile) GetFeaturesOfType(featureType feature.Type) []PlacedFeature {
	features := []PlacedFeature{}
	for _, feature := range placedTile.Features {
		if feature.FeatureType == featureType {
			features = append(features, feature)
		}
	}
	return features
}

// Return a list of all features of the given type that overlap the given side. The overlap does not need to be exact.
func (placedTile PlacedTile) GetPlacedFeaturesOverlappingSide(sideToCheck side.Side, featureType feature.Type) []PlacedFeature {
	features := []PlacedFeature{}
	for _, feature := range placedTile.Features {
		if sideToCheck.OverlapsSide(feature.Sides) && feature.FeatureType == featureType {
			features = append(features, feature)
		}
	}
	return features
}

// Return the feature of certain type on desired side
func (placedTile *PlacedTile) GetPlacedFeatureAtSide(sideToCheck side.Side, featureType feature.Type) *PlacedFeature {
	for i, feature := range placedTile.Features {
		if feature.Sides.HasSide(sideToCheck) && feature.FeatureType == featureType {
			return &placedTile.Features[i]
		}
	}
	return nil
}

func (placedTile PlacedTile) HasMeeple() bool {
	for _, feat := range placedTile.Features {
		if feat.Meeple.Type != NoneMeeple {
			return true
		}
	}
	return false
}

func (placedTile PlacedTile) Monastery() *PlacedFeature {
	for i, feat := range placedTile.Features {
		if feat.FeatureType == feature.Monastery {
			return &placedTile.Features[i]
		}
	}
	return nil
}

func NewStartingTile(tileSet tilesets.TileSet) PlacedTile {
	return ToPlacedTile(tileSet.StartingTile)
}
