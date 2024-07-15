package end_tests

import (
	"testing"

	gameMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

func MakeTurn(game *gameMod.Game, t *testing.T, tilePosition elements.Position, rotations uint, meeple elements.MeepleType, featureSide side.Side, featureType feature.Type) {
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

func CheckMeeplesAndScore(game *gameMod.Game, t *testing.T, playerScores []uint32, playerMeeples []uint8) {

	for i := range len(playerScores) {
		// load player
		var player = game.GetPlayerByID(elements.ID(i + 1))

		// check meeples
		if player.MeepleCount(elements.NormalMeeple) != playerMeeples[i] {
			t.Fatalf("meeples count does not match for player %d. Expected: %d  Got: %d", i+1, playerMeeples[i], player.MeepleCount(elements.NormalMeeple))
		}

		// check points
		if player.Score() != playerScores[i] {
			t.Fatalf("Player %d received wrong amount of points! Expected: %d  Got: %d ", i+1, playerScores[i], player.Score())
		}
	}
}
