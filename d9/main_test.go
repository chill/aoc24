package main

import (
	"slices"
	"testing"

	"aoc24/lib"
)

func TestCompactDefrag(t *testing.T) {
	const (
		in        = "2333133121414131402"
		expectOut = "00992111777.44.333....5555.6666.....8888.."
		expectSum = 2858
	)

	expectSlice := buildExpectedSlice(expectOut)
	m := loadMap(lib.StringRunes(in))
	res := compactDefrag(m)
	if !slices.Equal(res, expectSlice) {
		t.Errorf("expected to get compacted output of %v, got %v", expectSlice, res)
	}

	check := checksum(res)
	if check != expectSum {
		t.Errorf("expected checksum %d got %d", expectSum, check)
	}
}

func buildExpectedSlice(expectOut string) []int {
	expectSlice := make([]int, len(expectOut))
	for i, r := range slices.Collect(lib.StringRunes(expectOut)) {
		if r == '.' {
			expectSlice[i] = 0
			continue
		}
		expectSlice[i] = lib.Atoi(string(r)) + 1
	}

	return expectSlice
}
