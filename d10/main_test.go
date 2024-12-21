package main

import (
	"testing"

	"aoc24/lib"
)

func TestScoreTrailHeads_Complex(t *testing.T) {
	const (
		in = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
		expect = 36

		expectPaths = 81
	)

	g := buildGuide(lib.StringLines(in))
	if res := g.scoreTrailheads(false); res != expect {
		t.Errorf("expected %d got %d", expect, res)
	}

	if res := g.scoreTrailheads(true); res != expectPaths {
		t.Errorf("expected %d got %d", expectPaths, res)
	}
}

func TestScoreTrailHeads_Simple(t *testing.T) {
	const (
		in = `0123
1234
8765
9876`
		expect = 1
	)

	g := buildGuide(lib.StringLines(in))
	if res := g.scoreTrailheads(false); res != expect {
		t.Errorf("expected %d got %d", expect, res)
	}
}
