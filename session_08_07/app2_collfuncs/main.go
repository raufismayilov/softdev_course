package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	result := where(s, func(x int) bool {
		return x%2 == 0
	})

	result = transform(result, func(x int) int {
		return x * x
	})

	fmt.Println(result)

	gen := createGenerator(0, 5, 10)
	gen2 := createGenerator(time.Now().UnixNano(), 5, 10)

	for i := 0; i < 20; i++ {
		v1 := generate(60, 70)
		v2 := gen()
		v3 := gen2()

		fmt.Println(v1, v2, v3)
	}

	// fact1 := createFactGenerator()
	// fact2 := createFactGenerator()

	// for i := 0; i <= 10; i++ {
	// 	fmt.Println("1 ->", i, fact1())
	// 	if i%2 == 0 {
	// 		fmt.Println("2 ->", i/2, fact2())
	// 	}
	// }

	return

	//
	//
	//
	//
	//
	//
	//
	//
	//
	//

	sum := aggregate(s,
		func() int { return 0 },
		func(x int, acc int) int {
			return acc + x
		})

	min := aggregate(s,
		func() int { return math.MaxInt32 },
		func(x int, localMin int) int {
			if localMin > x {
				return x
			}

			return localMin
		})

	fmt.Println(sum)
	fmt.Println(min)
}

func foreach(s []int, handler func(int)) {
	for _, v := range s {
		handler(v)
	}
}

func transform(s []int, handler func(int) int) []int {
	result := make([]int, 0, len(s))

	for _, oldValue := range s {
		newValue := handler(oldValue)

		result = append(result, newValue)
	}

	return result
}

func where(s []int, condition func(int) bool) []int {
	result := []int{}

	for _, v := range s {
		if condition(v) {
			result = append(result, v)
		}
	}

	return result
}

func aggregate(
	s []int,
	initializer func() int,
	handler func(x int, localAcc int) int,
) int {
	acc := initializer()

	for _, v := range s {
		acc = handler(v, acc)
	}

	return acc
}

func createFactGenerator() func() int {
	var last, index int

	return func() int {
		if index == 0 {
			last, index = 1, 1
			return 1
		}

		last *= index
		index++

		return last
	}
}

func generate(min, max int) int {
	return rand.Intn(max-min) + min
}

func createGenerator(seed int64, min, max int) func() int {
	r := rand.New(rand.NewSource(seed))

	return func() int {
		return r.Intn(max-min) + min
	}
}

type generator struct {
	min, max int
	r        *rand.Rand
}

func newGenerator(seed int64, min, max int) *generator {
	return &generator{
		r:   rand.New(rand.NewSource(seed)),
		min: min,
		max: max,
	}
}

func (g *generator) generate() int {
	return g.r.Intn(g.max-g.min) + g.min
}
