package main

import (
	"fmt"

	"aoc24/d6/grid"
	"aoc24/lib"
)

func main() {
	lines := lib.ByLines("./input.txt")
	gr := grid.Parse(lines)
	walked := grid.Walk(gr)
	fmt.Println(len(walked))
	fmt.Println(len(grid.FindCycles(gr, walked)))
}
