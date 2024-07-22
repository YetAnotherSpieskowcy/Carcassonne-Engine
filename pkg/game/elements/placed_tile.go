package elements

import (
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

type TileWithMeeple struct {
	Features  []PlacedFeature
	HasShield bool
}

type PlacedFeature struct {
	feature.Feature
	Meeple
}

func (placedTile PlacedTile) Rotate(rotations uint) PlacedTile {
	_ = rotations
	panic("Rotate() not supported on PlacedTile")
}

func ToPlacedTile(tile tiles.Tile) PlacedTile {
	features := []PlacedFeature{}
	for _, n := range tile.Features {
		features = append(features, PlacedFeature{n, Meeple{NoneMeeple, NonePlayer}})
	}
	return PlacedTile{
		TileWithMeeple: TileWithMeeple{
			Features: features,
		},
		Position: position.New(0, 0),
	}
}

func ToTile(tile PlacedTile) tiles.Tile {
	features := []feature.Feature{}
	for _, n := range tile.Features {
		features = append(features, n.Feature)
	}
	return tiles.Tile{
		Features: features,
	}
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

/*
Return a list of all features of the given type that overlaps the given side. The overlap does not need to be exact.
*/
func (placedTile PlacedTile) GetPlacedFeaturesOverlappingSide(sideToCheck side.Side, featureType feature.Type) []PlacedFeature {
	features := []PlacedFeature{}
	for _, feature := range placedTile.Features {
		if sideToCheck&feature.Sides != 0 && feature.FeatureType == featureType {
			features = append(features, feature)
		}
	}
	return features
}

// represents a legal move (tile placement and meeple placement) on the board
type PlacedTile struct {
	TileWithMeeple
	Position position.Position
}

func NewStartingTile(tileSet tilesets.TileSet) PlacedTile {
	return ToPlacedTile(tileSet.StartingTile)
}

func (placedTile PlacedTile) Monastery() *PlacedFeature {
	for i, feat := range placedTile.Features {
		if feat.FeatureType == feature.Monastery {
			return &placedTile.Features[i]
		}
	}
	return nil
}

/*
Return the feature of certain type on desired side
*/
func (placedTile *PlacedTile) GetPlacedFeatureAtSide(sideToCheck side.Side, featureType feature.Type) *PlacedFeature {
	for i, feature := range placedTile.Features {
		if sideToCheck&feature.Sides == sideToCheck && feature.FeatureType == featureType {
			return &placedTile.Features[i]
		}
	}
	return nil
}
