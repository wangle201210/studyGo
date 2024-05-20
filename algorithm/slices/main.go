package main

import (
	"fmt"

	"slices"
)

func main() {
	x := []int{1, 2, 3, 7, 8, 9}
	i, found := slices.BinarySearch(x, 10000)
	fmt.Printf("found %v, i %v", found, i)
}
