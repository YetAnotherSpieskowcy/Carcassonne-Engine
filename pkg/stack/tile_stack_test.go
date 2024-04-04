package stack

import (
	"testing"
)

func TestStandardOrder(t *testing.T) {
	tiles := []Tile{{0}, {1}, {2}, {3}}
	stack := New(tiles)
	for i := range len(tiles) {
		tile, err := stack.Next()
		if err != nil {
			t.Fatal(err.Error())
		}
		if tile != tiles[i] {
			t.Fail()
		}

	}
}

func TestSetSeed(t *testing.T) {
	tiles := []Tile{{0}, {1}, {2}, {3}}
	expectedOrder := []int32{2, 3, 0, 1}
	stack := NewSeeded(tiles, 42)
	stack.Shuffle()
	for i := range len(tiles) {
		tile, err := stack.Next()
		if err != nil {
			t.Fatal(err.Error())
		}
		if int32(tile.id) != expectedOrder[i] {
			t.Fail()
		}

	}
}

func TestPeek(t *testing.T) {
	tiles := []Tile{{0}, {1}, {2}, {3}}
	stack := NewSeeded(tiles, 42)
	for range len(tiles) {
		tile_a, err := stack.Peek()
		if err != nil {
			t.Fatal(err.Error())
		}
		tile_b, err := stack.Next()
		if err != nil {
			t.Fatal(err.Error())
		}
		if tile_a != tile_b {
			t.Fail()
		}
	}
}

func TestOutOfBounds(t *testing.T) {
	tiles := []Tile{{0}}
	stack := NewSeeded(tiles, 42)
	stack.Next()
	_, err := stack.Peek()
	if err == nil {
		t.Fail()
	}
	_, err = stack.Next()
	if err == nil {
		t.Fail()
	}
}

func TestRemaining(t *testing.T) {
	tiles := []Tile{{0}, {1}, {2}, {3}}
	stack := New(tiles)
	for range 2 {
		stack.Next()
	}
	remainging := stack.GetRemaining()
	if remainging[0] != tiles[2] {
		t.Fail()
	}
	if remainging[1] != tiles[3] {
		t.Fail()
	}
}
