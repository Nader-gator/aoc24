package problems

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func Solve_11(stones []string) {
	p11_1(stones)
	p11_2(stones)
}

func replace(stone string) []string {
	if stone == "0" {
		return []string{"1"}
	} else if len(stone)%2 == 0 {

		mid := len(stone) / 2
		right := strings.TrimLeft(stone[mid:], "0")
		if right == "" {
			right = "0"
		}
		return []string{stone[:mid], right}
	} else {
		iv, err := strconv.Atoi(stone)
		if err != nil {
			panic(err)
		}
		v := iv * 2024
		return []string{strconv.Itoa(v)}
	}
}

type cache map[int]map[string]int

var calls = make(chan int, 10000)
var cUse = make(chan int, 10000)

func do_run(stone string, blinks int, blink_start int, c cache, c2 cache) int {
	if blinks == 0 {
		return 0
	}
	calls <- 1
	val, ok := c[blinks][stone]
	if ok {
		cUse <- 1
		cUse <- c2[blinks][stone]
		return val
	}
	v := replace(stone)
	if blinks == 1 {
		c[1][stone] = len(v)
		return len(v)
	}
	tot := 0
	for _, stone := range v {
		res := do_run(stone, blinks-1, blink_start, c, c2)
		c[blinks-1][stone] = res
		c2[blinks-1][stone] = len(v)
		tot += res
	}
	return tot
}

func p11_1(stones []string) {
	const blinks = 0

	tot := 0
	c := make(cache)
	for i := range blinks {
		c[i+1] = make(map[string]int)
	}
	c2 := make(cache)
	for i := range blinks {
		c[i+1] = make(map[string]int)
	}

	for _, stone := range stones {
		tot += do_run(stone, blinks, blinks, c, c2)
	}
	fmt.Println(tot)
}

func p11_2(stones []string) {
	const blinks = 75

	callsTot := 0
	cUseTot := 0
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for range calls {
			callsTot += 1
		}
		wg.Done()
	}()
	go func() {
		for range cUse {
			cUseTot += 1
		}
		wg.Done()
	}()

	tot := 0
	c := make(cache)
	c2 := make(cache)
	for i := range blinks {
		c[i+1] = make(map[string]int)
	}
	for i := range blinks {
		c2[i+1] = make(map[string]int)
	}

	for _, stone := range stones {
		tot += do_run(stone, blinks, blinks, c, c2)
	}
	close(calls)
	close(cUse)
	wg.Wait()
	fmt.Println("cache usage: ", float64(cUseTot)/float64(callsTot))
	fmt.Println(tot)
}
