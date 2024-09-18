# flake8: noqa ascci drawing makes a lot errors
""""""

"""
 diagonal edges represent cities, dots fields, straight lines roads.
 Player meeples will be represented as !@ signs ( you know, writing number but with shift!) 1->!, 2->@
 Final board: (each tile is represented by 5x5 ascii signs, at the center is the turn number :/)
 Tiles are played in order of: https://docs.google.com/spreadsheets/d/1TnPvB6oyisNGs7GZ0xpu-3LPp1V5-t0xH4vocCUPvsY/edit?gid=0#gid=0
 
 GIANT CITY!

       -3   -2   -1    0    1    2    3


                     -.|.-
                      \|/ 
4                      N    
                          
                                                      
                     |   |..|.--...--.|..    
                     .\ /...|/  \./  \|..  
3                    --B----I    K    J-- 
                     ..|.../         * \.
                     ..|..|             |
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   ||   ||   ||   ||   |
           ..|..|   |                   |
           ..|...\ /        *          /.
1          --5....F    O    L    M    G..
           ..|.....\       /.\  /|\  /...
           ..|......-     -...--.|.--....
      -......|.......|   |.....
       \.....|........\ /..[ ].
0       H....4----3----0---[2].
         \.................[ ].
          |....................
                     .........--....
                     .[ ]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..
               
"""
import logging
from pathlib import Path

from utils import make_turn, TurnParams, check_points, check_final_points

from carcassonne_engine import GameEngine, SerializedGame, tiletemplates
from carcassonne_engine._bindings.elements import MeepleType
from carcassonne_engine._bindings.feature import Type as FeatureType
from carcassonne_engine._bindings.side import Side
from carcassonne_engine.models import Position
from carcassonne_engine.tilesets import mini_tile_set


log = logging.getLogger(__name__)


def test_all_tiles_game(tmp_path: Path) -> None:
    engine = GameEngine(4, tmp_path)
    tile_set = mini_tile_set()

    game_id, game = engine.generate_ordered_game(tile_set)

    game_id, game = check_turn_1(engine, game, game_id)

    # TODO check final score
    game_id, game = finalize(engine, game, game_id)

    assert game.current_tile is None


"""
player1 places meeple (!) on a monastery 
       -3   -2   -1    0    1    2    3




4                           




3                                    




2          




1                              


                     |   |
                     .\ /.
0                    --0--
                     .....
                     .....
                     .....
                     .[!].  
-1                   .[1].       
                     .[ ]. 
                     .....

"""


def check_turn_1(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=0, y=-1),
        tile=tiletemplates.monastery_without_roads(),
        meepleType=MeepleType.NormalMeeple,
        side=Side.NoSide,
        featureType=FeatureType.Monastery,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [0, 0])

    return game_id, game


"""
player2 places meeple(@) on a monastery

       -3   -2   -1    0    1    2    3




4                           




3                                    




2          




1                              


                     |   |.....
                     .\ /..[@].
0                    --0---[2].
                     ......[ ].
                     ..........
                     .....
                     .[!].  
-1                   .[1].       
                     .[ ]. 
                     .....

"""


def check_turn_2(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=1, y=0),
        tile=tiletemplates.monastery_with_single_road().rotate(1),
        meepleType=MeepleType.NormalMeeple,
        side=Side.NoSide,
        featureType=FeatureType.Monastery,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [0, 0])

    return game_id, game


"""
player1 places meeple(!) on a road
       -3   -2   -1    0    1    2    3




4                           




3                                    




2          




1                              


                .....|   |.....
                ......\ /..[@].
0               !-3----0---[2].
                ...........[ ].
                ...............
                     .....
                     .[!].  
-1                   .[1].       
                     .[ ]. 
                     .....

"""


def check_turn_3(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=-1, y=0),
        tile=tiletemplates.straight_roads(),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Left,
        featureType=FeatureType.Road,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [0, 0])

    return game_id, game


"""
player2 places meeple on a field

       -3   -2   -1    0    1    2    3




4                           




3                                    




2          


                               
                               
1                              
                               
                               
           ..|.......|   |.....
           ..|........\ /..[@].
0          ..4--!-3----0---[2].
           ................[ ].
           ..@.................
                     .....
                     .[!].  
-1                   .[1].       
                     .[ ]. 
                     .....

"""


def check_turn_4(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=-2, y=0),
        tile=tiletemplates.roads_turn().rotate(2),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.Field,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [0, 0])

    return game_id, game


"""
player1 finished road, places meeple(!) on a top road
player1 scores 5 points
       -3   -2   -1    0    1    2    3




4                           




3                                    


           
           
2          
           
           
           ..!..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[@].
0          ..4----3----0---[2].
           ................[ ].
           ..@.................
                     .....
                     .[!].  
-1                   .[1].       
                     .[ ]. 
                     .....

"""


def check_turn_5(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=-2, y=1),
        tile=tiletemplates.t_cross_road().rotate(1),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Top,
        featureType=FeatureType.Road,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [5, 0])

    return game_id, game


"""
player 2 finishes road
player 2 places meeple on a right road
player 1 scores 2 points

       -3   -2   -1    0    1    2    3




4                           




3                                    


           ..|..
           ..|..
2          --6-@
           ..|..
           ..|..
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[@].
0          ..4----3----0---[2].
           ................[ ].
           ..@.................
                     .....
                     .[!].  
-1                   .[1].       
                     .[ ]. 
                     .....

"""


def check_turn_6(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=-2, y=2),
        tile=tiletemplates.x_cross_road(),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Right,
        featureType=FeatureType.Road,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [5, 0])

    return game_id, game


"""
player 1 places meeple(!) on a city

       -3   -2   -1    0    1    2    3




4                           




3                                    


           ..|..
           ..|..
2          --6-@
           ..|..
           ..|..
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[@].
0          ..4----3----0---[2].
           ................[ ].
           ..@.................
                     .........-
                     .[!]..../  
-1                   .[1]...7 !      
                     .[ ]....\ 
                     .........-

"""


def check_turn_7(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=2, y=-1),
        tile=tiletemplates.single_city_edge_no_roads().rotate(1),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Right,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [5, 0])

    return game_id, game


"""
player2 places meeple(@) on a city

       -3   -2   -1    0    1    2    3




4                           




3                                    


           ..|.......
           ..|.......
2          --6-@--8--
           ..|.../ \.
           ..|..| @ |
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[@].
0          ..4----3----0---[2].
           ................[ ].
           ..@.................
                     .........-
                     .[!]..../  
-1                   .[1]...7 !     
                     .[ ]....\ 
                     .........-

"""


def check_turn_8(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=-1, y=2),
        tile=tiletemplates.single_city_edge_straight_roads().rotate(2),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [5, 0])

    return game_id, game


"""
player1 places meeple(!) on a upper field
player1 closes city 
player1 scores 4 points
       -3   -2   -1    0    1    2    3




4                           




3                                    


           ..|.......
           ..|.......
2          --6-@--8--
           ..|.../ \.
           ..|..| @ |
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[@].
0          ..4----3----0---[2].
           ................[ ].
           ..@.................
                     .........--...!
                     .[!]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

"""


def check_turn_9(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=2, y=-1),
        tile=tiletemplates.single_city_edge_left_road_turn().rotate(3),
        meepleType=MeepleType.NormalMeeple,
        side=Side.TopRightEdge,
        featureType=FeatureType.Field,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [9, 0])

    return game_id, game


"""
player2 places meeple on a city
       -3   -2   -1    0    1    2    3




4                           


                                        
                                      
3                                    
                                    
                                    
           ..|.........|..
           ..|.........|..
2          --6-@--8----A..
           ..|.../ \../ \.
           ..|..| @ || @ |
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[@].
0          ..4----3----0---[2].
           ................[ ].
           ..@.................
                     .........--...!
                     .[!]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

"""


def check_turn_A(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=0, y=2),
        tile=tiletemplates.single_city_edge_right_road_turn().rotate(2),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [9, 0])

    return game_id, game


"""
player 1 places meeple(!) on a field
player 2 scores 4 points
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     ..|.!               
           ..|.........|..
           ..|.........|..
2          --6----8----A..
           ..|.../ \../ \.
           ..|..| @ || @ |
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[@].
0          ..4----3----0---[2].
           ................[ ].
           ..@.................
                     .........--...!
                     .[!]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

"""


def check_turn_B(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=0, y=3),
        tile=tiletemplates.single_city_edge_cross_road(),
        meepleType=MeepleType.NormalMeeple,
        side=Side.BottomRightEdge,
        featureType=FeatureType.Field,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [9, 0])

    return game_id, game


"""
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     ..|.!               
           ..|.........|..|   |
           ..|.........|...\ /.
2          --6----8----A....C..
           ..|.../ \../ \../ \.
           ..|..| @ || @ ||   |
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[@].
0          ..4----3----0---[2].
           ................[ ].
           ..@.................
                     .........--...!
                     .[!]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

"""


def check_turn_C(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=1, y=2),
        tile=tiletemplates.two_city_edges_up_and_down_not_connected(),
        meepleType=MeepleType.NoneMeeple,
        side=Side.NoSide,
        featureType=FeatureType.NoneType,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [9, 0])

    return game_id, game


"""
player 1 places meeple(!) on a city
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     ..|.!               
           ..|.........|..|   ||   |
           ..|.........|...\ /..\ /.
2          --6----8----A....C...|D|.
           ..|.../ \../ \../ \../ \.
           ..|..| @ || @ ||   || ! |
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[@].
0          ..4----3----0---[2].
           ................[ ].
           ..@.................
                     .........--...!
                     .[!]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

"""


def check_turn_D(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=2, y=2),
        tile=tiletemplates.two_city_edges_up_and_down_connected(),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [9, 0])

    return game_id, game


"""
player 2 places meeple(@) on a city
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     ..|.!               
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..| @ || @ ||   || ! || @ |
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[@].
0          ..4----3----0---[2].
           ................[ ].
           ..@.................
                     .........--...!
                     .[!]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

"""


def check_turn_E(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=3, y=2),
        tile=tiletemplates.two_city_edges_up_and_down_connected_shield(),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [9, 0])

    return game_id, game


"""
player1 places meeple(!) on a field
player2 scores 4 points

       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     ..|.!               
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |               
           ..|...\ /                
1          --5..!.F                 
           ..|.....\                
           ..|......-               
           ..|.......|   |.....
           ..|........\ /..[@].
0          ..4----3----0---[2].
           ................[ ].
           ..@.................
                     .........--...!
                     .[!]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

"""


def check_turn_F(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=-1, y=1),
        tile=tiletemplates.two_city_edges_corner_not_connected(),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Left,
        featureType=FeatureType.Field,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [9, 4])

    return game_id, game


"""
player2 places meeple(@) on a field
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     ..|.!               
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |                   |
           ..|...\ /                   /.
1          --5..!.F                   G..
           ..|.....\                 /...
           ..|......-               -..@.
           ..|.......|   |.....
           ..|........\ /..[@].
0          ..4----3----0---[2].
           ................[ ].
           ..@.................
                     .........--...!
                     .[!]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

"""


def check_turn_G(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=3, y=1),
        tile=tiletemplates.two_city_edges_corner_connected().rotate(3),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.Field,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [9, 4])

    return game_id, game


"""
player1 places meeple(!) on a city 
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     ..|.!               
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |                   |
           ..|...\ /                   /.
1          --5..!.F                   G..
           ..|.....\                 /...
           ..|......-               -..@.
      -......|.......|   |.....
       \.....|........\ /..[@].
0     * H....4----3----0---[2].
       ! \.................[ ].
          |..@.................
                     .........--...!
                     .[!]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

"""


def check_turn_H(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=-3, y=0),
        tile=tiletemplates.two_city_edges_corner_connected_shield().rotate(2),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [9, 4])

    return game_id, game


"""
player2 places meeple(@) on a road


       -3   -2   -1    0    1    2    3




4                           


                     |   |..@.-              
                     .\ /...|/             
3                    --B----I             
                     ..|.../             
                     ..|.!|              
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |                   |
           ..|...\ /                   /.
1          --5..!.F                   G..
           ..|.....\                 /...
           ..|......-               -..@.
      -......|.......|   |.....
       \.....|........\ /..[@].
0    *  H....4----3----0---[2].
       ! \.................[ ].
          |..@.................
                     .........--...!
                     .[!]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

"""


def check_turn_I(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=1, y=3),
        tile=tiletemplates.two_city_edges_corner_connected_road_turn().rotate(1),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Top,
        featureType=FeatureType.Road,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [9, 4])

    return game_id, game


"""
player1 places meeple(!) on a road

       -3   -2   -1    0    1    2    3




4                           


                     |   |..@.-     -.|..    
                     .\ /...|/       \|..  
3                    --B----I         J-! 
                     ..|.../         * \.
                     ..|.!|             |
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |                   |
           ..|...\ /                   /.
1          --5..!.F                   G..
           ..|.....\                 /...
           ..|......-               -..@.
      -......|.......|   |.....
       \.....|........\ /..[@].
0     * H....4----3----0---[2].
       ! \.................[ ].
          |..@.................
                     .........--...!
                     .[!]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

"""


def check_turn_J(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=3, y=3),
        tile=tiletemplates.two_city_edges_corner_connected_road_turn_shield().rotate(2),
        meepleType=MeepleType.NormalMeeple,
        side=Side.Right,
        featureType=FeatureType.Road,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [9, 4])

    return game_id, game


"""
       -3   -2   -1    0    1    2    3




4                           


                     |   |..@.--...--.|..    
                     .\ /...|/  \./  \|..  
3                    --B----I    K    J-! 
                     ..|.../         * \.
                     ..|.!|             |
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |                   |
           ..|...\ /                   /.
1          --5..!.F                   G..
           ..|.....\                 /...
           ..|......-               -..@.
      -......|.......|   |.....
       \.....|........\ /..[@].
0     * H....4----3----0---[2].
       ! \.................[ ].
          |..@.................
                     .........--...!
                     .[!]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

"""


def check_turn_K(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=2, y=3),
        tile=tiletemplates.three_city_edges_connected().rotate(2),
        meepleType=MeepleType.NoneMeeple,
        side=Side.NoSide,
        featureType=FeatureType.NoneType,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [9, 4])

    return game_id, game


"""
       -3   -2   -1    0    1    2    3




4                           


                     |   |..@.--...--.|..    
                     .\ /...|/  \./  \|..  
3                    --B----I    K    J-! 
                     ..|.../         * \.
                     ..|.!|             |
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |                   |
           ..|...\ /        *          /.
1          --5..!.F         L         G..
           ..|.....\       /.\       /...
           ..|......-     -...-     -..@.
      -......|.......|   |.....
       \.....|........\ /..[@].
0     * H....4----3----0---[2].
       ! \.................[ ].
          |..@.................
                     .........--...!
                     .[!]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

"""


def check_turn_L(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=1, y=1),
        tile=tiletemplates.three_city_edges_connected_shield(),
        meepleType=MeepleType.NoneMeeple,
        side=Side.NoSide,
        featureType=FeatureType.NoneType,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [9, 4])

    return game_id, game


"""
       -3   -2   -1    0    1    2    3


                          
                          
4                           


                     |   |..@.--...--.|..    
                     .\ /...|/  \./  \|..  
3                    --B----I    K    J-! 
                     ..|.../         * \.
                     ..|.!|             |
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |                   |
           ..|...\ /        *          /.
1          --5..!.F         L    M    G..
           ..|.....\       /.\  /|\  /...
           ..|......-     -...--.|.--..@.
      -......|.......|   |.....
       \.....|........\ /..[@].
0     * H....4----3----0---[2].
       ! \.................[ ].
          |..@.................
                     .........--...!
                     .[!]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

"""


def check_turn_M(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=2, y=1),
        tile=tiletemplates.three_city_edges_connected_road(),
        meepleType=MeepleType.NoneMeeple,
        side=Side.NoSide,
        featureType=FeatureType.NoneType,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [9, 4])

    return game_id, game


"""
       -3   -2   -1    0    1    2    3


                     -.|.-
                      \|/ 
4                    * N    


                     |   |..@.--...--.|..    
                     .\ /...|/  \./  \|..  
3                    --B----I    K    J-! 
                     ..|.../         * \.
                     ..|.!|             |
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |                   |
           ..|...\ /        *          /.
1          --5..!.F         L    M    G..
           ..|.....\       /.\  /|\  /...
           ..|......-     -...--.|.--..@.
      -......|.......|   |.....
       \.....|........\ /..[@].
0     * H....4----3----0---[2].
       ! \.................[ ].
          |..@.................
                     .........--...!
                     .[!]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

"""


def check_turn_N(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=0, y=4),
        tile=tiletemplates.three_city_edges_connected_road_shield().rotate(2),
        meepleType=MeepleType.NoneMeeple,
        side=Side.NoSide,
        featureType=FeatureType.NoneType,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [9, 4])

    return game_id, game


"""
player1 scores nothing because player2 has meeples in the city
player2 scores: 13*2 + 3*2 = 26+6=32
    13 tile city
    3 shield in a city

       -3   -2   -1    0    1    2    3


                     -.|.-
                      \|/ 
4                    * N    
                          
                                                      
                     |   |..@.--...--.|..    
                     .\ /...|/  \./  \|..  
3                    --B----I    K    J-! 
                     ..|.../         * \.
                     ..|.!|             |
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |                   |
           ..|...\ /        *          /.
1          --5..!.F    O    L    M    G..
           ..|.....\       /.\  /|\  /...
           ..|......-     -...--.|.--..@.
      -......|.......|   |.....
       \.....|........\ /..[@].
0    *  H....4----3----0---[2].
       ! \.................[ ].
          |..@.................
                     .........--...!
                     .[!]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..
               
"""


def check_turn_O(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=0, y=1),
        tile=tiletemplates.four_city_edges_connected_shield(),
        meepleType=MeepleType.NoneMeeple,
        side=Side.NoSide,
        featureType=FeatureType.NoneType,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [9, 36])

    return game_id, game


"""
player1 scores 20 points in total:
    - 1+1 points for a city with shield on H tile
    - 5 points for monastery on 1 tile
    - 3 points for farmer on 9 tile
    - 2*3points for farmer on F tile
    - 3 points for farmer on B tile
    - 1 point for road on J tile
    
player2 scores 19 points in total:   
    - 2*3points for farmer on 4 tile
    - 8 points for monastery on 2 tile
    - 3 points for farmer on G tile
    - 2 points for road on I tile 

       -3   -2   -1    0    1    2    3


                     -.|.-
                      \|/ 
4                    * N    


                     |   |..@.--...--.|..    
                     .\ /...|/  \./  \|..  
3                    --B----I    K    J-! 
                     ..|.../         * \.
                     ..|.!|             |
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   ||   ||   ||   ||   |
           ..|..|   |                   |
           ..|...\ /        *          /.
1          --5..!.F    O    L    M    G..
           ..|.....\       /.\  /|\  /...
           ..|......-     -...--.|.--..@.
      -......|.......|   |.....
       \.....|........\ /..[@].
0    *  H....4----3----0---[2].
       ! \.................[ ].
          |..@.................
                     .........--...!
                     .[!]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..
"""


def finalize(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    check_final_points([9 + 20, 36 + 19])
    return
