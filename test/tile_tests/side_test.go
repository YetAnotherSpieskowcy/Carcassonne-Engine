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
	if connection.TOP_LEFT_CORNER.Rotate(1) != connection.TOP_RIGHT_CORNER {
		t.Fatalf(`side topleft rotate 1 -> topright failed`)
	}

	if connection.TOP_RIGHT_CORNER.Rotate(1) != connection.BOTTOM_RIGHT_CORNER {
		t.Fatalf(`side topright rotate 1 -> bottomright failed`)
	}

	if connection.BOTTOM_RIGHT_CORNER.Rotate(1) != connection.BOTTOM_LEFT_CORNER {
		t.Fatalf(`side bottomright rotate 1 -> bottomleft failed`)
	}

	if connection.BOTTOM_LEFT_CORNER.Rotate(1) != connection.TOP_LEFT_CORNER {
		t.Fatalf(`side bottomleft rotate 1 -> topleft failed`)
	}

	if connection.CENTER.Rotate(1) != connection.CENTER {
		t.Fatalf(`side CENTER rotate 1 -> CENTER failed`)
	}

	if connection.TOP_LEFT_EDGE.Rotate(1) != connection.RIGHT_TOP_EDGE {
		t.Fatalf(`farm_side top rotate 1 -> RIGHT_TOP failed`)
	}

	if connection.TOP_RIGHT_EDGE.Rotate(1) != connection.RIGHT_BOTTOM_EDGE {
		t.Fatalf(`farm_side TOP_LEFT rotate 1 -> RIGHT_BOTTOM failed`)
	}

	if connection.RIGHT_TOP_EDGE.Rotate(1) != connection.BOTTOM_RIGHT_EDGE {
		t.Fatalf(`farm_side RIGHT_TOP rotate 1 -> BOTTOM_RIGHT failed`)
	}

	if connection.RIGHT_BOTTOM_EDGE.Rotate(1) != connection.BOTTOM_LEFT_EDGE {
		t.Fatalf(`farm_side RIGHT_BOTTOM rotate 1 -> BOTTOM_LEFT failed`)
	}

	if connection.BOTTOM_RIGHT_EDGE.Rotate(1) != connection.LEFT_BOTTOM_EDGE {
		t.Fatalf(`farm_side BOTTOM_RIGHT rotate 1 -> LEFT_BOTTOM failed`)
	}

	if connection.BOTTOM_LEFT_EDGE.Rotate(1) != connection.LEFT_TOP_EDGE {
		t.Fatalf(`farm_side BOTTOM_LEFT rotate 1 -> LEFT_TOP failed`)
	}

	if connection.LEFT_BOTTOM_EDGE.Rotate(1) != connection.TOP_LEFT_EDGE {
		t.Fatalf(`farm_side LEFT_BOTTOM rotate 1 -> TOP_LEFT failed`)
	}

	if connection.LEFT_TOP_EDGE.Rotate(1) != connection.TOP_RIGHT_EDGE {
		t.Fatalf(`farm_side LEFT_TOP rotate 1 -> TOP_RIGHT failed`)
	}

	if connection.NONE.Rotate(1) != connection.NONE {
		t.Fatalf(`side NONE rotate 1 -> NONE failed`)
	}

	if connection.Side(80).Rotate(1) != connection.NONE {
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

	if connection.TOP_LEFT_CORNER.ToString() != "TOP_LEFT_CORNER" {
		t.Fatalf(`side TOP_LEFT_CORNER to string failed`)
	}

	if connection.TOP_RIGHT_CORNER.ToString() != "TOP_RIGHT_CORNER" {
		t.Fatalf(`side TOP_RIGHT_CORNER to string failed`)
	}

	if connection.BOTTOM_LEFT_CORNER.ToString() != "BOTTOM_LEFT_CORNER" {
		t.Fatalf(`side BOTTOM_LEFT_CORNER to string failed`)
	}

	if connection.BOTTOM_RIGHT_CORNER.ToString() != "BOTTOM_RIGHT_CORNER" {
		t.Fatalf(`side BOTTOM_RIGHT_CORNER to string failed`)
	}

	if connection.CENTER.ToString() != "CENTER" {
		t.Fatalf(`side CENTER to string failed`)
	}

	if connection.TOP_LEFT_EDGE.ToString() != "TOP_LEFT_EDGE" {
		t.Fatalf(`side TOP_LEFT_EDGE to string failed`)
	}
	if connection.TOP_RIGHT_EDGE.ToString() != "TOP_RIGHT_EDGE" {
		t.Fatalf(`side TOP_RIGHT_EDGE to string failed`)
	}
	if connection.RIGHT_TOP_EDGE.ToString() != "RIGHT_TOP_EDGE" {
		t.Fatalf(`side RIGHT_TOP_EDGE to string failed`)
	}
	if connection.RIGHT_BOTTOM_EDGE.ToString() != "RIGHT_BOTTOM_EDGE" {
		t.Fatalf(`side RIGHT_BOTTOM_EDGE to string failed`)
	}
	if connection.LEFT_TOP_EDGE.ToString() != "LEFT_TOP_EDGE" {
		t.Fatalf(`side LEFT_TOP_EDGE to string failed`)
	}
	if connection.LEFT_BOTTOM_EDGE.ToString() != "LEFT_BOTTOM_EDGE" {
		t.Fatalf(`side LEFT_BOTTOM_EDGE to string failed`)
	}
	if connection.BOTTOM_LEFT_EDGE.ToString() != "BOTTOM_LEFT_EDGE" {
		t.Fatalf(`side BOTTOM_LEFT_EDGE to string failed`)
	}
	if connection.BOTTOM_RIGHT_EDGE.ToString() != "BOTTOM_RIGHT_EDGE" {
		t.Fatalf(`side BOTTOM_RIGHT_EDGE to string failed`)
	}

	if connection.NONE.ToString() != "NONE" {
		t.Fatalf(`side NONE to string failed`)
	}

	if connection.Side(20).ToString() != "ERROR" {
		t.Fatalf(`side ERROR to string failed`)
	}

}
