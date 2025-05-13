package problems

import (
	"fmt"
	"sync"
)

func Solve_10(heightMap [][]int) {
	p10_1(heightMap)
	p10_2(heightMap)
}

func (c1 CoOrds) addCoOrds(c2 CoOrds) (c3 CoOrds) {
	c3.x = c1.x + c2.x
	c3.y = c1.y + c2.y
	return
}

func getScore(
	heightMap [][]int,
	tot *int,
	seen map[int]map[int]bool,
	c CoOrds,
	xlimit,
	ylimit int,
	cha chan CoOrds,
	level int,
	distinct bool,
) {
	// fmt.Println(strings.Join(make([]string, level), "-"), heightMap[c.y][c.x], c)
	changes := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	if seen[c.x][c.y] == true {
		return
	}
	seen[c.x][c.y] = true
	old_loc := heightMap[c.y][c.x]
	if old_loc == -1 {
		return
	}
	for _, ch := range changes {
		dc := CoOrds{x: ch[0], y: ch[1]}
		newc := c.addCoOrds(dc)
		if newc.x < xlimit && newc.y < ylimit && newc.x > -1 && newc.y > -1 {
			loc := heightMap[newc.y][newc.x]
			if !(loc-old_loc == 1) {
				continue
			}
			if loc == 9 && ((!seen[newc.x][newc.y]) || distinct) {
				// fmt.Println("HIIT")
				seen[newc.x][newc.y] = true
				cha <- newc
			} else {
				getScore(heightMap, tot, seen, newc, xlimit, ylimit, cha, level+1, distinct)
			}
		}
	}
	seen[c.x][c.y] = false
}

func helper(heightMap [][]int, distinct bool) {
	tot := 0

	ch := make(chan CoOrds, 10000)
	var wg sync.WaitGroup
	ylimit := len(heightMap)
	xlimit := len(heightMap[0])
	wg.Add(1)
	for y, row := range heightMap {
		for x, height := range row {
			if height != 0 {
				continue
			}
			wg.Add(1)
			go func() {
				c := CoOrds{x, y}
				seen := make(map[int]map[int]bool)
				for idx := range ylimit {
					seen[idx] = make(map[int]bool)
				}
				getScore(heightMap, &tot, seen, c, xlimit, ylimit, ch, 0, distinct)
				wg.Done()
			}()
		}
	}
	done := make(chan int, 1)
	go func() {
		for range ch {
			tot += 1
		}
		done <- 1
	}()
	wg.Done()
	wg.Wait()
	close(ch)
	for range done {
		close(done)
		break
	}
	fmt.Println(tot)
}

func p10_1(heightMap [][]int) {
	helper(heightMap, false)
}
func p10_2(heightMap [][]int) {
	helper(heightMap, true)
}
