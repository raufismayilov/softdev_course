package main

import "fmt"

var hostname string = "hp"

func main() {
	x := 10 // addr 1000-1003, value 10

	y := x // addr 2000-2003, value 10

	var xptr *int // addr 3000-3007, value is 0 (nil)

	xptr = &x // addr 3000-3007, value 1000

	y = 100 // addr 2000-2003, value 100

	fmt.Println(x, y, xptr, *xptr)

	*xptr = 10000

	fmt.Println(x, y, xptr, *xptr)

	xptr = &y // addr 3000-3007, value 2000

	*xptr = 300000

	fmt.Println(x, y, xptr, *xptr)

	// if xptr holds the reference to x
	// 		then *xptr will jump to the x's address and will read its value

	var c int = 123456 // addr10000
	var d int = 123456 // addr20000

	fmt.Println("c is", c)
	fmt.Println("d is", d)

	add(10, 20, &c, d) // x: 10, y: 20, pres: addr10000, res: 123456

	fmt.Println("c became", c)
	fmt.Println("d is", d)

	sqr(10)
}

// lvalue

func add(x, y int, resptr *int, res int) int {
	*resptr = x + y
	res = x + y

	return x + y
}

func sqr(x float64) float64 {
	//fmt.Println(x * x)
	// z = f(x,y)
	return x * x
}

func hypot(a, b float64) float64 {
	// sqr(a), sqr(b)
	csq := sqr(a) + sqr(b)
}
