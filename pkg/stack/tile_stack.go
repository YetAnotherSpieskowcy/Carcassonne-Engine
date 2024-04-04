package stack

import (
	"fmt"
	"math/rand"
	"time"
)

type Tile struct { // To be replaced by #9
	id int
}

type Stack struct {
	seed    int64
	turn_no int32
	tiles   []Tile
	order   []int32
}

func New(tiles []Tile) Stack {
	return NewSeeded(tiles, time.Now().UnixNano())
}

func NewSeeded(tiles []Tile, seed int64) Stack {
	stack := Stack{
		seed:    seed,
		turn_no: 0,
		tiles:   tiles,
		order:   make([]int32, len(tiles)),
	}
	for i := range len(tiles) {
		stack.order[i] = int32(i)
	}
	return stack
}

func (s *Stack) Shuffle() {
	rng := rand.New(rand.NewSource(s.seed))
	rng.Shuffle(len(s.order), func(i, j int) {
		s.order[i], s.order[j] = s.order[j], s.order[i]
	})
}

func (s Stack) GetRemaining() []Tile {
	tiles := []Tile{}
	for _, i := range s.order[s.turn_no:] {
		tiles = append(tiles, s.tiles[i])
	}
	return tiles
}

func (s Stack) Get(n int32) (Tile, error) {
	if n >= int32(len(s.tiles)) {
		return Tile{}, fmt.Errorf("stack: out of bounds")
	}
	return s.tiles[s.order[n]], nil
}

func (s *Stack) Next() (Tile, error) {
	defer func() { s.turn_no += 1 }()
	return s.Get(s.turn_no)
}

func (s Stack) Peek() (Tile, error) {
	return s.Get(s.turn_no)
}
