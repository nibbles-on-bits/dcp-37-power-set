package main

import (
	"fmt"
)

func main() {
	fmt.Printf("dcp-37 power set\n")
	_ = getPowerSet([]int{1, 2, 3})
}

func getPowerSet(set []int) [][]int {
	ret := [][]int{}
	ret = append(ret, []int{})

	for x := 0; x <= len(set); x++ {
		for y := 0; y < len(set); y++ {
			for x := 0; x < len(set); x++ {
			}
		}
	}

	return ret
}
