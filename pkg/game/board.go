package game

import (
	"errors"
	"fmt"
	"maps"
	"slices"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/city"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/field"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

var (
	canBePlacedFunctions = map[feature.Type]func(*board, elements.PlacedTile, elements.PlacedFeature) bool{
		feature.Road:      (*board).roadCanBePlaced,
		feature.City:      (*board).cityCanBePlaced,
		feature.Field:     (*board).fieldCanBePlaced,
		feature.Monastery: (*board).monasteryCanBePlaced,
	}
	meepleTypes = []elements.MeepleType{elements.NormalMeeple}
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
	tilesMap map[position.Position]elements.PlacedTile

	placeablePositions []position.Position
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
		tilesMap: map[position.Position]elements.PlacedTile{
			position.New(0, 0): startingTile,
		},
		placeablePositions: []position.Position{
			position.New(0, 1),
			position.New(1, 0),
			position.New(0, -1),
			position.New(-1, 0),
		},
		cityManager: cityManager,
	}
}

func (board board) DeepClone() elements.Board {
	// note: skipped board.tileSet because TileSet is immutable

	// PlacedTile is not mutated after it's placed
	board.tiles = slices.Clone(board.tiles)
	board.tilesMap = maps.Clone(board.tilesMap)

	// Position is immutable
	board.placeablePositions = slices.Clone(board.placeablePositions)

	board.cityManager = board.cityManager.DeepClone()

	return &board
}

func (board *board) TileCount() int {
	return len(board.tilesMap)
}

func (board *board) Tiles() []elements.PlacedTile {
	return board.tiles
}

func (board *board) GetTileAt(pos position.Position) (elements.PlacedTile, bool) {
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

// Get legal moves that can be made from the given **valid** placement.
//
// This means that this function returns a slice with:
//   - the given placement as is (assumed to have no meeple)
//   - all variations of the given placement with meeple added to one of the features
//     that a meeple can be placed on considering the current situation on the board.
//
// Note: the placement given as input is assumed to have no meeples.
func (board *board) GetLegalMovesFor(basePlacement elements.PlacedTile) []elements.PlacedTile {
	// create initial move list without any meeple placed
	moves := []elements.PlacedTile{basePlacement}

	for i := range basePlacement.Features {
		for _, meepleType := range meepleTypes {
			placement := basePlacement.DeepClone()
			placement.Features[i].Meeple = elements.Meeple{Type: meepleType}
			feat := placement.Features[i]

			// Doing this for every meeple type may be suboptimal, if more meeple types
			// are added but that's not very likely at current time.
			if canBePlacedFunctions[feat.FeatureType](board, placement, feat) {
				moves = append(moves, placement)
			}
		}
	}
	return moves
}

/*
Returns true if the tile placement position is valid, i.e. if all existing neighbouring tiles have matching features.
(for example, city feature directly neighbouring road or field is not valid)

Only checks the validity of the tile placement based on the tile features and their neighbors. It does not take into account:
- The placement of meeples
- Whether the tile is being placed on an already occupied position
- Whether the tile is not neighboring any tiles
- Whether the tile has already been placed somewhere else on the board
*/
func (board *board) isPositionValid(tile elements.PlacedTile) bool {
	// Phase 1:
	// For all of the given tile's features, we need to check that all of its sides
	// have either a matching counterpart on the side's neighbouring tile
	// or are left unconnected.
	for _, tileFeature := range tile.Features {
		for _, side := range side.EdgeSides {
			if tileFeature.Sides.HasSide(side) {
				neighbourPosition := position.FromSide(side).Add(tile.Position)
				neighbouringTile, exists := board.GetTileAt(neighbourPosition)
				if exists && neighbouringTile.GetPlacedFeatureAtSide(side.Mirror(), tileFeature.FeatureType) == nil {
					return false
				}
			}

		}
	}
	// Phase 2:
	// Since some features may overlap other features, it is also necessary to check
	// that none of the neighbours have an (overlapping) feature that doesn't have
	// a matching counterpart on the given tile.
	// Currently, overlap can only occur between roads and fields. Since roads
	// are always accompanied by fields, we only need to check roads.
	//
	// TODO: rivers will probably have the same problems as roads, if they are implemented
	for _, side := range side.PrimarySides {
		neighbourPosition := position.FromSide(side).Add(tile.Position)
		neighbouringTile, exists := board.GetTileAt(neighbourPosition)
		if exists && neighbouringTile.GetPlacedFeatureAtSide(side.Mirror(), feature.Road) != nil {
			if tile.GetPlacedFeatureAtSide(side, feature.Road) == nil {
				return false
			}
		}
	}
	return true
}

func (board *board) CanBePlaced(tile elements.PlacedTile) bool {
	if !board.isPositionValid(tile) {
		return false
	}
	if !slices.Contains(board.placeablePositions, tile.Position) {
		return false
	}

	meepleCount := 0
	featuresWithMeeples := map[feature.Type]elements.PlacedFeature{}
	for _, feat := range tile.Features {
		if feat.Meeple.Type != elements.NoneMeeple {
			// Depending on use cases for this method, meeple counting may not be
			// strictly necessary.
			// One example where it's unnecessary is game.PlayTurn() which calls
			// player.PlaceTile() and that already validates meeple count
			// for other reasons.
			meepleCount++
			featuresWithMeeples[feat.FeatureType] = feat
			if meepleCount > 1 {
				return false
			}
		}
	}

	for featureType, feat := range featuresWithMeeples {
		if !canBePlacedFunctions[featureType](board, tile, feat) {
			return false
		}
	}

	return true
}

func (board *board) cityCanBePlaced(tile elements.PlacedTile, feat elements.PlacedFeature) bool {
	return board.cityManager.CanBePlaced(tile, feat)
}

func (board *board) fieldCanBePlaced(tile elements.PlacedTile, feat elements.PlacedFeature) bool {
	// While placing a tile, the player can only put a single meeple on a field.
	// This means we don't have to care about whether our field expands into
	// a different feature on our tile - we will only expand the feature
	// if it has a meeple and that only happens once.
	field := field.New(feat, tile)

	return field.IsFieldValid(board, 1)
}

func (board *board) monasteryCanBePlaced(_ elements.PlacedTile, _ elements.PlacedFeature) bool {
	// meeple can always be placed on a monastery
	return true
}

func (board *board) roadCanBePlaced(checkedTile elements.PlacedTile, checkedRoad elements.PlacedFeature) bool {
	// get the two sides connected by the road which we will use to
	// score roads on the neighbouring tiles (but not the tile itself)
	sides := []side.Side{
		checkedRoad.Sides.GetNthCardinalDirection(0), // 1st side
		checkedRoad.Sides.GetNthCardinalDirection(1), // 2nd side
	}
	for _, checkedRoadSide := range sides {
		neighbourTile, exists := board.GetTileAt(
			checkedTile.Position.Add(position.FromSide(checkedRoadSide)),
		)
		if !exists {
			// no existing tile found on this side of the road
			continue
		}

		// an existing tile found on this side of the road - we need to check,
		// if they have *any* meeple placed
		neighbourRoadSide := checkedRoadSide.Mirror()
		neighbourRoad := neighbourTile.GetPlacedFeatureAtSide(neighbourRoadSide, feature.Road)
		// score function checks the whole road for placed meeples
		// and reports any meeples that would be returned which we can use here
		scoreReport, _ := board.ScoreRoadCompletion(
			neighbourTile,
			neighbourRoad.Feature,
			true,
		)
		if len(scoreReport.ReturnedMeeples) != 0 {
			return false
		}
	}

	return true
}

func (board *board) PlaceTile(tile elements.PlacedTile) (elements.ScoreReport, error) {
	if board.TileCount() == cap(board.tiles) {
		return elements.ScoreReport{}, errors.New(
			"Board's tiles capacity exceeded, logic error?",
		)
	}

	if !board.CanBePlaced(tile) {
		return elements.ScoreReport{}, elements.ErrInvalidPosition
	}

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
	validNewPositions := []position.Position{
		position.New(tile.Position.X()+1, tile.Position.Y()),
		position.New(tile.Position.X()-1, tile.Position.Y()),
		position.New(tile.Position.X(), tile.Position.Y()+1),
		position.New(tile.Position.X(), tile.Position.Y()-1),
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
	scoreReport.Join(board.ScoreRoads(tile, false))
	scoreReport.Join(board.ScoreMonasteries(tile, false))
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
			_, ok := board.GetTileAt(position.New(x, y))
			if ok {
				score++
			}
		}
	}

	if score == 9 || forceScore {
		scoreReport := elements.NewScoreReport()
		scoreReport.ReceivedPoints[monasteryFeature.Meeple.PlayerID] = score
		scoreReport.ReturnedMeeples[monasteryFeature.Meeple.PlayerID] = []elements.MeepleWithPosition{
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
			adjacentTile, ok := board.GetTileAt(position.New(x, y))

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
func (board *board) CheckRoadInDirection(roadSide side.Side, startTile elements.PlacedTile) (bool, int, []elements.MeepleWithPosition, bool, side.Side, position.Position) {
	var meeples = []elements.MeepleWithPosition{}
	var tile = startTile
	var tileExists bool
	var score = 0
	var road *elements.PlacedFeature
	var finished bool
	var pos = startTile.Position
	startRoadSide := roadSide
	// check finished on way
	// do while loop
	for {
		pos = tile.Position.Add(position.FromSide(roadSide))
		tile, tileExists = board.GetTileAt(pos)
		roadSide = roadSide.Mirror()
		// check if tile exists
		if !tileExists {
			// tile does not exist
			break
		}

		// Get road feature
		road = tile.GetPlacedFeatureAtSide(roadSide, feature.Road)

		// check if loop
		if tile.Position == startTile.Position {
			// While the meeples on a start tile are already counted by the caller (ScoreRoadCompletion),
			// it only checks sides that together create that single feature on the tile.
			// If the start tile has two distinct (unconnected) roads, as is the case for crossroads,
			// a finished road may end up connecting them resulting in some sides being unchecked.
			// Therefore, we need to add any meeple present on the other side of the road,
			// if it is not part of the feature we started in.
			if !road.Sides.HasSide(startRoadSide) && road.Meeple.Type != elements.NoneMeeple {
				meeples = append(meeples, elements.NewMeepleWithPosition(
					road.Meeple,
					tile.Position),
				)
			}
			// We're back at the start tile which means we reached a loop or a crossroad.
			// Nothing more to do - the score for the start tile is counted by the caller
			// and the meeples have been counted appropriately by us and the caller already.
			break
		}

		score++

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

	return finished, score, meeples, looped, roadSide, pos
}

/*
Calculates score for road.

returns: ScoreReport, checked sides of the start tile (also including loop)
*/
func (board *board) ScoreRoadCompletion(tile elements.PlacedTile, road feature.Feature, forceScore bool) (elements.ScoreReport, side.Side) {
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
	roadFinishedResult, scoreResult, meeplesResult, loopResult, loopSide, finishedPosLeft := board.CheckRoadInDirection(leftSide, tile)
	score += scoreResult
	roadFinished = roadFinished && roadFinishedResult
	meeples = append(meeples, meeplesResult...)

	// check road in "right" direction
	if !loopResult && rightSide != side.NoSide {
		roadFinishedResult, scoreResult, meeplesResult, _, _, finishedPosRight := board.CheckRoadInDirection(rightSide, tile)
		score += scoreResult
		roadFinished = roadFinished && roadFinishedResult
		meeples = append(meeples, meeplesResult...)

		// Decrement the score to prevent counting the tile twice
		// when its road features (two different ones) are both the start
		// and the end of the road (as is the case for crossroads).
		//
		// Note that this is a different scenario than a literal loop (i.e. an actual circle
		// with no end or beginning) where we don't fall into this `if` branch at all
		// due to `loopResult` being `true`.
		if finishedPosLeft == finishedPosRight {
			score--
		}
	}

	// -------- start counting -------------
	if roadFinished || forceScore {
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
func (board *board) ScoreRoads(placedTile elements.PlacedTile, forceScore bool) elements.ScoreReport {
	scoreReport := elements.NewScoreReport()
	var tile = elements.ToTile(placedTile)
	var roads = tile.Roads()

	var checkedRoadSides side.Side

	for _, road := range roads {
		// check if the side of the tile was not already checked (special test case reference: TestBoardScoreRoadLoopCrossroad)
		if !checkedRoadSides.OverlapsSide(road.Sides) {
			scoreReportTemp, roadSide := board.ScoreRoadCompletion(placedTile, road, forceScore)
			scoreReport.Join(scoreReportTemp)
			checkedRoadSides |= roadSide
		}
		checkedRoadSides |= road.Sides
	}
	return scoreReport
}

func (board *board) ScoreMeeples(final bool) elements.ScoreReport {
	meeplesReport := elements.NewScoreReport()

	// score cities first (because they have their own manager)
	meeplesReport.Join(board.cityManager.ScoreCities(true))

	if final {
		// remove city meeples from board
		for _, returnedMeeples := range meeplesReport.ReturnedMeeples {
			for _, meeple := range returnedMeeples {
				board.RemoveMeeple(meeple.Position)
			}
		}
	}

	// score meeples left on the board (fields, monasteries, roads)
	for _, pTile := range board.Tiles() {
		for _, feat := range pTile.Features {
			miniReport := elements.NewScoreReport()
			if feat.Meeple.PlayerID != 0 {
				switch feat.FeatureType {
				case feature.Road:
					miniReport.Join(board.ScoreRoads(pTile, true))
				case feature.Field:
					field := field.New(feat, pTile)
					field.Expand(board, board.cityManager)
					miniReport.Join(field.GetScoreReport())
				case feature.Monastery:
					miniReport.Join(board.ScoreMonasteries(pTile, true))
				}
			}

			if final {
				// remove meeples from board
				for _, returnedMeeples := range miniReport.ReturnedMeeples {
					for _, meeple := range returnedMeeples {
						board.RemoveMeeple(meeple.Position)
					}
				}
			}
			meeplesReport.Join(miniReport)
		}
	}

	return meeplesReport
}
