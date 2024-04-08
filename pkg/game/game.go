package game

import (
	"errors"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
)


type Game struct {
	board         Board
	deck          *stack.Stack[Tile]
	players       []*Player
	currentPlayer int
	logger        *Logger
}

func NewGame(logger *Logger) (*Game, error) {
	deck := stack.New(BaseTileSet)
	return NewGameWithDeck(&deck, logger)
}

func NewGameWithDeck(deck *stack.Stack[Tile], logger *Logger) (*Game, error) {
	game := &Game{
		board:         NewBoard(deck.GetTotalTileCount()),
		deck:          deck,
		players:       []*Player{NewPlayer(0), NewPlayer(1)},
		currentPlayer: 0,
		logger:        logger,
	}

	// All tiles in base game can be placed on the first move but let's just check this
	// in case this isn't true for tiles from all of the expansions.
	game.ensureCurrentTileHasValidPlacement()
	if logger != nil {
		if err := logger.Start(game.deck, len(game.players)); err != nil {
			return nil, err
		}
	}

	return game, nil
}

func (game *Game) GetCurrentTile() (Tile, error) {
	return game.deck.Peek()
}

func (game *Game) CurrentPlayer() *Player {
	return game.players[game.currentPlayer]
}

func (game *Game) PlayerCount() int {
	return len(game.players)
}

func (game *Game) ensureCurrentTileHasValidPlacement() error {
	for {
		// Peek at the tile that will be returned by GetCurrentTile() next time
		// to see, if it can actually placed anywhere.
		nextTile, err := game.deck.Peek()
		if err != nil {
			if errors.Is(err, stack.StackOutOfBoundsError) {
				// out of bounds - not our concern at the end of a turn
				return nil
			}

			// unexpected error?
			return err
		}

		if game.board.HasValidPlacement(nextTile) {
			break
		}

		if _, err := game.deck.Next(); err != nil {
			// We already peeked and checked for out of bounds so that's unexpected...
			return err
		}
	}

	return nil
}

func (game *Game) PlayTurn(placedTile PlacedTile) error {
	// Get tile that the player is supposed to place.
	// This is guaranteed to return a tile that has at least one valid placement
	// or `OutOfBounds` error, if there's no tiles left in the deck and this turn
	// shouldn't be happening.
	currentTile, err := game.GetCurrentTile()
	if err != nil {
		return err
	}

	// TODO: This equality test needs to work with rotations, inner slices, etc.
	// How to do this depends on the final implementation of `Tile` (and `PlacedTile`)
	if currentTile != placedTile.Tile {
		return WrongTile
	}

	player := game.CurrentPlayer()
	defer func() { game.currentPlayer = (game.currentPlayer + 1) % game.PlayerCount() }()

	// In the class diagram, the `scoreReport` would be returned by
	// separate `CheckCompleted()` method but it's been abstracted by PlaceTile instead.
	scoreReport, err := player.PlaceTile(game.board, placedTile)
	if err != nil {
		return err
	}
	if game.logger != nil {
		if err := game.logger.PlaceTile(game.currentPlayer, placedTile); err != nil {
			return err
		}
	}

	// Score features and update meeple counts
	for playerId, receivedPoints := range scoreReport.ReceivedPoints {
		game.players[playerId].score += receivedPoints
	}
	for playerId, returnedMeeples := range scoreReport.ReturnedMeeples {
		game.players[playerId].meepleCount += returnedMeeples
	}

	// Pop from the stack after the move.
	if _, err := game.deck.Next(); err != nil {
		return err
	}

	game.ensureCurrentTileHasValidPlacement()

	return nil
}

func (game *Game) Finalize() ([]uint32, error) {
	playerScores := make([]uint32, len(game.players))

	if _, err := game.GetCurrentTile(); !errors.Is(err, stack.StackOutOfBoundsError) {
		return playerScores, GameIsNotFinished
	}

	for i, player := range game.players {
		playerScores[i] = player.Score()
	}
	if game.logger != nil {
		if err := game.logger.End(playerScores); err != nil {
			return playerScores, err
		}
	}

	return playerScores, nil
}
