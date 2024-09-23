from ._bindings import elements as _go_elements  # type: ignore[attr-defined] # no stubs

__all__ = ("Player",)


class Player:
    """
    State of the game the request is made for.

    This class is not meant to be instantiated by users directly
    and should be considered read-only.

    The instances of this class are provided by the `GameEngine` objects.
    """

    __slots__ = ("_go_obj", "_id", "_score", "_meeple_count")

    def __init__(self, go_obj: _go_elements.SerializedPlayer) -> None:
        self._go_obj = go_obj

    def _unwrap(self) -> _go_elements.SerializedPlayer:
        return self._go_obj

    @property
    def id(self) -> int:
        return self._go_obj.ID

    @property
    def score(self) -> int:
        return self._go_obj.Score

    @property
    def meeple_count(self) -> list[int]:
        meeples = []
        for amount in self._go_obj.MeepleCounts:
            meeples.append(amount)
        return meeples
