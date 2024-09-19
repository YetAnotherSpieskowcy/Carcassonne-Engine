# flake8: noqa ascci drawing makes a lot errors
# mypy: ignore-errors
""""""

"""
 diagonal edges represent cities, dots fields, straight lines roads.
 Final board: (each tile is represented by 5x5 ascii signs, at the center is the turn number in hex :/)

 
            -1    0    1    2    3
            
            
           ..|..............|..     
           ..|..............|..     
1          ..9----1----2....8..     
           ..|.../ \...|....|..     
           ..|..|   |..|....|..     
           ..|..|   |..|....|.......
           ..|...\ /...|....|.......
0          ..4----0----3....B----C--
           ..|......................
           ..|......................
           ..|..     -...--...--...-
           .[ ].      \./  \./  \./
-1         .[5].       6    7    A
           .[ ].      /.\  /.\  /.\ 
           .....     -...--   --...-
"""
import logging
from pathlib import Path

from utils import TurnParams, check_final_points, check_points, make_turn

from carcassonne_engine import GameEngine, SerializedGame, tiletemplates
from carcassonne_engine._bindings.elements import MeepleType
from carcassonne_engine._bindings.feature import Type as FeatureType
from carcassonne_engine._bindings.side import Side
from carcassonne_engine.models import Position
from carcassonne_engine.tilesets import mini_tile_set, ordered_mini_tile_set1

log = logging.getLogger(__name__)


def test_two_player_game(tmp_path: Path) -> None:
    engine = GameEngine(4, tmp_path)
    tile_set = ordered_mini_tile_set1()

    game_id, game = engine.generate_ordered_game(tile_set)

    game_id, game = check_first_turn(engine, game, game_id)
    game_id, game = check_second_turn(engine, game, game_id)
    game_id, game = check_third_turn(engine, game, game_id)
    game_id, game = check_fourth_turn(engine, game, game_id)
    game_id, game = check_fifth_turn(engine, game, game_id)
    game_id, game = check_sixth_turn(engine, game, game_id)
    game_id, game = check_seventh_turn(engine, game, game_id)
    game_id, game = check_eighth_turn(engine, game, game_id)
    game_id, game = check_nineth_turn(engine, game, game_id)
    game_id, game = check_tenth_turn(engine, game, game_id)
    game_id, game = check_eleventh_turn(engine, game, game_id)
    game_id, game = check_twelfth_turn(engine, game, game_id)

    # TODO check final score
    game_id, game = finalize(engine, game, game_id)

    assert game.current_tile is None


"""
// straight road with city edge
// player 1 places meeple on city, and closes it

            -1    0    1    2    3
            
            
                .....               
                .....               
1               --1--              
                ./ \.               
                |   |               
                |   |               
                .\ /.               
0               --0--               
                .....               
                .....               
                                    
                                   
-1                                
                                    
 
"""


def check_first_turn(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=0, y=1),
        tile=tiletemplates.single_city_edge_straight_roads().rotate(2),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [4, 0])

    return game_id, game


"""
// road turn
// player 2 places meeple (@) on a road
            -1    0    1    2    3
            
            
                ..........          
                ..........          
1               --1--@-2..         
                ./ \...|..          
                |   |..|..          
                |   |               
                .\ /.               
0               --0--               
                .....               
                .....               
                                    
                                   
-1                                
                                    
  
"""


def check_second_turn(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=1, y=1),
        tile=tiletemplates.roads_turn(),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Left,
        featureType=FeatureType.Road,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [4, 0])

    return game_id, game


"""
// road turn
// player 1 places meeple (!) on a field
            -1    0    1    2    3
            
            
                ..........          
                ..........          
1               --1--@-2..         
                ./ \...|..          
                |   |..|..          
                |   |..|..          
                .\ /.!.|..          
0               --0----3..          
                ..........          
                ..........          
                                    
                                   
-1                                
                                    
    
"""


def check_third_turn(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=1, y=0),
        tile=tiletemplates.roads_turn().rotate(1),
        meepleType=MeepleType.NormalMeeple,
        side=Side.TopLeftEdge,
        featureType=FeatureType.Field,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [4, 0])

    return game_id, game


"""
// T cross road
// player2 places meeple (@) on road going down
            -1    0    1    2    3
            
            
                ..........          
                ..........          
1               --1--@-2..         
                ./ \...|..          
                |   |..|..          
           ..|..|   |..|..          
           ..|...\ /.!.|..          
0          ..4----0----3..          
           ..|............          
           ..@............          
                                    
                                   
-1                                
                                    
                        
"""


def check_fourth_turn(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=-1, y=0),
        tile=tiletemplates.t_cross_road().rotate(3),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.Road,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [4, 0])

    return game_id, game


"""
// monastery with single road
// player1 places meeple (!) on a monastery
// road from 4 to 5 is finished, so player2 scores 2 points
            -1    0    1    2    3
            
            
                ..........          
                ..........          
1               --1--@-2..         
                ./ \...|..          
                |   |..|..          
           ..|..|   |..|..          
           ..|...\ /.!.|..          
0          ..4----0----3..          
           ..|............          
           ..|............          
           ..|..                    
           .[ ].                   
-1         .[5].                  
           .[!].                    
           .....            
"""


def check_fifth_turn(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=-1, y=-1),
        tile=tiletemplates.monastery_with_single_road().rotate(2),
        meepleType=MeepleType.NormalMeeple,
        side=Side.NoSide,
        featureType=FeatureType.Monastery,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [4, 2])

    return game_id, game


"""
// Two city edges not connected
// player 2 places meeple(@) on the right city
            -1    0    1    2    3
            
            
                ..........          
                ..........          
1               --1--@-2..         
                ./ \...|..          
                |   |..|..          
           ..|..|   |..|..          
           ..|...\ /.!.|..          
0          ..4----0----3..          
           ..|............          
           ..|............          
           ..|..     -...-          
           .[ ].      \./          
-1         .[5].       6 @        
           .[!].      /.\           
           .....     -...-      
"""


def check_sixth_turn(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=1, y=-1),
        tile=tiletemplates.two_city_edges_up_and_down_not_connected().rotate(1),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Right,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [4, 2])

    return game_id, game


"""
// Two city edges not connected
// player 1 places meeple (!) on the right city
// playey 2 scores points for finished city
            -1    0    1    2    3
            
            
                ..........          
                ..........          
1               --1--@-2..         
                ./ \...|..          
                |   |..|..          
           ..|..|   |..|..          
           ..|...\ /.!.|..          
0          ..4----0----3..          
           ..|............          
           ..|............          
           ..|..     -...--...-     
           .[ ].      \./  \./     
-1         .[5].       6    7 !   
           .[!].      /.\  /.\      
           .....     -...--   - 
"""


def check_seventh_turn(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=2, y=-1),
        tile=tiletemplates.two_city_edges_up_and_down_not_connected().rotate(1),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Right,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [4, 6])

    return game_id, game


"""
// straight road
// player 2 places meeple (@) on a bottom road
            -1    0    1    2    3
            
            
                ............|..     
                ............|..     
1               --1--@-2....8..     
                ./ \...|....|..     
                |   |..|....@..     
           ..|..|   |..|..          
           ..|...\ /.!.|..          
0          ..4----0----3..          
           ..|............          
           ..|............          
           ..|..     -...--...-     
           .[ ].      \./  \./     
-1         .[5].       6    7 !   
           .[!].      /.\  /.\      
           .....     -...--   -   
"""


def check_eighth_turn(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=2, y=1),
        tile=tiletemplates.straight_roads().rotate(1),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.Road,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [4, 6])

    return game_id, game


"""
// T cross road
// road is finished. Player 2 scores 6 points for a road
// player 1 places meeple (!) on a field
            -1    0    1    2    3
            
            
           ..|.!............|..     
           ..|..............|..     
1          ..9----1----2....8..     
           ..|.../ \...|....|..     
           ..|..|   |..|....@..     
           ..|..|   |..|..          
           ..|...\ /.!.|..          
0          ..4----0----3..          
           ..|............          
           ..|............          
           ..|..     -...--...-     
           .[ ].      \./  \./     
-1         .[5].       6    7 !    
           .[!].      /.\  /.\      
           .....     -...--   -     
"""


def check_nineth_turn(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=-1, y=1),
        tile=tiletemplates.t_cross_road().rotate(3),
        meepleType=MeepleType.NormalMeeple,
        side=Side.TopRightEdge,
        featureType=FeatureType.Field,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [4, 12])

    return game_id, game


"""
// Two city edges not connected
// player 2 places meeple (@) on the right city
// player 1 scores points for city
            -1    0    1    2    3
            
            
           ..|.!............|..     
           ..|..............|..     
1          ..9----1----2....8..     
           ..|.../ \...|....|..     
           ..|..|   |..|....@..     
           ..|..|   |..|..          
           ..|...\ /.!.|..          
0          ..4----0----3..          
           ..|............          
           ..|............          
           ..|..     -...--...--...-
           .[ ].      \./  \./  \./
-1         .[5].       6    7    A @
           .[!].      /.\  /.\  /.\ 
           .....     -...--   --...-
"""


def check_tenth_turn(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=3, y=-1),
        tile=tiletemplates.two_city_edges_up_and_down_not_connected().rotate(1),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Right,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [8, 12])

    return game_id, game


"""
// road turn
            -1    0    1    2    3
            
            
           ..|.!............|..     
           ..|..............|..     
1          ..9----1----2....8..     
           ..|.../ \...|....|..     
           ..|..|   |..|....@..     
           ..|..|   |..|....|..     
           ..|...\ /.!.|....|..     
0          ..4----0----3....B--     
           ..|.................     
           ..|.................     
           ..|..     -...--...--...-
           .[ ].      \./  \./  \./
-1         .[5].       6    7    A @
           .[!].      /.\  /.\  /.\ 
           .....     -...--   --...-
"""


def check_eleventh_turn(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=2, y=-0),
        tile=tiletemplates.roads_turn().rotate(2),
        meepleType=MeepleType.NoneMeeple,
        side=Side.NoSide,
        featureType=FeatureType.NoneType,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [8, 12])

    return game_id, game


"""
// straight road
            -1    0    1    2    3
            
            
           ..|.!............|..     
           ..|..............|..     
1          ..9----1----2....8..     
           ..|.../ \...|....|..     
           ..|..|   |..|....@..     
           ..|..|   |..|....|.......
           ..|...\ /.!.|....|.......
0          ..4----0----3....B----C--
           ..|......................
           ..|......................
           ..|..     -...--...--...-
           .[ ].      \|/  \./  \./
-1         .[5].       6    7    A @
           .[!].      /.\  /.\  /.\ 
           .....     -...--   --...-
"""


def check_twelfth_turn(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=3, y=0),
        tile=tiletemplates.straight_roads(),
        meepleType=MeepleType.NoneMeeple,
        side=Side.NoSide,
        featureType=FeatureType.NoneType,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [8, 12])

    return game_id, game


"""
player1 scores 11 points in total:
    - 2*3points for farmer on 9 tile
    - 3 points for farmer on 3 tile
    - 2 points for monastery on 5 tile  

player2 scores 4 points in total:
    - 3 points for road on 8 tile
    - 1 point for city on A tile    

            -1    0    1    2    3


           ..|.!............|..     
           ..|..............|..     
1          ..9----1----2....8..     
           ..|.../ \...|....|..     
           ..|..|   |..|....@..     
           ..|..|   |..|....|.......
           ..|...\ /.!.|....|.......
0          ..4----0----3....B----C--
           ..|......................
           ..|......................
           ..|..     -...--...--...-
           .[ ].      \|/  \./  \./
-1         .[5].       6    7    A @
           .[!].      /.\  /.\  /.\ 
           .....     -...--   --...-
"""


def finalize(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    check_final_points([8 + 11, 12 + 4])
    return
