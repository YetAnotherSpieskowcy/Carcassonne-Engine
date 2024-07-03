package city

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

func TestGetNeighbouringPositions(t *testing.T) {
	positions := GetNeighbouringPositions(elements.NewPosition(1, 1))

	topPosition := positions[side.Top]
	if topPosition.X() != 1 || topPosition.Y() != 2 {
		t.Fatalf("expected x=%#v y=%#v, got x=%#v y=%#v instead", 1, 2, topPosition.X(), topPosition.Y())
	}
}

func TestUpdateCitiesWhenNoCities(t *testing.T) {

}
