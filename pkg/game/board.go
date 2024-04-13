package game

import (
	"errors"
	"slices"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

// mutable type
// Position coordinates example on the board:
// (-1, +1)  (+0, +1)  (+1, +1)
// (-1, +0)  (+0, +0)  (+1, +0)
// (-1, -1)  (+0, -1)  (+1, -1)
//
// Starting tile is placed at (+0, +0) position.
type board struct {
	tileSet tilesets.TileSet
	// The information about the tile and its placement is stored sparsely
	// in a slice of size equal to the number of tiles in the set.
	// `tiles[0]` is always the starting tile.
	// Indexes in `tiles[1:]` correspond to positions of equivalent tiles in `tileSet`.
	tiles []elements.PlacedTile
	// tilesMap is used by the engine for faster lookups
	// but contains the same information as the `tiles` slice.
	tilesMap map[elements.Position]elements.PlacedTile
}

func NewBoard(tileSet tilesets.TileSet) elements.Board {
	tiles := make([]elements.PlacedTile, len(tileSet.Tiles)+1)
	startingTile := elements.NewStartingTile(tileSet)
	tiles[0] = startingTile
	return &board{
		tileSet: tileSet,
		tiles:   tiles,
		tilesMap: map[elements.Position]elements.PlacedTile{
			elements.NewPosition(0, 0): startingTile,
		},
	}
}

func (board *board) TileCount() int {
	return len(board.tilesMap)
}

func (board *board) Tiles() []elements.PlacedTile {
	return board.tiles
}

func (board *board) GetTileAt(pos elements.Position) (elements.PlacedTile, bool) {
	elem, ok := board.tilesMap[pos]
	return elem, ok
}

//revive:disable-next-line:unused-parameter Until the TODO is finished.
func (board *board) GetTilePlacementsFor(tile tiles.Tile) []elements.TilePlacement {
	// TODO for future tasks:
	// - implement generation of legal tile placements
	return []elements.TilePlacement{}
}

// early return variant of above
//
//revive:disable-next-line:unused-parameter Until the TODO is finished.
func (board *board) TileHasValidPlacement(tile tiles.Tile) bool {
	// TODO for future tasks:
	// - implement generation of legal tile placements
	return true
}

//revive:disable-next-line:unused-parameter Until the TODO is finished.
func (board *board) GetLegalMovesFor(tile elements.TilePlacement) []elements.LegalMove {
	// TODO for future tasks:
	// - implement generation of legal moves
	return []elements.LegalMove{}
}

//revive:disable-next-line:unused-parameter Until the TODO is finished.
func (board *board) CanBePlaced(tile elements.PlacedTile) bool {
	// TODO for future tasks:
	// - implement a way to check if a specified move is valid
	return true
}

func (board *board) PlaceTile(tile elements.PlacedTile) (elements.ScoreReport, error) {
	if board.TileCount() == cap(board.tiles) {
		return elements.ScoreReport{}, errors.New(
			"Board's tiles capacity exceeded, logic error?",
		)
	}
	// TODO for future tasks:
	// - determine if the tile can placed at a given position,
	//   or return ErrInvalidMove otherwise
	setTiles := board.tileSet.Tiles
	actualIndex := 1
	for {
		index := slices.IndexFunc(setTiles, func(candidate tiles.Tile) bool {
			return tile.Tile.Equals(candidate)
		})
		if index == -1 {
			return elements.ScoreReport{}, errors.New(
				"Placed tile not found in the tile set, logic error?",
			)
		}
		actualIndex += index
		if !board.tiles[actualIndex].Tile.Equals(tile.Tile) {
			break
		}
		// position already taken, gotta find another next matching tile
		actualIndex++
		setTiles = setTiles[index+1:]
	}
	board.tiles[actualIndex] = tile
	board.tilesMap[tile.Pos] = tile
	scoreReport, err := board.checkCompleted(tile)
	return scoreReport, err
}

func (board *board) checkCompleted(
	//revive:disable-next-line:unused-parameter Until the TODO is finished.
	tile elements.PlacedTile,
) (elements.ScoreReport, error) {
	// TODO for future tasks:
	// - identify all completed features
	// - resolve control of the completed features
	// - award points
	scoreReport := elements.ScoreReport{
		ReceivedPoints:  map[int]uint32{},
		ReturnedMeeples: map[int][]uint8{},
	}
	return scoreReport, nil
}

/*
Calculates score for road completion.
If argument forceScore is false, it won't force scoring if road is unfinished.
ForceScore is supposed to be true while counting unfinished roads.

returns: what should it return?
*/
func (board *board) ScoreRoadCompletion(tile elements.PlacedTile, road feature.Feature, forceScore bool) {

	// tuple type for saving meeples and their positions
	type MeepleTilePlacement struct {
		elements.MeeplePlacement
		elements.PlacedTile
	}

	var meeples = []MeepleTilePlacement{}
	var leftTile, rightTile elements.PlacedTile
	var leftRoad, rightRoad feature.Feature
	var leftSide, rightSide side.Side
	var score = 1
	var tileExists = false
	leftSide = road.Sides[0]
	rightSide = road.Sides[1]

	//check meeples on start tile
	if tile.Meeple.Side == leftSide || tile.Meeple.Side == rightSide {
		meeples = append(meeples, MeepleTilePlacement{tile.Meeple, tile})
	}

	// check if one end is already road end
	if leftSide == side.Center {
		leftTile = tile
		leftSide = side.Center
	} else if rightSide == side.Center {
		rightTile = tile
		rightSide = side.Center
	}

	// check finished on "left" way
	for leftSide != side.Center /* and check for loop */ {

		leftTile, tileExists = board.GetTileAt(leftTile.Pos.Add(elements.PositionFromSide(leftSide)))
		// check if tile exists
		if !tileExists {
			// tile does not exist
			// finish
			break
		}

		score++
		leftSide, _ = leftSide.ConnectedOpposite()
		// check error

		//check for meeple1
		if leftTile.Meeple.Side == leftSide {
			meeples = append(meeples, MeepleTilePlacement{leftTile.Meeple, leftTile})
		}

		// get road feature
		leftRoad = *leftTile.GetFeatureAtSide(leftSide) //check isn't needed because it was already checked at legalmoves

		// swap to other end of tile
		if leftRoad.Sides[0] == leftSide {
			leftSide = leftRoad.Sides[1]
		} else {
			leftSide = leftRoad.Sides[0]
		}

		//check for meeple2 (other end of road)
		if leftTile.Meeple.Side == leftSide {
			meeples = append(meeples, MeepleTilePlacement{leftTile.Meeple, leftTile})
		}

	}

	// check if loop
	if leftTile.Pos != tile.Pos {
		// no loop found, so check other side

		// check if leftRoad was finished (if the end is center), or check right side, because the end of game
		if leftSide == side.Center || forceScore {

			// check finished on "right" way
			for rightSide != side.Center /* and check for loop */ {
				rightTile, tileExists = board.GetTileAt(rightTile.Pos.Add(elements.PositionFromSide(rightSide)))
				// check if tile exists
				if !tileExists {
					// tile does not exist
					// finish
					break
				}

				rightSide, _ = rightSide.ConnectedOpposite()
				//check error

				//check for meeple1
				if rightTile.Meeple.Side == rightSide {
					meeples = append(meeples, MeepleTilePlacement{rightTile.Meeple, rightTile})
				}

				//get road feature
				rightRoad = *rightTile.GetFeatureAtSide(rightSide) //check isn't needed because it was already checked at legalmoves
				//swap to other end of tile
				if rightRoad.Sides[0] == rightSide {
					rightSide = rightRoad.Sides[1]
				} else {
					rightSide = rightRoad.Sides[0]
				}

				//check for meeple2 (other end of road)
				if rightTile.Meeple.Side == rightSide {
					meeples = append(meeples, MeepleTilePlacement{rightTile.Meeple, rightTile})
				}
			}
		}
	}

	// -------- start counting -------------

	//check if both ends are in center, or loop, or forces counting
	if (leftSide == side.Center && rightSide == side.Center) || leftTile.Pos == tile.Pos || forceScore {
		for _, meeple := range meeples {
			// remove meeples from tile
			// return meeples to player
			// add score to player	with most meeples
		}
	}

}
