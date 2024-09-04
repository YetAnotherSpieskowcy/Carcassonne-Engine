from ._bindings import engine as _go_engine  # type: ignore[attr-defined] # no stubs
from .models import GameState, PlacedTile, SerializedGame, Tile

__all__ = (
    "BaseResponse",
    "PlayTurnRequest",
    "PlayTurnResponse",
    "GetRemainingTilesRequest",
    "GetRemainingTilesResponse",
    "TileProbability",
    "GetLegalMovesRequest",
    "GetLegalMovesResponse",
    "MoveWithState",
)


class BaseResponse:
    """
    Base game engine response class.

    This is an abstract base class and should not be instantiated.

    Some of the objects returned by the `GameEngine` objects subclass this.
    """

    __slots__ = ("_go_obj", "game_id", "exception")

    def __init__(self, go_obj: _go_engine.BaseResponse) -> None:
        self._go_obj = go_obj
        self.game_id = go_obj.GameID()
        self.exception: Exception | None = None
        try:
            go_obj.Err()
        except RuntimeError as exc:
            self.exception = exc


class PlayTurnRequest:
    """
    Game engine request for placing specified tile on the game with specified ID.
    """

    __slots__ = ("_go_obj", "_game_id", "_move")

    def __init__(self, *, game_id: int, move: PlacedTile) -> None:
        self._go_obj = _go_engine.PlayTurnRequest(GameID=game_id, Move=move._unwrap())
        self._game_id = game_id
        self._move = move

    def _unwrap(self) -> _go_engine.PlayTurnRequest:
        return self._go_obj

    @property
    def game_id(self) -> int:
        return self._game_id

    @property
    def move(self) -> PlacedTile:
        return self._move


class PlayTurnResponse(BaseResponse):
    """
    Game engine response for `PlayTurnRequest` instances.

    This class is not meant to be instantiated by users directly
    and should be considered read-only.

    The instances of this class are provided by the `GameEngine` objects.
    """

    __slots__ = ("game", "final_scores")

    def __init__(self, go_obj: _go_engine.PlayTurnResponse) -> None:
        super().__init__(go_obj)
        self.game = SerializedGame(go_obj.Game) if not self.exception else None
        self.final_scores: dict[int, int] | None = None
        if go_obj.FinalScores:
            self.final_scores = {k: v for k, v in go_obj.FinalScores.items()}


class GetRemainingTilesRequest:
    """
    Game engine request for getting the remaining tiles
    in the game with specified ID and state.
    """

    __slots__ = ("_go_obj", "_base_game_id", "_state_to_check")

    def __init__(
            self, *, base_game_id: int, state_to_check: GameState | None = None
    ) -> None:
        if state_to_check is not None:
            self._go_obj = _go_engine.GetRemainingTilesRequest(
                BaseGameID=base_game_id,
                StateToCheck=state_to_check._unwrap(),
            )
        else:
            # gopy bindings don't consider None as Go's nil for pointers
            self._go_obj = _go_engine.GetRemainingTilesRequest(
                BaseGameID=base_game_id,
            )
        self._base_game_id = base_game_id
        self._state_to_check = state_to_check

    def _unwrap(self) -> _go_engine.GetRemainingTilesRequest:
        return self._go_obj

    @property
    def base_game_id(self) -> int:
        return self._base_game_id

    @property
    def state_to_check(self) -> GameState | None:
        return self._state_to_check


class GetRemainingTilesResponse(BaseResponse):
    """
    Game engine response for `GetRemainingTilesRequest` instances.

    This class is not meant to be instantiated by users directly
    and should be considered read-only.

    The instances of this class are provided by the `GameEngine` objects.
    """

    __slots__ = ("tile_probabilities",)

    def __init__(self, go_obj: _go_engine.GetRemainingTilesResponse) -> None:
        super().__init__(go_obj)
        self.tile_probabilities = (
            [
                TileProbability(go_probability)
                for go_probability in go_obj.TileProbabilities
            ]
            if not self.exception
            else None
        )


class TileProbability:
    """
    A tile and its probability to be drawn from the deck.

    This class is not meant to be instantiated by users directly
    and should be considered read-only.

    The instances of this class are provided by the `GameEngine` objects.
    """

    __slots__ = ("tile", "probability")

    def __init__(self, go_obj: _go_engine.TileProbability) -> None:
        self.tile = Tile(go_obj.Tile)
        self.probability = go_obj.Probability


class GetLegalMovesRequest:
    """
    Game engine request for getting the legal moves for the placeable tile
    in the game with specified ID and state.
    """

    __slots__ = ("_go_obj", "_base_game_id", "_state_to_check", "_tile_to_place")

    def __init__(
            self,
            *,
            base_game_id: int,
            state_to_check: GameState | None = None,
            tile_to_place: Tile,
    ) -> None:
        if state_to_check is not None:
            self._go_obj = _go_engine.GetLegalMovesRequest(
                BaseGameID=base_game_id,
                StateToCheck=state_to_check._unwrap(),
                TileToPlace=tile_to_place._unwrap(),
            )
        else:
            # gopy bindings don't consider None as Go's nil for pointers
            self._go_obj = _go_engine.GetLegalMovesRequest(
                BaseGameID=base_game_id,
                TileToPlace=tile_to_place._unwrap(),
            )
        self._base_game_id = base_game_id
        self._state_to_check = state_to_check
        self._tile_to_place = tile_to_place

    def _unwrap(self) -> _go_engine.GetLegalMovesRequest:
        return self._go_obj

    @property
    def base_game_id(self) -> int:
        return self._base_game_id

    @property
    def state_to_check(self) -> GameState | None:
        return self._state_to_check

    @property
    def tile_to_place(self) -> Tile:
        return self._tile_to_place


class GetLegalMovesResponse(BaseResponse):
    """
    Game engine response for `GetLegalMovesRequest` instances.

    This class is not meant to be instantiated by users directly
    and should be considered read-only.

    The instances of this class are provided by the `GameEngine` objects.
    """

    __slots__ = ("moves",)

    def __init__(self, go_obj: _go_engine.GetLegalMovesResponse) -> None:
        super().__init__(go_obj)
        self.moves = (
            [MoveWithState(go_move) for go_move in go_obj.Moves]
            if not self.exception
            else None
        )


class MoveWithState:
    """
    A legal move and the game state it would result in.

    This class is not meant to be instantiated by users directly
    and should be considered read-only.

    The instances of this class are provided by the `GameEngine` objects.
    """

    __slots__ = ("move", "state")

    def __init__(self, go_obj: _go_engine.MoveWithState) -> None:
        self.move = PlacedTile(go_obj.Move)
        self.state = GameState(go_obj.State)


class GetMidGameScoreRequest:
    """
   Game engine request for getting points as if the game just finished
   in the game with specified ID and state.
   """
    __slots__ = ("_go_obj", "_base_game_id", "_state_to_check")

    def __init__(
            self, *, base_game_id: int, state_to_check: GameState | None = None
    ) -> None:
        if state_to_check is not None:
            self._go_obj = _go_engine.GetMidGameScoreRequest(
                BaseGameID=base_game_id,
                StateToCheck=state_to_check._unwrap(),
            )
        else:
            # gopy bindings don't consider None as Go's nil for pointers
            self._go_obj = _go_engine.GetMidGameScoreRequest(
                BaseGameID=base_game_id,
            )
        self._base_game_id = base_game_id
        self._state_to_check = state_to_check

    def _unwrap(self) -> _go_engine.GetMidGameScoreRequest:
        return self._go_obj

    @property
    def base_game_id(self) -> int:
        return self._base_game_id

    @property
    def state_to_check(self) -> GameState | None:
        return self._state_to_check


class GetMidGameScoreResponse(BaseResponse):
    """
    Game engine response for `GetMidGameScoreRequest` instances.

    This class is not meant to be instantiated by users directly
    and should be considered read-only.

    The instances of this class are provided by the `GameEngine` objects.
    """

    __slots__ = ("player_scores",)

    def __init__(self, go_obj: _go_engine.GetMidGameScoreResponse) -> None:
        super().__init__(go_obj)
        self.player_scores = (
            {
                score[0]: score[1]
                for score in go_obj.Scores
            }
            if not self.exception
            else None
        )
