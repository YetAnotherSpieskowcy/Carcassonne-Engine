package game

import (
	"fmt"
	"slices"
	"testing"
)


func TestPositionMarshalTextWithPositiveCoords(t *testing.T) {
	pos := Position{1, 3}
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
	pos := Position{-31, -5}
	expected := []byte("-31,-5")
	actual, err := pos.MarshalText()
	if err != nil {
		t.Fatal(err.Error())
	}
	if !slices.Equal(actual, expected) {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}
