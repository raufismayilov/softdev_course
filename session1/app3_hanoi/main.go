package main

// var tower = [][]int{
// 	[]int{5, 4, 3, 2, 1},
// 	[]int{},
// 	[]int{},
// }

var tower [][]int

var tower1 []int
var tower2 []int
var tower3 []int

// the app should log actions in the below manner
// mov disc "4" from '1' to '3'
// mov disc "5" from '2' to '3'

func main() {
	initHanoi(5)
	solveHanoi()
}

// initialize the first tower with the given discs
func initHanoi(numberOfDiscs int) {
}

// moves all the discs from first tower to the third one
func solveHanoi() {
	firstTowerDiscs := len(tower[0])

	mov(firstTowerDiscs, 1, 3)
}

func mov(discs int, from, to int) {
}

func movSingle(from, to int) {
}

// #1: 3, 2, 1
// #2: 5, 4
// #3:

// remove(2) -> 4
// add(4, 3)

// #1: 3, 2, 1
// #2: 5
// #3: 4

func remove(from int) int {
	return -1
}

func add(discNo int, to int) {
}
