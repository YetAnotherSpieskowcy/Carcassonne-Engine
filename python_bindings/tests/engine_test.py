from pathlib import Path

import pytest
from pytest import approx

from carcassonne_engine import (
    GetRemainingTilesRequest,
    GetLegalMovesRequest,
    PlayTurnRequest,
    Slice_tiles_Tile,
    Slice_Ptr_engine_PlayTurnRequest,
    Slice_Ptr_engine_GetLegalMovesRequest,
    Slice_Ptr_engine_GetRemainingTilesRequest,
    StartGameEngine,
    tiletemplates,
)
from carcassonne_engine.tilesets import StandardTileSet, TileSet


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


def test_game_engine_send_get_remaining_tiles_batch_returns_remaining_tiles(
    tmp_path: Path
) -> None:
    t1 = tiletemplates.MonasteryWithSingleRoad()
    t2 = tiletemplates.RoadsTurn()
    tiles = [t1, t2, t1]
    tile_set = TileSet(
        StartingTile=tiletemplates.SingleCityEdgeStraightRoads(),
        Tiles=Slice_tiles_Tile(tiles),
    )

    engine = StartGameEngine(1, str(tmp_path))
    game = engine.GenerateGame(tile_set)

    request = GetRemainingTilesRequest(BaseGameID=game.ID)
    resp, = engine.SendGetRemainingTilesBatch(
        Slice_Ptr_engine_GetRemainingTilesRequest([request])
    )
    resp.Err()  # the binding would raise if this returned an error
    engine.Shutdown()

    for tile_prob in resp.TileProbabilities:
        for tile in tiles:
            if tile_prob.Tile.Equals(tile):
                assert tile_prob.Probability == approx(tiles.count(tile) / len(tiles))
                break
        else:
            assert False, (
                f"could not find a tile matching the move probability {tile_prob}"
            )


def test_game_engine_send_get_legal_moves_batch_returns_no_duplicates(
    tmp_path: Path
) -> None:
    tile = tiletemplates.MonasteryWithoutRoads()
    tile_set = TileSet(
        StartingTile=tiletemplates.SingleCityEdgeStraightRoads(),
        Tiles=Slice_tiles_Tile([tile]),
    )

    engine = StartGameEngine(1, str(tmp_path))
    game = engine.GenerateGame(tile_set)

    request = GetLegalMovesRequest(BaseGameID=game.ID, TileToPlace=tile)
    resp, = engine.SendGetLegalMovesBatch(
        Slice_Ptr_engine_GetLegalMovesRequest([request])
    )
    resp.Err()  # the binding would raise if this returned an error
    engine.Shutdown()

    # Monastery with no roads is symmetrical both horizontally and vertically
    assert len(resp.Moves) == 1
    assert resp.Moves[0].Move.Position.X() == 0
    assert resp.Moves[0].Move.Position.Y() == -1
    for a, b in zip(tile.Features, resp.Moves[0].Move.Features, strict=True):
        assert a.Equals(b)


def test_game_engine_send_get_legal_moves_batch_returns_all_legal_rotations(
    tmp_path: Path
) -> None:
    tile = tiletemplates.MonasteryWithSingleRoad()
    tile_set = TileSet(
        StartingTile=tiletemplates.SingleCityEdgeStraightRoads(),
        Tiles=Slice_tiles_Tile([tile]),
    )

    engine = StartGameEngine(1, str(tmp_path))
    game = engine.GenerateGame(tile_set)

    request = GetLegalMovesRequest(BaseGameID=game.ID, TileToPlace=tile)
    resp, = engine.SendGetLegalMovesBatch(
        Slice_Ptr_engine_GetLegalMovesRequest([request])
    )
    resp.Err()  # the binding would raise if this returned an error
    engine.Shutdown()

    # Monastery with single road can only be placed at (0, -1)
    # but in 3 different rotations (only field connected with road is invalid)
    assert len(resp.Moves) == 3
    for move_state in resp.Moves:
        assert move_state.Move.Position.X() == 0
        assert move_state.Move.Position.Y() == -1

    # starting tile has a field at the bottom so expected values are:
    expected = [
        # road at the bottom
        tile,
        # road on the left
        tile.Rotate(1),
        # road on the right
        tile.Rotate(3),
    ]
    for idx, move_state in enumerate(resp.Moves):
        assert len(move_state.Move.Features) == len(tile.Features)
        for a, b in zip(expected[idx].Features, move_state.Move.Features, strict=True):
            assert a.Equals(b)
