import logging
from pathlib import Path

import pytest
from pytest import approx

from carcassonne_engine import GameEngine, SerializedGameWithID

from carcassonne_engine.tilesets import standard_tile_set

log = logging.getLogger(__name__)


def test_basic_params(tmp_path: Path) -> None:
    engine = GameEngine(4, tmp_path)
    tile_set = standard_tile_set()

    serialized_game: SerializedGameWithID = engine.generate_game(tile_set)
    game = serialized_game.game

    assert game.current_tile is not None
    assert game.valid_tile_placements is not None
    assert len(game.tiles) == 72
    assert game.tile_set is not None
    assert len(game.binary_tiles) == 72
    assert game.binary_tiles[1] == 0  # not placed tile




def test_player_params(tmp_path: Path) -> None:
    engine = GameEngine(4, tmp_path)
    tile_set = standard_tile_set()

    serialized_game: SerializedGameWithID = engine.generate_game(tile_set)
    game = serialized_game.game

    assert game.players is not None
    assert game.player_count == 2

    assert game.players[0].score == 0
    assert game.players[1].score == 0

    assert game.players[1].meeple_count == [0, 7]
