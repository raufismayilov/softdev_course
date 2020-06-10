package main

import (
	"fmt"
	"time"
)

type intslice []int // intslice is created on top of []int
type intptr *int
type mystring string

// type <newtype> <existing_type>

// geometry point 3d {x, y, z}
type point struct {
	x, y, z int
}

type car struct {
	make, model  string
	creationdate time.Time
	color        string
}

func newPoint() point {
	return point{x: 10, y: 20, z: 30}
}

type triangle struct {
	a1, a2, a3 point
}

func main() {
	var arr [10]int

	fmt.Println(arr)

	add(&arr[5], 20)

	fmt.Println(arr)

	return

	var x []int
	var y intslice

	x = []int{10, 20, 30}
	y = intslice(x)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	var p point

	p.x = 10
	p.y = 20
	p.z = 30

	p1 := point{x: 10, y: 20, z: 30}

	p2 := point{10, 20, 30}

	p3 := point{x: 10}

	showPoint(p)
	showPoint(p1)
	showPoint(p2)
	showPoint(p3)

	t := triangle{
		a1: point{x: 0, y: 0, z: 30},
		a2: point{x: 10, y: 0, z: 30},
		a3: point{x: 0, y: 20, z: 30},
	}

	fmt.Println()
	fmt.Println(triangleToString(t))

	t = offsetTriangleByZ(t, -30)

	fmt.Println()
	fmt.Println(triangleToString(t))

	v := 10

	fmt.Println()
	fmt.Println(v)

	add(&v, 100)

	fmt.Println(v)

	a := int(10) // 4 bytes
	acopy := a
	b := int64(10) // 8 bytes

	var p10 point // 12 bytes

	_, _, _, _ = a, b, p10, acopy

	c := point{x: 10, y: 20, z: 30}
	d := c // 12 bytes of value "c" will be copied to variable "d"

	c.z = 3333333

	fmt.Println(pointToString(c))
	fmt.Println(pointToString(d))
}

func showPoint(p point) {
	fmt.Printf("{x:%d, y:%d, z:%d}\n", p.x, p.y, p.z)
}

func pointToString(p point) string {
	return fmt.Sprintf("{x:%d, y:%d, z:%d}", p.x, p.y, p.z)
}

func triangleToString(t triangle) string {
	return fmt.Sprintf(
		"{a1:%s, a2:%s, a3:%s}",
		pointToString(t.a1),
		pointToString(t.a2),
		pointToString(t.a3))
}

func offsetTriangleByZ(t triangle, offset int) triangle {
	t.a1.z += offset
	t.a2.z += offset
	t.a3.z += offset

	return t
}

// passing func arguments by value
// passing func arguments by reference
func sum(x, y int) int {
	r := x + y

	x, y = 0, 0

	return r
}

// func main
// {
//	a, b := 10, 20
//  c := sum(a, b)

// a, b := 10, 20
// x := a
// y := b
// sum()
// }

// func main
// {
//	v := 10
//  vptr *int
//  set vptr = addressof(v)
// 	set diff = 100
//  add()
// }

func add(vptr *int, diff int) {
	*vptr += diff
}
