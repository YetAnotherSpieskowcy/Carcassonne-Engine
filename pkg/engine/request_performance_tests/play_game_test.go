package requestperformancetests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/engine"
)

// coverage test
func TestPlayGame(t *testing.T) {
	b := testing.B{}
	PlayGame(1, &b, func(games *[]engine.SerializedGameWithID, eng *engine.GameEngine, b *testing.B) {})
}
