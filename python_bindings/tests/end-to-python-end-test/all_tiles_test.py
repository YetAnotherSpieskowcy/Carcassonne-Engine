''''''
'''
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
               
'''
import logging
from pathlib import Path

from carcassonne_engine import GameEngine, SerializedGame
from carcassonne_engine import tiletemplates
from carcassonne_engine.models import Position
from carcassonne_engine.tilesets import mini_tile_set
from utils import make_turn, TurnParams

log = logging.getLogger(__name__)


def test_four_player_game(tmp_path: Path) -> None:
    engine = GameEngine(4, tmp_path)
    tile_set = mini_tile_set()

    game_id, game = engine.generate_ordered_game(tile_set)

    game_id, game = check_first_turn(engine, game, game_id)

    assert game.current_tile is None


'''
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
                     .[ ].  
-1                   .[1].       
                     .[ ]. 
                     .....

'''


def check_turn_1(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=0, y=-1),
        tile=tiletemplates.monastery_without_roads())

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           




3                                    




2          




1                              


                     |   |.....
                     .\ /..[ ].
0                    --0---[2].
                     ......[ ].
                     ..........
                     .....
                     .[ ].  
-1                   .[1].       
                     .[ ]. 
                     .....

'''


def check_turn_2(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=1, y=0),
        tile=tiletemplates.monastery_with_single_road())

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           




3                                    




2          




1                              


                .....|   |.....
                ......\ /..[ ].
0               --3----0---[2].
                ...........[ ].
                ...............
                     .....
                     .[ ].  
-1                   .[1].       
                     .[ ]. 
                     .....

'''


def check_turn_3(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=-1, y=0),
        tile=tiletemplates.straight_roads(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           




3                                    




2          


                               
                               
1                              
                               
                               
           ..|.......|   |.....
           ..|........\ /..[ ].
0          ..4----3----0---[2].
           ................[ ].
           ....................
                     .....
                     .[ ].  
-1                   .[1].       
                     .[ ]. 
                     .....

'''


def check_turn_4(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=-2, y=0),
        tile=tiletemplates.roads_turn(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           




3                                    


           
           
2          
           
           
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[ ].
0          ..4----3----0---[2].
           ................[ ].
           ....................
                     .....
                     .[ ].  
-1                   .[1].       
                     .[ ]. 
                     .....

'''


def check_turn_5(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=-2, y=1),
        tile=tiletemplates.t_cross_road(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           




3                                    


           ..|..
           ..|..
2          --6--
           ..|..
           ..|..
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[ ].
0          ..4----3----0---[2].
           ................[ ].
           ....................
                     .....
                     .[ ].  
-1                   .[1].       
                     .[ ]. 
                     .....

'''


def check_turn_6(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=-2, y=2),
        tile=tiletemplates.x_cross_road(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           




3                                    


           ..|..
           ..|..
2          --6--
           ..|..
           ..|..
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[ ].
0          ..4----3----0---[2].
           ................[ ].
           ....................
                     .........-
                     .[ ]..../  
-1                   .[1]...7       
                     .[ ]....\ 
                     .........-

'''


def check_turn_7(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=2, y=-1),
        tile=tiletemplates.single_city_edge_no_roads(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           




3                                    


           ..|.......
           ..|.......
2          --6----8--
           ..|.../ \.
           ..|..|   |
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[ ].
0          ..4----3----0---[2].
           ................[ ].
           ....................
                     .........-
                     .[ ]..../  
-1                   .[1]...7       
                     .[ ]....\ 
                     .........-

'''


def check_turn_8(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=-1, y=2),
        tile=tiletemplates.single_city_edge_straight_roads(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           




3                                    


           ..|.......
           ..|.......
2          --6----8--
           ..|.../ \.
           ..|..|   |
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[ ].
0          ..4----3----0---[2].
           ................[ ].
           ....................
                     .........--....
                     .[ ]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

'''


def check_turn_9(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=2, y=-1),
        tile=tiletemplates.single_city_edge_left_road_turn(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           


                                        
                                      
3                                    
                                    
                                    
           ..|.........|..
           ..|.........|..
2          --6----8----A..
           ..|.../ \../ \.
           ..|..|   ||   |
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[ ].
0          ..4----3----0---[2].
           ................[ ].
           ....................
                     .........--....
                     .[ ]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

'''


def check_turn_A(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=0, y=2),
        tile=tiletemplates.single_city_edge_right_road_turn(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     ..|..               
           ..|.........|..
           ..|.........|..
2          --6----8----A..
           ..|.../ \../ \.
           ..|..|   ||   |
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[ ].
0          ..4----3----0---[2].
           ................[ ].
           ....................
                     .........--....
                     .[ ]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

'''


def check_turn_B(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=0, y=3),
        tile=tiletemplates.single_city_edge_cross_road(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     ..|..               
           ..|.........|..|   |
           ..|.........|...\ /.
2          --6----8----A....C..
           ..|.../ \../ \../ \.
           ..|..|   ||   ||   |
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[ ].
0          ..4----3----0---[2].
           ................[ ].
           ....................
                     .........--....
                     .[ ]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

'''


def check_turn_C(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=1, y=2),
        tile=tiletemplates.two_city_edges_up_and_down_not_connected(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     ..|..               
           ..|.........|..|   ||   |
           ..|.........|...\ /..\ /.
2          --6----8----A....C...|D|.
           ..|.../ \../ \../ \../ \.
           ..|..|   ||   ||   ||   |
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[ ].
0          ..4----3----0---[2].
           ................[ ].
           ....................
                     .........--....
                     .[ ]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

'''


def check_turn_D(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=2, y=2),
        tile=tiletemplates.two_city_edges_up_and_down_connected(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     ..|..               
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   ||   ||   ||   ||   |
           ..|..                    
           ..|..                    
1          --5..                    
           ..|..                    
           ..|..                    
           ..|.......|   |.....
           ..|........\ /..[ ].
0          ..4----3----0---[2].
           ................[ ].
           ....................
                     .........--....
                     .[ ]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

'''


def check_turn_E(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=3, y=2),
        tile=tiletemplates.two_city_edges_up_and_down_connected_shield(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     ..|..               
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   ||   ||   ||   ||   |
           ..|..|   |               
           ..|...\ /                
1          --5....F                 
           ..|.....\                
           ..|......-               
           ..|.......|   |.....
           ..|........\ /..[ ].
0          ..4----3----0---[2].
           ................[ ].
           ....................
                     .........--....
                     .[ ]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

'''


def check_turn_F(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=-1, y=1),
        tile=tiletemplates.two_city_edges_corner_not_connected(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     ..|..               
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   ||   ||   ||   ||   |
           ..|..|   |                   |
           ..|...\ /                   /.
1          --5....F                   G..
           ..|.....\                 /...
           ..|......-               -....
           ..|.......|   |.....
           ..|........\ /..[ ].
0          ..4----3----0---[2].
           ................[ ].
           ....................
                     .........--....
                     .[ ]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

'''


def check_turn_G(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=3, y=1),
        tile=tiletemplates.two_city_edges_corner_connected(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           


                     |   |                   
                     .\ /.                 
3                    --B--                
                     ..|..               
                     ..|..               
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   ||   ||   ||   ||   |
           ..|..|   |                   |
           ..|...\ /                   /.
1          --5....F                   G..
           ..|.....\                 /...
           ..|......-               -....
      -......|.......|   |.....
       \.....|........\ /..[ ].
0     * H....4----3----0---[2].
         \.................[ ].
          |....................
                     .........--....
                     .[ ]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

'''


def check_turn_H(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=-3, y=0),
        tile=tiletemplates.two_city_edges_corner_connected_shield(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           


                     |   |..|.-              
                     .\ /...|/             
3                    --B----I             
                     ..|.../             
                     ..|..|              
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   ||   ||   ||   ||   |
           ..|..|   |                   |
           ..|...\ /                   /.
1          --5....F                   G..
           ..|.....\                 /...
           ..|......-               -....
      -......|.......|   |.....
       \.....|........\ /..[ ].
0    *  H....4----3----0---[2].
         \.................[ ].
          |....................
                     .........--....
                     .[ ]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

'''


def check_turn_I(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=1, y=3),
        tile=tiletemplates.two_city_edges_corner_connected_road_turn(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           


                     |   |..|.-     -.|..    
                     .\ /...|/       \|..  
3                    --B----I         J-- 
                     ..|.../         * \.
                     ..|..|             |
           ..|.........|..|   ||   || * |
           ..|.........|...\ /..\ /..\ /.
2          --6----8----A....C...|D|..|E|.
           ..|.../ \../ \../ \../ \../ \.
           ..|..|   ||   ||   ||   ||   |
           ..|..|   |                   |
           ..|...\ /                   /.
1          --5....F                   G..
           ..|.....\                 /...
           ..|......-               -....
      -......|.......|   |.....
       \.....|........\ /..[ ].
0     * H....4----3----0---[2].
         \.................[ ].
          |....................
                     .........--....
                     .[ ]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

'''


def check_turn_J(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=3, y=3),
        tile=tiletemplates.two_city_edges_corner_connected_road_turn_shield(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           


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
           ..|...\ /                   /.
1          --5....F                   G..
           ..|.....\                 /...
           ..|......-               -....
      -......|.......|   |.....
       \.....|........\ /..[ ].
0     * H....4----3----0---[2].
         \.................[ ].
          |....................
                     .........--....
                     .[ ]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

'''


def check_turn_K(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=2, y=3),
        tile=tiletemplates.three_city_edges_connected(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3




4                           


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
1          --5....F         L         G..
           ..|.....\       /.\       /...
           ..|......-     -...-     -....
      -......|.......|   |.....
       \.....|........\ /..[ ].
0     * H....4----3----0---[2].
         \.................[ ].
          |....................
                     .........--....
                     .[ ]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

'''


def check_turn_L(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=1, y=1),
        tile=tiletemplates.three_city_edges_connected_shield(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3


                          
                          
4                           


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
1          --5....F         L    M    G..
           ..|.....\       /.\  /|\  /...
           ..|......-     -...--.|.--....
      -......|.......|   |.....
       \.....|........\ /..[ ].
0     * H....4----3----0---[2].
         \.................[ ].
          |....................
                     .........--....
                     .[ ]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

'''


def check_turn_M(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=2, y=1),
        tile=tiletemplates.three_city_edges_connected_road(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3


                     -.|.-
                      \|/ 
4                    * N    


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
1          --5....F         L    M    G..
           ..|.....\       /.\  /|\  /...
           ..|......-     -...--.|.--....
      -......|.......|   |.....
       \.....|........\ /..[ ].
0     * H....4----3----0---[2].
         \.................[ ].
          |....................
                     .........--....
                     .[ ]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..

'''


def check_turn_N(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=0, y=4),
        tile=tiletemplates.three_city_edges_connected_road_shield(),
        )

    return make_turn(engine, game, game_id, turn_params)


'''
       -3   -2   -1    0    1    2    3


                     -.|.-
                      \|/ 
4                    * N    
                          
                                                      
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
0    *  H....4----3----0---[2].
         \.................[ ].
          |....................
                     .........--....
                     .[ ]..../  \... 
-1                   .[1]...7    9--     
                     .[ ]....\  /|..
                     .........--.|..
               
'''


def check_turn_O(engine: GameEngine, game, game_id) -> (int, SerializedGame):
    turn_params = TurnParams(
        pos=Position(x=0, y=1),
        tile=tiletemplates.four_city_edges_connected_shield(),
        )

    return make_turn(engine, game, game_id, turn_params)
