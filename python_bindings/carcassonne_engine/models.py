from ._bindings import (
    elements as _go_elements,
    engine as _go_engine,
    position as _go_position,
    tiles as _go_tiles,
)

__all__ = ("GameState", "PlacedTile", "SerializedGame", "Tile")


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
    A serialized game consisting of its ID and serialized state.

    This class is not meant to be instantiated by users directly
    and should be considered read-only.

    The instances of this class are provided by the `GameEngine` objects.
    """

    __slots__ = ("_go_obj",)

    def __init__(self, go_obj: _go_engine.SerializedGameWithID) -> None:
        self._go_obj = go_obj

    @property
    def id(self) -> int:
        return self._go_obj.ID


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


class Position:
    """
    Position of a tile on the board.
    """

    __slots__ = ("_go_obj", "_x", "_y")

    def __init__(self, go_obj: _go_position.Position) -> None:
        self._go_obj = go_obj
        self._x = go_obj.X()
        self._y = go_obj.Y()

    @property
    def x(self) -> int:
        return self._x

    @property
    def y(self) -> int:
        return self._y


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
        self._position = Position(go_obj.Position)

    def _unwrap(self) -> _go_elements.PlacedTile:
        return self._go_obj

    @property
    def position(self) -> Position:
        return self._position

    def to_tile(self) -> Tile:
        return Tile(_go_elements.ToTile(self._go_obj))
