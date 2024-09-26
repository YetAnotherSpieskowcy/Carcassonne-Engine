package game

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

type MeepleParams struct {
	MeepleType  elements.MeepleType
	FeatureSide side.Side
	FeatureType feature.Type
}

func NoneMeeple() MeepleParams {
	return MeepleParams{
		MeepleType:  elements.NoneMeeple,
		FeatureSide: side.NoSide,
		FeatureType: feature.NoneType,
	}
}

func MakeTurn(gamee *Game, t *testing.T, tilePosition position.Position, meepleParams MeepleParams) {
	tile, err := gamee.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}

	var player = gamee.CurrentPlayer()

	ptile := elements.ToPlacedTile(tile)
	ptile.Position = tilePosition
	if meepleParams.MeepleType != elements.NoneMeeple {
		ptile.GetPlacedFeatureAtSide(meepleParams.FeatureSide, meepleParams.FeatureType).Meeple = elements.Meeple{
			Type:     meepleParams.MeepleType,
			PlayerID: player.ID(),
		}
	}

	err = gamee.PlayTurn(ptile)

	if err != nil {
		t.Fatal(err.Error())
	}
}

func MakeTurnValidCheck(gamee *Game, t *testing.T, tilePosition position.Position, meepleParams MeepleParams, correctMove bool, turnNumber uint) {
	tile, err := gamee.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}

	var player = gamee.CurrentPlayer()

	ptile := elements.ToPlacedTile(tile)
	ptile.Position = tilePosition
	if meepleParams.MeepleType != elements.NoneMeeple {
		ptile.GetPlacedFeatureAtSide(meepleParams.FeatureSide, meepleParams.FeatureType).Meeple = elements.Meeple{
			Type:     meepleParams.MeepleType,
			PlayerID: player.ID(),
		}
	}

	err = gamee.PlayTurn(ptile)

	if err != nil && correctMove {
		t.Fatal(err.Error())
	} else if err == nil && !correctMove {
		t.Fatalf("Turn %d: Wrongly placed meeple wasn't detected by engine!", turnNumber)
	}
}

func CheckMeeplesAndScore(gamee *Game, t *testing.T, playerScores []uint32, playerMeeples []uint8, turnNumber uint) {

	for i := range len(playerScores) {
		// load player
		var player = gamee.GetPlayerByID(elements.ID(i + 1))

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

func VerifyMeepleExistence(t *testing.T, gamee *Game, pos position.Position, s side.Side, featureType feature.Type, meepleExist bool, turnNumber uint) {
	board := gamee.GetBoard()
	placedTile, tileExists := board.GetTileAt(pos)
	if !tileExists {
		t.Fatalf("Turn %d: There is no tile on desired positon: %#v", turnNumber, pos)
	}
	placedFeature := placedTile.GetPlacedFeatureAtSide(s, featureType)
	if meepleExist {
		if placedFeature.Meeple.Type != elements.NormalMeeple {
			t.Fatalf("Turn %d: Missing meeple on a tile!", turnNumber)
		}
	} else {
		if placedFeature.Meeple.Type != elements.NoneMeeple {
			t.Fatalf("Turn %d: Meeple hasn't been removed!", turnNumber)
		}
	}

}
