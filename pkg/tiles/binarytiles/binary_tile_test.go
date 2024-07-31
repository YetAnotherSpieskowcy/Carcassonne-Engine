package binarytiles

import (
	"fmt"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/test"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature/modifier"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

func TestFromFeatures(t *testing.T) {
	tile := test.GetTestCustomPlacedTile(tiletemplates.TwoCityEdgesCornerConnectedRoadTurn())

	binaryTile := FromPlacedTile(tile)
	fmt.Printf("%064b\n", binaryTile)

	city := tile.GetFeaturesOfType(feature.City)[0]

	city.ModifierType = modifier.Shield

	city.Meeple.PlayerID = 5
	city.Meeple.Type = elements.NormalMeeple

	binaryTile = FromPlacedTile(tile)
	fmt.Printf("%064b\n", binaryTile)

	tile.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple =
		elements.Meeple{PlayerID: 1, Type: elements.NormalMeeple}
	tile.GetPlacedFeatureAtSide(side.Top, feature.City).ModifierType = modifier.Shield

	binaryTile = FromPlacedTile(tile)
	fmt.Printf("%064b", binaryTile)

}
