package grid

import (
	"iter"

	"aoc24/lib"
)

func Walk(g *grid) lib.Set[coord] {
	nodes := lib.NewSet[node]()
	visited := lib.NewSet[coord]()
	hatStart, hatDir := g.hat, g.hatDir
	bounds := true
	for bounds {
		walked, inBounds, cycle := g.walkHat(nodes)
		if cycle {
			panic("cycle detected")
		}

		nodes.Add(walked...)
		for _, n := range walked {
			visited.Add(n.c)
		}
		bounds = inBounds
	}

	g.reset(hatStart, hatDir)
	return visited
}

func FindCycles(g *grid, walked lib.Set[coord]) lib.Set[coord] {
	// assume fresh grid, zeroed
	startHat, startDir := g.hat, g.hatDir
	cycles := lib.NewSet[coord]()
	nodes := lib.NewSet[node]()
	for c := range walked.Values() {
		if c == g.hat {
			// don't put a block on the start tile
			continue
		}

		// add the new block in
		g.blocks.Add(c)
		oldV := g.g[c.Y][c.X].val
		g.g[c.Y][c.X].val = '#'

		bounds := true
		iter := 0
		for bounds {
			visited, inBounds, cycle := g.walkHat(nodes)
			iter++
			if cycle {
				cycles.Add(c)
				break
			}
			nodes.Add(visited...)
			bounds = inBounds
		}

		g.blocks.Delete(c)
		g.g[c.Y][c.X].val = oldV
		g.reset(startHat, startDir)
		nodes = lib.NewSet[node]()
	}

	return cycles
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
// it returns the number of tiles walked, whether the hat stays in bounds, and if a cycle was detected
func (g *grid) walkHat(cycleCheck lib.Set[node]) ([]node, bool, bool) {
	prev := g.hat
	visited := []node{
		{
			c: g.hat,
			d: g.hatDir,
		},
	}

	for {
		current := addCoords(prev, g.hatDir)
		if !lib.InBounds(current.Y, current.X, g.g) {
			g.hat = prev
			return visited, false, false
		}

		n := node{
			c: current,
			d: g.hatDir,
		}
		if cycleCheck.Contains(n) {
			// been here before in this direction, cycle
			return visited, false, true
		}

		if g.blocks.Contains(current) {
			g.hat = prev
			g.hatDir = quarterTurn(g.hatDir, true)
			return visited, true, false
		}

		visited = append(visited, node{
			c: current,
			d: g.hatDir,
		})

		prev = current
		// loop will always terminate because we will always hit grid-edge in the worst case
	}
}

func (g *grid) reset(h, d coord) {
	// reset the grid
	lib.ApplyMatrix(g.g, func(p pos) pos {
		return pos{
			val:     p.val,
			visited: false,
		}
	})

	g.g[h.Y][h.X].visited = true
	g.hat = h
	g.hatDir = d
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

type node struct {
	c coord
	d coord
}

type pos struct {
	val     rune
	visited bool
}
