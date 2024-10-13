from collections.abc import Iterator
from typing import Self

from ._bindings import (  # type: ignore[attr-defined] # no stubs
    engine as _go_engine,
    tilesets as _go_tilesets,
)
from .placed_tile import Tile

__all__ = (
    "TileSet",
    "standard_tile_set",
)


class TileSet:
    """
    A set of tiles that will be put into the deck and the starting tile
    that is automatically placed at (0, 0) position on the board.

    This class is not meant to be instantiated by users directly
    and should be considered read-only.

    If you want to get an instance of it, call the appropriate method
    for a predefined set such as `standard_tile_set()` or use the
    `from_tiles()` factory method.
    """

    __slots__ = ("_go_obj",)

    def __init__(self, go_obj: _go_tilesets.TileSet) -> None:
        self._go_obj = go_obj

    def __len__(self) -> int:
        return len(self._go_obj.Tiles) + 1

    def __iter__(self) -> Iterator[Tile]:
        for go_tile in self._go_obj.Tiles:
            yield Tile(go_tile)

    @classmethod
    def from_tiles(cls, tiles: list[Tile], *, starting_tile: Tile) -> Self:
        go_obj = _go_tilesets.TileSet(
            Tiles=_go_engine.Slice_tiles_Tile(tile._unwrap() for tile in tiles),
            StartingTile=starting_tile._unwrap(),
        )
        return cls(go_obj)

    def _unwrap(self) -> _go_tilesets.TileSet:
        return self._go_obj


def standard_tile_set() -> TileSet:
    return TileSet(_go_tilesets.StandardTileSet())
