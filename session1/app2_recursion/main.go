package main

import (
	"fmt"
)

// Fn = Fn-1 + Fn-2
// F0 = 0
// F1 = 1
// F2 = F1 + F0 = 1 + 0
// F3 = F2 + F1 = 1 + 1 = 2
// F4 = F3 + F2 = 2 + 1 = 3

// mov(6, 1->3)
// mov(5, 1->3)

// mov(4, 1->3) =>
// 		mov(3, 1->2) + move_single(4, 1->3) + mov(3, 2->3)

// mov(3, 1->2) =>
//		mov(2, 1->3) + mov_single(3, 1->2) + mov(2, 3->2)

// mov(2, 1->3) =>
//		mov(1, 1->2) + mov_single(2, 1->3) + mov(1, 2->3)

// mov(1, 1->2) => mov_single(1, 1->2)

// func mov(discs int, from, to int) {
// 		if discs == 1 {
// 			remove(from)
//			add(to)
//			return
// 		}
//
//		var subTo int
// }

func main() {
	fmt.Println("test fibonacci")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d -> %d\n", i, fib(i))
	}

	fmt.Println()
	fmt.Println("test slice sum")
	fmt.Println(sliceSum([]int{1, 2, 3, 4, 5}))
}

func sliceSum(s []int) int {
	// 1, 2, 3, 4, 5
	// sum(1, 2, 3, 4, 5) = 1 + sum(2, 3, 4, 5)
	// sum(2, 3, 4, 5) = 2 + sum(3, 4, 5)
	// ...
	// sum(4, 5) = 4 + sum(5)
	// sum(5) = 5
	if len(s) == 0 {
		return 0
	}

	return s[0] + sliceSum(s[1:])
}

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fib(n-2) + fib(n-1)
}

func printNumbers(n int) {
	if n < 0 {
		return
	}

	fmt.Println(n)

	printNumbers(n - 1)
}

// n! = 1 * 2 * 3 * ... * (n-2) * (n-1) * n

// 5! = 1 * 2 * 3 * 4 * 5
// 5! = (1 * 2 * 3 * 4) * 5 = 4! * 5
// 4! = 1 * 2 * 3 * 4 = (1 * 2 * 3) * 4 = 3! * 4

// n! = (n-1)! * n
// 0! = 1! = 1

// factRec(5)
//    5 * factRec(4)
//        4 * factRec(3)
//            3 * factRec(2)
//                2 * factRec(1)
//

func f(n int) int {
	return n * (n - 1)
}

func factRec(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * factRec(n-1) // left and right should be calculated
	// left is n
	// right is factRec(n-1)
}

func f1() int {
	return 1
}

func f2() int {
	return f1() * 2
}

func f3() int {
	return f2() * 3
}

func f4() int {
	return f3() * 4
}

func f5() int {
	return f4() * 5
}

func factLoop(n int) int {
	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}
