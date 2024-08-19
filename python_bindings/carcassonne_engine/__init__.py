import os
import warnings
from types import TracebackType
from typing import Self

from . import requests
from ._bindings import engine as _go_engine
from .models import SerializedGame
from .tilesets import TileSet

__all__ = ("GameEngine",)


class GameEngine:
    __slots__ = ("_go_game_engine",)

    def __init__(self, worker_count: int, log_dir: os.PathLike) -> None:
        try:
            self._go_game_engine = _go_engine.StartGameEngine(
                worker_count, os.fspath(log_dir)
            )
        except RuntimeError as exc:
            # it's hard to make a more exact mapping of exceptions when gopy
            # just flattens them into RuntimeErrors but we at least know
            # that only Go's os.MkdirAll() can return an error here
            # so it's definitely some kind of OSError
            raise OSError(str(exc)) from None

    def __del__(self) -> None:
        if self.closed:
            return
        warnings.warn(f"Unclosed game engine {self!r}", ResourceWarning, source=self)
        self.close()

    @property
    def closed(self) -> bool:
        return self._go_game_engine.IsClosed()

    def close(self) -> None:
        self._go_game_engine.Close()

    def _check_closed(self) -> None:
        if self.closed:
            raise RuntimeError("The game engine has already been closed.")

    def __enter__(self) -> Self:
        self._check_closed()
        return self

    def __exit__(
        self,
        exc_type: type[BaseException] | None,
        value: BaseException | None,
        tb: TracebackType | None,
    ) -> None:
        self.close()

    def generate_game(self, tileset: TileSet) -> SerializedGame:
        self._check_closed()
        try:
            go_obj = self._go_game_engine.GenerateGame(tileset._unwrap())
        except RuntimeError as exc:
            # We want to raise IOError (or its subclasses) or engine-specific
            # exceptions depending on what error is returned here but since gopy
            # flattens these, let's just raise generic Exception to not bind ourselves
            # to a tighter API contract.
            # TODO: map exceptions once we migrate from gopy to manually-written bindings
            raise Exception(str(exc)) from None
        return SerializedGame(go_obj)

    def clone_game(self, game_id: int, count: int) -> list[int]:
        self._check_closed()
        try:
            ret = self._go_game_engine.CloneGame(game_id, count)
        except RuntimeError as exc:
            # We want to raise IOError (or its subclasses) or engine-specific
            # exceptions depending on what error is returned here but since gopy
            # flattens these, let's just raise generic Exception to not bind ourselves
            # to a tighter API contract.
            # TODO: map exceptions once we migrate from gopy to manually-written bindings
            raise Exception(str(exc)) from None
        return ret

    def sub_clone_game(self, game_id: int, count: int) -> list[int]:
        self._check_closed()
        try:
            ret = self._go_game_engine.SubCloneGame(game_id, count)
        except RuntimeError as exc:
            # We want to raise IOError (or its subclasses) or engine-specific
            # exceptions depending on what error is returned here but since gopy
            # flattens these, let's just raise generic Exception to not bind ourselves
            # to a tighter API contract.
            # TODO: map exceptions once we migrate from gopy to manually-written bindings
            raise Exception(str(exc)) from None
        return ret

    def delete_games(self, game_ids: list[int]) -> None:
        self._check_closed()
        self._go_game_engine.DeleteGames(game_ids)

    def send_play_turn_batch(
        self, concrete_requests: list[requests.PlayTurnRequest]
    ) -> list[requests.PlayTurnResponse]:
        self._check_closed()
        go_requests = _go_engine.Slice_Ptr_engine_PlayTurnRequest(
            req._unwrap() for req in concrete_requests
        )
        go_obj = self._go_game_engine.SendPlayTurnBatch(go_requests)
        return [requests.PlayTurnResponse(go_resp) for go_resp in go_obj]

    def send_get_remaining_tiles_batch(
        self, concrete_requests: list[requests.GetRemainingTilesRequest]
    ) -> list[requests.GetRemainingTilesResponse]:
        self._check_closed()
        go_requests = _go_engine.Slice_Ptr_engine_GetRemainingTilesRequest(
            req._unwrap() for req in concrete_requests
        )
        go_obj = self._go_game_engine.SendGetRemainingTilesBatch(go_requests)
        return [requests.GetRemainingTilesResponse(go_resp) for go_resp in go_obj]

    def send_get_legal_moves_batch(
        self, concrete_requests: list[requests.GetLegalMovesRequest]
    ) -> list[requests.GetLegalMovesResponse]:
        self._check_closed()
        go_requests = _go_engine.Slice_Ptr_engine_GetLegalMovesRequest(
            req._unwrap() for req in concrete_requests
        )
        go_obj = self._go_game_engine.SendGetLegalMovesBatch(go_requests)
        return [requests.GetLegalMovesResponse(go_resp) for go_resp in go_obj]
