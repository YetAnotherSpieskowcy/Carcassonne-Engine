from typing import NamedTuple, Self

from ._bindings import (  # type: ignore[attr-defined] # no stubs
    elements as _go_elements,
    position as _go_position,
    tiles as _go_tiles,
)

__all__ = ("PlacedTile", "Position", "Tile")


class Tile:
    """
    Representation of a Carcassonne tile with specific feature configuration
    (and orientation).

    This class is not meant to be instantiated by users directly
    and should be considered read-only.

    The instances of this class are provided by the `GameEngine` objects.
    """

    __slots__ = ("_go_obj",)

    def __init__(self, go_obj: _go_tiles.Tile) -> None:
        self._go_obj = go_obj

    def _unwrap(self) -> _go_tiles.Tile:
        return self._go_obj

    def __eq__(self, other: object) -> bool:
        if not isinstance(other, self.__class__):
            return NotImplemented
        return self._go_obj.Equals(other._go_obj)

    def exact_equals(self, other: object) -> bool:
        """
        Unlike ``self == other``, this method also checks that
        the compared tiles are in the same orientation.
        """
        if not isinstance(other, self.__class__):
            return NotImplemented
        return self._go_obj.ExactEquals(other._go_obj)

    def to_bits(self) -> int:
        return _go_elements.BinaryTileFromTile(self._go_obj)


class Position(NamedTuple):
    """Position of a tile on the board."""

    x: int
    y: int

    @classmethod
    def _from_go_obj(cls, go_obj: _go_position.Position) -> Self:
        return cls(go_obj.X(), go_obj.Y())


class PlacedTile:
    """
    Represents a tile (to be) placed on the board.

    This class is not meant to be instantiated by users directly
    and should be considered read-only.

    The instances of this class are provided by the `GameEngine` objects.
    """

    __slots__ = ("_go_obj", "_position")

    def __init__(self, go_obj: _go_elements.PlacedTile) -> None:
        self._go_obj = go_obj
        self._position = Position._from_go_obj(go_obj.Position)

    def _unwrap(self) -> _go_elements.PlacedTile:
        return self._go_obj

    @property
    def position(self) -> Position:
        return self._position

    def to_tile(self) -> Tile:
        return Tile(_go_elements.ToTile(self._go_obj))

    def to_bits(self) -> int:
        return _go_elements.BinaryTileFromPlacedTile(self._go_obj)
