package city

import (
	"maps"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/binarytiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
)

// Represents cities on board
type City struct {
	completed bool
	scored    bool
	features  map[position.Position][]binarytiles.BinaryTileFeature
	shields   uint8
}

func NewCity(tile binarytiles.BinaryTile, cityFeatures []binarytiles.BinaryTileSide) City {
	features := make([]binarytiles.BinaryTileFeature, len(cityFeatures))
	var shields = uint8(0)
	for i, feature := range cityFeatures {
		if tile.HasShieldAtSide(feature) {
			shields++
		}
		features[i] = binarytiles.BinaryTileFeature{Side: feature, Tile: tile}
	}
	return City{
		completed: false,
		scored:    false,
		features: map[position.Position][]binarytiles.BinaryTileFeature{
			tile.Position(): features,
		},
		shields: shields,
	}
}

func (city City) DeepClone() City {
	city.features = maps.Clone(city.features)
	// shields number is already copied for being uint8
	return city
}

func (city City) IsCompleted() bool {
	return city.completed
}

// Checks if city is closed and sets city.completed.
func (city *City) checkCompleted() bool {
	city.completed = true
	for pos, features := range city.features {
		for _, feature := range features {
			for _, side := range binarytiles.OrthogonalSides {
				if feature.Side.OverlapsSide(side) {
					neighbouringPosition := side.PositionFromSide()
					_, ok := city.GetFeaturesFromTile(pos.Add(neighbouringPosition))
					if !ok {
						city.completed = false
						return false
					}
				}
			}
		}
	}
	return true
}

func (city *City) SetScored(scored bool) {
	city.scored = scored
}

// Calculates score value of the city and
// determines players that should receive points.
func (city *City) GetScoreReport() elements.ScoreReport {
	var returnedMeeples = []elements.MeepleWithPosition{}
	var totalScore uint8
	// calculate total value of the city and get all meeples
	for pos, features := range city.features {
		for _, feat := range features {
			playerID := feat.Tile.GetMeepleIDAtSide(feat.Side, feature.City)
			if playerID != elements.NonePlayer {
				returnedMeeples = append(returnedMeeples, elements.NewMeepleWithPosition(
					elements.Meeple{Type: elements.NormalMeeple, PlayerID: playerID},
					pos,
				))
			}
		}
		totalScore += 2
	}
	totalScore += city.shields * 2

	if !city.completed {
		totalScore /= 2
	}

	return elements.CalculateScoreReportOnMeeples(int(totalScore), returnedMeeples)
}

// Returns all features from a tile at a given position that are part of a city
// and whether such a tile is in the city.
func (city City) GetFeaturesFromTile(pos position.Position) ([]binarytiles.BinaryTileFeature, bool) {
	cities, ok := city.features[pos]
	return cities, ok
}

func (city *City) AddTile(tile binarytiles.BinaryTile, cityFeatures []binarytiles.BinaryTileSide) {
	hasShield := false
	features := make([]binarytiles.BinaryTileFeature, len(cityFeatures))
	for i, feature := range cityFeatures {
		if !hasShield && tile.HasShieldAtSide(feature) {
			hasShield = true
		}
		features[i] = binarytiles.BinaryTileFeature{Side: feature, Tile: tile}
	}

	_, tileInCity := city.GetFeaturesFromTile(tile.Position())
	if !tileInCity {
		city.features[tile.Position()] = features
		if hasShield {
			city.shields++
		}

	} else {
		city.features[tile.Position()] = append(city.features[tile.Position()], features...)
		// If the tile being checked is already in the city, it must have the shield already counted,
		// because tiles with shields, has only one city feature, so it's impossible
		// that there are multiple city features on that tile.
	}
	city.checkCompleted()
}

// Merges two cities when they are connected.
// Other city must be deleted after to avoid problems
func (city *City) JoinCities(other City) {
	for pos, otherFeature := range other.features {
		feature, ok := city.GetFeaturesFromTile(pos)
		if ok {
			features := feature
			features = append(features, otherFeature...)
			city.features[pos] = features
		} else {
			city.features[pos] = otherFeature
		}
	}
	city.shields += other.shields
	city.checkCompleted()
}
