package main

import (
	"testing"

	"aoc24/lib"
)

var example = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestExample(t *testing.T) {
	const expectXmas = 18
	matrix := lib.BuildMatrix(lib.StringLines(example), lib.Runes)
	if res := xmasSearch(matrix); res != expectXmas {
		t.Errorf("expected %d XMAS got %d", expectXmas, res)
	}

	const expectCross = 9
	if res := crossSearch(matrix); res != 9 {
		t.Errorf("expected %d X-MAS got %d", expectCross, res)
	}
}
