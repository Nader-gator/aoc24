package problems

import (
	"fmt"
	"math"
	"slices"
)

func Solve_12(groups [][]string) {
	p12_1(groups)
	p12_2(groups)
}
func p12_1(groups [][]string) {
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
	xEdges, yEdges map[int][]int
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
	ret.xEdges = make(map[int][]int)
	ret.yEdges = make(map[int][]int)
	for len(toCrawl) > 0 {
		c := toCrawl[0]
		ret.area += 1
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
			}
			if ch.x == 0 {
				ret.yEdges[nc.y] = append(ret.yEdges[nc.y], nc.x)
			} else {
				ret.xEdges[nc.x] = append(ret.xEdges[nc.x], nc.y)

			}
			ret.perim += 1
		}
		toCrawl = toCrawl[1:]

	}
	return
}

func countConnections(nums []int) (ret int) {
	fmt.Println(nums)
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
	for midx := 1; midx < len(nums) && midx < len(nums); midx++ {
		fmt.Println(sidx, midx)
		curr := nums[sidx]
		next := nums[midx]
		if curr == next {
			fmt.Println("1!")
			continue
		} else if !(curr == next-1) {
			fmt.Println("2!")
			ret += 1
			sidx = midx + 1
			continue
		} else if curr == next-1 {
			midx += 1
			fmt.Println("3!")
			continue
		}
	}

	return
}

func p12_2(groups [][]string) {
	fmt.Println(countConnections([]int{0, 1, 2, 3, 4}))
	tot := 0
	seen := make(map[CoOrds]bool)
	for y, row := range groups {
		for x := range row {
			c := CoOrds{x, y}
			chTot := 0
			if _, ok := seen[c]; !ok {
				res := crawl(groups, seen, c)
				ch := groups[c.y][c.x]

				fmt.Println(ch)
				for _, v := range res.xEdges {
					res := countConnections(v)
					fmt.Println("xres:", res)
					chTot += res
					fmt.Println(res)
				}
				for _, v := range res.yEdges {
					res := countConnections(v)
					fmt.Println("yres:", res)
					chTot += res
					fmt.Println(res)
				}
				fmt.Println(ch, chTot, res.area)
				tot += (chTot * res.area)
			}
		}
	}
	fmt.Println(tot)
}
