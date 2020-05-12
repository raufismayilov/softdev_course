package main

import (
	"fmt"
	"math"
)

func main() {
	s := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(getMinRec(s))
}



func getMin(s []int) int {
	min := math.MaxInt32

	for i := 0; i < len(s); i++ {
		if min > s[i] {
			min = s[i]
		}
	}

	return min
}

func getMinRec(s []int) int {
	fmt.Println("getMinRec", s)

	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return s[0]
	}

	minRest := getMinRec(s[1:])

	if minRest < s[0] {
		return minRest
	}

	return s[0]
}

// s -> 10, 9, 8, 7, 6, 5
// s[0] -> 10, s[1:] -> 9, 8, 7, 6
// s[0] -> 9, s[1:] -> 8, 7, 6, 5
