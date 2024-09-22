import logging
from pathlib import Path

import pytest
from pytest import approx

from carcassonne_engine import GameEngine, models, tiletemplates
from carcassonne_engine._bindings.elements import MeepleType
from carcassonne_engine._bindings.side import Side
from carcassonne_engine._bindings.feature import Type as FeatureType
from carcassonne_engine.models import Position
from carcassonne_engine.tilesets import TileSet
from carcassonne_engine.requests import (
    GetLegalMovesRequest,
    GetMidGameScoreRequest,
    GetRemainingTilesRequest,
    PlayTurnRequest,
)
from carcassonne_engine.tilesets import TileSet, standard_tile_set
from carcassonne_engine.utils import format_binary_tile_bits
from tests.utils import make_turn, TurnParams

log = logging.getLogger(__name__)


def test_full_game(tmp_path: Path) -> None:
    engine = GameEngine(4, tmp_path)
    tile_set = standard_tile_set()

    game_id, game = engine.generate_ordered_game(tile_set)

    for i, expected_tile in enumerate(tile_set):
        assert game.current_tile == expected_tile
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
        assert legal_moves_resp.moves is not None
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
        assert play_turn_resp.game is not None
        log.info("iteration %s played turn", i)

        game = play_turn_resp.game
        game_id = play_turn_resp.game_id
        log.info("iteration %s end", i)

    assert game.current_tile is None


def test_concurrent_read_requests(tmp_path: Path) -> None:
    engine = GameEngine(4, tmp_path)
    tile_set = standard_tile_set()

    game_id, game = engine.generate_game(tile_set)
    assert game.current_tile is not None

    legal_moves_req = GetLegalMovesRequest(
        base_game_id=game_id, tile_to_place=game.current_tile
    )
    (legal_moves_resp,) = engine.send_get_legal_moves_batch([legal_moves_req])
    assert legal_moves_resp.exception is None
    assert legal_moves_resp.moves is not None

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
        engine.send_play_turn_batch(requests)


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
        assert resp.tile_probabilities is not None

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
        assert resp.moves is not None

    # Monastery with no roads is symmetrical both horizontally and vertically
    # and there's just one valid position (below starting tile, i.e. (0, -1))
    # so there's three legal moves: without meeple and with meeple on each of
    # the 2 features
    assert len(resp.moves) == 3
    for move_with_state in resp.moves:
        assert move_with_state.move.position.x == 0
        assert move_with_state.move.position.y == -1
        assert tile.exact_equals(move_with_state.move.to_tile())


def test_game_engine_send_get_legal_moves_batch_returns_all_legal_rotations(
    tmp_path: Path,
) -> None:
    tile = tiletemplates.monastery_with_single_road()
    tile_set = TileSet.from_tiles(
        [tile],
        # non-default starting tile - limits number of possible positions to one
        starting_tile=tiletemplates.three_city_edges_connected(),
    )

    with GameEngine(1, tmp_path) as engine:
        game_id, game = engine.generate_game(tile_set)

        request = GetLegalMovesRequest(base_game_id=game_id, tile_to_place=tile)
        (resp,) = engine.send_get_legal_moves_batch([request])
        assert resp.exception is None
        assert resp.moves is not None

    # Monastery with single road can only be placed at (0, -1)
    # but in 3 different orientations (only field connected with road is invalid)
    # For each orientation, there are 4 valid meeple placements:
    # - no meeple
    # - meeple on monastery
    # - meeple on field
    # - meeple on road
    expected_move_count = 1 * 3 * 4

    assert len(resp.moves) == expected_move_count
    for move_state in resp.moves:
        assert move_state.move.position.x == 0
        assert move_state.move.position.y == -1

    # starting tile has a field at the bottom so expected values are:
    expected_tiles = [
        # road at the bottom
        tile,
        # road on the left
        models.Tile(tile._go_obj.Rotate(1)),
        # road on the right
        models.Tile(tile._go_obj.Rotate(3)),
    ]
    move_idx = 0
    for expected in expected_tiles:
        for _ in range(4):
            move_state = resp.moves[move_idx]
            assert expected.exact_equals(move_state.move.to_tile())
            move_idx += 1


def test_mid_game_score_request_at_start(tmp_path: Path) -> None:
    engine = GameEngine(4, tmp_path)
    tile_set = standard_tile_set()

    game_id, game = engine.generate_game(tile_set)

    mid_game_score_request = GetMidGameScoreRequest(
        base_game_id=game_id,
    )
    (mid_game_score_response,) = engine.send_get_mid_game_score_batch(
        [mid_game_score_request]
    )

    assert mid_game_score_response.player_scores == {1: 0, 2: 0}

def test_mid_game_score_request_mid_game(tmp_path: Path) -> None:
    engine = GameEngine(4, tmp_path)
    tiles = [
        tiletemplates.four_city_edges_connected_shield(),
        tiletemplates.straight_roads(),
        tiletemplates.straight_roads(),
    ]
    tile_set = TileSet.from_tiles(tiles=tiles,starting_tile= tiletemplates.single_city_edge_straight_roads())

    game_id, game = engine.generate_ordered_game(tile_set)

    # play turns
    turn_params = TurnParams(
        pos=Position(x=0, y=1),
        tile=tiletemplates.four_city_edges_connected_shield(),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Left,
        featureType=FeatureType.City,
    )
    game_id, game = make_turn(engine, game, game_id, turn_params)

    turn_params = TurnParams(
        pos=Position(x=1, y=0),
        tile=tiletemplates.straight_roads(),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Left,
        featureType=FeatureType.Road,
    )
    game_id, game = make_turn(engine, game, game_id, turn_params)

    # check scores
    mid_game_score_request = GetMidGameScoreRequest(
        base_game_id=game_id,
    )
    (mid_game_score_response,) = engine.send_get_mid_game_score_batch(
        [mid_game_score_request]
    )

    assert mid_game_score_response.player_scores == {1: 3, 2: 2}
