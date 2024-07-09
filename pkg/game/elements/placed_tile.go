package elements

import (
	"fmt"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	featureMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	sideMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

type Position struct {
	// int8 would be fine for base game (72 tiles) but let's be a bit more generous
	x int16
	y int16
}

func NewPosition(x int16, y int16) Position {
	return Position{x, y}
}

func (pos Position) X() int16 {
	return pos.x
}

func (pos Position) Y() int16 {
	return pos.y
}

func (pos Position) Add(other Position) Position {
	return NewPosition(pos.x+other.x, pos.y+other.y)
}

/*
Returns relative position directed by the side.
Caution! It is supposed to be used with side directing only one cardinal direction (or two edges connected by corner)!
Otherwise it will return undesired value!
*/
func PositionFromSide(side sideMod.Side) Position {
	switch {
	case side&sideMod.Top != 0:
		return NewPosition(0, 1)
	case side&sideMod.Right != 0:
		return NewPosition(1, 0)
	case side&sideMod.Left != 0:
		return NewPosition(-1, 0)
	case side&sideMod.Bottom != 0:
		return NewPosition(0, -1)

	default:
		return NewPosition(0, 0)
	}
}

func (pos Position) MarshalText() ([]byte, error) {
	return fmt.Appendf([]byte{}, "%v,%v", pos.x, pos.y), nil
}

func (pos *Position) UnmarshalText(text []byte) error {
	_, err := fmt.Sscanf(string(text), "%v,%v", &pos.x, &pos.y)
	return err
}

func AddPositions(pos1 Position, pos2 Position) Position {
	return Position{
		x: pos1.x + pos2.x,
		y: pos1.y + pos2.y,
	}
}

// https://wikicarpedia.com/car/Game_Figures
type MeepleType uint8

const (
	NoneMeeple MeepleType = iota
	NormalMeeple

	MeepleTypeCount int = iota
)

type Meeple struct {
	MeepleType
	PlayerID ID
}

type TileWithMeeple struct {
	Features  []PlacedFeature
	HasShield bool
}

func (placedTile PlacedTile) Rotate(_ uint) PlacedTile {
	panic("Rotate() not supported on PlacedTile")
}

type PlacedFeature struct {
	featureMod.Feature
	Meeple
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
		Position: NewPosition(0, 0),
	}
}

func ToTile(tile PlacedTile) tiles.Tile {
	features := []featureMod.Feature{}
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
		if f.FeatureType == featureMod.City {
			cityFeatures = append(cityFeatures, f)
		}
	}
	return cityFeatures
}

// represents a legal move (tile placement and meeple placement) on the board
type PlacedTile struct {
	TileWithMeeple
	Position Position
}

func NewStartingTile(tileSet tilesets.TileSet) PlacedTile {
	return ToPlacedTile(tileSet.StartingTile)
}

func (placedTile PlacedTile) Monastery() *PlacedFeature {
	for i, feature := range placedTile.Features {
		if feature.FeatureType == featureMod.Monastery {
			return &placedTile.Features[i]
		}
	}
	return nil
}

/*
Return the feature of certain type on desired side
*/
func (placedTile *PlacedTile) GetPlacedFeatureAtSide(sideToCheck sideMod.Side, featureType featureMod.Type) *PlacedFeature {
	for i, feature := range placedTile.TileWithMeeple.Features {
		if sideToCheck&feature.Sides == sideToCheck && feature.FeatureType == featureType {
			return &placedTile.TileWithMeeple.Features[i]
		}
	}
	return nil
}
