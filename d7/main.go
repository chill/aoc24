package main

import (
	"fmt"
	"slices"
	"strconv"

	"aoc24/lib"
)

func main() {
	lr := lib.ByLines("./input.txt")
	lines := slices.Collect(lr)
	fmt.Println(validTotal(lines))
	fmt.Println(validConcatTotal(lines))
}

func validTotal(lines []string) int {
	var total int
	for _, line := range lines {
		r, v := parseLine(line)
		if valid(r, v) {
			total += r
		}
	}
	return total
}

func validConcatTotal(lines []string) int {
	var total int
	for _, line := range lines {
		r, v := parseLine(line)
		if validConcat(r, v) {
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

func validConcat(res int, vals []int) bool {
	added := vals[0] + vals[1]
	multiplied := vals[0] * vals[1]
	concat := lib.Atoi(strconv.Itoa(vals[0]) + strconv.Itoa(vals[1]))

	if len(vals) == 2 {
		// base case
		return added == res || multiplied == res || concat == res
	}

	return validConcat(res, append([]int{added}, vals[2:]...)) ||
		validConcat(res, append([]int{multiplied}, vals[2:]...)) ||
		validConcat(res, append([]int{concat}, vals[2:]...))
}

func parseLine(line string) (int, []int) {
	words := lib.Words(line)
	words[0] = words[0][:len(words[0])-1] // slice off last element of first word, as it's the colon
	vals := lib.IntsSlice(words)
	return vals[0], vals[1:]
}
