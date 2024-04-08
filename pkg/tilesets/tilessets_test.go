package tileSets

import (
	"testing"
)

// reference for sets tiles amount https://docs.google.com/spreadsheets/d/1TnPvB6oyisNGs7GZ0xpu-3LPp1V5-t0xH4vocCUPvsY/edit#gid=0

func TestSetStandardTiles(t *testing.T) {
	var set = GetStandardTiles()

	if len(set) != 72 {
		t.Fatalf("got %#v tiles, should be %#v", len(set), 72)
	}

}
