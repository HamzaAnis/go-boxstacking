// package BoxStack
package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/fatih/color"
)

type box struct {
	l float64
	w float64
	h float64
}

//A variable to store the height
var maxHeightR float64 = 0.0

//A struct to sort
type ByArea []box

func (c ByArea) Len() int      { return len(c) }
func (c ByArea) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c ByArea) Less(i, j int) bool {
	a1 := c[i].area()
	a2 := c[j].area()
	return a1 < a2
}

func (b box) area() float64 {
	return b.l * b.w
}

func (b box) compare(compare box) float64 {
	return b.area() - compare.area()
}

// BoxStack is a function to solve the Box stacking problem
func BoxStack(b [][]float64) float64 {
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
	fmt.Println("Before Sorting")
	display(boxes, len(b))
	sort.Sort(ByArea(boxes))
	fmt.Println("After Sorting")
	display(boxes, len(b))

	boxStackHelper(boxes, len(boxes)-1)
	return maxHeightR
}
func boxStackHelper(b []box, pos int) float64 {
	if pos < 0 {
		return 0
	}
	maxHeight := b[pos].h
	for i := 0; i < pos; i++ {
		a := boxStackHelper(b, i)
		if b[i].l < b[pos].l && b[i].w < b[pos].w {
			maxHeight = math.Max(maxHeight, a+b[pos].h)
		}
	}
	maxHeightR = math.Max(maxHeight, maxHeightR)
	// maxHeightR = findMax(maxHeight, maxHeightR)
	return maxHeight
}

func display(abox []box, a int) {
	for i := 0; i < a; i++ {
		fmt.Println(abox[i*3].l, abox[i*3].w, abox[i*3].h)
		fmt.Println(abox[i*3+1].l, abox[i*3+1].w, abox[i*3+1].h)
		fmt.Println(abox[i*3+2].l, abox[i*3+2].w, abox[i*3+2].h)
	}
}
func main() {
	var test1 = [][]float64{
		{5, 5, 1},
		{4, 5, 2},
	}
	var test2 = [][]float64{
		{4, 7, 9},
		{5, 8, 9},
		{11, 20, 40},
		{1, 2, 3},
	}

	var test3 = [][]float64{
		{4, 6, 7},
		{1, 2, 3},
		{4, 5, 6},
		{10, 12, 32},
	}
	result := BoxStack(test1)
	color.Red("%v", result)
	result = BoxStack(test2)
	color.Red("%v", result)
	result = BoxStack(test3)
	color.Red("%v", result)
}
