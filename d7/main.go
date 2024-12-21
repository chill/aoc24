package main

import (
	"fmt"
	"iter"

	"aoc24/lib"
)

func main() {
	lines := lib.ByLines("./input.txt")
	fmt.Println(validTotal(lines))
}

func validTotal(lines iter.Seq[string]) int {
	var total int
	for line := range lines {
		r, v := parseLine(line)
		if valid(r, v) {
			total += r
		}
	}
	return total
}

func valid(res int, vals []int) bool {
	added := vals[0] + vals[1]
	multiplied := vals[0] * vals[1]
	if len(vals) == 2 {
		// base case
		return added == res || multiplied == res
	}

	return valid(res, append([]int{added}, vals[2:]...)) || valid(res, append([]int{multiplied}, vals[2:]...))
}

func parseLine(line string) (int, []int) {
	words := lib.Words(line)
	words[0] = words[0][:len(words[0])-1] // slice off last element of first word, as it's the colon
	vals := lib.IntsSlice(words)
	return vals[0], vals[1:]
}
