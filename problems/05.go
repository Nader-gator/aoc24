package problems

import (
	"fmt"
)

type P5_q struct {
	Left  int
	Right int
}

type ID = int

type cacheNode struct {
	Left, Right map[ID]bool
}

type cacheCluster map[ID]cacheNode

func (c cacheCluster) addRecord(q P5_q) {
	if _, ok := c[q.Left]; !ok {
		c[q.Left] = cacheNode{
			Left:  make(map[ID]bool),
			Right: make(map[ID]bool),
		}
	}
	c[q.Left].Right[q.Right] = true

	if _, ok := c[q.Right]; !ok {
		c[q.Right] = cacheNode{
			Left:  make(map[ID]bool),
			Right: make(map[ID]bool),
		}
	}
	c[q.Right].Left[q.Left] = true
}

func (cc cacheCluster) challengerAllowedSmaller(challenger, defender ID) bool {
	var chWins, chLoses, defWins, defLoses bool

	chNode, chOk := cc[challenger]
	defNode, defOk := cc[defender]

	if chOk {
		_, chWins = chNode.Right[defender]
		_, chLoses = chNode.Left[defender]
	}

	if defOk {
		_, defWins = defNode.Right[challenger]
		_, defLoses = defNode.Left[challenger]
	}

	if chWins && defWins {
		panic(fmt.Sprintf("two nodes are saying they should be bigger and smaller than each other: %b - %b", challenger, defender))
	}
	if chWins || defLoses {
		return true
	}
	if chLoses || defWins {
		return false
	}
	return true
}

func checkRow(cc cacheCluster, row []int) bool {
	for idx, valLeft := range row {
		for _, valRight := range row[idx+1:] {
			if !cc.challengerAllowedSmaller(valLeft, valRight) {
				return false
			}
		}
	}
	return true
}

func fixRow(cc cacheCluster, row []int) {
	for idx := 0; idx < len(row); idx++ {
		for swap_idx := idx + 1; swap_idx < len(row); swap_idx++ {
			if !cc.challengerAllowedSmaller(row[idx], row[swap_idx]) {
				row[idx], row[swap_idx] = row[swap_idx], row[idx]
			}
		}
	}
}

func Solve_5(queue []P5_q, updates [][]int) {
	p05_1(queue, updates)
	p05_2(queue, updates)
}

func p05_1(queue []P5_q, updates [][]int) {

	cc := cacheCluster(make(map[ID]cacheNode))
	for _, q := range queue {
		cc.addRecord(q)
	}
	var rows [][]int
	for _, row := range updates {
		if checkRow(cc, row) {
			rows = append(rows, row)
		}
	}
	sum := 0
	for _, row := range rows {
		sum += (row[len(row)/2])

	}
	fmt.Println(sum)

}

func p05_2(queue []P5_q, updates [][]int) {
	cc := cacheCluster(make(map[ID]cacheNode))
	for _, q := range queue {
		cc.addRecord(q)
	}
	var rows [][]int
	for _, row := range updates {
		if !checkRow(cc, row) {
			rows = append(rows, row)
		}
	}

	for _, row := range rows {
		fixRow(cc, row)
	}

	sum := 0
	for _, row := range rows {
		sum += (row[len(row)/2])

	}
	fmt.Println(sum)
}
