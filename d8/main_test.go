package main

import (
	"testing"

	"aoc24/lib"
)

func TestCalc(t *testing.T) {
	const (
		in = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

		expectP1 = 14
		expectP2 = 34
	)

	if res := calc(lib.StringLines(in), false); res != expectP1 {
		t.Errorf("expected %d got %d", expectP1, res)
	}

	if res := calc(lib.StringLines(in), true); res != expectP2 {
		t.Errorf("expected %d got %d", expectP2, res)
	}
}
