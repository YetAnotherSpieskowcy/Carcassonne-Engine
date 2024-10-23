package test

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

// equivalent to Game type from game package
type Game interface {
	GetCurrentTile() (tiles.Tile, error)
	CurrentPlayer() elements.Player
	PlayTurn(move elements.PlacedTile) error
	GetPlayerByID(playerID elements.ID) elements.Player
	GetBoard() elements.Board
	GetMidGameScore() elements.ScoreReport
}

type T interface {
	Fatal(args ...any)
	Fatalf(format string, args ...any)
}

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

type MakeTurn struct {
	Game         Game
	TestingT     T
	Position     position.Position
	MeepleParams MeepleParams
	TurnNumber   uint
	WrongTurn    bool
}

func (turn MakeTurn) Run() {
	tile, err := turn.Game.GetCurrentTile()
	if err != nil {
		turn.TestingT.Fatal(err.Error())
	}

	var player = turn.Game.CurrentPlayer()

	ptile := elements.ToPlacedTile(tile)
	ptile.Position = turn.Position
	if turn.MeepleParams.MeepleType != elements.NoneMeeple {
		ptile.GetPlacedFeatureAtSide(turn.MeepleParams.FeatureSide, turn.MeepleParams.FeatureType).Meeple = elements.Meeple{
			Type:     turn.MeepleParams.MeepleType,
			PlayerID: player.ID(),
		}
	}

	err = turn.Game.PlayTurn(ptile)

	if turn.WrongTurn {
		if err == nil {
			turn.TestingT.Fatalf("Turn %d: Wrongly placed tile wasn't detected by engine!", turn.TurnNumber)
		}
	} else {
		if err != nil {
			turn.TestingT.Fatal(err.Error())
		}
	}

}

type CheckMeeplesAndScore struct {
	Game          Game
	TestingT      T
	PlayerScores  []uint32
	PlayerMeeples []uint8
	TurnNumber    uint
}

func (turn CheckMeeplesAndScore) Run() {
	for i := range len(turn.PlayerScores) {
		// load player
		var player = turn.Game.GetPlayerByID(elements.ID(i + 1))

		// check meeples
		if player.MeepleCount(elements.NormalMeeple) != turn.PlayerMeeples[i] {
			turn.TestingT.Fatalf("Turn %d: meeples count does not match for player %d. Expected: %d  Got: %d", turn.TurnNumber, i+1, turn.PlayerMeeples[i], player.MeepleCount(elements.NormalMeeple))
		}

		// check points
		if player.Score() != turn.PlayerScores[i] {
			turn.TestingT.Fatalf("Turn %d: Player %d received wrong amount of points! Expected: %d  Got: %d ", turn.TurnNumber, i+1, turn.PlayerScores[i], player.Score())
		}
	}
}

type VerifyMeepleExistence struct {
	Game         Game
	TestingT     T
	Position     position.Position
	Side         side.Side
	FeatureType  feature.Type
	MeepleExists bool
	TurnNumber   uint
}

func (turn VerifyMeepleExistence) Run() {
	board := turn.Game.GetBoard()
	placedTile, tileExists := board.GetTileAt(turn.Position)
	if !tileExists {
		turn.TestingT.Fatalf("Turn %d: There is no tile on desired positon: %#v", turn.TurnNumber, turn.Position)
	}
	placedFeature := placedTile.GetPlacedFeatureAtSide(turn.Side, turn.FeatureType)
	if turn.MeepleExists {
		if placedFeature.Meeple.Type != elements.NormalMeeple {
			turn.TestingT.Fatalf("Turn %d: Missing meeple on a tile!", turn.TurnNumber)
		}
	} else {
		if placedFeature.Meeple.Type != elements.NoneMeeple {
			turn.TestingT.Fatalf("Turn %d: Meeple hasn't been removed!", turn.TurnNumber)
		}
	}
}

type CheckMidGameScore struct {
	Game         Game
	TestingT     T
	PlayerScores []uint32
	TurnNumber   uint
}

func (turn CheckMidGameScore) Run() {

	mid_score := turn.Game.GetMidGameScore()

	for id, score := range mid_score.ReceivedPoints {
		// check points
		if score != turn.PlayerScores[int(id)-1] {
			turn.TestingT.Fatalf("Turn %d: Player %d received wrong amount of points! Expected: %d  Got: %d ", turn.TurnNumber, id, turn.PlayerScores[int(id)-1], score)
		}
	}
}
