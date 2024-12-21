package lib

import (
	"bufio"
	"io"
	"iter"
	"os"
	"slices"
	"strings"
)

func ByRunes(path string) iter.Seq[rune] { // ignores newlines
	lines := ByLines(path)
	return func(yield func(rune) bool) {
		for line := range lines {
			runes := StringRunes(line)
			for r := range runes {
				if !yield(r) {
					return
				}
			}
		}
	}

}

func ByLines(path string) iter.Seq[string] {
	return scan(fopen(path), bufio.ScanLines, stringer)
}

func fopen(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return f
}

func StringLines(s string) iter.Seq[string] {
	r := io.NopCloser(strings.NewReader(s))
	return scan(r, bufio.ScanLines, stringer)
}

func StringRunes(s string) iter.Seq[rune] {
	r := io.NopCloser(strings.NewReader(s))
	return scan(r, bufio.ScanRunes, runer)
}

func Lines(path string) []string {
	return slices.Collect(ByLines(path))
}

func Words(line string) []string {
	r := io.NopCloser(strings.NewReader(line))
	return slices.Collect(scan(r, bufio.ScanWords, stringer))
}

func Runes(line string) []rune {
	return slices.Collect(slices.Values([]rune(line)))
}

func IntsWords(line string) []int {
	r := io.NopCloser(strings.NewReader(line))
	return slices.Collect(scan(r, bufio.ScanWords, inter))
}

func IntsSlice(words []string) []int {
	return ConvSlice(words, inter)
}

func scan[V any](r io.ReadCloser, splitFunc bufio.SplitFunc, convert func(s string) V) iter.Seq[V] {
	scanner := bufio.NewScanner(r)
	scanner.Split(splitFunc)

	return func(yield func(V) bool) {
		defer r.Close()
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
	return rune(s[0])
}
