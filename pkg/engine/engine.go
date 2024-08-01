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

// Internal struct with communication tools used by the game engine to send requests
// and cooperatively shutdown the workers.
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
	if comm.closed {
		comm.workGroup.Wait()
		return
	}
	comm.closed = true
	close(comm.inputBuffer)
	comm.workGroup.Wait()
}

type SerializedGameWithID struct {
	ID   int
	Game game.SerializedGame
}

// The entry point for Python side of things - the engine keeps track of created games,
// sends requests to its workers and returns back responses received from them.
type GameEngine struct {
	comm          *communicator
	logDir        string
	games         map[int]*game.Game
	nextGameID    int
	nextRequestID int
}

func StartGameEngine(workerCount int, logDir string) (*GameEngine, error) {
	if err := os.MkdirAll(logDir, 0700); err != nil {
		return nil, err
	}
	comm := newCommunicator()
	engine := &GameEngine{
		comm:          comm,
		logDir:        logDir,
		games:         map[int]*game.Game{},
		nextGameID:    1,
		nextRequestID: 1,
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

// Generate a random game from the given tileset.
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

// Due to limitations of Python bindings generator with []interface return type,
// this wraps sendBatch() and limits the return type to only one Response type.
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

// Due to limitations of Python bindings generator with []interface return type,
// this wraps sendBatch() and limits the return type to only one Response type.
func (engine *GameEngine) SendGetRemainingTilesBatch(concreteRequests []*GetRemainingTilesRequest) []*GetRemainingTilesResponse {
	requests := make([]Request, len(concreteRequests))
	for i := range concreteRequests {
		requests[i] = concreteRequests[i]
	}
	responses := engine.sendBatch(requests)
	concreteResponses := make([]*GetRemainingTilesResponse, len(responses))
	for i := range responses {
		var ok bool
		concreteResponses[i], ok = responses[i].(*GetRemainingTilesResponse)
		if !ok {
			// we can get a SyncResponse here, if the request didn't reach
			// a worker due to failure during prepareWorkerInput
			// this *is* stupid but it's what we have to deal with due to
			// a limitation with auto-generated bindings breaking on
			// a `[]Interface` return:
			// https://github.com/go-python/gopy/issues/357
			concreteResponses[i] = &GetRemainingTilesResponse{
				BaseResponse: responses[i].(*SyncResponse).BaseResponse,
			}
		}
	}
	return concreteResponses
}

// Due to limitations of Python bindings generator with []interface return type,
// this wraps sendBatch() and limits the return type to only one Response type.
func (engine *GameEngine) SendGetLegalMovesBatch(concreteRequests []*GetLegalMovesRequest) []*GetLegalMovesResponse {
	requests := make([]Request, len(concreteRequests))
	for i := range concreteRequests {
		requests[i] = concreteRequests[i]
	}
	responses := engine.sendBatch(requests)
	concreteResponses := make([]*GetLegalMovesResponse, len(responses))
	for i := range responses {
		var ok bool
		concreteResponses[i], ok = responses[i].(*GetLegalMovesResponse)
		if !ok {
			// we can get a SyncResponse here, if the request didn't reach
			// a worker due to failure during prepareWorkerInput
			// this *is* stupid but it's what we have to deal with due to
			// a limitation with auto-generated bindings breaking on
			// a `[]Interface` return:
			// https://github.com/go-python/gopy/issues/357
			concreteResponses[i] = &GetLegalMovesResponse{
				BaseResponse: responses[i].(*SyncResponse).BaseResponse,
			}
		}
	}
	return concreteResponses
}

// API for handling the sent requests using background workers.
// The order and types of returned responses correspond to the requests slice.
//
// If a request fails before reaching a worker, a `SyncResponse` type will be
// be returned in place of the worker response.
//
// Concurrent calls to this function can be made but no more than
// one request for a *single* game (ID) can be performed at the same time
// to avoid concurrent writes by the workers on different threads.
// You will receive `ErrGameNotFound` error, if you try doing so.
func (engine *GameEngine) sendBatch(requests []Request) []Response {
	outputReqIndexes := map[int]int{}
	outputReqIndexesLock := sync.RWMutex{}

	games := map[int]*game.Game{}
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
		games[req.gameID()] = input.game
		outputReqIndexesLock.Lock()
		outputReqIndexes[input.requestID] = i
		outputReqIndexesLock.Unlock()
		engine.send(input)
	}

	// Wait for all workers request to finish running on the workers.
	waitGroup.Wait()
	// Close the output buffer to let the response-handling goroutine know
	// that there will be no more requests and wait for it to finish.
	close(outputBuffer)
	outputWaitGroup.Wait()

	// prepareWorkerInput removes games from the map to avoid concurrent
	// requests to the same game. Now it's time to add them back.
	for gameID, game := range games {
		engine.games[gameID] = game
	}

	return responses
}

// Prepare the input that will be sent through engine's input buffer
// to a worker. This mutates engine's games map (deleting the game from
// the request) and updates the next free request ID.
// If this function did not return an error, the caller needs to readd
// the game from the request after the worker is done with it.
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
	// Delete the game to avoid concurrent requests to it
	// This needs to be added back by the caller when the requests finishes.
	delete(engine.games, gameID)

	requestID := engine.nextRequestID
	engine.nextRequestID++
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
