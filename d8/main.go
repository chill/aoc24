package main

import (
	"fmt"
	"iter"
	"slices"

	"aoc24/lib"
)

func main() {
	// for an antenna group size n, there are n! antinodes, ew
	r := lib.ByLines("./input.txt")
	lines := slices.Collect(r)
	fmt.Println(calc(slices.Values(lines), false))
	fmt.Println(calc(slices.Values(lines), true))
}

func calc(lines iter.Seq[string], part2 bool) int {
	matrix, antennae := buildMatrix(lines)
	anti := lib.NewSet[lib.Vec]()
	for _, antennae := range antennae {
		for ant := range antennae.Values() {
			nodes := antinodes(ant, antennae, len(matrix), len(matrix[0]), part2)
			anti.Merge(nodes)
		}
	}
	return len(anti)
}

func antinodes(a lib.Vec, group lib.Set[lib.Vec], yLen, xLen int, repeat bool) lib.Set[lib.Vec] {
	res := lib.NewSet[lib.Vec]()
	for b := range group.Values() {
		if b == a {
			continue
		}

		// mirror in both axes to find where an antinode can be positioned
		vec := a.Sub(b) // vector to get from b to a
		anti := a.Add(vec)
		bounds := lib.InBounds(anti.Y, anti.X, yLen, xLen)
		for bounds {
			res.Add(anti)
			if !repeat { // part 1 logic, no repeats
				break
			}

			// part 2 logic, run to the edge of the grid
			anti = anti.Add(vec)
			bounds = lib.InBounds(anti.Y, anti.X, yLen, xLen)
		}

		if !repeat {
			continue
		}

		// part 2 logic, reverse too
		vec = vec.Invert()
		res.Add(a)
		anti = a.Add(vec)
		bounds = lib.InBounds(anti.Y, anti.X, yLen, xLen)
		for bounds {
			res.Add(anti)
			anti = anti.Add(vec)
			bounds = lib.InBounds(anti.Y, anti.X, yLen, xLen)
		}
	}
	return res
}

func buildMatrix(lines iter.Seq[string]) (lib.Matrix[rune], map[rune]lib.Set[lib.Vec]) {
	var yPos int
	antennae := map[rune]lib.Set[lib.Vec]{}

	matrix := lib.BuildMatrix(lines, func(line string) []rune {
		runes := lib.Runes(line)

		for i, r := range runes {
			if r == '.' {
				continue
			}

			v := lib.Vec{
				Y: yPos,
				X: i,
			}

			if _, ok := antennae[r]; !ok {
				antennae[r] = lib.NewSet(v)
			} else {
				antennae[r].Add(v)
			}
		}

		yPos++
		return runes
	})

	return matrix, antennae
}
