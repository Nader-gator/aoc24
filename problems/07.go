package problems

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Op string

const (
	Add  Op = "+"
	Mult    = "*"
	Eq      = "="
)

type P7_ops struct {
	Value int
	Nums  []int
}

type Node struct {
	op    Op
	value *string
	left  *Node
	right *Node
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	if n.value != nil {
		return *n.value
	}
	return fmt.Sprintf("(%s %s %s)", n.op, n.left.String(), n.right.String())
}

type op func(int, int) int

var ops = map[string]op{
	"+": func(a, b int) int { return a + b },
	"*": func(a, b int) int { return a * b },
	"||": func(a, b int) int {
		v, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
		if err != nil {
			panic("BAD MATH!")
		}
		return v
	},
}

func permutations(nums []int, ops []string) (res []string) {
	if len(nums) == 1 {
		return []string{fmt.Sprintf("%d", nums[0])}
	}
	v := fmt.Sprintf("%d", nums[0])
	for _, p := range permutations(nums[1:], ops) {
		for _, op := range ops {
			res = append(res, fmt.Sprintf("%s %s %s", v, op, p))
		}
	}
	return
}

func parse(s string) Node {
	l := strings.Split(s, " ")
	if len(l) < 3 {
		panic("too short!")
	}
	n := Node{
		op:    Op(l[1]),
		left:  &Node{value: &(l[0]), op: "="},
		right: &Node{value: &(l[2]), op: "="},
	}
	if len(l) == 3 {
		return n
	}
	idx := 3
	var leaf = func(nextNum string) *Node {
		return &Node{value: &nextNum, op: "="}
	}

	for idx+1 < len(l) {
		nextOp, nextNum := l[idx], l[idx+1]
		if nextOp == Mult {
			oldRight := n.right
			n.right = &Node{
				op:    "*",
				left:  oldRight,
				right: leaf(nextNum),
			}
		}
		if nextOp == "+" {
			oldN := n
			n = Node{right: &oldN, left: leaf(nextNum), op: "+"}
		}
		idx += 2

	}
	return n
}

func eval(n Node) int {
	fmt.Println("Evalating", &n)
	var v int
	if n.op == Eq {
		v, err := strconv.Atoi(*n.value)
		if err != nil {
			panic("1")
		}
		fmt.Println("EQ sent", v)
		return v
	}
	if n.op == Add {
		r := eval(*n.right)
		v := eval(*n.left)
		res := r + v
		fmt.Println("Add sent", res, "r:", r, v)
		return res
	}
	if n.op == Mult {
		r := eval(*n.right)
		v := eval(*n.left)
		res := r * v
		fmt.Println("Mult sent", "r:", res, "r:", r)
		return v
	}

	return v
}

func tryForNum(nums []int, target int, ops []string) (res []string) {
	for _, p := range permutations(nums, ops) {
		vp := parse(p)
		v := eval(vp)
		if v == target {
			res = append(res, p)
		}
	}
	return res
}

func calculate(s string) int {
	l := strings.Split(s, " ")
	sum, err := strconv.Atoi(l[0])
	idx := 1
	if err != nil {
		panic("bad math!")
	}
	for idx < len(l) {
		op, num := l[idx], l[idx+1]
		n, err := strconv.Atoi(num)
		if err != nil {
			panic("bad math")
		}
		sum = ops[op](sum, n)
		idx += 2
	}
	return sum

}

func tryForNumNoPrec(nums []int, target int, ops []string) (cum int) {
	for _, p := range permutations(nums, ops) {
		if res := calculate(p); res == target {
			cum += res
			break
		}
	}
	return cum
}

func Solve_7(inputs []P7_ops) {
	p07_1(inputs)
	p07_2(inputs)
}

func p07_1(inputs []P7_ops) {
	cum := 0
	for _, input := range inputs {
		cum += tryForNumNoPrec(input.Nums, input.Value, []string{"*", "+"})
	}
	fmt.Println("1:", cum)
}

func p07_2(inputs []P7_ops) {
	ch := make(chan int, 1000000)
	var wg sync.WaitGroup
	wg.Add(1)
	fn := func(input P7_ops) {
		wg.Add(1)
		ch <- tryForNumNoPrec(input.Nums, input.Value, []string{"*", "+", "||"})
		wg.Done()
	}
	for _, input := range inputs {
		go fn(input)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	wg.Done()
	cum := 0
	for i := range ch {
		cum += i
	}
	fmt.Println("2:", cum)
}
