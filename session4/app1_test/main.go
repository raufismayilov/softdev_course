package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// input and output

// standard input, output - console
// files
// something that can be used for input and output

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Please enter your input in the below format or 'quit' to exit")
		fmt.Println("format is '<a> <op> <b> (operation can be '+', '-', '*', '/'")

		if !scanner.Scan() {
			break
		}

		cmd := scanner.Text()

		if strings.ToLower(cmd) == "quit" {
			break
		}

		var a, b int
		var op string

		fmt.Sscanln(cmd, &a, &op, &b)

		fmt.Println()

		res, err := calc(a, b, op)

		if err != "" {
			fmt.Println("Operation failed: ", err)
		} else {
			fmt.Println("Result: ", res)
		}

		fmt.Println()
	}
}

func calc(a, b int, op string) (int, string) {
	switch op {
	case "+":
		return a + b, ""
	case "-":
		return a - b, ""
	case "*":
		return a * b, ""
	case "/":
		if b == 0 {
			return 0, "division by zero"
		}
		return a / b, ""
	default:
		return 0, "unrecognized command: " + op
	}
}

func testMultipleScan() {
	fmt.Println("Please enter data in the following format: <name> <last name> <age>")
	var name, lastname string
	var age int
	fmt.Scanln(&name, &lastname, &age)

	fmt.Printf("Hello %s %s, your are %d year(s) old\n", name, lastname, age)
}

func testSimpleScan() {
	var name, lastname string
	var age int

	fmt.Println("Please enter your name")

	fmt.Print("name: ")
	fmt.Scanln(&name)

	fmt.Print("last name: ")
	fmt.Scanln(&lastname)

	fmt.Print("age: ")
	fmt.Scanln(&age)

	fmt.Println()
	fmt.Printf("Hello %s %s, you are %d year(s) old\n", name, lastname, age)
}
