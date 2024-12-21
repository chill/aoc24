package lib

import "iter"

type Matrix[V any] [][]V

func (m Matrix[V]) InBounds(y, x int) bool {
	return InBounds(y, x, len(m), len(m[0]))
}

func InBounds(y, x, lenY, lenX int) bool {
	return y >= 0 && x >= 0 && y < lenY && x < lenX
}

func BuildMatrix[V any](lines iter.Seq[string], conv func(string) []V) Matrix[V] {
	matrix := make([][]V, 0)
	for line := range lines {
		matrix = append(matrix, conv(line))
	}

	return matrix
}

func (m Matrix[V]) Apply(fn func(V) V) {
	for _, r := range m {
		for i, v := range r {
			r[i] = fn(v)
		}
	}
}

type Vec struct {
	Y, X int
}

func QuarterTurn(c Vec, clockwise bool) Vec {
	// y' = x sin θ + y cos θ
	// x' = x cos θ − y sin θ
	// cos 90/270 = 0, sin 90 = 1, sin 270 = -1
	if clockwise {
		return Vec{
			Y: c.X,
			X: c.Y * -1,
		}
	}

	return Vec{
		Y: c.X * -1,
		X: c.Y,
	}
}

func (v Vec) Add(d Vec) Vec {
	return Vec{
		Y: v.Y + d.Y,
		X: v.X + d.X,
	}
}

func (v Vec) Sub(d Vec) Vec {
	return Vec{
		Y: v.Y - d.Y,
		X: v.X - d.X,
	}
}

func (v Vec) Invert() Vec {
	return Vec{
		Y: -v.Y,
		X: -v.X,
	}
}
