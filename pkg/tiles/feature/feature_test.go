package feature

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

func TestFeatureType(t *testing.T) {
	cityFeature := New(side.Right|side.Top, City, true)
	fieldFeature := New(side.All, Field, false)
	roadFeature := New(side.Right|side.Bottom, Road, false)
	monasteryFeature := New(side.NoSide, Monastery, false)

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
	cityFeature1 := New(side.Right|side.Top, City, true)
	cityFeature2 := New(side.Right|side.Top, City, false)

	if !cityFeature1.HasModifier() {
		t.Fatalf("expected: %#v\ngot: %#v", true, cityFeature1.HasModifier())
	}
	if cityFeature1.HasModifier() {
		t.Fatalf("expected: %#v\ngot: %#v", false, cityFeature2.HasModifier())
	}
}

func TestMeeples(t *testing.T) {
	feature1 := New(side.Right|side.Top, City, true)
	feature2 := NewWithMeeple(side.Right|side.Top, City, true, 0)
	feature3 := NewWithMeeple(side.Right|side.Top, City, false, 1)
	feature4 := NewWithMeeple(side.Right|side.Top, City, false, 2)

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
