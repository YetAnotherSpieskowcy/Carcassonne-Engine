from pathlib import Path

import pytest

from carcassonne_engine import (
    PlayTurnRequest,
    Slice_Ptr_engine_PlayTurnRequest,
    StartGameEngine,
)
from carcassonne_engine.tilesets import StandardTileSet


def test_game_engine_send_batch_receives_correct_responses_after_worker_requests(
    tmp_path: Path
) -> None:
    engine = StartGameEngine(4, str(tmp_path))
    request_count = 100

    tile_set = StandardTileSet()
    games = [engine.GenerateGame(tile_set) for _ in range(request_count)]
    requests = [
        PlayTurnRequest(GameID=game.ID, Move=game.Game.ValidTilePlacements[0])
        for game in games
    ]

    responses = engine.SendPlayTurnBatch(Slice_Ptr_engine_PlayTurnRequest(requests))
    for idx, resp in enumerate(responses):
        resp.Err()  # the binding would raise if this returned an error
        expected = requests[idx].GameID
        actual = resp.GameID()
        assert actual == expected
    engine.Shutdown()


def test_game_engine_send_batch_returns_failure_when_game_id_not_found(
    tmp_path: Path
) -> None:
    engine = StartGameEngine(4, str(tmp_path))
    requests = []

    game = engine.GenerateGame(StandardTileSet())

    successful_req = PlayTurnRequest(
        GameID=game.ID, Move=game.Game.ValidTilePlacements[0]
    )
    requests.append(successful_req)

    wrong_id = game.ID + 2
    failed_req = PlayTurnRequest(GameID=wrong_id, Move=game.Game.ValidTilePlacements[0])
    requests.append(failed_req)

    responses = engine.SendPlayTurnBatch(Slice_Ptr_engine_PlayTurnRequest(requests))
    responses[0].Err()  # the binding would raise if this returned an error
    expected = game.ID
    actual = responses[0].GameID()
    assert expected == actual

    with pytest.raises(RuntimeError):
        responses[1].Err()
    expected = wrong_id
    actual = responses[1].GameID()
    assert expected == actual

    engine.Shutdown()


def test_game_engine_send_batch_returns_failures_when_communicator_closed(
    tmp_path: Path
) -> None:
    engine = StartGameEngine(4, str(tmp_path))
    request_count = 5

    tile_set = StandardTileSet()
    games = [engine.GenerateGame(tile_set) for _ in range(request_count)]
    requests = [
        PlayTurnRequest(GameID=game.ID, Move=game.Game.ValidTilePlacements[0])
        for game in games
    ]
    engine.Shutdown()

    responses = engine.SendPlayTurnBatch(Slice_Ptr_engine_PlayTurnRequest(requests))
    for idx, resp in enumerate(responses):
        with pytest.raises(RuntimeError):
            resp.Err()
        expected = requests[idx].GameID
        actual = resp.GameID()
        assert actual == expected
