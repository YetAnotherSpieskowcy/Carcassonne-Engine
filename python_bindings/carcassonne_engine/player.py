from ._bindings import elements as _go_elements  # type: ignore[attr-defined] # no stubs

__all__ = ("SerializedPlayer",)


class SerializedPlayer:
    """
    State of the game the request is made for.

    This class is not meant to be instantiated by users directly
    and should be considered read-only.

    The instances of this class are provided by the `GameEngine` objects.
    """

    __slots__ = ("_id", "_score", "_meeple_counts")

    def __init__(self, go_obj: _go_elements.SerializedPlayer) -> None:
        self._id = go_obj.ID
        self._score = go_obj.Score
        self._meeple_counts = list(go_obj.MeepleCounts)

    @property
    def id(self) -> int:
        return self._id

    @property
    def score(self) -> int:
        return self._score

    @property
    def meeple_counts(self) -> list[int]:
        return self._meeple_counts
