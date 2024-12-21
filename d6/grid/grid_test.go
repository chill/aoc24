package grid

import (
	"testing"

	"aoc24/lib"
)

func TestParseGrid(t *testing.T) {
	const in = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	expectHat := lib.Vec{
		Y: 6,
		X: 4,
	}

	expectBlocks := lib.NewSet([]lib.Vec{
		{Y: 0, X: 4},
		{Y: 1, X: 9},
		{Y: 3, X: 2},
		{Y: 4, X: 7},
		{Y: 6, X: 1},
		{Y: 7, X: 8},
		{Y: 8, X: 0},
		{Y: 9, X: 6},
	}...)

	g := Parse(lib.StringLines(in))

	if len(g.g) != 10 {
		t.Errorf("expected 10 rows, got %d", len(g.g))
	}

	rowLen := len(g.g[0])
	for _, r := range g.g[1:] {
		if len(r) != rowLen {
			t.Errorf("got row length %d, all should be the same: %d", len(r), rowLen)
		}
	}

	if g.hat != expectHat {
		t.Errorf("expected hat %+v, got %+v", expectHat, g.hat)
	}

	if !g.blocks.Equals(expectBlocks) {
		t.Errorf("expected blocks %+v, got %+v", expectBlocks, g.blocks)
	}

	walked := Walk(g)
	if len(walked) != 41 {
		t.Errorf("expected 41 positions to be visited, got %d", len(walked))
	}

	cyclesAt := FindCycles(g, walked)
	if len(cyclesAt) != 6 {
		t.Errorf("expected 6 positions to generate cycles, got %d", len(cyclesAt))
	}
}

func TestPart1(t *testing.T) {
	res := Walk(Parse(lib.ByLines("../input.txt")))
	if len(res) != 5318 {
		t.Errorf("expected 5318 positions to be visited, got %d", len(res))
	}
}
