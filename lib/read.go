package lib

import (
	"bufio"
	"io"
	"iter"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ByLines(path string) iter.Seq[string] {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return scan(f, bufio.ScanLines, stringer)
}

func StringLines(s string) iter.Seq[string] {
	return scan(strings.NewReader(s), bufio.ScanLines, stringer)
}

func Lines(path string) []string {
	return slices.Collect(ByLines(path))
}

func Words(line string) []string {
	return slices.Collect(scan(strings.NewReader(line), bufio.ScanWords, stringer))
}

func Runes(line string) []rune {
	return slices.Collect(scan(strings.NewReader(line), bufio.ScanRunes, runer))
}

func IntsWords(line string) []int {
	return slices.Collect(scan(strings.NewReader(line), bufio.ScanWords, inter))
}

func IntsSlice(words []string) []int {
	return convSlice(words, inter)
}

func convSlice[V any, B any](in []V, conv func(V) B) []B {
	res := make([]B, len(in))
	for i, val := range in {
		res[i] = conv(val)
	}
	return res
}

func scan[V any](r io.Reader, splitFunc bufio.SplitFunc, convert func(s string) V) iter.Seq[V] {
	scanner := bufio.NewScanner(r)
	scanner.Split(splitFunc)

	return func(yield func(V) bool) {
		for scanner.Scan() {
			if !yield(convert(scanner.Text())) {
				return
			}
		}
		if scanner.Err() != nil {
			panic(scanner.Err())
		}
	}
}

func stringer(s string) string {
	return s
}

func inter(s string) int {
	return Atoi(s)
}

func runer(s string) rune {
	return []rune(s)[0]
}

func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

func Abs(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}

func SafeDel[V any](vs []V, i int) []V {
	res := make([]V, len(vs)-1)
	n := copy(res, vs[:i])
	copy(res[n:], vs[i+1:])
	return res
}
