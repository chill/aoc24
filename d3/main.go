package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"aoc24/lib"
)

func main() {
	lines := lib.ByLines("./input.txt")
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)

	total := 0
	skip := false
	for line := range lines {
		matches := re.FindAllString(line, -1)
		for _, m := range matches {
			fmt.Println(m)
			switch m[:4] {
			case "mul(":
				total += mul(m, skip)
			case "do()":
				skip = false
			case "don'":
				skip = true
			default:
				log.Panicf("unknown match %s", m)
			}
		}
	}

	fmt.Printf("total of %d\n", total)
}

func mul(s string, skip bool) int {
	if skip {
		return 0
	}

	s = s[4 : len(s)-1] // slice off first 4 and final char - assume format is mul(XXX,YYY)
	vs := strings.Split(s, ",")
	if len(vs) != 2 {
		panic("bad string")
	}

	nums := lib.IntsSlice(vs)
	return nums[0] * nums[1]
}
