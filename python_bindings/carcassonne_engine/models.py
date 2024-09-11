from typing import NamedTuple, Self

from ._bindings import (  # type: ignore[attr-defined] # no stubs
    binarytiles as _go_binarytiles,
    elements as _go_elements,
    engine as _go_engine,
    game as _go_game,
    position as _go_position,
    tiles as _go_tiles,
)

__all__ = ("GameState", "SerializedGame")

from ._bindings.engine import Slice_elements_PlacedTile
from .placed_tile import Tile
from .player import Player
from .tilesets import TileSet


class GameState:
    """
    State of the game the request is made for.

    This class is not meant to be instantiated by users directly
    and should be considered read-only.

    The instances of this class are provided by the `GameEngine` objects.
    """

    __slots__ = ("_go_obj",)

    def __init__(self, go_obj: _go_engine.GameState) -> None:
        self._go_obj = go_obj

    def _unwrap(self) -> _go_engine.GameState:
        return self._go_obj


class SerializedGame:
    """
    A serialized state of a game.

    This class is not meant to be instantiated by users directly
    and should be considered read-only.

    The instances of this class are provided by the `GameEngine` objects.
    """

    __slots__ = ("_go_obj",
                 "_current_tile",
                 "_valid_tile_placements",
                 "_current_player_id",
                 "_players",
                 "_player_count",
                 "_tiles",
                 "_tile_set",
                 "_binary_tiles",
                 )

    def __init__(self, go_obj: _go_game.SerializedGame) -> None:
        self._go_obj = go_obj
        go_tile_obj = go_obj.CurrentTile
        self._current_tile: Tile | None = None
        if go_tile_obj.Features:
            self._current_tile = Tile(go_tile_obj)
        self._valid_tile_placements = go_obj.ValidTilePlacements
        self._current_player_id = go_obj.CurrentPlayerID
        self._players = [Player(x) for x in go_obj.Players]
        self._player_count = go_obj.PlayerCount
        self._tiles = go_obj.Tiles
        self._tile_set = go_obj.TileSet
        self._binary_tiles = go_obj.BinaryTiles

    @property
    def current_tile(self) -> Tile | None:
        return self._current_tile

    @property
    def valid_tile_placements(self) -> Slice_elements_PlacedTile:
        return self._valid_tile_placements

    @property
    def current_player_id(self) -> int:
        return self._current_player_id

    @property
    def players(self) -> list[Player]:
        return self._players

    @property
    def player_count(self) -> int:
        return self._player_count

    @property
    def tiles(self) -> Slice_elements_PlacedTile:
        return self._tiles

    @property
    def tile_set(self) -> TileSet:
        return self._tile_set

    @property
    def binary_tiles(self) -> list[int]:
        tiles = []
        #for tile in self._binary_tiles:
        #    tiles.append(tile._unwrap())
        return self._binary_tiles

class SerializedGameWithID(NamedTuple):
    """
    A serialized game consisting of its ID and serialized state.

    This class is not meant to be instantiated by users directly
    and should be considered read-only.

    The instances of this class are provided by the `GameEngine` objects.
    """

    id: int
    game: SerializedGame
