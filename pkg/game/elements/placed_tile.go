package elements

import (
	"fmt"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
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
func PositionFromSide(sideToCheck side.Side) Position {
	position := NewPosition(0, 0)

	if sideToCheck&side.Top != 0 {
		position = position.Add(NewPosition(0, 1))
	}

	if sideToCheck&side.Right != 0 {
		position = position.Add(NewPosition(1, 0))
	}

	if sideToCheck&side.Bottom != 0 {
		position = position.Add(NewPosition(0, -1))
	}

	if sideToCheck&side.Left != 0 {
		position = position.Add(NewPosition(-1, 0))
	}

	return position
}

func (pos Position) MarshalText() ([]byte, error) {
	return fmt.Appendf([]byte{}, "%v,%v", pos.x, pos.y), nil
}

func (pos *Position) UnmarshalText(text []byte) error {
	_, err := fmt.Sscanf(string(text), "%v,%v", &pos.x, &pos.y)
	return err
}

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

func (placedTile PlacedTile) Rotate(rotations uint) PlacedTile {
	_ = rotations
	panic("Rotate() not supported on PlacedTile")
}

type PlacedFeature struct {
	feature.Feature
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
	Position Position
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
