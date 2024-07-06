package engine

import (
	"errors"
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/logger"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

var (
	ErrCommunicatorClosed = errors.New("communicator is closed")
	ErrGameNotFound       = errors.New("game with the given ID was not found")
)

func worker(comm *communicator) {
	defer comm.workGroup.Done()

	for input := range comm.inputBuffer {
		resp := input.request.Execute(input.game)
		input.outputBuffer <- workerOutput{
			requestID: input.requestID,
			resp:      resp,
		}
	}
}

type communicator struct {
	workGroup   sync.WaitGroup
	inputBuffer chan workerInput
	closed      bool
}

func newCommunicator() *communicator {
	return &communicator{
		inputBuffer: make(chan workerInput),
	}
}

func (comm *communicator) Shutdown() {
	comm.closed = true
	close(comm.inputBuffer)
	comm.workGroup.Wait()
}

type SerializedGameWithID struct {
	ID   int
	Game game.SerializedGame
}

type GameEngine struct {
	comm       *communicator
	logDir     string
	games      map[int]*game.Game
	nextGameID int
	nextReqID  int
}

func StartGameEngine(workerCount int, logDir string) (*GameEngine, error) {
	if err := os.MkdirAll(logDir, 0700); err != nil {
		return nil, err
	}
	comm := newCommunicator()
	engine := &GameEngine{
		comm:       comm,
		logDir:     logDir,
		games:      map[int]*game.Game{},
		nextGameID: 1,
		nextReqID:  1,
	}

	for range workerCount {
		comm.workGroup.Add(1)
		go worker(comm)
	}

	return engine, nil
}

func (engine *GameEngine) Shutdown() {
	engine.comm.Shutdown()
}

func (engine *GameEngine) GenerateGame(tileSet tilesets.TileSet) (SerializedGameWithID, error) {
	id := engine.nextGameID
	engine.nextGameID++

	logFile := path.Join(engine.logDir, fmt.Sprintf("%v.jsonl", id))
	logger, err := logger.NewFromFile(logFile)
	if err != nil {
		return SerializedGameWithID{}, err
	}

	g, err := game.NewFromTileSet(tileSet, &logger)
	if err != nil {
		return SerializedGameWithID{}, err
	}

	engine.games[id] = g
	return SerializedGameWithID{id, g.Serialized()}, nil
}

func (engine *GameEngine) SendBatch(requests []Request) []Response {
	outputReqIndexes := map[int]int{}
	responses := make([]Response, len(requests))
	outputBuffer := make(chan workerOutput)

	for i, req := range requests {
		requestID, err := engine.send(outputBuffer, req)
		if err != nil {
			responses[i] = &SyncResponse{baseResponse{gameID: req.GameID(), err: err}}
		}
		outputReqIndexes[requestID] = i
	}

	for output := range outputBuffer {
		i := outputReqIndexes[output.requestID]
		responses[i] = output.resp
	}

	return responses
}

func (engine *GameEngine) send(
	outputBuffer chan workerOutput,
	req Request,
) (int, error) {
	if engine.comm.closed {
		return 0, ErrCommunicatorClosed
	}
	game, ok := engine.games[req.GameID()]
	if !ok {
		return 0, ErrGameNotFound
	}
	delete(engine.games, req.GameID())

	requestID := engine.nextReqID
	engine.nextReqID++
	input := workerInput{
		requestID:    requestID,
		outputBuffer: outputBuffer,
		game:         game,
		request:      req,
	}
	engine.comm.inputBuffer <- input
	return requestID, nil
}
