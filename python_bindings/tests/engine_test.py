from pathlib import Path

from carcassonne_engine import (
    PlayTurnRequest,
    Slice_Ptr_engine_PlayTurnRequest,
    StartGameEngine,
)
from carcassonne_engine.tilesets import StandardTileSet


def test_game_engine_send_batch_receives_correct_responses_after_worker_requests(
    tmp_path: Path
) -> None:
    engine = StartGameEngine(4, str(tmp_path))
    request_count = 100

    tile_set = StandardTileSet()
    games = [engine.GenerateGame(tile_set) for _ in range(request_count)]
    requests = [
        PlayTurnRequest(GameID=game.ID, Move=game.Game.ValidTilePlacements[0])
        for game in games
    ]
    responses = engine.SendPlayTurnBatch(Slice_Ptr_engine_PlayTurnRequest(requests))
    for idx, resp in enumerate(responses):
        resp.Err()  # the binding would raise if this returned an error
        expected = requests[idx].GameID
        actual = resp.GameID()
        assert actual == expected
    engine.Shutdown()
