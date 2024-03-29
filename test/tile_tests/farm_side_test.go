package tile_tests

import (
	"reflect"
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

	if farm_connection.NONE.Rotate(1) != farm_connection.NONE {
		t.Fatalf(`farm_side NONE rotate 1 -> CENTER failed`)
	}

	if farm_connection.FarmSide(20).Rotate(1) != farm_connection.NONE {
		t.Fatalf(`farm_side ERROR rotate 1 -> NONE failed`)
	}
}

func TestFarmConnection(t *testing.T) {

	var farm = farm_connection.FarmConnection{[]farm_connection.FarmSide{farm_connection.TOP_LEFT, farm_connection.LEFT_TOP}}
	var rotated = farm.Rotate(1)
	var result = farm_connection.FarmConnection{[]farm_connection.FarmSide{farm_connection.RIGHT_TOP, farm_connection.TOP_RIGHT}}

	if !reflect.DeepEqual(rotated, result) {
		println("got ", rotated.ToString())
		println("shoulde be ", result.ToString())
		t.Fatalf(`farm connection rotation failed`)
	}
}

func TestFarmSideToString(t *testing.T) {
	if farm_connection.TOP_LEFT.ToString() != "TOP_LEFT" {
		t.Fatalf(`side TOP_LEFT to string failed`)
	}
	if farm_connection.TOP_RIGHT.ToString() != "TOP_RIGHT" {
		t.Fatalf(`side TOP_RIGHT to string failed`)
	}
	if farm_connection.RIGHT_TOP.ToString() != "RIGHT_TOP" {
		t.Fatalf(`side RIGHT_TOP to string failed`)
	}
	if farm_connection.RIGHT_BOTTOM.ToString() != "RIGHT_BOTTOM" {
		t.Fatalf(`side RIGHT_BOTTOM to string failed`)
	}
	if farm_connection.LEFT_TOP.ToString() != "LEFT_TOP" {
		t.Fatalf(`side LEFT_TOP to string failed`)
	}
	if farm_connection.LEFT_BOTTOM.ToString() != "LEFT_BOTTOM" {
		t.Fatalf(`side LEFT_BOTTOM to string failed`)
	}
	if farm_connection.BOTTOM_LEFT.ToString() != "BOTTOM_LEFT" {
		t.Fatalf(`side BOTTOM_LEFT to string failed`)
	}
	if farm_connection.BOTTOM_RIGHT.ToString() != "BOTTOM_RIGHT" {
		t.Fatalf(`side BOTTOM_RIGHT to string failed`)
	}
	if farm_connection.CENTER.ToString() != "CENTER" {
		t.Fatalf(`side CENTER to string failed`)
	}
	if farm_connection.NONE.ToString() != "NONE" {
		t.Fatalf(`side NONE to string failed`)
	}
	if farm_connection.FarmSide(20).ToString() != "ERROR" {
		t.Fatalf(`side NONE to string failed`)
	}
}

func TestFarmConnectionToString(t *testing.T) {
	var connection = farm_connection.FarmConnection{[]farm_connection.FarmSide{farm_connection.BOTTOM_LEFT, farm_connection.TOP_LEFT}}
	if connection.ToString() != farm_connection.BOTTOM_LEFT.ToString()+" "+farm_connection.TOP_LEFT.ToString()+" " {
		println("got")
		println(connection.ToString())
		println("should be")
		println(farm_connection.BOTTOM_LEFT.ToString() + " " + farm_connection.TOP_LEFT.ToString() + " ")

		t.Fatalf(`FarmConnectionToString failed`)
	}
}

func TestNewFarmConnection(t *testing.T) {
	var connection = farm_connection.FarmConnection{[]farm_connection.FarmSide{farm_connection.BOTTOM_LEFT, farm_connection.TOP_LEFT}}
	if !reflect.DeepEqual(connection.Sides, []farm_connection.FarmSide{farm_connection.BOTTOM_LEFT, farm_connection.TOP_LEFT}) {
		t.Fatalf(`NewFarmConnection failed`)
	}
}
