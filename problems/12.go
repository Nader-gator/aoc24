package problems

import (
	"fmt"
	"math"
	"slices"
)

func Solve_12(groups [][]string) {
	// p12_1(groups)
	p12_2(groups)
}
func _p12_1(groups [][]string) {
	tot := 0
	seen := make(map[CoOrds]bool)
	for y, row := range groups {
		for x := range row {
			c := CoOrds{x, y}
			if _, ok := seen[c]; !ok {
				res := crawl(groups, seen, c)
				tot += (res.area * res.perim)
			}
		}

	}
	fmt.Println(tot)
}

type Garden struct {
	area, perim    int
	xEdges, yEdges [][]int
}

func crawl(groups [][]string, seen map[CoOrds]bool, c CoOrds) (ret Garden) {
	changes := []CoOrds{
		{x: 1, y: 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	toCrawl := []CoOrds{c}
	seen[c] = true
	v := groups[c.y][c.x]
	const BUF_SIZE = 1000
	ret.xEdges = make([][]int, BUF_SIZE)
	ret.yEdges = make([][]int, BUF_SIZE)
	for idx := range BUF_SIZE {
		ret.xEdges[idx] = make([]int, BUF_SIZE)
		ret.yEdges[idx] = make([]int, BUF_SIZE)
	}
	for len(toCrawl) > 0 {
		c := toCrawl[0]
		ret.area++
		for _, ch := range changes {
			nc := ch.addCoOrds(c)
			if nc.x >= 0 && nc.y >= 0 && nc.x < len(groups[0]) && nc.y < len(groups) {
				v2 := groups[nc.y][nc.x]
				if v == v2 {
					if !seen[nc] {
						toCrawl = append(toCrawl, nc)
						seen[nc] = true
					}
					continue
				}
				fmt.Println(v2)
			}
			fmt.Println(c, ch, v)
			if ch.x == 0 {
				ret.yEdges[nc.y+1] = append(ret.yEdges[nc.y+1], nc.x)
			} else {
				ret.xEdges[nc.x+1] = append(ret.xEdges[nc.x+1], nc.y)
			}
			ret.perim += 1
		}
		toCrawl = toCrawl[1:]

	}
	return
}

func countConnections(nums []int) (ret int) {
	if len(nums) == 1 {
		return 1
	}
	if len(nums) == 2 {
		if math.Abs(float64(nums[0]-nums[1])) == 1.0 {
			return 1
		} else {
			return 2
		}
	}
	slices.Sort(nums)
	ret = 1
	sidx := 0
	midx := 1
	// [1 2 3 3 4 4]
	for midx < len(nums) {
		if len(nums) == 1 {
			ret += 1
		}
		fmt.Println("NOTE", sidx, midx, nums, ret)
		fmt.Println(nums)
		curr := nums[sidx]
		next := nums[midx]
		original_midx := midx
		if curr == next {
			fmt.Println("1!")
			fmt.Println("midx", midx, "LN", (nums))
			sidx++
			midx++
		} else if !(curr == next-1) {
			fmt.Println("2!")
			fmt.Println("midx", midx, "LN", (nums))
			sidx++
			midx++
			ret += 1
		} else {
			for curr == (next - 1) {
				new_next := nums[original_midx] - 1
				if new_next-1 == curr {
					original_midx++
					continue
				} else {
					ret += 1
					fmt.Println("3!", nums, sidx, midx+1)
					fmt.Println("midx", midx, "LN", (nums))
					nums = slices.Delete(nums, sidx, original_midx)
					fmt.Println("DONE")
				}

			}
		}
	}
	// fmt.Println("slicee", "lciencee")
	return
}

func isSafe(c CoOrds, group [][]string) bool {
	if c.y < 0 || c.x < 0 {
		return false
	}
	if c.y >= len(group) || c.x >= len(group[0]) {
		return false
	}
	return true
}

func getC(c CoOrds, group [][]string) string {
	return group[c.y][c.x]
}

func safeGet(c CoOrds, group [][]string) (string, bool) {
	if isSafe(c, group) {
		return getC(c, group), true
	} else {
		return "", false
	}

}

type CheapGarden struct {
	area, sides, id int
	letter          string
	coords          map[CoOrds]bool
}

func addYCoords(c CoOrds) (neightbors []CoOrds) {
	neightbors = []CoOrds{
		CoOrds{x: 0, y: -1}.addCoOrds(c),
		CoOrds{x: 0, y: 1}.addCoOrds(c),
	}
	return
}

func checkRowSides(groups [][]string, seen map[string]bool) (cheapGardens []CheapGarden) {

	return
}

func TopMost(a, b CoOrds) CoOrds {
	if a.y < b.y {
		return a
	}
	if a.x < b.x {
		return a
	}
	return a
}

func crawlV(
	groups [][]string,
	cheapGardens map[int]CheapGarden,
	seen map[CoOrds]bool,
) {
	changes := []CoOrds{
		{x: 1, y: 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	start := &CoOrds{x: 0, y: 0}
	id := 1
	for start != nil {
		// time.Sleep(time.Duration(500) * time.Millisecond)
		// fmt.Println("uh", start)
		// fmt.Println(seen)
		ch := getC(*start, groups)
		cheapGarden := CheapGarden{letter: ch, id: id, coords: make(map[CoOrds]bool), area: 0, sides: 0}
		toCrawl := []CoOrds{*start}
		// fmt.Println("to crawl", toCrawl)
		for len(toCrawl) > 0 {
			// fmt.Println("INNER to crawl", toCrawl)
			p := toCrawl[0]
			toCrawl = toCrawl[1:]
			if v, ok := seen[p]; v {
				if !ok {
					panic("Nah")
				}
				// fmt.Println("skipping cont", p)
				continue
			}
			seen[p] = true
			// fmt.Println("changes", changes)
			for _, change := range changes {
				newC := change.addCoOrds(p)
				// fmt.Println("newC", newC)
				v, ok := safeGet(newC, groups)
				// fmt.Println("v-ok", v, ok, ch)
				if !ok {
					continue
				} else if v == ch {
					// fmt.Println("inner v==ch", v, ok, ch, newC)
					toCrawl = append(toCrawl, newC)
					// fmt.Println("crawl?", toCrawl)
					cheapGarden.coords[newC] = true
				}
			}
			// fmt.Println("done with change", toCrawl)
		}
		cheapGardens[id] = cheapGarden
		id += 1
		start = nil
		for y, seen := range seen {
			if !seen {
				start = &y
			}
		}
	}
	return
}

// func countSides(garden CheapGarden, groups [][]string) (count int) {
//
//		for k, v := range garden.coords {
//			if !v {
//				continue
//			}
//			a := make([][]string, len(groups))
//			a[k.y][k.x] = garden.letter
//
//		}
//		checkRowSidesiInner := func(c CoOrds, side CoOrds) {
//			for c.x < len(groups[0]) {
//				v, ok := safeGet(side.addCoOrds(c), groups)
//
//				if !ok {
//					garden.sides += 1
//				} else if ok && garden.letter != v {
//					garden.sides += 1
//
//				} else {
//					continue
//				}
//			}
//		}
//		c := CoOrds{x: 0, y: 0}
//		for idx, _ := range groups[0] {
//			checkRowSidesInner(c, "", CoOrds{x: 0, y: -1}, 1)
//			checkRowSidesInner(c, "", CoOrds{x: 0, y: 1}, 1)
//		}
//	}
func p12_2(groups [][]string) {
	seen := make(map[CoOrds]bool)
	gardens := make(map[int]CheapGarden)
	for y := range groups {
		for x := range groups[y] {
			c := CoOrds{x: x, y: y}
			seen[c] = false
		}
	}
	crawlV(groups, gardens, seen)
	// fmt.Println(seen)
	a := make([][]string, len(groups))
	for idx := range a {
		a[idx] = make([]string, len(groups))
	}
	// for _, v := range gardens {
	// 	for _, c := range v.coords {
	// 		a[c.y][c.x] = v.letter
	//
	// 	}
	// }
	// for _, v := range a {
	// 	fmt.Println(v)
	//
	// }
}
