package problems

import (
	"fmt"
	// "slices"
	// "math"
)

func Solve_2(lines [][]int) {
	p02_1(lines)
	p02_2(lines)
}

func between(x1, x2, v int) bool {
	return v >= x1 && v <= x2
}

func hasIssues(report []int, i int, j int, increasing bool) bool {
	fmt.Println(i, j)
	if j >= len(report) {
		return false
	}

	first, second := report[i], report[j]
	switch {
	case
		increasing && (first > second),
		increasing && !between(1, 3, second-first),
		!increasing && (first < second),
		!increasing && !between(1, 3, first-second):
		println("HI")
		return true
	default:
		println("BIU")
		return false
	}
}

func p02_1(lines [][]int) {
	safe_reports := 0
	for _, report := range lines {
		is_safe := true
		increasing := report[0] < report[len(report)-1]

		for i := 0; i < (len(report) - 1); i++ {
			if hasIssues(report, i, i+1, increasing) {
				is_safe = false
				break
			}
		}
		if is_safe {
			safe_reports += 1
		}
	}
	fmt.Println(safe_reports)
}

func p02_2(lines [][]int) {
	safe_reports := 0
	for _, report := range lines {
		is_safe := true
		increasing := report[0] < report[len(report)-1]
		forgiveness := 1

		for i := 0; i < (len(report) - 1); i++ {
			if hasIssues(report, i, i+1, increasing) {
				if forgiveness < 1 {
					is_safe = false
					break
				}
				if hasIssues(report, i, i+2, increasing) && i+2 >= len(report) {
					is_safe = false
					break
				} else {
					forgiveness = 0
					i++
				}
			}
		}
		if is_safe {
			safe_reports += 1
		}
	}
	fmt.Println(safe_reports)
}
