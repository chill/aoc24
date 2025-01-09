package main

import (
	"iter"

	"aoc24/lib"
)

func main() {
	in := lib.ByLines("./input.txt")
	calcCosts(in)
}

func calcCosts(lines iter.Seq[string]) int {
	matrix := lib.BuildMatrix(lines, lib.Runes)
	// BFS
	// mark nodes as visited within current set, so we can skip BFS-ing them again
	// calculate perimeter and store against set key
	// check whether fully enclosed by another area, if so add to perimeter of that set too
	
	// must be a way to efficiently check for enclosure, maybe:
	// pick a tile within the box, transpose to 0, apply same transpose to all coords in other box
	// then for the quadrants around 0, check if one pos is greater than the other for all pos ???
	// still compares all tiles, can we do better? the transpose achieves nothing than making it easier to think about
	return len(matrix) // TODO
}
