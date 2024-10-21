package feature

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature/modifier"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

func TestFeatureType(t *testing.T) {
	cityFeature := New(City, side.Right|side.Top, modifier.Shield)
	fieldFeature := New(Field, side.All)
	roadFeature := New(Road, side.Right|side.Bottom)
	monasteryFeature := New(Monastery, side.NoSide)

	if cityFeature.Type() != City {
		t.Fatalf("expected: %#v\ngot: %#v", City, cityFeature.Type())
	}
	if fieldFeature.Type() != Field {
		t.Fatalf("expected: %#v\ngot: %#v", City, fieldFeature.Type())
	}
	if roadFeature.Type() != Road {
		t.Fatalf("expected: %#v\ngot: %#v", City, roadFeature.Type())
	}
	if monasteryFeature.Type() != Monastery {
		t.Fatalf("expected: %#v\ngot: %#v", City, monasteryFeature.Type())
	}
}

func TestModifier(t *testing.T) {
	cityFeature1 := New(City, side.Right|side.Top, modifier.Shield)
	cityFeature2 := New(City, side.Right|side.Top)

	if !cityFeature1.HasModifier() {
		t.Fatalf("expected: %#v\ngot: %#v", true, cityFeature1.HasModifier())
	}
	if cityFeature2.HasModifier() {
		t.Fatalf("expected: %#v\ngot: %#v", false, cityFeature2.HasModifier())
	}
}

func TestMeeples(t *testing.T) {
	feature1 := New(City, side.Right|side.Top, modifier.Shield)
	feature2 := NewWithMeeple(City, side.Right|side.Top, 0, modifier.Shield)
	feature3 := NewWithMeeple(City, side.Right|side.Top, 1)
	feature4 := NewWithMeeple(City, side.Right|side.Top, 2)

	if feature1.HasMeeple() || feature1.OwnerID() != 0 {
		t.Fatalf("expected: %#v, %#v\ngot: %#v, %#v", false, 0, feature1.HasMeeple(), feature1.OwnerID())
	}
	if feature2.HasMeeple() || feature2.OwnerID() != 0 {
		t.Fatalf("expected: %#v, %#v\ngot: %#v, %#v", false, 0, feature2.HasMeeple(), feature2.OwnerID())
	}
	if !feature3.HasMeeple() || feature3.OwnerID() != 1 {
		t.Fatalf("expected: %#v, %#v\ngot: %#v, %#v", false, 0, feature3.HasMeeple(), feature3.OwnerID())
	}
	if !feature4.HasMeeple() || feature4.OwnerID() != 2 {
		t.Fatalf("expected: %#v, %#v\ngot: %#v, %#v", false, 0, feature4.HasMeeple(), feature4.OwnerID())
	}
}
