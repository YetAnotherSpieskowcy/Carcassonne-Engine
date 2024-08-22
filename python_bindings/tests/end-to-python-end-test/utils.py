from typing import NamedTuple

from carcassonne_engine import GameEngine, SerializedGame
from carcassonne_engine.models import PlacedTile, Tile, Position
from carcassonne_engine.requests import PlayTurnRequest, GetLegalMovesRequest
from carcassonne_engine._bindings.side import Side
from carcassonne_engine._bindings.elements import MeepleType
from carcassonne_engine._bindings.feature import Type as FeatureType


class TurnParams(NamedTuple):
    pos: Position
    tile: Tile
    meepleType: MeepleType
    side: Side
    featureType: FeatureType


def get_placed_tile(placed_tiles: list[PlacedTile], turnParams:TurnParams) -> PlacedTile:
    for ptile in placed_tiles:
        # if exact placement
        if ptile.to_tile().exact_equals(turnParams.tile) and ptile.position == turnParams.pos:
            if turnParams.meepleType == MeepleType.NoneMeeple:
                meepleExists = False
                # check if there is no meeple on a tile
                for feature in ptile._go_obj.Features:
                    if feature.Meeple.Type != MeepleType.NoneMeeple:
                        meepleExists = True
                        break
                if not meepleExists:
                    return ptile
            else:
                # check if meeple in desired position
                if ptile._go_obj.GetPlacedFeatureAtSide(sideToCheck=turnParams.side, featureType=turnParams.featureType).Meeple.Type == turnParams.meepleType:
                    return ptile

    raise KeyError("did not find the specified tile")


def make_turn(engine: GameEngine, game: SerializedGame, game_id: int, turn_params: TurnParams) -> (int, SerializedGame):
    # get legal moves
    legal_moves_req = GetLegalMovesRequest(
        base_game_id=game_id, tile_to_place=game.current_tile
    )
    (legal_moves_resp,) = engine.send_get_legal_moves_batch([legal_moves_req])

    # find the desired one
    move = get_placed_tile(legal_moves_resp.moves, turn_params.tile, turn_params.pos)

    # play turn
    play_turn_req = PlayTurnRequest(game_id=game_id, move=move)
    (play_turn_resp,) = engine.send_play_turn_batch([play_turn_req])
    if turn_params.valid_turn:
        assert play_turn_resp.exception is None
    else:
        assert play_turn_resp.exception is not None

    return play_turn_resp.game_id, play_turn_resp.game


def check_points(engine: GameEngine, game_id: int):
    # TODO access scores
    return
