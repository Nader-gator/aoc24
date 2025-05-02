package problems

import (
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func endOfIn(values string, idx int) int {
	var offset int
	var r rune
	for offset, r = range values[idx:] {
		if !unicode.IsDigit(r) {
			break
		}
	}
	return offset + idx
}

func findClose(values string, idx int, sum *int) int {
	maybeCommaIdx := endOfIn(values, idx)
	if maybeCommaIdx == idx {
		return idx + 1
	}
	r, _ := utf8.DecodeRuneInString(values[maybeCommaIdx:])
	if r != ',' {
		return maybeCommaIdx
	}
	maybeCloseIdx := endOfIn(values, maybeCommaIdx+1)
	if maybeCloseIdx == (maybeCommaIdx + 1) {
		return maybeCloseIdx
	}
	r, _ = utf8.DecodeRuneInString(values[maybeCloseIdx:])
	if r != ')' {
		return maybeCloseIdx
	}
	num1, err := strconv.Atoi(values[idx:maybeCommaIdx])
	if err != nil {
		panic(fmt.Sprintf("bad value num1 %b %b", idx, maybeCommaIdx))
	}
	num2, err := strconv.Atoi(values[maybeCommaIdx+1 : maybeCloseIdx])
	if err != nil {
		panic(fmt.Sprintf("bad value num2 %b %b", maybeCommaIdx+1, maybeCloseIdx))
	}
	*sum += (num1 * num2)
	return maybeCloseIdx + 1
}

func hasMulParen(values string, idx int, isOn *bool, check bool) bool {
	if !check {
		goto skip
	}
	if !*isOn {
		if len(values) < idx+4 {
			return false
		}
		if values[idx:idx+4] == "do()" {
			fmt.Println("turn on!", idx)
			*isOn = true
		} else {
			return false
		}
	} else {
		if len(values) < idx+7 {
			return false
		}
		if values[idx:idx+7] == "don't()" {
			fmt.Println("turn off!", idx)
			*isOn = false
			return false
		}
	}
skip:
	if len(values) < idx+4 {
		return false
	}
	if values[idx:idx+4] == "mul(" {
		return true
	} else {
		return false
	}
}

func p03_1(values string) {
	sum := 0
	isOn := true
	for idx := 0; idx < len(values); {
		if hasMulParen(values, idx, &isOn, false) {
			idx = findClose(values, idx+4, &sum)
		} else {
			idx++
		}

	}
	fmt.Println(sum)
}

func p03_2(values string) {
	sum := 0
	isOn := true
	for idx := 0; idx < len(values); {
		if hasMulParen(values, idx, &isOn, true) {
			idx = findClose(values, idx+4, &sum)
		} else {
			idx++
		}

	}
	fmt.Println(sum)
}
func Solve_3(lines string) {
	p03_1(lines)
	p03_2(lines)
}
