package main

import (
	"fmt"
	"iter"
	"strings"

	"aoc24/lib"
)

func main() {
	lines := lib.ByLines("./input.txt")
	input := parseInput(lines)

	var total int
	invalid := make([][]int, 0)
	for _, update := range input.updates {
		if !validUpdate(update, input.rules) {
			invalid = append(invalid, update)
			continue
		}

		mid := update[len(update)/2]
		total += mid
	}
	fmt.Printf("part one: %d\n", total)

	total = 0
	for _, update := range invalid {
		fixed := fixUpdate(update, input.rules)

		mid := fixed[len(fixed)/2]
		total += mid
	}

	fmt.Printf("part two: %d\n", total)
}

func validUpdate(update []int, rules map[int]map[int]struct{}) bool {
	seen := map[int]struct{}{}
	for _, n := range update {
		seen[n] = struct{}{}

		rulesForN, ok := rules[n]
		if !ok {
			// no rules for this number
			continue
		}

		for rule := range rulesForN {
			if _, ok := seen[rule]; ok {
				// fmt.Printf("update: %+v\ninvalid due to %d happens before %d rule\n", update, n, rule)
				return false
			}
		}
	}

	return true
}

func fixUpdate(update []int, rules map[int]map[int]struct{}) []int {
	// we only need to move elements left
	// but we also need to backtrack when we shift elements left
	seenAt := map[int]int{}
	for i := 0; i < len(update); i++ {
		n := update[i]
		seenAt[n] = i

		rulesForN, ok := rules[n]
		if !ok {
			// no rules for this number
			continue
		}

		moveTo := i
		for rule := range rulesForN {
			if at, ok := seenAt[rule]; ok && at < moveTo {
				// found a bad element
				// move to the index where we saw the rule, shifting everything already there right
				moveTo = at
			}
		}

		if moveTo == i {
			continue
		}

		// insert n at index moveTo
		update = append(update[:moveTo], append([]int{n}, update[moveTo:]...)...)
		// we now know the original n value exists at i+1, let's delete it
		update = append(update[:i+1], update[i+2:]...)
		// now anything between moveTo+1 and i has not been seen before (i because we just cut one element out)
		for _, v := range update[moveTo+1 : i+1] {
			delete(seenAt, v)
		}
		seenAt[n] = moveTo
		// and we finally need to reset i to continue from the next element at the end of this iteration
		i = moveTo
	}

	return update
}

type input struct {
	rules   map[int]map[int]struct{}
	updates [][]int
}

func parseInput(lines iter.Seq[string]) *input {
	var upd bool
	rules := map[int]map[int]struct{}{}
	updates := make([][]int, 0)
	for line := range lines {
		if strings.TrimSpace(line) == "" {
			upd = true
			continue
		}

		if upd {
			updates = append(updates, parseLine(line, ","))
		} else {
			r := parseLine(line, "|")
			key := rules[r[0]]
			if key == nil {
				key = map[int]struct{}{}
				rules[r[0]] = key
			}
			key[r[1]] = struct{}{}
		}
	}

	return &input{
		rules:   rules,
		updates: updates,
	}
}

func parseLine(line, sep string) []int {
	nums := strings.Split(line, sep)
	res := make([]int, len(nums))
	for i, num := range nums {
		res[i] = lib.Atoi(num)
	}
	return res
}
