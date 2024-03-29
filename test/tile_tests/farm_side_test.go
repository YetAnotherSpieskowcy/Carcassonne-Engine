package tile_tests

import (
	"testing"

	farm_connection "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/farm_connection"
)

func TestFarmSideRotate(t *testing.T) {

	if farm_connection.TOP_LEFT.Rotate(1) != farm_connection.RIGHT_TOP {
		t.Fatalf(`farm_side top rotate 1 -> RIGHT_TOP failed`)
	}

	if farm_connection.TOP_RIGHT.Rotate(1) != farm_connection.RIGHT_BOTTOM {
		t.Fatalf(`farm_side TOP_LEFT rotate 1 -> RIGHT_BOTTOM failed`)
	}

	if farm_connection.RIGHT_TOP.Rotate(1) != farm_connection.BOTTOM_RIGHT {
		t.Fatalf(`farm_side RIGHT_TOP rotate 1 -> BOTTOM_RIGHT failed`)
	}

	if farm_connection.RIGHT_BOTTOM.Rotate(1) != farm_connection.BOTTOM_LEFT {
		t.Fatalf(`farm_side RIGHT_BOTTOM rotate 1 -> BOTTOM_LEFT failed`)
	}

	if farm_connection.BOTTOM_RIGHT.Rotate(1) != farm_connection.LEFT_BOTTOM {
		t.Fatalf(`farm_side BOTTOM_RIGHT rotate 1 -> LEFT_BOTTOM failed`)
	}

	if farm_connection.BOTTOM_LEFT.Rotate(1) != farm_connection.LEFT_TOP {
		t.Fatalf(`farm_side BOTTOM_LEFT rotate 1 -> LEFT_TOP failed`)
	}

	if farm_connection.LEFT_BOTTOM.Rotate(1) != farm_connection.TOP_LEFT {
		t.Fatalf(`farm_side LEFT_BOTTOM rotate 1 -> TOP_LEFT failed`)
	}

	if farm_connection.LEFT_TOP.Rotate(1) != farm_connection.TOP_RIGHT {
		t.Fatalf(`farm_side LEFT_TOP rotate 1 -> TOP_RIGHT failed`)
	}
	if farm_connection.CENTER.Rotate(1) != farm_connection.CENTER {
		t.Fatalf(`farm_side CENTER rotate 1 -> CENTER failed`)
	}
}

func TestFarmConnection(t *testing.T) {

	var farm = farm_connection.FarmConnection{farm_connection.TOP_LEFT, farm_connection.LEFT_TOP}
	var rotated = farm.Rotate(1)
	var result = farm_connection.FarmConnection{farm_connection.RIGHT_TOP, farm_connection.TOP_RIGHT}

	if rotated != result {
		println("got ", rotated.ToString())
		println("shoulde be ", result.ToString())
		t.Fatalf(`farm connection rotation failed`)
	}
}
