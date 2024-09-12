import logging
from pathlib import Path

from carcassonne_engine import GameEngine
from carcassonne_engine.tilesets import standard_tile_set

log = logging.getLogger(__name__)


def test_serialized_properties(tmp_path: Path) -> None:
    engine = GameEngine(4, tmp_path)
    tile_set = standard_tile_set()

    serialized_game_with_id = engine.generate_game(tile_set)
    serialized_game = serialized_game_with_id.game

    assert serialized_game.current_tile is not None
    assert len(serialized_game.tiles) == 72
    assert serialized_game.tile_set is not None
    assert len(serialized_game.binary_tiles) == 72
    assert serialized_game.binary_tiles[1] == 0  # not placed tile is 0


def test_serialized_player_properties(tmp_path: Path) -> None:
    engine = GameEngine(4, tmp_path)
    tile_set = standard_tile_set()

    serialized_game_with_id = engine.generate_game(tile_set)
    serialized_game = serialized_game_with_id.game

    assert serialized_game.current_player_id == 1

    assert len(serialized_game.players) == 2
    assert serialized_game.player_count == 2

    assert serialized_game.players[0].id == 1
    assert serialized_game.players[1].id == 2

    assert serialized_game.players[0].score == 0
    assert serialized_game.players[1].score == 0

    assert serialized_game.players[0].meeple_count == [0, 7]
    assert serialized_game.players[1].meeple_count == [0, 7]
