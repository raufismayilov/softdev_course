package main

import "fmt"

type intptr *int
type intptrptr *intptr
type intslice []int

type myint int

type peopleages map[string]int

type matrix2DRow []int
type matrix2D []matrix2DRow

type person struct {
	name     string
	lastname string
	age      int
	gender   string
}

func main() {
	var x intslice = intslice{0, 1, 2, 3, 4}

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	ages := peopleages{
		"Rauf":    37,
		"Ashraf":  25,
		"Ismayil": 14,
	}

	fmt.Printf("%T %[1]v\n", ages)

	m2d := matrix2D{
		matrix2DRow{0, 1, 2, 3, 4},
		matrix2DRow{0, 1, 2, 3, 4},
		matrix2DRow{0, 1, 2, 3, 4},
		matrix2DRow{0, 1, 2, 3, 4},
		matrix2DRow{0, 1, 2, 3, 4},
	}

	fmt.Println(m2d)

	// explicit/implicit
	// x := 10
	// var x int64 = 10

	// casting -> bring value of a type to another type
	// int32 x = 10
	// int64 y = int64(x) // casting

	// c# implicit upcast is OK (upcast is from lower precision to higher precision: int8 -> int32)
	// short x = 10;
	// int y = x; // OK
	// c# implicit downcast is NOT OK (downcast is from higher precision to lower precision: int32 -> int8)
	// int x = 10;
	// short y = x; // NOT OK, should be short y = (short)x; // explicit casting

	// go doesn't allow any implicit casting
	// var x int8 = 10
	// var y int32 = x // NOT OK
	// var y int32 = int32(x) // OK
	// var y = int32(x)

	fmt.Println()

	a := myint(100)

	fmt.Printf("%d %[1]T\n", a)

	b := int(a)

	fmt.Printf("%d %[1]T\n", b)

	c := myint(b)

	fmt.Printf("%d %[1]T\n", c)

	var p = person{
		name:     "Farid",
		lastname: "Maharramov",
		age:      109,
		gender:   "M",
	}

	fmt.Printf("%v %[1]T\n", p)

	fmt.Printf("name: %s", p.name)
	fmt.Printf("lastname: %s ", p.lastname)
	fmt.Printf("age: %d ", p.age)
	fmt.Printf("gender: %s\n", p.gender)

	pp := &person{
		name:     "Ismayil",
		lastname: "Isayilzade",
		age:      14,
		gender:   "M",
	}

	fmt.Printf("name: %s ", pp.name)
	fmt.Printf("lastname: %s ", pp.lastname)
	fmt.Printf("age: %d ", pp.age)
	fmt.Printf("gender: %s\n", pp.gender)
}

// show(name, lastname, age)
// show(age:27, name:"Rauf", lastname:"ismayilov")
// show(a,b,c,d,e=10,f=false)

// show(10, 2, "x", "a")
