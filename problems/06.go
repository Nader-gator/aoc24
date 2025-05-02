package problems

import (
	"fmt"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
)

type Cell int

type P6_grid [][]Cell

type p6Cache map[CoOrds]map[Cell]bool

const (
	Visted Cell = iota + 1
	Empty
	Obstacle
	GaurdR
	GaurdL
	GaurdU
	GaurdD
)

var cell_str = func() map[Cell]string {
	m := make(map[Cell]string)
	for k, v := range cell_type {
		m[v] = k
	}
	return m
}()

func (c Cell) String() string {
	return cell_str[c]
}

func (c Cell) turn() Cell {
	switch c {
	case GaurdR:
		return GaurdD
	case GaurdD:
		return GaurdL
	case GaurdL:
		return GaurdU
	case GaurdU:
		return GaurdR
	default:
		panic("Turning non gaurd cell")
	}
}

var cell_type = map[string]Cell{
	"X": Visted,
	".": Empty,
	"#": Obstacle,
	">": GaurdR,
	"<": GaurdL,
	"^": GaurdU,
	"v": GaurdD,
}

func ToCell(s string) Cell {
	c, ok := cell_type[s]
	if !ok {
		panic(fmt.Sprintf("invalid letter, %s", s))
	}
	return c
}

func (co CoOrds) String() string {
	return fmt.Sprintf("{x: %v, y: %v}", co.x, co.y)
}
func rowString(r []Cell) string {
	s := strings.Builder{}
	for _, c := range r {
		s.WriteString(c.String())
	}
	return s.String()
}

func (cache p6Cache) add(co CoOrds, c Cell) (added bool) {
	l, ok := cache[co]
	if !ok {
		l = make(map[Cell]bool)
		cache[co] = l
		added = true
	}
	if _, ok := l[c]; !ok {
		l[c] = true
		added = true
	}
	return
}

func (g P6_grid) String() string {
	var lines []string
	for _, row := range g {
		lines = append(lines, rowString(row))
	}
	return strings.Join(lines, "\n")
}

func (g P6_grid) get(co CoOrds) Cell {
	return g[co.y][co.x]
}

func (g P6_grid) set(co CoOrds, v Cell) {
	g[co.y][co.x] = v
}

func (g P6_grid) countV() int {
	return len(g.vList())
}

func (g P6_grid) vList() (c []CoOrds) {
	for y, row := range g {
		for x, el := range row {
			if el == Visted {
				c = append(c, CoOrds{x, y})
			}
		}
	}
	return
}

func (g P6_grid) isGaurd(c CoOrds) bool {
	return slices.Contains([]Cell{GaurdD, GaurdU, GaurdR, GaurdL}, g.get(c))
}

func (g P6_grid) gaurd_present(gaurd_co CoOrds) bool {
	if gaurd_co.y >= 0 && gaurd_co.y < len(g) &&
		gaurd_co.x >= 0 && gaurd_co.x < len(g[0]) {
		return true
	} else {
		return false
	}
}

func (g P6_grid) findGaurd() (c CoOrds, success bool) {
	for y, row := range g {
		for x := range row {
			coords := CoOrds{x, y}
			if g.isGaurd(coords) {
				return coords, true
			}
		}
	}
	return
}
func (g P6_grid) copy() (ng P6_grid) {
	for _, row := range g {
		nr := make([]Cell, len(row))
		copy(nr, row)
		ng = append(ng, nr)
	}
	return ng
}

func (g P6_grid) walkGaurd(gaurd_co CoOrds, seen *p6Cache) (loop bool) {
	var newCoords CoOrds
	var g_dir Cell

	switch g_dir = g.get(gaurd_co); g_dir {
	case GaurdR:
		newCoords = CoOrds{x: gaurd_co.x + 1, y: gaurd_co.y}
	case GaurdL:
		newCoords = CoOrds{x: gaurd_co.x - 1, y: gaurd_co.y}
	case GaurdU:
		newCoords = CoOrds{x: gaurd_co.x, y: gaurd_co.y - 1}
	case GaurdD:
		newCoords = CoOrds{x: gaurd_co.x, y: gaurd_co.y + 1}
	default:
		fmt.Println("BAD: ", g_dir)
		panic("Not looking at gaurd")
	}
	if !g.gaurd_present(newCoords) {
		g.set(gaurd_co, Visted)
		return
	}
	if g.get(newCoords) == Obstacle {
		g.set(gaurd_co, g_dir.turn())
		return g.walkGaurd(gaurd_co, seen)
	} else {
		if seen != nil && !seen.add(newCoords, g_dir) {
			return true
		}
		g.set(gaurd_co, Visted)
		g.set(newCoords, g_dir)
		return g.walkGaurd(newCoords, seen)
	}
}

func Solve_6(g P6_grid) {
	p06_1(g.copy())
	p06_2(g.copy())
}

func p06_1(g P6_grid) {
	gaurd_co, ok := g.findGaurd()
	if !ok {
		panic("Gaurd not found")
	}
	cache := make(p6Cache)
	g.walkGaurd(gaurd_co, &cache)
	fmt.Println("1:", g.countV())
}

func p06_2(g P6_grid) {
	msg := make(chan CoOrds)
	// msg_res := make(chan int)
	og := g.copy()
	gaurd_co, ok := g.findGaurd()
	if !ok {
		panic("Gaurd not found")
	}
	g.walkGaurd(gaurd_co, nil)

	var wg sync.WaitGroup

	var ops atomic.Uint64

	for range 22 {
		go func() {
			wg.Add(1)
			for co := range msg {
				if co == gaurd_co {
					continue
				}
				new_g := og.copy()
				new_g.set(co, Obstacle)
				cache := make(p6Cache)
				if new_g.walkGaurd(gaurd_co, &cache) {
					ops.Add(1)
				}
			}
			wg.Done()
		}()
	}
	for _, co := range g.vList() {
		msg <- co
	}
	close(msg)
	wg.Wait()
	fmt.Println("2:", ops.Load())
}
