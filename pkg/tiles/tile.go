package tiles

import (
	connection "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/connection"
	farm_connection "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/farm_connection"
)

type Tile struct {
	Cities 
	Roads  
	Fields 
	HasShield bool
	Building   Bulding

	//dać Building po prostu by skomponować

	//not sure how to include undefied/null?
	//meeple    Meeple
}

func (tile *Tile) Rotate(rotations int) Tile {
	var t Tile
	//rotate cities	
	for _, cityConnection := tile.Cities{
		append(t.cities, cityConnection.Rotate(rotations))
	}

	//rotate roads
	for _, road := tile.roads{
		append(t.roads, road.Rotate(rotations))
	}

	//rotate fields
	for _, field := tile.field{
		append(t.fields, field.Rotate(rotations))
	}

	//copy parameters
	t.hasShield = tile.hasShield
	t.Building = tile.Building
	return t	
}


