package main

import (
	"fmt"

	"github.com/mohammadmohebi/beehive/sort"
)

func printArray(arr []int) {
	for _, x := range arr {
		fmt.Print(x, " ")
	}
	fmt.Println("")
}

func demoSort() {
	// Demo sorts
	in := []int{8, 1, 1, 9, 10, 0, 3, 4, 2, 3, 4}
	var out [3][]int

	out[0] = sort.Bubble(in)
	out[1] = sort.Selection(in)
	out[2] = sort.Quicksort(in)

	for i, x := range out {
		fmt.Println("Sort algo", i)
		printArray(in)
		printArray(x)
		fmt.Print("\n\n")
	}
}

func main() {
	//Demo sorts
	demoSort()
}
