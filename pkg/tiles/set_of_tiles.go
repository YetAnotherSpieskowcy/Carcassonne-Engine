package tiles

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/buildings"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/connection"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/farm_connection"
)

func GetStandardTiles() []Tile {
	var tiles []Tile
	//from wikipedia from left to right, up to down

	//monastery without roads
	for i := 0; i < 5; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{}},
				Roads{[]connection.Connection{}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.TOP_LEFT,
							farm_connection.TOP_RIGHT,

							farm_connection.RIGHT_TOP,
							farm_connection.RIGHT_BOTTOM,

							farm_connection.LEFT_TOP,
							farm_connection.LEFT_BOTTOM,

							farm_connection.BOTTOM_LEFT,
							farm_connection.BOTTOM_RIGHT,
						},
					},
				}},
				false,
				buildings.MONASTERY,
			})
	}

	//monastery with single road
	for i := 0; i < 2; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{}},
				Roads{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.CENTER,
						connection.BOTTOM,
					},
				}}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.TOP_LEFT,
							farm_connection.TOP_RIGHT,

							farm_connection.RIGHT_TOP,
							farm_connection.RIGHT_BOTTOM,

							farm_connection.LEFT_TOP,
							farm_connection.LEFT_BOTTOM,

							farm_connection.BOTTOM_LEFT,
							farm_connection.BOTTOM_RIGHT,
						},
					},
				}},
				false,
				buildings.MONASTERY,
			})
	}

	//straight roads
	for i := 0; i < 8; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{}},
				Roads{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.LEFT,
						connection.RIGHT,
					},
				}}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_BOTTOM,
							farm_connection.BOTTOM_LEFT,
							farm_connection.BOTTOM_RIGHT,
							farm_connection.RIGHT_BOTTOM,
						},
					},
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_TOP,
							farm_connection.TOP_LEFT,
							farm_connection.TOP_RIGHT,
							farm_connection.RIGHT_TOP,
						},
					},
				}},
				false,
				buildings.NONE_BULDING,
			})
	}

	//roads turns
	for i := 0; i < 9; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{}},
				Roads{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.LEFT,
						connection.BOTTOM,
					},
				}}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_BOTTOM,
							farm_connection.BOTTOM_LEFT,
						},
					},
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_TOP,
							farm_connection.TOP_LEFT,
							farm_connection.TOP_RIGHT,
							farm_connection.RIGHT_TOP,
							farm_connection.RIGHT_BOTTOM,
							farm_connection.BOTTOM_RIGHT,
						},
					},
				}},
				false,
				buildings.NONE_BULDING,
			})
	}

	//T cross
	for i := 0; i < 9; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{}},
				Roads{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.LEFT,
						connection.CENTER,
					},
				},
					connection.Connection{
						[]connection.Side{
							connection.BOTTOM,
							connection.CENTER,
						},
					},

					connection.Connection{
						[]connection.Side{
							connection.RIGHT,
							connection.CENTER,
						},
					},
				}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_BOTTOM,
							farm_connection.BOTTOM_LEFT,
						},
					},
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.RIGHT_BOTTOM,
							farm_connection.BOTTOM_RIGHT,
						},
					},
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_TOP,
							farm_connection.TOP_LEFT,
							farm_connection.TOP_RIGHT,
							farm_connection.RIGHT_TOP,
						},
					},
				}},
				false,
				buildings.NONE_BULDING,
			})
	}

	// + cross
	for i := 0; i < 1; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{}},
				Roads{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.LEFT,
						connection.CENTER,
					},
				},
					connection.Connection{
						[]connection.Side{
							connection.BOTTOM,
							connection.CENTER,
						},
					},

					connection.Connection{
						[]connection.Side{
							connection.RIGHT,
							connection.CENTER,
						},
					},
					connection.Connection{
						[]connection.Side{
							connection.TOP,
							connection.CENTER,
						},
					},
				}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_BOTTOM,
							farm_connection.BOTTOM_LEFT,
						},
					},
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.RIGHT_BOTTOM,
							farm_connection.BOTTOM_RIGHT,
						},
					},
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_TOP,
							farm_connection.TOP_LEFT,
						},
					},

					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.TOP_RIGHT,
							farm_connection.RIGHT_TOP,
						},
					},
				}},
				false,
				buildings.NONE_BULDING,
			})
	}

	//1 city edge no roads
	for i := 0; i < 5; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.TOP,
					},
				}}},
				Roads{[]connection.Connection{}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_TOP,
							farm_connection.RIGHT_TOP,
							farm_connection.RIGHT_BOTTOM,
							farm_connection.BOTTOM_RIGHT,
							farm_connection.LEFT_BOTTOM,
							farm_connection.BOTTOM_LEFT,
						},
					},
				}},
				false,
				buildings.NONE_BULDING,
			})
	}

	//1 city edge straight road
	for i := 0; i < 4; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.TOP,
					},
				}}},
				Roads{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.RIGHT,
						connection.LEFT,
					},
				}}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.RIGHT_BOTTOM,
							farm_connection.BOTTOM_RIGHT,
							farm_connection.LEFT_BOTTOM,
							farm_connection.BOTTOM_LEFT,
						},
					},
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_TOP,
							farm_connection.RIGHT_TOP,
						},
					},
				}},
				false,
				buildings.NONE_BULDING,
			})
	}

	//1 city edge -| turn
	for i := 0; i < 4; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.TOP,
					},
				}}},
				Roads{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.LEFT,
						connection.BOTTOM,
					},
				}}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.RIGHT_BOTTOM,
							farm_connection.BOTTOM_RIGHT,
							farm_connection.LEFT_TOP,
							farm_connection.RIGHT_TOP,
						},
					},
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.BOTTOM_LEFT,
							farm_connection.LEFT_BOTTOM,
						},
					},
				}},
				false,
				buildings.NONE_BULDING,
			})
	}

	//1 city edge |- turn
	for i := 0; i < 4; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.TOP,
					},
				}}},
				Roads{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.RIGHT,
						connection.BOTTOM,
					},
				}}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_TOP,
							farm_connection.RIGHT_TOP,
							farm_connection.BOTTOM_LEFT,
							farm_connection.LEFT_BOTTOM,
						},
					},
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.RIGHT_BOTTOM,
							farm_connection.BOTTOM_RIGHT,
						},
					},
				}},
				false,
				buildings.NONE_BULDING,
			})
	}

	//1 city edge, road cross
	for i := 0; i < 4; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.TOP,
					},
				}}},
				Roads{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.RIGHT,
						connection.CENTER,
					},
				},
					connection.Connection{
						[]connection.Side{
							connection.LEFT,
							connection.CENTER,
						},
					},
					connection.Connection{
						[]connection.Side{
							connection.BOTTOM,
							connection.CENTER,
						},
					}}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_TOP,
							farm_connection.RIGHT_TOP,
						},
					},
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.RIGHT_BOTTOM,
							farm_connection.BOTTOM_RIGHT,
						},
					},
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.BOTTOM_LEFT,
							farm_connection.LEFT_BOTTOM,
						},
					},
				}},
				false,
				buildings.NONE_BULDING,
			})
	}

	//2 city edges (up and down)
	for i := 0; i < 3; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.TOP,
					},
				},
					connection.Connection{
						[]connection.Side{
							connection.BOTTOM,
						},
					}}},
				Roads{[]connection.Connection{}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_TOP,
							farm_connection.RIGHT_TOP,
							farm_connection.LEFT_BOTTOM,
							farm_connection.RIGHT_BOTTOM,
						},
					},
				}},
				false,
				buildings.NONE_BULDING,
			})
	}

	//2 city edges (up and right)
	for i := 0; i < 2; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.TOP,
					},
				},
					connection.Connection{
						[]connection.Side{
							connection.RIGHT,
						},
					}}},
				Roads{[]connection.Connection{}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_TOP,
							farm_connection.LEFT_BOTTOM,
							farm_connection.BOTTOM_LEFT,
							farm_connection.BOTTOM_RIGHT,
						},
					},
				}},
				false,
				buildings.NONE_BULDING,
			})
	}

	//2 city edges (left and right but connected)
	for i := 0; i < 1; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.TOP,
						connection.RIGHT,
						connection.CENTER,
					},
				},
				}},
				Roads{[]connection.Connection{}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_TOP,
							farm_connection.LEFT_BOTTOM,
							farm_connection.BOTTOM_LEFT,
							farm_connection.BOTTOM_RIGHT,
						},
					},
				}},
				false,
				buildings.NONE_BULDING,
			})
	}

	//2 city edges (left and right but connected but also shields)
	for i := 0; i < 2; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.TOP,
						connection.RIGHT,
						connection.CENTER,
					},
				},
				}},
				Roads{[]connection.Connection{}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_TOP,
							farm_connection.LEFT_BOTTOM,
							farm_connection.BOTTOM_LEFT,
							farm_connection.BOTTOM_RIGHT,
						},
					},
				}},
				true,
				buildings.NONE_BULDING,
			})
	}

	//2 city edges (up and right but connected)
	for i := 0; i < 3; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.TOP,
						connection.RIGHT,
					},
				},
				}},
				Roads{[]connection.Connection{}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_TOP,
							farm_connection.LEFT_BOTTOM,
							farm_connection.BOTTOM_LEFT,
							farm_connection.BOTTOM_RIGHT,
						},
					},
				}},
				false,
				buildings.NONE_BULDING,
			})
	}

	//2 city edges (up and right but connected but with shield)
	for i := 0; i < 2; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{connection.Connection{
					[]connection.Side{
						connection.TOP,
						connection.RIGHT,
					},
				},
				}},
				Roads{[]connection.Connection{}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_TOP,
							farm_connection.LEFT_BOTTOM,
							farm_connection.BOTTOM_LEFT,
							farm_connection.BOTTOM_RIGHT,
						},
					},
				}},
				true,
				buildings.NONE_BULDING,
			})
	}

	//2 city edges (up and right but connected but road)
	for i := 0; i < 3; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{
					connection.Connection{
						[]connection.Side{
							connection.TOP,
							connection.RIGHT,
						},
					},
				}},
				Roads{[]connection.Connection{
					connection.Connection{
						[]connection.Side{
							connection.LEFT,
							connection.BOTTOM,
						},
					},
				}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_BOTTOM,
							farm_connection.BOTTOM_LEFT,
						},
					},
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_TOP,
							farm_connection.BOTTOM_RIGHT,
						},
					},
				}},
				false,
				buildings.NONE_BULDING,
			})
	}

	//2 city edges (up and right but connected but road but shield)
	for i := 0; i < 2; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{
					connection.Connection{
						[]connection.Side{
							connection.TOP,
							connection.RIGHT,
						},
					},
				}},
				Roads{[]connection.Connection{
					connection.Connection{
						[]connection.Side{
							connection.LEFT,
							connection.BOTTOM,
						},
					},
				}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_BOTTOM,
							farm_connection.BOTTOM_LEFT,
						},
					},
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_TOP,
							farm_connection.BOTTOM_RIGHT,
						},
					},
				}},
				true,
				buildings.NONE_BULDING,
			})
	}

	//3 city edges ( but connected)
	for i := 0; i < 3; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{
					connection.Connection{
						[]connection.Side{
							connection.TOP,
							connection.RIGHT,
							connection.CENTER,
							connection.LEFT,
						},
					},
				}},
				Roads{[]connection.Connection{}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_BOTTOM,
							farm_connection.BOTTOM_LEFT,
						},
					},
				}},
				false,
				buildings.NONE_BULDING,
			})
	}

	//3 city edges (but connected but shield)
	for i := 0; i < 1; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{
					connection.Connection{
						[]connection.Side{
							connection.TOP,
							connection.RIGHT,
							connection.CENTER,
							connection.LEFT,
						},
					},
				}},
				Roads{[]connection.Connection{}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_BOTTOM,
							farm_connection.BOTTOM_LEFT,
						},
					},
				}},
				true,
				buildings.NONE_BULDING,
			})
	}

	//3 city edges (but connected but road)
	for i := 0; i < 1; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{
					connection.Connection{
						[]connection.Side{
							connection.TOP,
							connection.RIGHT,
							connection.CENTER,
							connection.LEFT,
						},
					},
				}},
				Roads{[]connection.Connection{
					connection.Connection{
						[]connection.Side{
							connection.CENTER,
							connection.BOTTOM,
						},
					},
				}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_BOTTOM,
						},
					},
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.BOTTOM_LEFT,
						},
					},
				}},
				false,
				buildings.NONE_BULDING,
			})
	}

	//3 city edges (but connected but road but shield)
	for i := 0; i < 2; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{
					connection.Connection{
						[]connection.Side{
							connection.TOP,
							connection.RIGHT,
							connection.CENTER,
							connection.LEFT,
						},
					},
				}},
				Roads{[]connection.Connection{
					connection.Connection{
						[]connection.Side{
							connection.CENTER,
							connection.BOTTOM,
						},
					},
				}},
				Fields{[]farm_connection.FarmConnection{
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.LEFT_BOTTOM,
						},
					},
					farm_connection.FarmConnection{
						[]farm_connection.FarmSide{
							farm_connection.BOTTOM_LEFT,
						},
					},
				}},
				true,
				buildings.NONE_BULDING,
			})
	}

	//4 city edges (but shield)
	for i := 0; i < 1; i++ {
		tiles = append(tiles,
			Tile{
				Cities{[]connection.Connection{
					connection.Connection{
						[]connection.Side{
							connection.TOP,
							connection.RIGHT,
							connection.CENTER,
							connection.LEFT,
							connection.BOTTOM,
						},
					},
				}},
				Roads{[]connection.Connection{}},
				Fields{[]farm_connection.FarmConnection{}},
				true,
				buildings.NONE_BULDING,
			})
	}
	return tiles
}
