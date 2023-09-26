package main

import "testing"

func TestSimple(t *testing.T) {
	if 0 == 1 {
		t.Fatal("impossible")
	}
}