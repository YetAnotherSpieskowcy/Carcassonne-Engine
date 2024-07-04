package elements

import (
	"slices"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

func TestPositionMarshalTextWithPositiveCoords(t *testing.T) {
	pos := NewPosition(1, 3)
	expected := []byte("1,3")
	actual, err := pos.MarshalText()
	if err != nil {
		t.Fatal(err.Error())
	}
	if !slices.Equal(actual, expected) {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestPositionMarshalTextWithNegativeCoords(t *testing.T) {
	pos := NewPosition(-31, -5)
	expected := []byte("-31,-5")
	actual, err := pos.MarshalText()
	if err != nil {
		t.Fatal(err.Error())
	}
	if !slices.Equal(actual, expected) {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestPositionUnmarshalTextWithPositiveCoords(t *testing.T) {
	text := []byte("1,3")
	expectedX := int16(1)
	expectedY := int16(3)

	actual := Position{}
	err := actual.UnmarshalText(text)
	if err != nil {
		t.Fatal(err.Error())
	}
	if actual.X() != expectedX {
		t.Fatalf("expected %#v, got %#v instead", expectedX, actual)
	}
	if actual.Y() != expectedY {
		t.Fatalf("expected %#v, got %#v instead", expectedY, actual)
	}
}

func TestPositionUnmarshalTextWithNegativeCoords(t *testing.T) {
	text := []byte("-31,-5")
	expectedX := int16(-31)
	expectedY := int16(-5)

	actual := Position{}
	err := actual.UnmarshalText(text)
	if err != nil {
		t.Fatal(err.Error())
	}
	if actual.X() != expectedX {
		t.Fatalf("expected %#v, got %#v instead", expectedX, actual)
	}
	if actual.Y() != expectedY {
		t.Fatalf("expected %#v, got %#v instead", expectedY, actual)
	}
}

func TestTilePlacementRotate(t *testing.T) {
	move := ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected TilePlacement.Rotate() to panic")
		}
	}()

	move.Rotate(1)
}

func TestPlacedTileFeatureGet(t *testing.T) {
	move := ToPlacedTile(tiletemplates.MonasteryWithSingleRoad())
	move.Monastery().Meeple.MeepleType = NormalMeeple
	move.Monastery().Meeple.PlayerID = 1

	expectedMonastery := tiletemplates.MonasteryWithSingleRoad().Monastery()

	if move.Monastery().Feature != *expectedMonastery {
		t.Fatalf("got\n %#v \nshould be \n%#v", move.Monastery().Feature, *expectedMonastery)
	}
	if move.Monastery().Meeple.MeepleType != NormalMeeple {
		t.Fatalf("got\n %#v \nshould be \n%#v", move.Monastery().Meeple.MeepleType, NormalMeeple)
	}
	if MeepleType(move.Monastery().Meeple.PlayerID) != 1 {
		t.Fatalf("got\n %#v \nshould be \n%#v", move.Monastery().Meeple.PlayerID, 1)
	}
}

func TestGetCityFeatures(t *testing.T) {
	var expectedLen int = 1
	var expectedSide side.Side = side.Top

	tile := ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())

	cityFeatures := tile.GetCityFeatures()

	if len(cityFeatures) != expectedLen {
		t.Fatalf("expected %#v, got %#v instead", expectedLen, len(cityFeatures))
	}

	actualSide := cityFeatures[0].Sides
	if actualSide != expectedSide {
		t.Fatalf("expected side %#v, got %#v instead", expectedSide, actualSide)
	}
}
