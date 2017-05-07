// package BoxStack
package main

import "github.com/fatih/color"

type box struct {
	l int
	w int
	h int
}

var maxHeightR = 0

type boxesS []box

func (c boxesS) Len() int      { return len(c) }
func (c boxesS) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

// func (c boxesS) Less(i, j int) bool { return c[i].area > c[j].area }

func (b box) area() int {
	return b.l * b.w
}

func (b box) compare(compare box) int {
	return b.area() - compare.area()
}

func BoxStack(b [][]int) int {
	if len(b) <= 0 {
		color.Red("Not a valid length")
	}
	boxes := make([]box, len(b)*3)

	for i := 0; i < len(b); i++ {
		boxes[i*3] = box{b[i][0], b[i][1], b[i][2]}
		boxes[i*3+1] = box{b[i][1], b[i][2], b[i][0]}
		boxes[i*3+2] = box{b[i][2], b[i][0], b[i][1]}
	}
	maxHeightR = 0
	boxStackHelper(boxes, len(boxes)-1)
	return maxHeightR
	// sort.Sort(boxesS(boxes))
}
func boxStackHelper(b []box, pos int) int {
	if pos < 0 {
		return 0
	}
	maxHeight := b[pos].h
	for i := 0; i < pos; i++ {
		a := boxStackHelper(b, i)
		if b[i].l < b[pos].l && b[i].w < b[pos].w {
			maxHeight = findMax(maxHeight, a+b[pos].h)
		}
	}
	maxHeightR = findMax(maxHeight, maxHeightR)
	return maxHeight
}

func findMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var test1 = [][]int{
		{5, 5, 1},
		{4, 5, 2},
	}
	var test2 = [][]int{
		{4, 6, 7},
		{1, 2, 3},
		{4, 5, 6},
		{10, 12, 32},
	}

	// color.Red("%v", len(test1))
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Print(test1[i][j])
	// 		fmt.Print(" ")
	// 	}
	// 	fmt.Println()
	// }
	result := BoxStack(test1)
	color.Red("%v", result)
}
