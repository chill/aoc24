package main

import (
	"fmt"
	"iter"

	"aoc24/lib"
)

func main() {
	lines := lib.ByLines("./input.txt")
	guide := buildGuide(lines)
	fmt.Println(guide.scoreTrailheads(false))
	fmt.Println(guide.scoreTrailheads(true))
}

type guide struct {
	trailheads lib.Set[lib.Vec]
	matrix     lib.Matrix[int]
}

func buildGuide(lines iter.Seq[string]) guide {
	trailheads := lib.NewSet[lib.Vec]()
	var yPos int
	matrix := lib.BuildMatrix(lines, func(s string) []int {
		r := make([]int, len(s))
		for i, v := range s {
			val := int(v) - 48
			r[i] = val
			if val == 0 {
				trailheads.Add(lib.Vec{
					Y: yPos,
					X: i,
				})
			}
		}

		yPos++
		return r
	})

	return guide{
		trailheads: trailheads,
		matrix:     matrix,
	}
}

func (g guide) scoreTrailheads(paths bool) int {
	var score int
	for head := range g.trailheads.Values() {
		score += g.scoreNode(head, lib.NewSet[lib.Vec](), paths)
	}
	return score
}

func (g guide) scoreNode(node lib.Vec, seen lib.Set[lib.Vec], path bool) int {
	val := g.matrix.At(node)
	if val == 9 {
		if !path { // if we're evaluating unique paths, ignore the seen set
			seen.Add(node)
		}
		return 1
	}

	next := val + 1
	var score int
	for _, dir := range lib.UDLR() {
		newNode := node.Add(dir)
		if !seen.Contains(newNode) && g.matrix.InBoundsVec(newNode) && g.matrix.At(newNode) == next {
			score += g.scoreNode(newNode, seen, path)
		}
	}

	return score
}
