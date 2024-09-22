from typing import NamedTuple, Tuple

from carcassonne_engine import GameEngine, SerializedGame
from carcassonne_engine._bindings.elements import MeepleType
from carcassonne_engine._bindings.feature import Type as FeatureType
from carcassonne_engine._bindings.side import Side
from carcassonne_engine.models import PlacedTile, Position, Tile
from carcassonne_engine.requests import (
    GetLegalMovesRequest,
    MoveWithState,
    PlayTurnRequest,
)


class TurnParams(NamedTuple):
    pos: Position
    tile: Tile
    meepleType: MeepleType
    side: Side
    featureType: FeatureType


def get_placed_tile(moves: list[MoveWithState], turnParams: TurnParams) -> PlacedTile:
    for move in moves:
        # if exact placement
        if (
            move.move.to_tile().exact_equals(turnParams.tile)
            and move.move.position == turnParams.pos
        ):
            if turnParams.meepleType == MeepleType.NoneMeeple:
                meepleExists = False
                # check if there is no meeple on a tile
                for feature in move.move._go_obj.Features:
                    if feature.Meeple.Type != MeepleType.NoneMeeple:
                        meepleExists = True
                        break
                if not meepleExists:
                    return move.move
            else:
                # check if meeple in desired position
                feature = move.move._go_obj.GetPlacedFeatureAtSide(
                    sideToCheck=turnParams.side.value,
                    featureType=turnParams.featureType.value,
                )
                if (
                    feature is not None
                    and feature.Meeple.Type != turnParams.meepleType.NoneMeeple.value
                ):
                    return move.move

    raise KeyError("did not find the specified tile")


def make_turn(
    engine: GameEngine, game: SerializedGame, game_id: int, turn_params: TurnParams
) -> Tuple[int, SerializedGame] | None:
    if game.current_tile is None:
        return None

    # get legal moves
    legal_moves_req = GetLegalMovesRequest(
        base_game_id=game_id, tile_to_place=game.current_tile
    )
    (legal_moves_resp,) = engine.send_get_legal_moves_batch([legal_moves_req])

    # find the desired one
    if legal_moves_resp.moves is None:
        return None
    move = get_placed_tile(legal_moves_resp.moves, turn_params)

    # play turn
    play_turn_req = PlayTurnRequest(game_id=game_id, move=move)
    (play_turn_resp,) = engine.send_play_turn_batch([play_turn_req])
    assert play_turn_resp.exception is None
    assert play_turn_resp.game is not None

    return play_turn_resp.game_id, play_turn_resp.game
