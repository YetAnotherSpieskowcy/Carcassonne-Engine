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

	if connection.CENTER.Rotate(1) != connection.CENTER {
		t.Fatalf(`side CENTER rotate 1 -> CENTER failed`)
	}

	if connection.NONE.Rotate(1) != connection.NONE {
		t.Fatalf(`side NONE rotate 1 -> NONE failed`)
	}

	if connection.Side(20).Rotate(1) != connection.NONE {
		t.Fatalf(`side ERROR rotate 1 -> NONE failed`)
	}
}

func TestSideToString(t *testing.T) {
	if connection.TOP.ToString() != "TOP" {
		t.Fatalf(`side TOP to string failed`)
	}

	if connection.RIGHT.ToString() != "RIGHT" {
		t.Fatalf(`side RIGHT to string failed`)
	}

	if connection.LEFT.ToString() != "LEFT" {
		t.Fatalf(`side LEFT to string failed`)
	}

	if connection.BOTTOM.ToString() != "BOTTOM" {
		t.Fatalf(`side BOTTOM to string failed`)
	}

	if connection.TOPLEFT.ToString() != "TOPLEFT" {
		t.Fatalf(`side TOPLEFT to string failed`)
	}

	if connection.TOPRIGHT.ToString() != "TOPRIGHT" {
		t.Fatalf(`side TOPRIGHT to string failed`)
	}

	if connection.BOTTOMLEFT.ToString() != "BOTTOMLEFT" {
		t.Fatalf(`side BOTTOMLEFT to string failed`)
	}

	if connection.BOTTOMRIGHT.ToString() != "BOTTOMRIGHT" {
		t.Fatalf(`side BOTTOMRIGHT to string failed`)
	}

	if connection.CENTER.ToString() != "CENTER" {
		t.Fatalf(`side CENTER to string failed`)
	}

	if connection.NONE.ToString() != "NONE" {
		t.Fatalf(`side NONE to string failed`)
	}

	if connection.Side(20).ToString() != "ERROR" {
		t.Fatalf(`side ERROR to string failed`)
	}

}
