# flake8: noqa ascci drawing makes a lot errors
# mypy: ignore-errors

from carcassonne_engine._bindings.elements import MeepleType
from carcassonne_engine._bindings.side import Side
from carcassonne_engine.placed_tile import Position

"""
 diagonal edges represent cities, dots fields, straight lines roads.
 Player meeples will be represented as !@ signs (you know, writing number but with shift!) 1->!, 2->@
 Final board: (each tile is represented by 5x5 ascii signs, at the center is the turn number in hex :/)

|                 0    1    2    3
|
|
|               |   |...............
|               .\ /............[ ].
|1              ..2....4----9---[A].
|               ./ \...|.../ \..[ ].
|               |   |..|..|   |.....
|          .....|   |..|..|   |
|          ......\ /...|...\ /.
|0         ..8----0----1....3..
|          ..|.........|.../ \.
|          ..|.........|..|   |
|          ..|.........|..|   |
|          ..|.........|...\ /
|-1        ..B----6----5....7..
|          ..|............./ \.
|          ..|............|   |
|               .....
|               .....
|-2             --C--
|               .....
|               .....
"""
import logging
from pathlib import Path

from end_utils import TurnParams, check_points, make_turn

from carcassonne_engine import GameEngine, tiletemplates
from carcassonne_engine._bindings.feature import Type as FeatureType
from carcassonne_engine.models import SerializedGame
from carcassonne_engine.tilesets import TileSet

log = logging.getLogger(__name__)


def test_two_player_game2(tmp_path: Path) -> None:
    print(tmp_path)
    engine = GameEngine(4, tmp_path)
    tile_set = create_tileset()

    game_id, game = engine.generate_ordered_game(tile_set)

    game_id, game = check_first_turn(engine, game, game_id)
    game_id, game = check_second_turn(engine, game, game_id)
    game_id, game = check_third_turn(engine, game, game_id)
    game_id, game = check_fourth_turn(engine, game, game_id)
    game_id, game = check_fifth_turn(engine, game, game_id)
    game_id, game = check_sixth_turn(engine, game, game_id)
    game_id, game = check_seventh_turn(engine, game, game_id)
    game_id, game = check_eighth_turn(engine, game, game_id)
    game_id, game = check_ninth_turn(engine, game, game_id)
    game_id, game = check_tenth_turn(engine, game, game_id)
    game_id, game = check_eleventh_turn(engine, game, game_id)
    game_id, game = check_twelfth_turn(engine, game, game_id)

    assert game.current_tile is None


def create_tileset() -> TileSet:
    tiles = [
        tiletemplates.t_cross_road().rotate(1),
        tiletemplates.two_city_edges_up_and_down_not_connected(),
        tiletemplates.two_city_edges_up_and_down_not_connected(),
        tiletemplates.roads_turn().rotate(3),
        tiletemplates.roads_turn().rotate(1),
        tiletemplates.straight_roads(),
        tiletemplates.two_city_edges_up_and_down_not_connected(),
        tiletemplates.roads_turn().rotate(3),
        tiletemplates.single_city_edge_straight_roads().rotate(2),
        tiletemplates.monastery_with_single_road().rotate(1),
        tiletemplates.t_cross_road().rotate(3),
        tiletemplates.straight_roads(),
    ]

    return TileSet.from_tiles(
        tiles, starting_tile=tiletemplates.single_city_edge_straight_roads()
    )

"""
// player1 places T Cross road with meeple on a bottom road
|                 0    1    2    3
|
|
|
|
|1
|
|
|               |   |..|..
|               .\ /...|..
|0              --0----1..
|               .......!..
|               .......|..
|
|
|-1
|
|
|
|
|-2
|
|
"""


def check_first_turn(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=1, y=0),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.Road,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [0, 0])

    return game_id, game


"""
// player2 places Two city edges not connected, and a meeple on a closing city
|                 0    1    2    3
|
|
|               |   |
|               .\ /.
|1              ..2..
|               ./ \.
|               |   |
|               |   |..|..
|               .\ /...|..
|0              --0----1..
|               .......!..
|               .......|..
|
|
|-1
|
|
|
|
|-2
|
|
"""


def check_second_turn(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=0, y=1),
        tile=tiletemplates.two_city_edges_up_and_down_not_connected(),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [0, 0 + 4])

    return game_id, game


"""
// player1 places Two city edges not connected and a meeple on bottom city
|                 0    1    2    3
|
|
|               |   |
|               .\ /.
|1              ..2..
|               ./ \.
|               |   |
|               |   |..|..|   |
|               .\ /...|...\ /.
|0              --0----1....3..
|               .......!.../ \.
|               .......|..| ! | 
|
|
|-1
|
|
|
|
|-2
|
|
"""


def check_third_turn(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=2, y=0),
        tile=tiletemplates.two_city_edges_up_and_down_not_connected(),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [0, 4])

    return game_id, game


"""
// player2 places Road turn with a meeple(@) on a road
|                 0    1    2    3
|
|
|               |   |.....
|               .\ /......
|1              ..2....4-@
|               ./ \...|..
|               |   |..|..
|               |   |..|..|   |
|               .\ /...|...\ /.
|0              --0----1....3..
|               .......!.../ \.
|               .......|..| ! |
|
|
|-1
|
|
|
|
|-2
|
|
"""


def check_fourth_turn(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=1, y=1),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.Right,
        featureType=FeatureType.Road,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [0, 4])

    return game_id, game


"""
// player1 places Road turn with a farmer(!) on the right bottom
|                 0    1    2    3
|
|
|               |   |.....
|               .\ /......
|1              ..2....4-@
|               ./ \...|..
|               |   |..|..
|               |   |..|..|   |
|               .\ /...|...\ /.
|0              --0----1....3..
|               .......!.../ \.
|               .......|..| ! |
|                    ..|..
|                    ..|..
|-1                  --5..
|                    ...!.
|                    .....
|
|
|-2
|
|
"""


def check_fifth_turn(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=1, y=-1),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.BottomRightEdge,
        featureType=FeatureType.Field,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [0, 4])

    return game_id, game


"""
// player2 places Straight Road with a meeple on a top field
|                 0    1    2    3
|
|
|               |   |.....
|               .\ /......
|1              ..2....4-@
|               ./ \...|..
|               |   |..|..
|               |   |..|..|   |
|               .\ /...|...\ /.
|0              --0----1....3..
|               .......!.../ \.
|               .......|..| ! |
|               ..@....|..
|               .......|..
|-1             --6----5..
|               ........!.
|               ..........
|
|
|-2
|
|
"""


def check_sixth_turn(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=0, y=-1),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.Top,
        featureType=FeatureType.Field,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [0, 4])

    return game_id, game


"""
// player1 places Two city edges not connected, finishing own city, and placing a meeple on a new one
// player1 scores 4 points for a city
|                 0    1    2    3
|
|
|               |   |.....
|               .\ /......
|1              ..2....4-@
|               ./ \...|..
|               |   |..|..
|               |   |..|..|   |
|               .\ /...|...\ /.
|0              --0----1....3..
|               .......!.../ \.
|               .......|..|   |
|               ..@....|..|   |
|               .......|...\ /
|-1             --6----5....7..
|               ........!../ \.
|               ..........| ! |
|
|
|-2
|
|
"""


def check_seventh_turn(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=2, y=-1),
        tile=tiletemplates.two_city_edges_up_and_down_not_connected(),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [0 + 4, 4])

    return game_id, game


"""
// player2 places Road turn and a meeple on a road
|                 0    1    2    3
|
|
|               |   |.....
|               .\ /......
|1              ..2....4-@
|               ./ \...|..
|               |   |..|..
|          .....|   |..|..|   |
|          ......\ /...|...\ /.
|0         ..8----0----1....3..
|          ..|.........!.../ \.
|          ..@.........|..|   |
|               ..@....|..|   |
|               .......|...\ /
|-1             --6----5....7..
|               ........!../ \.
|               ..........| ! |
|
|
|-2
|
|
"""


def check_eighth_turn(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=-1, y=0),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.Road,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [4, 4])

    return game_id, game


"""
// player1 places Straight road with city edge and places a meeple on the upper field.
No one scores for the finished city
|                 0    1    2    3
|
|
|               |   |.......!..
|               .\ /...........
|1              ..2....4-@--9--
|               ./ \...|.../ \.
|               |   |..|..|   |
|          .....|   |..|..|   |
|          ......\ /...|...\ /.
|0         ..8----0----1....3..
|          ..|.........!.../ \.
|          ..@.........|..|   |
|               ..@....|..|   |
|               .......|...\ /
|-1             --6----5....7..
|               ........!../ \.
|               ..........| ! |
|
|
|-2
|
|
"""


def check_ninth_turn(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=2, y=1),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.TopRightEdge,
        featureType=FeatureType.Field,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [4, 4])

    return game_id, game


"""
// player2 places Monastery with road, with a meeple on a monastery.
player2 scores 3 points for finished road
|                 0    1    2    3
|
|
|               |   |.......!.......
|               .\ /............[ ].
|1              ..2....4----9---[A].
|               ./ \...|.../ \..[@].
|               |   |..|..|   |.....
|          .....|   |..|..|   |
|          ......\ /...|...\ /.
|0         ..8----0----1....3..
|          ..|.........!.../ \.
|          ..@.........|..|   |
|               ..@....|..|   |
|               .......|...\ /
|-1             --6----5....7..
|               ........!../ \.
|               ..........| ! |
|
|
|-2
|
|
"""


def check_tenth_turn(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=3, y=1),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.NoSide,
        featureType=FeatureType.Monastery,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [4, 4 + 4])

    return game_id, game


"""
// player1 places T cross road, with a meeple on a bottom road
player1 and player2 score 4 points for their roads
|                 0    1    2    3
|
|
|               |   |.......!.......
|               .\ /............[ ].
|1              ..2....4----9---[A].
|               ./ \...|.../ \..[@].
|               |   |..|..|   |.....
|          .....|   |..|..|   |
|          ......\ /...|...\ /.
|0         ..8----0----1....3..
|          ..|.........|.../ \.
|          ..|.........|..|   |
|          ..|....@....|..|   |
|          ..|.........|...\ /
|-1        ..B----6----5....7..
|          ..|..........!../ \.
|          ..!............| ! |
|
|
|-2
|
|
"""


def check_eleventh_turn(
    engine: GameEngine, game, game_id
) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=-1, y=-1),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.Road,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [4 + 4, 8 + 4])

    return game_id, game


"""
// player2 places Straight road with a meeple on a road
|                 0    1    2    3
|
|
|               |   |.......!.......
|               .\ /............[ ].
|1              ..2....4----9---[A].
|               ./ \...|.../ \..[@].
|               |   |..|..|   |.....
|          .....|   |..|..|   |
|          ......\ /...|...\ /.
|0         ..8----0----1....3..
|          ..|.........|.../ \.
|          ..|.........|..|   |
|          ..|....@....|..|   |
|          ..|.........|...\ /
|-1        ..B----6----5....7..
|          ..|..........!../ \.
|          ..!............| ! |
|               .....
|               .....
|-2             --C-@
|               .....
|               .....
"""


def check_twelfth_turn(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=0, y=-2),
        tile=tiletemplates.straight_roads(),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Right,
        featureType=FeatureType.Road,
    )

    game_id, game = make_turn(
        engine, game, game_id, turn_params, {1: 8 + 11, 2: 12 + 4}
    )
    check_points(game, [8, 12])

    return game_id, game


"""
player1 scores 11 points in total:
    - 3*3points for farmer on 9 and 5 tile
    - 1 point for road on B tile
    - 1 point for city on 7 tile        

player2 scores 4 points in total:    
    - 1 point for road on C tile
    - 3 points for monastery on A tile    

|                 0    1    2    3
|
|
|               |   |.......!.......
|               .\ /............[ ].
|1              ..2....4----9---[A].
|               ./ \...|.../ \..[@].
|               |   |..|..|   |.....
|          .....|   |..|..|   |
|          ......\ /...|...\ /.
|0         ..8----0----1....3..
|          ..|.........|.../ \.
|          ..|.........|..|   |
|          ..|....@....|..|   |
|          ..|.........|...\ /
|-1        ..B----6----5....7..
|          ..|..........!../ \.
|          ..!............| ! |
|               .....
|               .....
|-2             --C-@
|               .....
|               .....
"""
