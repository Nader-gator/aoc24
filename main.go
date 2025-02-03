package main

import (
	p "aoc24/problems"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		panic("bad args")
	}
	problem_num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	f, err := os.Open(fmt.Sprintf("inputs/%02d.txt", problem_num))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := make([][]int, 0)
	scanner := bufio.NewScanner(f)
	line_idx := 0
	for ; scanner.Scan(); line_idx++ {
		lines = append(lines, []int{})
		lines[line_idx] = []int{}
		for _, line := range strings.Fields(scanner.Text()) {
			val, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			lines[line_idx] = append(lines[line_idx], val)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	switch problem_num {
	case 0:
		panic("problems start at 1")
	case 1:
		p.Solve_1(lines)
	case 2:
		p.Solve_2(lines)
	}
}
