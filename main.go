package main

import (
	p "aoc24/problems"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twoDNums(problem_num int) (lines [][]int) {
	f, err := os.Open(fmt.Sprintf("inputs/%02d.txt", problem_num))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines = make([][]int, 0)
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
	return
}

func longStr(problem_num int) (values string) {
	f, err := os.ReadFile(fmt.Sprintf("inputs/%02d.txt", problem_num))
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(string(f), "\n", "")
}
func linesStr(problem_num int) []string {
	f, err := os.ReadFile(fmt.Sprintf("inputs/%02d.txt", problem_num))
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimSpace(string(f)), "\n")
}

func p5(problem_num int) (queue []p.P5_q, ordering [][]int) {
	f, err := os.ReadFile(fmt.Sprintf("inputs/%02d.txt", problem_num))
	if err != nil {
		panic(err)
	}

	s := strings.Split(string(f), "\n\n")
	queueOrder := s[0]
	updates := s[1]

	for _, line := range strings.Split(string(queueOrder), "\n") {

		parts := strings.Split(strings.TrimSpace(string(line)), "|")
		p1, err1 := strconv.Atoi(parts[0])
		p2, err2 := strconv.Atoi(parts[1])
		if (err1 != nil) || (err2 != nil) {
			panic("bad number formatting")
		}
		queue = append(queue, p.P5_q{Left: p1, Right: p2})
	}

	for lineIdx, line := range strings.Split(strings.TrimSpace(string(updates)), "\n") {
		ordering = append(ordering, []int{})
		for _, num := range strings.Split(strings.TrimSpace(string(line)), ",") {
			p, err := strconv.Atoi(num)
			if err != nil {
				panic("bad number formatting")
			}
			ordering[lineIdx] = append(ordering[lineIdx], p)
		}
	}
	return
}

func p6(problem_num int) (grid p.P6_grid) {

	for _, line := range linesStr(problem_num) {
		var row []p.Cell
		for _, letter := range strings.Split(line, "") {
			row = append(row, p.ToCell(letter))
		}
		// grid.Grid = append(grid.Grid, row)
		grid = append(grid, row)
	}
	return
}

func p7(problem_num int) (ops []p.P7_ops) {

	for _, line := range linesStr(problem_num) {
		s := strings.Split(line, ":")
		v := s[0]
		value, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		nums := []int{}
		for _, nv := range strings.Split(strings.TrimSpace(s[1]), " ") {
			v, err := strconv.Atoi(nv)
			if err != nil {
				panic(err)
			}
			nums = append(nums, v)
		}
		ops = append(ops, p.P7_ops{
			Value: value,
			Nums:  nums,
		})
	}
	return
}

func p8(problem_num int) (grid p.P8Grid) {

	for _, line := range linesStr(problem_num) {

		var row []string
		for _, v := range strings.Split(line, "") {

			row = append(row, v)
		}
		grid = append(grid, row)
	}
	return
}

func main() {
	if len(os.Args) != 2 {
		panic("bad args")
	}
	problem_num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	switch problem_num {
	case 0:
		panic("problems start at 1")
	case 1:
		lines := twoDNums(problem_num)
		p.Solve_1(lines)
	case 2:
		lines := twoDNums(problem_num)
		p.Solve_2(lines)
	case 3:
		input := longStr(problem_num)
		p.Solve_3(input)
	case 4:
		input := linesStr(problem_num)
		p.Solve_4(input)
	case 5:
		queue, order := p5(problem_num)
		p.Solve_5(queue, order)
	case 6:
		grid := p6(problem_num)
		p.Solve_6(grid)
	case 7:
		ops := p7(problem_num)
		p.Solve_7(ops)
	case 8:
		grid := p8(problem_num)
		p.Solve_8(grid)
	case 9:
		line := longStr(9)
		p.Solve_9(line)
	default:
		panic("Problem not solved yet")
	}
}
