from typing import NamedTuple, Self

from . import _bindings
from ._bindings import (  # type: ignore[attr-defined] # no stubs
    binarytiles as _go_binarytiles,
    elements as _go_elements,
    engine as _go_engine,
    game as _go_game,
    position as _go_position,
    tiles as _go_tiles,
)

__all__ = ("GameState", "SerializedGame")

from ._bindings.engine import Slice_elements_PlacedTile, Slice_elements_Player
from .placed_tile import Tile
from .tilesets import TileSet


class Player:
    """
    State of the game the request is made for.

    This class is not meant to be instantiated by users directly
    and should be considered read-only.

    The instances of this class are provided by the `GameEngine` objects.
    """

    __slots__ = ("_go_obj",
                 "_id",
                 "_score",
                 "_meeple_count"
                 )

    def __init__(self, go_obj: _go_elements.Player) -> None:
        self._go_obj = go_obj

    def _unwrap(self) -> _go_elements.Player:
        return self._go_obj

    @property
    def id(self) -> int:
        return self._go_obj.ID()

    @property
    def score(self) -> int:
        return self._go_obj.Score()

    @property
    def meeple_count(self) -> int:
        return self._go_obj.MeepleCount()