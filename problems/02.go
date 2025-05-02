package problems

import (
	"fmt"
	"slices"
)

func between(x1, x2, v int) bool {
	return v >= x1 && v <= x2
}

func hasIssues(report []int, i int, j int, increasing bool) bool {
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
		return true
	default:
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
		rev_done := false
	do_rev:
		is_safe := true
		increasing := report[0] < report[len(report)-1]
		forgiveness := 1

		for i := 0; i < (len(report) - 1); i++ {
			if hasIssues(report, i, i+1, increasing) {
				if forgiveness < 1 {
					is_safe = false
					break
				} else if hasIssues(report, i, i+2, increasing) {
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
		} else if !rev_done {
			slices.Reverse(report)
			rev_done = true
			goto do_rev
		}
	}
	fmt.Println(safe_reports)
}
func Solve_2(lines [][]int) {
	p02_1(lines)
	p02_2(lines)
}
