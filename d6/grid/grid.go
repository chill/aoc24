package grid

import (
	"iter"

	"aoc24/lib"
)

func Walk(g *grid) int {
	visited := 1
	bounds := true
	for bounds {
		walked, inBounds := g.walkHat()
		visited += walked
		bounds = inBounds
	}
	return visited
}

func Parse(lines iter.Seq[string]) *grid {
	res := &grid{
		g:      make([][]pos, 0),
		blocks: lib.NewSet[coord](),
	}

	var lineCount int
	for line := range lines {
		// rather than a nicely extracted func I wanted a closure over the coords to get them efficiently
		var xCount int
		ps := lib.ConvSlice([]rune(line), func(v rune) pos {
			visited := false
			switch v {
			case '^', '>', '<', 'v':
				res.hat.Y = lineCount
				res.hat.X = xCount
				visited = true
				res.hatDir = buildDir(v)
			case '#':
				res.blocks.Add(coord{
					Y: lineCount,
					X: xCount,
				})
			default:
			}

			xCount++
			return pos{
				val:     v,
				visited: visited,
			}
		})

		res.g = append(res.g, ps)
		lineCount++
	}

	return res
}

func buildDir(r rune) coord {
	var res coord
	switch r {
	case '^':
		res.Y = -1
		res.X = 0
	case '>':
		res.Y = 0
		res.X = 1
	case 'v':
		res.Y = 1
		res.X = 0
	case '<':
		res.Y = 0
		res.X = -1
	default:
		panic("this should never happen")
	}
	return res
}

type grid struct {
	g      [][]pos
	hat    coord
	hatDir coord // not really a coord, actually dy,dx
	blocks lib.Set[coord]
}

// walkHat walks the hat in hatDir until it hits a blocker or the edge
// it returns the number of tiles walked, and whether the hat stays in bounds
func (g *grid) walkHat() (int, bool) {
	prev := g.hat
	visited := 0
	for {
		current := addCoords(prev, g.hatDir)
		if !lib.InBounds(current.Y, current.X, g.g) {
			g.hat = prev
			return visited, false
		}

		if g.blocks.Contains(current) {
			g.hat = prev
			g.hatDir = quarterTurn(g.hatDir, true)
			return visited, true
		}

		if !g.g[current.Y][current.X].visited {
			visited++
			g.g[current.Y][current.X].visited = true
		}

		prev = current
		// loop will always terminate because we will always hit grid-edge in the worst case
	}
}

type coord struct {
	Y, X int
}

func quarterTurn(c coord, clockwise bool) coord {
	// y' = x sin θ + y cos θ
	// x' = x cos θ − y sin θ
	// cos 90/270 = 0, sin 90 = 1, sin 270 = -1
	if clockwise {
		return coord{
			Y: c.X,
			X: c.Y * -1,
		}
	}

	return coord{
		Y: c.X * -1,
		X: c.Y,
	}
}

func addCoords(c, d coord) coord {
	return coord{
		Y: c.Y + d.Y,
		X: c.X + d.X,
	}
}

type pos struct {
	val     rune
	visited bool
}
