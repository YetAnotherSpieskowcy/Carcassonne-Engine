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

func (placedTile PlacedTile) GetCityFeatures() []PlacedFeature {
	cityFeatures := []PlacedFeature{}
	for _, f := range placedTile.Features {
		if f.FeatureType == feature.City {
			cityFeatures = append(cityFeatures, f)
		}
	}
	return cityFeatures
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
	for i, feat := range placedTile.TileWithMeeple.Features {
		if sideToCheck&feat.Sides == sideToCheck && feat.FeatureType == featureType {
			return &placedTile.TileWithMeeple.Features[i]
		}
	}
	return nil
}
