# flake8: noqa ascci drawing makes a lot errors
# mypy: ignore-errors

import logging
from pathlib import Path

from end_utils import TurnParams, check_points, make_turn

from carcassonne_engine import GameEngine
from carcassonne_engine._bindings.elements import MeepleType
from carcassonne_engine._bindings.feature import Type as FeatureType
from carcassonne_engine._bindings.side import Side
from carcassonne_engine.models import SerializedGame
from carcassonne_engine.placed_tile import Position
from carcassonne_engine.tilesets import every_tile_once_tile_set

"""
 diagonal edges represent cities, dots fields, straight lines roads.
 Player meeples will be represented as !@ signs (you know, writing number but with shift!) 1->!, 2->@
 Final board: (each tile is represented by 5x5 ascii signs, at the center is the turn number :/)
 Tiles are played in order of: 
 https://docs.google.com/spreadsheets/d/1TnPvB6oyisNGs7GZ0xpu-3LPp1V5-t0xH4vocCUPvsY/edit?gid=0#gid=0
 
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

log = logging.getLogger(__name__)


def test_all_tiles_game(tmp_path: Path) -> None:
    engine = GameEngine(4, tmp_path)
    tile_set = every_tile_once_tile_set()

    game_id, game = engine.generate_ordered_game(tile_set)

    game_id, game = check_turn_1(engine, game, game_id)
    game_id, game = check_turn_2(engine, game, game_id)
    game_id, game = check_turn_3(engine, game, game_id)
    game_id, game = check_turn_4(engine, game, game_id)
    game_id, game = check_turn_5(engine, game, game_id)
    game_id, game = check_turn_6(engine, game, game_id)
    game_id, game = check_turn_7(engine, game, game_id)
    game_id, game = check_turn_8(engine, game, game_id)
    game_id, game = check_turn_9(engine, game, game_id)
    game_id, game = check_turn_A(engine, game, game_id)
    game_id, game = check_turn_B(engine, game, game_id)
    game_id, game = check_turn_C(engine, game, game_id)
    game_id, game = check_turn_D(engine, game, game_id)
    game_id, game = check_turn_E(engine, game, game_id)
    game_id, game = check_turn_F(engine, game, game_id)
    game_id, game = check_turn_G(engine, game, game_id)
    game_id, game = check_turn_H(engine, game, game_id)
    game_id, game = check_turn_I(engine, game, game_id)
    game_id, game = check_turn_J(engine, game, game_id)
    game_id, game = check_turn_K(engine, game, game_id)
    game_id, game = check_turn_L(engine, game, game_id)
    game_id, game = check_turn_M(engine, game, game_id)
    game_id, game = check_turn_N(engine, game, game_id)
    game_id, game = check_turn_O(engine, game, game_id)

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


def check_turn_1(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=0, y=-1),
        tile=game.current_tile,
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


def check_turn_2(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=1, y=0),
        tile=game.current_tile,
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


def check_turn_3(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=-1, y=0),
        tile=game.current_tile,
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


def check_turn_4(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=-2, y=0),
        tile=game.current_tile,
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


def check_turn_5(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=-2, y=1),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.Top,
        featureType=FeatureType.Road,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [0 + 5, 0])

    return game_id, game


"""
player2 finishes road
player2 places meeple on a right road
player1 scores 2 points

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


def check_turn_6(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=-2, y=2),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.Right,
        featureType=FeatureType.Road,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [5 + 2, 0])

    return game_id, game


"""
player1 places meeple(!) on a city

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


def check_turn_7(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=1, y=-1),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.Right,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [7, 0])

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


def check_turn_8(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=-1, y=2),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [7, 0])

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


def check_turn_9(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=2, y=-1),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.TopRightEdge,
        featureType=FeatureType.Field,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [7 + 4, 0])

    return game_id, game


"""
player2 places meeple(@) on a city
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


def check_turn_A(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=0, y=2),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [11, 0])

    return game_id, game


"""
player1 places meeple(!) on a field
player1 closes road
player2 scores 4 points
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     !.|..               
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


def check_turn_B(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=0, y=3),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.BottomLeftEdge,
        featureType=FeatureType.Field,
    )
    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [11, 0 + 4])

    return game_id, game


"""
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     !.|..               
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


def check_turn_C(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=1, y=2),
        tile=game.current_tile,
        meepleType=MeepleType.NoneMeeple,
        side=Side.NoSide,
        featureType=FeatureType.NoneType,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [11, 4])

    return game_id, game


"""
player1 places meeple(!) on a city
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     !.|..               
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


def check_turn_D(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=2, y=2),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [11, 4])

    return game_id, game


"""
player2 places meeple(@) on a city
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     !.|..               
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


def check_turn_E(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=3, y=2),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [11, 4])

    return game_id, game


"""
player1 places meeple(!) on a field
player1 closes city
player2 scores 4 points

       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     !.|..               
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |               
           ..|...\ /                
1          --5....F !                
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


def check_turn_F(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=-1, y=1),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.Right,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [11, 4 + 4])

    return game_id, game


"""
player2 places meeple(@) on a field
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     !.|..              
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |                   |
           ..|...\ /                   /.
1          --5....F !                 G..
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


def check_turn_G(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=3, y=1),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.Field,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [11, 8])

    return game_id, game


"""
player1 places meeple(!) on a city 
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     !.|..               
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |                   |
           ..|...\ /                   /.
1          --5....F !                 G..
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


def check_turn_H(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=-3, y=0),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.Bottom,
        featureType=FeatureType.City,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [11, 8])

    return game_id, game


"""
player2 places meeple(@) on a road


       -3   -2   -1    0    1    2    3




4                           


                     |   |..@.-              
                     .\ /...|/             
3                    --B----I             
                     ..|.../             
                     !.|..|              
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |                   |
           ..|...\ /                   /.
1          --5....F !                 G..
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


def check_turn_I(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=1, y=3),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.Top,
        featureType=FeatureType.Road,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [11, 8])

    return game_id, game


"""
player1 places meeple(!) on a road

       -3   -2   -1    0    1    2    3




4                           


                     |   |..@.-     -.|..    
                     .\ /...|/       \|..  
3                    --B----I         J-! 
                     ..|.../         * \.
                     !.|..|             |
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |                   |
           ..|...\ /                   /.
1          --5....F !                 G..
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


def check_turn_J(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=3, y=3),
        tile=game.current_tile,
        meepleType=MeepleType.NormalMeeple,
        side=Side.Right,
        featureType=FeatureType.Road,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [11, 8])

    return game_id, game


"""
       -3   -2   -1    0    1    2    3




4                           


                     |   |..@.--...--.|..    
                     .\ /...|/  \./  \|..  
3                    --B----I    K    J-! 
                     ..|.../         * \.
                     !.|..|             |
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |                   |
           ..|...\ /                   /.
1          --5....F !                 G..
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


def check_turn_K(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=2, y=3),
        tile=game.current_tile,
        meepleType=MeepleType.NoneMeeple,
        side=Side.NoSide,
        featureType=FeatureType.NoneType,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [11, 8])

    return game_id, game


"""
       -3   -2   -1    0    1    2    3




4                           


                     |   |..@.--...--.|..    
                     .\ /...|/  \./  \|..  
3                    --B----I    K    J-! 
                     ..|.../         * \.
                     !.|..|             |
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |                   |
           ..|...\ /        *          /.
1          --5....F !       L         G..
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


def check_turn_L(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=1, y=1),
        tile=game.current_tile,
        meepleType=MeepleType.NoneMeeple,
        side=Side.NoSide,
        featureType=FeatureType.NoneType,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [11, 8])

    return game_id, game


"""
       -3   -2   -1    0    1    2    3


                          
                          
4                           


                     |   |..@.--...--.|..    
                     .\ /...|/  \./  \|..  
3                    --B----I    K    J-! 
                     ..|.../         * \.
                     !.|..|             |
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |                   |
           ..|...\ /        *          /.
1          --5....F !       L    M    G..
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


def check_turn_M(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=2, y=1),
        tile=game.current_tile,
        meepleType=MeepleType.NoneMeeple,
        side=Side.NoSide,
        featureType=FeatureType.NoneType,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [11, 8])

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
                     !.|..|             |
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |                   |
           ..|...\ /        *          /.
1          --5....F !       L    M    G..
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


def check_turn_N(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=0, y=4),
        tile=game.current_tile,
        meepleType=MeepleType.NoneMeeple,
        side=Side.NoSide,
        featureType=FeatureType.NoneType,
    )

    game_id, game = make_turn(engine, game, game_id, turn_params)
    check_points(game, [11, 8])

    return game_id, game


"""
player2 closes city
both players score: 13*2 + 3*2 = 26+6=32 (huge city)
    14 tile city
    3 shield in a city

       -3   -2   -1    0    1    2    3


                     -.|.-
                      \|/ 
4                    * N    
                          
                                                      
                     |   |..@.--...--.|..    
                     .\ /...|/  \./  \|..  
3                    --B----I    K    J-! 
                     ..|.../         * \.
                     !.|..|             |
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   || @ ||   || ! || @ |
           ..|..|   |                   |
           ..|...\ /        *          /.
1          --5....F !  O    L    M    G..
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


def check_turn_O(engine: GameEngine, game, game_id) -> tuple[int, SerializedGame]:
    turn_params = TurnParams(
        pos=Position(x=0, y=1),
        tile=game.current_tile,
        meepleType=MeepleType.NoneMeeple,
        side=Side.NoSide,
        featureType=FeatureType.NoneType,
    )

    game_id, game = make_turn(
        engine, game, game_id, turn_params, {1: 11 + 34 + 11, 2: 8 + 34 + 22}
    )
    check_points(game, [11 + 34, 8 + 34])

    return game_id, game


"""
player1 scores 11 points in total:
    - 1+1 points for a city with shield on H tile
    - 5 points for monastery on 1 tile
    - 3 points for farmer on 9 tile    
    - 0 points for farmer on B tile
    - 1 point for road on J tile
    
player2 scores 22 points in total:   
    - 3*3points for farmer on 4 tile
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
                     !.|..|             |
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   ||   ||   ||   ||   |
           ..|..|   |                   |
           ..|...\ /        *          /.
1          --5....F    O    L    M    G..
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
