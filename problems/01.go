package problems

import (
	"fmt"
	"slices"
)

func p01_1(lines_1 []int, lines_2 []int) {
	slices.Sort(lines_1)
	slices.Sort(lines_2)
	if len(lines_1) != len(lines_2) {
		panic("different line lengths")
	}
	distances := 0
	for i := range len(lines_1) {
		var diff int
		if diff = lines_1[i] - lines_2[i]; diff < 0 {
			diff = -diff
		}
		distances += diff

	}
	fmt.Println(distances)
}

func p01_2(lines_1 []int, lines_2 []int) {

	lines_2_map := make(map[int]int)
	for _, val := range lines_2 {
		if count, ok := lines_2_map[val]; ok {
			lines_2_map[val] = count + 1
		} else {

			lines_2_map[val] = 1
		}
	}
	similarity := 0
	for _, val := range lines_1 {
		count := lines_2_map[val]
		similarity += val * count
	}
	fmt.Println(similarity)
}

func Solve_1(lines [][]int) {
	var lines_1 []int
	var lines_2 []int
	for _, line := range lines {
		if len(line) != 2 {
			panic(line)
		}
		lines_1 = append(lines_1, line[0])
		lines_2 = append(lines_2, line[1])
	}
	p01_1(lines_1, lines_2)
	p01_2(lines_1, lines_2)
}
