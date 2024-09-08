from ._bindings import (  # type: ignore[attr-defined] # no stubs
    tiletemplates as _go_tiletemplates,
)
from .models import Tile

__all__ = (
    "four_city_edges_connected_shield",
    "single_city_edge_left_road_turn",
    "single_city_edge_straight_roads",
    "three_city_edges_connected_road",
    "two_city_edges_corner_not_connected",
    "two_city_edges_corner_connected_shield",
    "two_city_edges_up_and_down_connected_shield",
    "single_city_edge_no_roads",
    "straight_roads",
    "t_cross_road",
    "test_only_field",
    "two_city_edges_corner_connected",
    "two_city_edges_corner_connected_road_turn_shield",
    "roads_turn",
    "three_city_edges_connected_road_shield",
    "three_city_edges_connected_shield",
    "two_city_edges_up_and_down_not_connected",
    "two_city_edges_up_and_down_connected",
    "x_cross_road",
    "monastery_with_single_road",
    "monastery_without_roads",
    "single_city_edge_cross_road",
    "single_city_edge_right_road_turn",
    "three_city_edges_connected",
    "two_city_edges_corner_connected_road_turn",
)


def four_city_edges_connected_shield() -> Tile:
    return Tile(_go_tiletemplates.FourCityEdgesConnectedShield())


def single_city_edge_left_road_turn() -> Tile:
    return Tile(_go_tiletemplates.SingleCityEdgeLeftRoadTurn())


def single_city_edge_straight_roads() -> Tile:
    return Tile(_go_tiletemplates.SingleCityEdgeStraightRoads())


def three_city_edges_connected_road() -> Tile:
    return Tile(_go_tiletemplates.ThreeCityEdgesConnectedRoad())


def two_city_edges_corner_not_connected() -> Tile:
    return Tile(_go_tiletemplates.TwoCityEdgesCornerNotConnected())


def two_city_edges_corner_connected_shield() -> Tile:
    return Tile(_go_tiletemplates.TwoCityEdgesCornerConnectedShield())


def two_city_edges_up_and_down_connected_shield() -> Tile:
    return Tile(_go_tiletemplates.TwoCityEdgesUpAndDownConnectedShield())


def single_city_edge_no_roads() -> Tile:
    return Tile(_go_tiletemplates.SingleCityEdgeNoRoads())


def straight_roads() -> Tile:
    return Tile(_go_tiletemplates.StraightRoads())


def t_cross_road() -> Tile:
    return Tile(_go_tiletemplates.TCrossRoad())


def test_only_field() -> Tile:
    return Tile(_go_tiletemplates.TestOnlyField())


def two_city_edges_corner_connected() -> Tile:
    return Tile(_go_tiletemplates.TwoCityEdgesCornerConnected())


def two_city_edges_corner_connected_road_turn_shield() -> Tile:
    return Tile(_go_tiletemplates.TwoCityEdgesCornerConnectedRoadTurnShield())


def roads_turn() -> Tile:
    return Tile(_go_tiletemplates.RoadsTurn())


def three_city_edges_connected_road_shield() -> Tile:
    return Tile(_go_tiletemplates.ThreeCityEdgesConnectedRoadShield())


def three_city_edges_connected_shield() -> Tile:
    return Tile(_go_tiletemplates.ThreeCityEdgesConnectedShield())


def two_city_edges_up_and_down_not_connected() -> Tile:
    return Tile(_go_tiletemplates.TwoCityEdgesUpAndDownNotConnected())


def two_city_edges_up_and_down_connected() -> Tile:
    return Tile(_go_tiletemplates.TwoCityEdgesUpAndDownConnected())


def x_cross_road() -> Tile:
    return Tile(_go_tiletemplates.XCrossRoad())


def monastery_with_single_road() -> Tile:
    return Tile(_go_tiletemplates.MonasteryWithSingleRoad())


def monastery_without_roads() -> Tile:
    return Tile(_go_tiletemplates.MonasteryWithoutRoads())


def single_city_edge_cross_road() -> Tile:
    return Tile(_go_tiletemplates.SingleCityEdgeCrossRoad())


def single_city_edge_right_road_turn() -> Tile:
    return Tile(_go_tiletemplates.SingleCityEdgeRightRoadTurn())


def three_city_edges_connected() -> Tile:
    return Tile(_go_tiletemplates.ThreeCityEdgesConnected())


def two_city_edges_corner_connected_road_turn() -> Tile:
    return Tile(_go_tiletemplates.TwoCityEdgesCornerConnectedRoadTurn())
