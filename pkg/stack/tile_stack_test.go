package stack

import (
	"errors"
	"testing"
)

type Tile struct {
	id int
}

func TestStandardOrder(t *testing.T) {
	tiles := []Tile{{0}, {1}, {2}, {3}}
	stack := NewOrdered(tiles)
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
	stack := NewOrdered(tiles)
	for range len(tiles) {
		tileA, err := stack.Peek()
		if err != nil {
			t.Fatal(err.Error())
		}
		tileB, err := stack.Next()
		if err != nil {
			t.Fatal(err.Error())
		}
		if tileA != tileB {
			t.Fail()
		}
	}
}

func TestOutOfBounds(t *testing.T) {
	tiles := []Tile{{0}}
	stack := NewOrdered(tiles)
	_, err := stack.Next()
	if err != nil {
		t.Fail()
	}
	_, err = stack.Peek()
	if err == nil {
		t.Fail()
	}
	if err == nil || !errors.Is(err, ErrStackOutOfBounds) {
		t.Fail()
	}
	_, err = stack.Next()
	if err == nil || !errors.Is(err, ErrStackOutOfBounds) {
		t.Fail()
	}
}

func TestRemaining(t *testing.T) {
	tiles := []Tile{{0}, {1}, {2}, {3}}
	stack := NewOrdered(tiles)
	for range 2 {
		_, err := stack.Next()
		if err != nil {
			t.Fail()
		}
	}
	remaining := stack.GetRemaining()
	if remaining[0] != tiles[2] {
		t.Fail()
	}
	if remaining[1] != tiles[3] {
		t.Fail()
	}
}

func TestTotalTileCount(t *testing.T) {
	tiles := []Tile{{0}, {1}, {2}, {3}}
	tileCount := int32(len(tiles))
	stack := NewOrdered(tiles)
	if stack.GetTotalTileCount() != tileCount {
		t.Fail()
	}
	for range 2 {
		stack.Next()
	}
	if stack.GetTotalTileCount() != tileCount {
		t.Fail()
	}
}

func TestRemainingTileCount(t *testing.T) {
	tiles := []Tile{{0}, {1}, {2}, {3}}
	tileCount := int32(len(tiles))
	stack := NewOrdered(tiles)
	if stack.GetTotalTileCount() != tileCount {
		t.Fail()
	}
	for range 2 {
		stack.Next()
	}
	if stack.GetTotalTileCount() != tileCount - 2 {
		t.Fail()
	}
}
