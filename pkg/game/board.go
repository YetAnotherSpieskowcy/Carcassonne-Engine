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
		ReceivedPoints:  map[uint8]uint32{},
		ReturnedMeeples: map[uint8][]uint8{},
	}
	return scoreReport, nil
}

/*
It doesn't analyze starting tile.
It analyzes road directed by roadSide parameter.
returns: road_finished, score, [meeples on road], loop
*/
func (board *board) CheckRoadInDirection(roadSide side.Side, startTile elements.PlacedTile) (bool, int, []elements.MeepleTilePlacement, bool) {
	var meeples = []elements.MeepleTilePlacement{}
	var tile = startTile
	var tileExists bool
	var score = 0
	var road feature.Feature
	var finished bool
	var singleIterationMade = false // to prevent ending before entering loop (f.e.: placed tile is a monastery with a road, so one side is Center from the beginning but it's not loop)
	// check finished on way
	// do while loop
	for {
		singleIterationMade = true
		tile, tileExists = board.GetTileAt(tile.Pos.Add(elements.PositionFromSide(roadSide)))
		// check if tile exists or loop
		if !tileExists || tile.Pos == startTile.Pos {
			// tile does not exist
			// finish

			break
		}

		score++
		roadSide = roadSide.ConnectedOpposite()

		// check for meeple1
		if tile.Meeple.Side == roadSide {
			meeples = append(meeples, elements.MeepleTilePlacement{MeeplePlacement: tile.Meeple, PlacedTile: tile})
		}

		// get road feature
		road = *tile.GetFeatureAtSide(roadSide) // check isn't needed because it was already checked at legalmoves

		// swap to other end of tile
		if road.Sides.GetNthCardinalDirection(0) == roadSide {
			roadSide = road.Sides.GetNthCardinalDirection(1)
		} else {
			roadSide = road.Sides.GetNthCardinalDirection(0)
		}

		// check for meeple2 (other end of road)
		if tile.Meeple.Side == roadSide {
			meeples = append(meeples, elements.MeepleTilePlacement{MeeplePlacement: tile.Meeple, PlacedTile: tile})
		}

		if road.Sides.GetCardinalDirectionsLength() == 1 {

			break
		}
	}

	finished = (road.Sides.GetCardinalDirectionsLength() == 1) || (tile.Pos == startTile.Pos)
	finished = finished && tileExists

	looped := (tile.Pos == startTile.Pos) && singleIterationMade
	return finished, score, meeples, looped
}

/*
Calculates score for road.

returns: ScoreReport
*/
func (board *board) ScoreRoadCompletion(tile elements.PlacedTile, road feature.Feature) elements.ScoreReport {
	var meeples = []elements.MeepleTilePlacement{}
	var leftSide, rightSide side.Side
	var score = 1
	leftSide = road.Sides.GetNthCardinalDirection(0)
	rightSide = road.Sides.GetNthCardinalDirection(1)
	var roadFinished = true

	var roadFinishedResult bool
	var scoreResult int
	var meeplesResult []elements.MeepleTilePlacement
	var loopResult bool

	// check meeples on start tile
	if tile.Meeple.Side&leftSide == leftSide || (tile.Meeple.Side&rightSide == rightSide && rightSide != side.None) {
		meeples = append(meeples, elements.MeepleTilePlacement{MeeplePlacement: tile.Meeple, PlacedTile: tile})
	}

	roadFinishedResult, scoreResult, meeplesResult, loopResult = board.CheckRoadInDirection(leftSide, tile)
	score += scoreResult
	roadFinished = roadFinished && roadFinishedResult
	meeples = append(meeples, meeplesResult...)

	if !loopResult && rightSide != side.None {
		roadFinishedResult, scoreResult, meeplesResult, _ = board.CheckRoadInDirection(rightSide, tile)
		score += scoreResult
		roadFinished = roadFinished && roadFinishedResult
		meeples = append(meeples, meeplesResult...)
	}

	// -------- start counting -------------
	if roadFinished {

		return elements.CalculateScoreReportOnMeeples(score, meeples)
	}

	// return empty report
	return elements.NewScoreReport()

}

/*
Calculates summary score report from all roads on a tile
*/
func (board *board) ScoreRoads(tile elements.PlacedTile) elements.ScoreReport {
	scoreReport := elements.NewScoreReport()
	for _, road := range tile.Roads() {
		scoreReportTemp := board.ScoreRoadCompletion(tile, road)
		scoreReport.JoinReport(scoreReportTemp)
	}
	return scoreReport
}
