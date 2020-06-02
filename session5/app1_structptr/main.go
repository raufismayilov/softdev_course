package main

import (
	"fmt"
	"time"
)

// struct - value type
// class - reference type

type point struct {
	x, y, z int // 4B x 3 => 12B: 0B..11B: x->0..3B, y->4B..7B, z:8B..11B
}

func newPoint(x, y, z int) point {
	return point{x, y, z}
}

func newPointHeap(x, y, z int) *point {
	p := newPoint(x, y, z)

	return &p
}

func main() {
	pp := &point{x: 1, y: 2, z: 3}

	pp = nil

	_ = pp

	p := newPoint(10, 20, 30)

	pp = newPointHeap(10, 20, 30)

	fmt.Println(p)

	now := time.Now()
	x := int64(1)
	show(x)
	fmt.Println(time.Since(now))

	var arr [1000000000]byte

	now = time.Now()
	showBytes(&arr)
	fmt.Println(time.Since(now))

	a, b, c := 10, 20, 30

	p1 := newPoint(a, b, c)

	_ = p1
}

func show(x int64) {}

func showBytes(arr *[1000000000]byte) {}

func A1() point {
	return point{}
}

// A() (allocate 4B for x, shift stack cursor 4B below)
// A -> B (allocate 4B for y, shift stack cursor 4B below)
//		B -> C // (allocate 4B for z, shift stack cursor 4B below)
//		C -> B // (deallocate 4B for z, shift stack cursor 4B above)
// B -> A // (deallocate 4B for y, shift stack cursor 4B above)

func A() {
	x := 10

	B()

	_ = x
}

func B() {
	y := C()

	_ = y
}

func C() int {
	z := 30

	return z
}

func client() { // res2 local for client
	var res1 int // addr -> 50, value -> 0
	var res2 int // addr -> 100

	changer(10, 20, &res1, res2)
}

func changer(x, y int, res1 *int, res2 int) {
	// res2 local for changer addr -> 200
	// res1 local for changer-> addr 150, value -> 50
	*res1 = x + y
	res2 = x + y
}
