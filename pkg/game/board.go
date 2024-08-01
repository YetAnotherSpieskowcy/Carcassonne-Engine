package game

import (
	"errors"
	"fmt"
	"slices"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/city"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	positionMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
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
	tilesMap map[positionMod.Position]elements.PlacedTile

	placeablePositions []positionMod.Position
	cityManager        city.Manager
}

func NewBoard(tileSet tilesets.TileSet) elements.Board {
	tiles := make([]elements.PlacedTile, len(tileSet.Tiles)+1)
	startingTile := elements.NewStartingTile(tileSet)
	tiles[0] = startingTile
	cityManager := city.NewCityManager()
	cityManager.UpdateCities(startingTile)
	return &board{
		tileSet: tileSet,
		tiles:   tiles,
		tilesMap: map[positionMod.Position]elements.PlacedTile{
			positionMod.New(0, 0): startingTile,
		},
		placeablePositions: []positionMod.Position{
			positionMod.New(0, 1),
			positionMod.New(1, 0),
			positionMod.New(0, -1),
			positionMod.New(-1, 0),
		},
		cityManager: cityManager,
	}
}

func (board *board) TileCount() int {
	return len(board.tilesMap)
}

func (board *board) Tiles() []elements.PlacedTile {
	return board.tiles
}

func (board *board) GetTileAt(pos positionMod.Position) (elements.PlacedTile, bool) {
	elem, ok := board.tilesMap[pos]
	return elem, ok
}

func (board *board) GetTilePlacementsFor(tile tiles.Tile) []elements.PlacedTile {
	valid := []elements.PlacedTile{}
	rotations := tile.GetTileRotations()
	for _, currentTile := range rotations {
		for _, placeable := range board.placeablePositions {
			tilePlacement := elements.ToPlacedTile(currentTile)
			tilePlacement.Position = placeable
			if board.isPositionValid(tilePlacement) {
				valid = append(valid, tilePlacement)
			}
		}
	}
	return valid
}

// Verifies if a certain side of a adjacent tile on the board matches an expected feature type.
// The method takes a board, a position, the expected side, and the expected feature type.
// Returns a boolean indicating whether the tile has an expected feature on specified side.
func (board *board) testSide(position positionMod.Position, expectedSide side.Side, expectedFeatureType feature.Type) bool {
	var tile elements.PlacedTile
	var ok bool
	switch expectedSide {
	case side.Bottom:
		tile, ok = board.tilesMap[positionMod.New(position.X(), position.Y()+1)]
	case side.Top:
		tile, ok = board.tilesMap[positionMod.New(position.X(), position.Y()-1)]
	case side.Left:
		tile, ok = board.tilesMap[positionMod.New(position.X()+1, position.Y())]
	case side.Right:
		tile, ok = board.tilesMap[positionMod.New(position.X()-1, position.Y())]
	}
	if !ok {
		return true
	}
	for _, tileFeature := range tile.Features {
		if tileFeature.FeatureType == expectedFeatureType {
			if tileFeature.Sides.HasSide(expectedSide) {
				return true
			}
		}
	}
	return false
}

func (board *board) TileHasValidPlacement(tile tiles.Tile) bool {
	rotations := tile.GetTileRotations()
	for _, currentTile := range rotations {
		for _, placeable := range board.placeablePositions {
			tilePlacement := elements.ToPlacedTile(currentTile)
			tilePlacement.Position = placeable
			if board.isPositionValid(tilePlacement) {
				return true
			}
		}
	}
	return false
}

//revive:disable-next-line:unused-parameter Until the TODO is finished.
func (board *board) GetLegalMovesFor(tile elements.PlacedTile) []elements.PlacedTile {
	// TODO for future tasks:
	// - implement generation of legal moves
	// - to be implemented after #18, #19 and #20 to avoid code duplication
	return []elements.PlacedTile{}
}

func (board *board) isPositionValid(tile elements.PlacedTile) bool {
	for _, f := range tile.Features {
		s := side.Top
		for range 4 {
			if f.Sides.HasSide(s) && !board.testSide(tile.Position, s.Rotate(2), f.FeatureType) {
				return false
			}
			s = s.Rotate(1)
		}
	}
	return true
}

//revive:disable-next-line:unused-parameter Until the TODO is finished.
func (board *board) CanBePlaced(tile elements.PlacedTile) bool {
	// TODO for future tasks:
	// - implement generation of legal moves
	// - to be implemented after #18, #19 and #20 to avoid code duplication
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
			return elements.ToTile(tile).Equals(candidate)
		})
		if index == -1 {
			return elements.ScoreReport{}, errors.New(
				"Placed tile not found in the tile set, logic error?",
			)
		}
		actualIndex += index
		if !elements.ToTile(board.tiles[actualIndex]).Equals(elements.ToTile(tile)) {
			break
		}
		// position already taken, gotta find another next matching tile
		actualIndex++
		setTiles = setTiles[index+1:]
	}

	board.updateValidPlacements(tile)
	board.tiles[actualIndex] = tile
	board.tilesMap[tile.Position] = tile
	scoreReport, err := board.checkCompleted(tile)
	return scoreReport, err
}

func (board *board) RemoveMeeple(pos position.Position) {
	placedTile := board.tilesMap[pos]
	for featureIndex, feature := range placedTile.Features {
		if feature.Meeple.Type != elements.NoneMeeple {
			placedTile.Features[featureIndex].Meeple = elements.Meeple{Type: elements.NoneMeeple, PlayerID: elements.ID(0)}
			break
		}
	}
}

func (board *board) updateValidPlacements(tile elements.PlacedTile) {
	tileIndex := slices.Index(board.placeablePositions, tile.Position)
	if tileIndex == -1 {
		panic(fmt.Sprintf("Invalid move was played: %v", tile.Position))
	}
	board.placeablePositions = slices.Delete(board.placeablePositions, tileIndex, tileIndex+1)
	validNewPositions := []positionMod.Position{
		positionMod.New(tile.Position.X()+1, tile.Position.Y()),
		positionMod.New(tile.Position.X()-1, tile.Position.Y()),
		positionMod.New(tile.Position.X(), tile.Position.Y()+1),
		positionMod.New(tile.Position.X(), tile.Position.Y()-1),
	}
	for _, position := range validNewPositions {
		_, ok := board.tilesMap[position]
		if !ok && !slices.Contains(board.placeablePositions, position) {
			board.placeablePositions = append(board.placeablePositions, position)
		}
	}
}

func (board *board) checkCompleted(
	//revive:disable-next-line:unused-parameter Until the TODO is finished.
	tile elements.PlacedTile,
) (elements.ScoreReport, error) {
	// TODO for future tasks:
	// - identify all completed features
	// - resolve control of the completed features
	// - award points
	scoreReport := elements.NewScoreReport()
	board.cityManager.UpdateCities(tile)
	scoreReport.Join(board.cityManager.ScoreCities(false))
	scoreReport.Join(board.ScoreRoads(tile))
	return scoreReport, nil
}

/*
Calculates score for a single monastery.
If the monastery is finished and has a meeple, returns a ScoreReport with 9 points and the meeple that was in the monastery.
Otherwise, returns an empty ScoreReport.

'forceScore' can be set to true to score unfinished monasteries at the end of the game.
In other cases, 'forceScore' should be false

returns: ScoreReport (with one player at most)
*/
func (board *board) ScoreSingleMonastery(tile elements.PlacedTile, forceScore bool) (elements.ScoreReport, error) {
	var monasteryFeature = tile.Monastery()
	if monasteryFeature == nil {
		return elements.ScoreReport{}, errors.New("ScoreSingleMonastery() called on a tile without a monastery")
	}
	if monasteryFeature.Meeple.Type == elements.NoneMeeple {
		return elements.ScoreReport{}, errors.New("ScoreSingleMonastery() called on a tile without a meeple")
	}

	var score uint32
	for x := tile.Position.X() - 1; x <= tile.Position.X()+1; x++ {
		for y := tile.Position.Y() - 1; y <= tile.Position.Y()+1; y++ {
			_, ok := board.GetTileAt(positionMod.New(x, y))
			if ok {
				score++
			}
		}
	}

	if score == 9 || forceScore {
		scoreReport := elements.NewScoreReport()
		scoreReport.ReceivedPoints[monasteryFeature.PlayerID] = score
		scoreReport.ReturnedMeeples[monasteryFeature.PlayerID] = []elements.MeepleWithPosition{
			elements.MeepleWithPosition{
				Meeple:   monasteryFeature.Meeple,
				Position: tile.Position,
			},
		}

		return scoreReport, nil
	}

	return elements.NewScoreReport(), nil
}

/*
Finds all tiles with a monastery and a meeple in it adjacent to 'tile' (and 'tile' itself) and calls ScoreSingleMonastery on each of them.
This function should be called after the placement of each tile, in case it neighbours a monastery.

returns: ScoreReport
*/
func (board *board) ScoreMonasteries(tile elements.PlacedTile, forceScore bool) elements.ScoreReport {
	var finalReport = elements.NewScoreReport()

	for x := tile.Position.X() - 1; x <= tile.Position.X()+1; x++ {
		for y := tile.Position.Y() - 1; y <= tile.Position.Y()+1; y++ {
			adjacentTile, ok := board.GetTileAt(positionMod.New(x, y))

			if ok {
				report, err := board.ScoreSingleMonastery(adjacentTile, forceScore)
				if err == nil {
					finalReport.Join(report)
				}
			}
		}
	}
	return finalReport
}

/*
It analyzes road directed by roadSide parameter.
It doesn't analyze starting tile.
param roadSide: always indicates only one cardinal direction!
returns: road_finished, score, [meeples on road], loop, sideFinishedOn
sideFinishedOn matters only if loop is True. Variable used to prevent checking the same road twice in ScoreRoads function
*/
func (board *board) CheckRoadInDirection(roadSide side.Side, startTile elements.PlacedTile) (bool, int, []elements.MeepleWithPosition, bool, side.Side) {
	var meeples = []elements.MeepleWithPosition{}
	var tile = startTile
	var tileExists bool
	var score = 0
	var road *elements.PlacedFeature
	var finished bool
	// check finished on way
	// do while loop
	for {
		tile, tileExists = board.GetTileAt(tile.Position.Add(positionMod.FromSide(roadSide)))
		roadSide = roadSide.Mirror()
		// check if tile exists or loop
		if !tileExists || tile.Position == startTile.Position {
			// tile does not exist
			// or loop found
			// so finish
			break
		}

		score++

		// Get road feature
		road = tile.GetPlacedFeatureAtSide(roadSide, feature.Road)

		// check if there is meeple on the feature
		if road.Meeple.Type != elements.NoneMeeple {
			meeples = append(meeples, elements.NewMeepleWithPosition(
				road.Meeple,
				tile.Position),
			)
		}

		if road.Sides.GetCardinalDirectionsLength() == 1 {
			// found the end of a road
			break
		}
		// swap to other end of the road on the same tile
		roadSide = road.Sides.GetConnectedOtherCardinalDirection(roadSide)

	}

	looped := (tile.Position == startTile.Position)
	finished = tileExists && (road.Sides.GetCardinalDirectionsLength() == 1 || looped)

	return finished, score, meeples, looped, roadSide
}

/*
Calculates score for road.

returns: ScoreReport, checked sides of the start tile (also including loop)
*/
func (board *board) ScoreRoadCompletion(tile elements.PlacedTile, road feature.Feature) (elements.ScoreReport, side.Side) {
	var meeples = []elements.MeepleWithPosition{}
	var leftSide, rightSide side.Side
	var score = 1
	leftSide = road.Sides.GetNthCardinalDirection(0)  // first side
	rightSide = road.Sides.GetNthCardinalDirection(1) // second side
	var roadFinished = true

	var roadFinishedResult bool
	var scoreResult int
	var meeplesResult []elements.MeepleWithPosition
	var loopResult bool
	var loopSide side.Side

	// check meeples on start tile
	var roadLeft = tile.GetPlacedFeatureAtSide(leftSide, feature.Road)
	var roadRight = tile.GetPlacedFeatureAtSide(rightSide, feature.Road)
	if roadLeft.Meeple.Type != elements.NoneMeeple {
		meeples = append(meeples, elements.NewMeepleWithPosition(
			roadLeft.Meeple,
			tile.Position),
		)
	} else if roadRight != nil && roadRight.Meeple.Type != elements.NoneMeeple {
		meeples = append(meeples, elements.NewMeepleWithPosition(
			roadRight.Meeple,
			tile.Position),
		)
	}

	// check road in "left" direction
	roadFinishedResult, scoreResult, meeplesResult, loopResult, loopSide = board.CheckRoadInDirection(leftSide, tile)
	score += scoreResult
	roadFinished = roadFinished && roadFinishedResult
	meeples = append(meeples, meeplesResult...)

	// check road in "right" direction
	if !loopResult && rightSide != side.NoSide {
		roadFinishedResult, scoreResult, meeplesResult, _, _ = board.CheckRoadInDirection(rightSide, tile)
		score += scoreResult
		roadFinished = roadFinished && roadFinishedResult
		meeples = append(meeples, meeplesResult...)
	}

	// -------- start counting -------------
	if roadFinished {
		if loopResult {
			return elements.CalculateScoreReportOnMeeples(score, meeples), leftSide | rightSide | loopSide
		}
		return elements.CalculateScoreReportOnMeeples(score, meeples), leftSide | rightSide

	}

	// return empty report
	return elements.NewScoreReport(), leftSide | rightSide

}

/*
Calculates summary score report from all roads on a tile
*/
func (board *board) ScoreRoads(placedTile elements.PlacedTile) elements.ScoreReport {
	scoreReport := elements.NewScoreReport()
	var tile = elements.ToTile(placedTile)
	var roads = tile.Roads()

	var checkedRoadSides side.Side

	for _, road := range roads {
		// check if the side of the tile was not already checked (special test case reference: TestBoardScoreRoadLoopCrossroad)
		if !checkedRoadSides.OverlapsSide(road.Sides) {
			scoreReportTemp, roadSide := board.ScoreRoadCompletion(placedTile, road)
			scoreReport.Join(scoreReportTemp)
			checkedRoadSides |= roadSide
		}
		checkedRoadSides |= road.Sides
	}
	return scoreReport
}
