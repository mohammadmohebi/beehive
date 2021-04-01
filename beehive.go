package main

import (
	"fmt"

	"github.com/mohammadmohebi/beehive/sort"
)

func demoSort() {
	// Demo bubble sort
	in := []int{1, 0, 3, 4}
	var out [2][]int
	out[0] = sort.Bubble(in)
	out[1] = sort.Selection(in)
	for i, x := range out {
		fmt.Println("Sort algo", i)
		for _, y := range x {
			fmt.Print(y, " ")
		}
		fmt.Print("\n\n")
	}
}

func main() {
	//Demo sort
	demoSort()
}
