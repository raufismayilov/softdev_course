package main

import (
	"fmt"
	"time"
)

type newtype struct { // 50 bytes
	i int
	f float64
}

type slice struct {
	length, capacity int       // 8
	ptr              *[100]int // 8
}

func main() {
	x := newtype{i: 10, f: 20}

	x.i = 1000

	px := &newtype{i: 1000, f: 2000}

	px.f = 40000 // (*px).f = 40000

	fmt.Println(x)
	fmt.Println(px)

	y := x // 50 bytes will be copied

	_ = y

	arr := [100]int{10, 20, 30, 40, 50}

	s := slice{length: 10, capacity: 20, ptr: &arr}
	s1 := s

	_ = s1

	// slice
	a := []int{10, 20, 30, 40, 50, 60}
	b := a

	fmt.Println(a) // 10 20 30 40 50 60

	b[3] = 0

	fmt.Println(a) // 10 20 30 0 50 60

	// array
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := arr1

	fmt.Println(arr1) // 1 2 3 4 5

	arr2[3] = 0

	fmt.Println(arr1) // 1 2 3 4 5

	arrslice := arr1[3:4] // slice which points to arr1

	arrslice[0] = 0

	fmt.Println(arr1)

	val := 0

	add(&val, 100)

	fmt.Println(val)

	bigarr := [100000000]int8{} // this will take ~1GB in memory

	now := time.Now()

	showArrInfo(bigarr)

	fmt.Println("arr by value", time.Since(now))

	now = time.Now()

	showArrPtrInfo(&bigarr)

	fmt.Println("arr by ref", time.Since(now))
}

func add(ptr *int, diff int) {
	*ptr += diff
}

func showArrInfo(a [100000000]int8) [100000000]int {
	fmt.Println(len(a))
	fmt.Println(a[:10])
}

func showArrPtrInfo(a *[100000000]int8) {
	fmt.Println(len(*a))
	fmt.Println((*a)[:10])
}
