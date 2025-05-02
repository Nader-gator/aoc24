package problems

import (
	"fmt"
	"slices"
	"strings"
)

type Scanner struct {
	grid      P8Grid
	signals   []Signal
	signalLoc map[Signal][]CoOrds
	antiNodes map[CoOrds]bool
}

type Signal string
type P8Grid [][]string

type CoOrds struct {
	x, y int
}

func (g P8Grid) copy() (ng P8Grid) {
	for _, row := range g {
		nr := make([]string, len(row))
		copy(nr, row)
		ng = append(ng, nr)
	}
	return ng
}

func (s Scanner) copy() Scanner {
	scanner := Scanner{
		grid:      s.grid.copy(),
		signals:   []Signal{},
		signalLoc: make(map[Signal][]CoOrds),
		antiNodes: make(map[CoOrds]bool),
	}
	return scanner
}

func (c1 CoOrds) antiNode(c2 CoOrds) (c3 CoOrds) {
	dx := c1.x - c2.x
	c3.x = c2.x - dx
	dy := c1.y - c2.y
	c3.y = c2.y - dy
	return
}

func (s Scanner) inBound(c CoOrds) bool {
	return c.x > -1 && c.x < len(s.grid[0]) && c.y > -1 && c.y < len(s.grid)

}

func (s Scanner) String() string {
	sb := strings.Builder{}
	sb.WriteString("\n")
	for y, row := range s.grid {
		for x, v := range row {
			if s.antiNodes[CoOrds{x, y}] {
				sb.WriteString("#")
			} else {
				sb.WriteString(v)
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func Solve_8(grid P8Grid) {
	scanner := Scanner{
		grid:      grid,
		signals:   []Signal{},
		signalLoc: make(map[Signal][]CoOrds),
		antiNodes: make(map[CoOrds]bool),
	}
	for y, row := range grid {
		for x, v := range row {
			if v == "." {
				continue
			}
			sig := Signal(v)
			l, ok := scanner.signalLoc[sig]
			if !ok {
				scanner.signals = append(scanner.signals, sig)
				l = []CoOrds{}
			}
			scanner.signalLoc[sig] = append(l, CoOrds{x, y})
		}
	}
	slices.Sort(scanner.signals)

	p08_1(scanner)
	p08_2(scanner)
}

func updateMap(center CoOrds, friend CoOrds, s Scanner) {
	for aNode := center.antiNode(friend); s.inBound(aNode); aNode = center.antiNode(friend) {
		s.antiNodes[aNode] = true
		center, friend = friend, aNode
	}

}

func checkFriends(s Scanner, sig Signal, selfCheck bool) {
	co := s.signalLoc[sig]
	for baseIdx := 0; baseIdx < len(co); baseIdx++ {
		c := co[baseIdx]
		for idx := baseIdx + 1; idx < len(co); idx++ {
			friend := co[idx]
			updateMap(c, friend, s)
			updateMap(friend, c, s)
			if selfCheck {
				s.antiNodes[c] = true
				s.antiNodes[friend] = true

			}
		}
	}
}

func p08_1(scanner Scanner) {
	for _, signal := range scanner.signals {
		checkFriends(scanner, signal, false)
	}
	fmt.Println("1:", len(scanner.antiNodes))
}

func p08_2(scanner Scanner) {
	for _, signal := range scanner.signals {
		checkFriends(scanner, signal, true)
	}
	// fmt.Println(scanner)
	fmt.Println("2:", len(scanner.antiNodes))
}
