package main

import "testing"

func TestMain(t *testing.T) {
	main()
	tmp := 2
	if tmp != 2 {
		t.Fatal("FAILED")
	}
}
