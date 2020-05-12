package main

import (
	"flag"
	"fmt"
)

func returnInt() int {
	return 10
}

func main() {
	px := flag.Int("x", 0, "value of x")
	py := flag.Int("y", 0, "value of y")
	pOper := flag.String("op", "", "operation to be executed")

	flag.Parse()

	if !flag.Parsed() || *pOper == "" {
		flag.Usage()
		return
	}

	x, y, op := *px, *py, *pOper

	fmt.Print(x, y, " ", op, " ")

	switch op {
	case "+":
		fmt.Println(x + y)
	case "-":
		fmt.Println(x - y)
	case "*":
		fmt.Println(x * y)
	case "/":
		fmt.Println(x / y)
	default:
		panic("unexpected operation: " + op)
	}
}
