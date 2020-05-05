package main

import "fmt"

func main() {
	s := "Привет друзья"
	b := []byte(s)

	printSlice(b)

	fmt.Println(s)

	// ascii -128..127 - 0..255 1B
	// 2B - 0..65535 Unicode
	// UTF7, UTF8, UTF16

	fmt.Println("start")
	fmt.Println(s, len(s))
	fmt.Println(b, len(b))

	r := []rune(s)

	fmt.Println(r, len(r))

	s1 := string(b)
	fmt.Println(s1, len(s1))
}

func printSlice(s []byte) {
	fmt.Println(s)
}
