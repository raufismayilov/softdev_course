package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	var d int
	var e int

	ptrs := []*int{&a, &b, &c, &d, &e}

	showPtrValues(ptrs...)

	a = 20

	showPtrValues(ptrs...)

	var opFunc func(x, y int) int

	s := "+"

	if s == "+" {
		opFunc = add
	} else {
		opFunc = sub
	}

	fmt.Println(opFunc(10, 20))
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func showPtrValues(ptrs ...*int) {
	values := make([]int, len(ptrs))

	for i, v := range ptrs {
		values[i] = *v
	}

	fmt.Println(values)
}
