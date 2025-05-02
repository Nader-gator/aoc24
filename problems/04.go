package problems

import (
	"fmt"
	"strings"
)

func addWord(word string, words []string) (words_ret []string) {
	words_ret = append(words, word)
	s := strings.Builder{}
	for idx := len(word) - 1; idx >= 0; idx-- {
		_, err := s.WriteString(string(word[idx]))
		if err != nil {
			panic("writing failed")
		}
	}
	return append(words_ret, s.String())

}

func countXmas(x, y int, lines []string) (count int) {
	y_max := len(lines)
	x_max := len(lines[0])
	var words []string

	topFree := (y - 3) >= 0
	bottomFree := (y_max - y) > 3
	leftFree := (x - 3) >= 0
	rightFree := (x_max - x) > 3

	if topFree {
		s := strings.Join([]string{string(lines[y][x]), string(lines[y-1][x]), string(lines[y-2][x]), string(lines[y-3][x])}, "")
		words = addWord(s, words)
	}
	if bottomFree {
		s := strings.Join([]string{string(lines[y][x]), string(lines[y+1][x]), string(lines[y+2][x]), string(lines[y+3][x])}, "")
		words = addWord(s, words)
	}
	if leftFree {
		s := strings.Join([]string{string(lines[y][x]), string(lines[y][x-1]), string(lines[y][x-2]), string(lines[y][x-3])}, "")
		words = addWord(s, words)
	}
	if rightFree {
		s := strings.Join([]string{string(lines[y][x]), string(lines[y][x+1]), string(lines[y][x+2]), string(lines[y][x+3])}, "")
		words = addWord(s, words)
	}
	var offSets [][]int
	//[
	// [x, y]
	//]
	if leftFree && topFree {
		offSets = append(offSets, []int{-1, -1})
	}
	if leftFree && bottomFree {
		offSets = append(offSets, []int{-1, 1})
	}
	if rightFree && topFree {
		offSets = append(offSets, []int{1, -1})
	}
	if rightFree && bottomFree {
		offSets = append(offSets, []int{1, 1})
	}

	for _, coord := range offSets {
		x_off := coord[0]
		y_off := coord[1]
		s := strings.Join([]string{string(lines[y][x]), string(lines[y+y_off][x+x_off]), string(lines[y+(2*y_off)][x+(2*x_off)]), string(lines[y+(3*y_off)][x+(3*x_off)])}, "")
		words = addWord(s, words)
	}
	for _, word := range words {
		if word == "XMAS" {
			count += 1
		}
	}

	return
}

func countX_mas(x, y int, lines []string) (count int) {
	y_max := len(lines)
	x_max := len(lines[0])

	topFree := (y - 1) >= 0
	bottomFree := (y_max - y) > 1
	leftFree := (x - 1) >= 0
	rightFree := (x_max - x) > 1

	if leftFree && topFree && rightFree && bottomFree {
		tLtoBr := strings.Join([]string{string(lines[y-1][x-1]), string(lines[y][x]), string(lines[y+1][x+1])}, "")
		bLtoTr := strings.Join([]string{string(lines[y+1][x-1]), string(lines[y][x]), string(lines[y-1][x+1])}, "")
		lSeen := false
		c1 := addWord(tLtoBr, []string{})
		for _, str := range c1 {
			if str == "MAS" {
				lSeen = true
			}
		}
		rSeen := false
		c2 := addWord(bLtoTr, []string{})
		for _, str := range c2 {
			if str == "MAS" {
				rSeen = true
			}
		}
		if lSeen && rSeen {
			count += 1
		}
	}

	return
}

func p04_1(lines []string) {
	count := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if string(lines[y][x]) == "X" {
				count += countXmas(x, y, lines)
			}
		}
	}
	fmt.Println(count)
}

func p04_2(lines []string) {
	count := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if string(lines[y][x]) == "A" {
				count += countX_mas(x, y, lines)
			}
		}
	}
	fmt.Println(count)
}

func Solve_4(lines []string) {
	p04_1(lines)
	p04_2(lines)
}
