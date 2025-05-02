package problems

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve_9(files string) {
	p09_1(files)
	p09_2(files)
}

func decodeLayout(files string, ids ...int) (fileBlocks []string) {
	if len(ids) == 0 {
		for i := range len(files) {
			ids = append(ids, i)
		}
	}
	id := 0
	for idx, ch := range strings.Split(files, "") {
		size, err := strconv.Atoi(ch)
		if err != nil {
			panic("Bad math!")
		}
		if idx%2 == 0 {
			for range size {
				fileBlocks = append(fileBlocks, strconv.Itoa(ids[id]))
			}
			id++
		} else {
			for range size {
				fileBlocks = append(fileBlocks, ".")
			}
		}
	}
	return
}

func checkSum(fileBlocks []string) (cum int) {

	for idx, v := range fileBlocks {
		if v == "." {
			continue
		}
		num, err := strconv.Atoi(v)
		if err != nil {
			panic("bad math")
		}

		cum += (idx * num)
	}

	return
}

func p09_1(files string) {

	fileBlocks := decodeLayout(files)
	endPtr := len(fileBlocks) - 1
	startPtr := 0

	for endPtr > startPtr {
		for fileBlocks[startPtr] != "." {
			startPtr++
		}
		for fileBlocks[endPtr] == "." {
			endPtr--
		}
		for fileBlocks[startPtr] == "." && fileBlocks[endPtr] != "." && endPtr > startPtr {
			fileBlocks[startPtr], fileBlocks[endPtr] = fileBlocks[endPtr], "."
			startPtr++
			endPtr--
		}
	}
	fmt.Println(checkSum(fileBlocks))

}

type file = struct {
	size, id, idx int
}

func p09_2(files string) {
	var fileBlocks []int
	var fileInfo []file
	id := 0
	fmap := make(map[int]*file)
	mapIdx := 0
	for idx := 0; len(files) > idx; idx++ {
		if idx%2 == 0 {
			f := file{size: int(files[idx] - '0'), id: id, idx: mapIdx}
			fileInfo = append(fileInfo, f)
			fmap[id] = &f
			for range int(files[idx] - '0') {
				fileBlocks = append(fileBlocks, id)
				mapIdx++
			}
			id++
		} else {
			for range int(files[idx] - '0') {
				fileBlocks = append(fileBlocks, -1)
				mapIdx++
			}
		}
	}

	for idx := range fileInfo {
		id := len(fileInfo) - 1 - idx
		f := fmap[id]
		top := f.idx
		start := -1
		count := 0
	out:
		for i := range top {
			if fileBlocks[i] < 0 {
				start = i
				for range f.size {
					if fileBlocks[start+count] < 0 {
						count++
						if count == f.size {
							break out
						}
					} else {
						count = 0
						start = -1
						break
					}
				}
			} else {
				count = 0
				start = -1
			}
		}
		if start > -1 {
			for i := 0; i < f.size; i++ {
				fileBlocks[start+i] = f.id
			}
			for i := 0; i < f.size; i++ {
				fileBlocks[f.idx+i] = -1
			}
			f.idx = start
		}
	}
	cum := 0
	for i, v := range fileBlocks {
		if v > -1 {
			cum += (i * v)
		}
	}
	fmt.Println(cum)
}
