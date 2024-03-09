package test

import "testing"

func TestMain(t *testing.T) {
	tmp := 2
	var tmp2 int = 2
	if tmp != tmp2 {
		t.Fatalf("FAILED")
	}
}
