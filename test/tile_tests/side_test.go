package tile_tests

import (
	"testing"

	connection "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/connection"
)

func TestSideRotate(t *testing.T) {

	if connection.TOP.Rotate(1) != connection.RIGHT {
		t.Fatalf(`side top rotate 1 -> right failed`)
	}

	if connection.RIGHT.Rotate(1) != connection.BOTTOM {
		t.Fatalf(`side right rotate 1 -> bottom failed`)
	}

	if connection.BOTTOM.Rotate(1) != connection.LEFT {
		t.Fatalf(`side bottom rotate 1 -> left failed`)
	}

	if connection.LEFT.Rotate(1) != connection.TOP {
		t.Fatalf(`side left rotate 1 -> top failed`)
	}
	// corners
	if connection.TOPLEFT.Rotate(1) != connection.TOPRIGHT {
		t.Fatalf(`side topleft rotate 1 -> topright failed`)
	}

	if connection.TOPRIGHT.Rotate(1) != connection.BOTTOMRIGHT {
		t.Fatalf(`side topright rotate 1 -> bottomright failed`)
	}

	if connection.BOTTOMRIGHT.Rotate(1) != connection.BOTTOMLEFT {
		t.Fatalf(`side bottomright rotate 1 -> bottomleft failed`)
	}

	if connection.BOTTOMLEFT.Rotate(1) != connection.TOPLEFT {
		t.Fatalf(`side bottomleft rotate 1 -> topleft failed`)
	}
}
