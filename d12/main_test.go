package main

import (
	"fmt"
	"testing"

	"aoc24/lib"
)

func TestRegionCost(t *testing.T) {
	const (
		one = `AAAA
BBCD
BBCC
EEEC`
		oneCost = 140
		two     = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`
		twoCost = 772
		three   = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
		threeCost = 1930
	)

	tests := map[int]string{
		oneCost:   one,
		twoCost:   two,
		threeCost: three,
	}

	for expect, test := range tests {
		t.Run(fmt.Sprintf("Expect_%d", expect), func(t *testing.T) {
			res := calcCosts(lib.StringLines(test))
			if res != expect {
				t.Errorf("expected %d, got %d", expect, res)
			}
		})
	}
}
