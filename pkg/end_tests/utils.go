package end_tests

import (
	"testing"

	gameMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

func MakeTurn(game *gameMod.Game, t *testing.T, tilePosition position.Position, rotations uint, meeple elements.MeepleType, featureSide side.Side, featureType feature.Type) {
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}

	var player = game.CurrentPlayer()

	ptile := elements.ToPlacedTile(tile.Rotate(rotations))
	ptile.Position = tilePosition
	if meeple != elements.NoneMeeple {
		ptile.GetPlacedFeatureAtSide(featureSide, featureType).Meeple = elements.Meeple{
			MeepleType: meeple,
			PlayerID:   player.ID(),
		}
	}

	err = game.PlayTurn(ptile)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func CheckMeeplesAndScore(game *gameMod.Game, t *testing.T, playerScores []uint32, playerMeeples []uint8, turnNumber uint) {

	for i := range len(playerScores) {
		// load player
		var player = game.GetPlayerByID(elements.ID(i + 1))

		// check meeples
		if player.MeepleCount(elements.NormalMeeple) != playerMeeples[i] {
			t.Fatalf("Turn %d: meeples count does not match for player %d. Expected: %d  Got: %d", turnNumber, i+1, playerMeeples[i], player.MeepleCount(elements.NormalMeeple))
		}

		// check points
		if player.Score() != playerScores[i] {
			t.Fatalf("Turn %d: Player %d received wrong amount of points! Expected: %d  Got: %d ", turnNumber, i+1, playerScores[i], player.Score())
		}
	}
}

func VerifyMeepleExistence(t *testing.T, game *gameMod.Game, pos position.Position, side side.Side, featureType feature.Type, meepleExist bool, turnNumber uint) {
	board := game.GetBoard() // nolint: typecheck
	placedTile, tileExists := board.GetTileAt(pos)
	if !tileExists {
		t.Fatalf("Turn %d: There is no tile on desired positon: %#v", turnNumber, pos)
	}
	placedFeature := placedTile.GetPlacedFeatureAtSide(side, featureType)
	if meepleExist {
		if placedFeature.MeepleType != elements.NormalMeeple {
			t.Fatalf("Turn %d: Missing meeple on a tile!", turnNumber)
		}
	} else {
		if placedFeature.MeepleType != elements.NoneMeeple {
			t.Fatalf("Turn %d: Meeple hasn't been removed!", turnNumber)
		}
	}

}
