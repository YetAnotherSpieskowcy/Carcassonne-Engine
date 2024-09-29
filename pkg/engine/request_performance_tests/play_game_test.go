package requestperformancetests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/engine"
)

// coverage test
func TestPlayGame(_ *testing.T) {
	b := testing.B{}
	PlayGame(1, &b, func(_ *[]engine.SerializedGameWithID, _ *engine.GameEngine, _ *testing.B) {})
}
