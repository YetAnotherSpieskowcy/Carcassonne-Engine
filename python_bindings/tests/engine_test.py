import logging
from pathlib import Path

import pytest
from pytest import approx

from carcassonne_engine import GameEngine
from carcassonne_engine import models, tiletemplates
from carcassonne_engine.requests import (
    GetLegalMovesRequest,
    GetRemainingTilesRequest,
    PlayTurnRequest,
)
from carcassonne_engine.tilesets import standard_tile_set, TileSet
from carcassonne_engine.utils import format_binary_tile_bits

log = logging.getLogger(__name__)


@pytest.mark.skip(
    reason=(
        "fails due to unfinished GetLegalMovesFor() method"
        " - this causes panic taking whole interpreter with it"
    ),
)
def test_full_game(tmp_path: Path) -> None:
    engine = GameEngine(4, tmp_path)
    tile_set = standard_tile_set()

    game_id, game = engine.generate_game(tile_set)

    for i in range(len(tile_set) - 1):
        log.info(
            "iteration %s start: %s",
            i,
            format_binary_tile_bits(game.current_tile.to_bits()),
        )
        legal_moves_req = GetLegalMovesRequest(
            base_game_id=game_id, tile_to_place=game.current_tile
        )
        (legal_moves_resp,) = engine.send_get_legal_moves_batch([legal_moves_req])
        assert legal_moves_resp.exception is None
        log.info("iteration %s got moves", i)

        move = legal_moves_resp.moves[0].move
        log.info(
            "iteration %s selecting move: %s at position %s",
            i,
            format_binary_tile_bits(move.to_bits()),
            move.position,
        )
        play_turn_req = PlayTurnRequest(game_id=game_id, move=move)
        (play_turn_resp,) = engine.send_play_turn_batch([play_turn_req])
        assert play_turn_resp.exception is None
        log.info("iteration %s played turn", i)

        game = play_turn_resp.game
        game_id = play_turn_resp.game_id
        log.info("iteration %s end", i)

    assert game.current_tile is None


def test_concurrent_read_requests(tmp_path: Path) -> None:
    engine = GameEngine(4, tmp_path)
    tile_set = standard_tile_set()

    game_id, game = engine.generate_game(tile_set)

    legal_moves_req = GetLegalMovesRequest(
        base_game_id=game_id, tile_to_place=game.current_tile
    )
    (legal_moves_resp,) = engine.send_get_legal_moves_batch([legal_moves_req])
    assert legal_moves_resp.exception is None

    requests = []
    for move_with_state in legal_moves_resp.moves:
        requests.append(
            GetRemainingTilesRequest(
                base_game_id=game_id, state_to_check=move_with_state.state
            )
        )
    responses = engine.send_get_remaining_tiles_batch(requests)

    for resp in responses:
        assert resp.exception is None


def test_game_engine_send_batch_receives_correct_responses_after_worker_requests(
    tmp_path: Path,
) -> None:
    engine = GameEngine(4, tmp_path)
    request_count = 100

    tile_set = standard_tile_set()
    games = [engine.generate_game(tile_set) for _ in range(request_count)]
    requests = [
        PlayTurnRequest(
            game_id=generated_game.id,
            move=models.PlacedTile(generated_game.game._go_obj.ValidTilePlacements[0]),
        )
        for generated_game in games
    ]

    responses = engine.send_play_turn_batch(requests)
    for idx, resp in enumerate(responses):
        assert resp.exception is None
        expected = requests[idx].game_id
        actual = resp.game_id
        assert actual == expected
    engine.close()


def test_game_engine_send_batch_returns_failure_when_game_id_not_found(
    tmp_path: Path,
) -> None:
    engine = GameEngine(4, tmp_path)
    requests = []

    game_id, game = engine.generate_game(standard_tile_set())

    successful_req = PlayTurnRequest(
        game_id=game_id,
        move=models.PlacedTile(game._go_obj.ValidTilePlacements[0]),
    )
    requests.append(successful_req)

    wrong_id = game_id + 2
    failed_req = PlayTurnRequest(
        game_id=wrong_id,
        move=models.PlacedTile(game._go_obj.ValidTilePlacements[0]),
    )
    requests.append(failed_req)

    responses = engine.send_play_turn_batch(requests)
    assert responses[0].exception is None
    expected = game_id
    actual = responses[0].game_id
    assert expected == actual

    assert isinstance(responses[1].exception, Exception)
    expected = wrong_id
    actual = responses[1].game_id
    assert expected == actual

    engine.close()


def test_game_engine_send_batch_raises_when_communicator_closed(
    tmp_path: Path,
) -> None:
    engine = GameEngine(4, tmp_path)
    request_count = 5

    tile_set = standard_tile_set()
    games = [engine.generate_game(tile_set) for _ in range(request_count)]
    requests = [
        PlayTurnRequest(
            game_id=generated_game.id,
            move=models.PlacedTile(generated_game.game._go_obj.ValidTilePlacements[0]),
        )
        for generated_game in games
    ]
    engine.close()

    with pytest.raises(RuntimeError):
        responses = engine.send_play_turn_batch(requests)


def test_game_engine_send_get_remaining_tiles_batch_returns_remaining_tiles(
    tmp_path: Path,
) -> None:
    t1 = tiletemplates.monastery_with_single_road()
    t2 = tiletemplates.roads_turn()
    tiles = [t1, t2, t1]
    tile_set = TileSet.from_tiles(
        tiles,
        starting_tile=tiletemplates.single_city_edge_straight_roads(),
    )

    with GameEngine(1, tmp_path) as engine:
        game_id, game = engine.generate_game(tile_set)

        request = GetRemainingTilesRequest(base_game_id=game_id)
        (resp,) = engine.send_get_remaining_tiles_batch([request])
        assert resp.exception is None

    for tile_prob in resp.tile_probabilities:
        for tile in tiles:
            if tile_prob.tile == tile:
                assert tile_prob.probability == approx(tiles.count(tile) / len(tiles))
                break
        else:
            assert (
                False
            ), f"could not find a tile matching the move probability {tile_prob}"


def test_game_engine_send_get_legal_moves_batch_returns_no_duplicates(
    tmp_path: Path,
) -> None:
    tile = tiletemplates.monastery_without_roads()
    tile_set = TileSet.from_tiles(
        [tile],
        starting_tile=tiletemplates.single_city_edge_straight_roads(),
    )

    with GameEngine(1, tmp_path) as engine:
        game_id, game = engine.generate_game(tile_set)

        request = GetLegalMovesRequest(base_game_id=game_id, tile_to_place=tile)
        (resp,) = engine.send_get_legal_moves_batch([request])
        assert resp.exception is None

    # Monastery with no roads is symmetrical both horizontally and vertically
    assert len(resp.moves) == 1
    assert resp.moves[0].move.position.x == 0
    assert resp.moves[0].move.position.y == -1
    assert tile.exact_equals(resp.moves[0].move.to_tile())


def test_game_engine_send_get_legal_moves_batch_returns_all_legal_rotations(
    tmp_path: Path,
) -> None:
    tile = tiletemplates.monastery_with_single_road()
    tile_set = TileSet.from_tiles(
        [tile],
        starting_tile=tiletemplates.single_city_edge_straight_roads(),
    )

    with GameEngine(1, tmp_path) as engine:
        game_id, game = engine.generate_game(tile_set)

        request = GetLegalMovesRequest(base_game_id=game_id, tile_to_place=tile)
        (resp,) = engine.send_get_legal_moves_batch([request])
        assert resp.exception is None

    # Monastery with single road can only be placed at (0, -1)
    # but in 3 different rotations (only field connected with road is invalid)
    assert len(resp.moves) == 3
    for move_state in resp.moves:
        assert move_state.move.position.x == 0
        assert move_state.move.position.y == -1

    # starting tile has a field at the bottom so expected values are:
    expected = [
        # road at the bottom
        tile,
        # road on the left
        models.Tile(tile._go_obj.Rotate(1)),
        # road on the right
        models.Tile(tile._go_obj.Rotate(3)),
    ]
    for idx, move_state in enumerate(resp.moves):
        assert expected[idx].exact_equals(move_state.move.to_tile())
