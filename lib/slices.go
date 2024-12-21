package lib

func ConvSlice[V any, B any](in []V, conv func(V) B) []B {
	res := make([]B, len(in))
	for i, val := range in {
		res[i] = conv(val)
	}
	return res
}

func SafeDel[V any](vs []V, i int) []V {
	res := make([]V, len(vs)-1)
	n := copy(res, vs[:i])
	copy(res[n:], vs[i+1:])
	return res
}

func RepeatSlice[V any](v V, count int) []V {
	res := make([]V, count)
	for i := 0; i < count; i++ {
		res[i] = v
	}

	return res
}
