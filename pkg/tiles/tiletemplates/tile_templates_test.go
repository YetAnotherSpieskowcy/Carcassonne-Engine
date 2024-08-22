package tiletemplates_test

import (
	"reflect"
	"runtime"
	"slices"
	"strings"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

func TestTileTemplateSidesOverlapInValidWays(t *testing.T) {
	tiles := []func() tiles.Tile{
		tiletemplates.MonasteryWithoutRoads,
		tiletemplates.MonasteryWithSingleRoad,
		tiletemplates.StraightRoads,
		tiletemplates.RoadsTurn,
		tiletemplates.TCrossRoad,
		tiletemplates.XCrossRoad,
		tiletemplates.SingleCityEdgeNoRoads,
		tiletemplates.SingleCityEdgeStraightRoads,
		tiletemplates.SingleCityEdgeLeftRoadTurn,
		tiletemplates.SingleCityEdgeRightRoadTurn,
		tiletemplates.SingleCityEdgeCrossRoad,
		tiletemplates.TwoCityEdgesUpAndDownNotConnected,
		tiletemplates.TwoCityEdgesCornerNotConnected,
		tiletemplates.TwoCityEdgesUpAndDownConnected,
		tiletemplates.TwoCityEdgesUpAndDownConnectedShield,
		tiletemplates.TwoCityEdgesCornerConnected,
		tiletemplates.TwoCityEdgesCornerConnectedShield,
		tiletemplates.TwoCityEdgesCornerConnectedRoadTurn,
		tiletemplates.TwoCityEdgesCornerConnectedRoadTurnShield,
		tiletemplates.ThreeCityEdgesConnected,
		tiletemplates.ThreeCityEdgesConnectedShield,
		tiletemplates.ThreeCityEdgesConnectedRoad,
		tiletemplates.ThreeCityEdgesConnectedRoadShield,
		tiletemplates.FourCityEdgesConnectedShield,
		tiletemplates.TestOnlyField,
	}
	validFeatureTypeCombinations := [][]feature.Type{
		{feature.Road, feature.Field},
	}
	for _, tileTemplateFunc := range tiles {
		funcNameParts := strings.Split(
			runtime.FuncForPC(reflect.ValueOf(tileTemplateFunc).Pointer()).Name(),
			".",
		)
		templateName := funcNameParts[len(funcNameParts)-1]

		t.Run(templateName, func(t *testing.T) {
			tile := tileTemplateFunc()
			featureTypesPerSide := map[side.Side][]feature.Type{}
			for _, tileFeature := range tile.Features {
				for _, side := range side.EdgeSides {
					if tileFeature.Sides.HasSide(side) {
						featureTypesPerSide[side] = append(
							featureTypesPerSide[side], tileFeature.FeatureType,
						)
					}
				}
			}

			for side, featureTypes := range featureTypesPerSide {
				if len(featureTypes) == 1 {
					continue
				}
				found := false
				for _, combination := range validFeatureTypeCombinations {
					found = found || slices.Equal(combination, featureTypes)
				}
				if !found {
					t.Errorf(
						"features on side %v overlap in unsupported way: %#v\n",
						side,
						featureTypes,
					)
				}
			}
		})
	}
}
