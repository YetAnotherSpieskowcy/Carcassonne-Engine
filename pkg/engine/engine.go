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

const inputBufferSize = 10_000_000

func worker(comm *communicator) {
	defer comm.workGroup.Done()

	for input := range comm.inputBuffer {
		resp := input.request.execute(input.game)
		input.outputBuffer <- workerOutput{
			requestID: input.requestID,
			resp:      resp,
		}
		input.waitGroup.Done()
	}
}

type communicator struct {
	workGroup   sync.WaitGroup
	inputBuffer chan workerInput
	closed      bool
}

func newCommunicator() *communicator {
	return &communicator{
		inputBuffer: make(chan workerInput, inputBufferSize),
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

func (engine *GameEngine) SendPlayTurnBatch(concreteRequests []*PlayTurnRequest) []*PlayTurnResponse {
	requests := make([]Request, len(concreteRequests))
	for i := range concreteRequests {
		requests[i] = concreteRequests[i]
	}
	responses := engine.sendBatch(requests)
	concreteResponses := make([]*PlayTurnResponse, len(responses))
	for i := range responses {
		var ok bool
		concreteResponses[i], ok = responses[i].(*PlayTurnResponse)
		if !ok {
			// we can get a SyncResponse here, if the request didn't reach
			// a worker due to failure during prepareWorkerInput
			// this *is* stupid but it's what we have to deal with due to
			// a limitation with auto-generated bindings breaking on
			// a `[]Interface` return:
			// https://github.com/go-python/gopy/issues/357
			concreteResponses[i] = &PlayTurnResponse{
				BaseResponse: responses[i].(*SyncResponse).BaseResponse,
			}
		}
	}
	return concreteResponses
}

func (engine *GameEngine) SendGameTreeBatch(concreteRequests []*GameTreeRequest) []*GameTreeResponse {
	requests := make([]Request, len(concreteRequests))
	for i := range concreteRequests {
		requests[i] = concreteRequests[i]
	}
	responses := engine.sendBatch(requests)
	concreteResponses := make([]*GameTreeResponse, len(responses))
	for i := range responses {
		var ok bool
		concreteResponses[i], ok = responses[i].(*GameTreeResponse)
		if !ok {
			// we can get a SyncResponse here, if the request didn't reach
			// a worker due to failure during prepareWorkerInput
			// this *is* stupid but it's what we have to deal with due to
			// a limitation with auto-generated bindings breaking on
			// a `[]Interface` return:
			// https://github.com/go-python/gopy/issues/357
			concreteResponses[i] = &GameTreeResponse{
				BaseResponse: responses[i].(*SyncResponse).BaseResponse,
			}
		}
	}
	return concreteResponses
}

func (engine *GameEngine) sendBatch(requests []Request) []Response {
	outputReqIndexes := map[int]int{}
	outputReqIndexesLock := sync.RWMutex{}

	responses := make([]Response, len(requests))
	outputBuffer := make(chan workerOutput, len(requests))
	waitGroup := sync.WaitGroup{}
	outputWaitGroup := sync.WaitGroup{}

	outputWaitGroup.Add(1)
	go func() {
		for output := range outputBuffer {
			outputReqIndexesLock.RLock()
			i := outputReqIndexes[output.requestID]
			outputReqIndexesLock.RUnlock()
			responses[i] = output.resp
		}
		outputWaitGroup.Done()
	}()

	for i, req := range requests {
		input, err := engine.prepareWorkerInput(&waitGroup, outputBuffer, req)
		if err != nil {
			responses[i] = &SyncResponse{BaseResponse{gameID: req.gameID(), err: err}}
			continue
		}
		outputReqIndexesLock.Lock()
		outputReqIndexes[input.requestID] = i
		outputReqIndexesLock.Unlock()
		engine.send(input)
	}

	waitGroup.Wait()
	close(outputBuffer)

	outputWaitGroup.Wait()

	return responses
}

func (engine *GameEngine) prepareWorkerInput(
	waitGroup *sync.WaitGroup,
	outputBuffer chan workerOutput,
	req Request,
) (workerInput, error) {
	if engine.comm.closed {
		return workerInput{}, ErrCommunicatorClosed
	}
	gameID := req.gameID()
	game, ok := engine.games[gameID]
	if !ok {
		return workerInput{}, ErrGameNotFound
	}
	delete(engine.games, gameID)

	requestID := engine.nextReqID
	engine.nextReqID++
	return workerInput{
		requestID:    requestID,
		waitGroup:    waitGroup,
		outputBuffer: outputBuffer,
		game:         game,
		request:      req,
	}, nil
}

func (engine *GameEngine) send(input workerInput) {
	input.waitGroup.Add(1)
	engine.comm.inputBuffer <- input
}
