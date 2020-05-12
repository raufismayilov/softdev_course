package main

import (
	"fmt"
)

var tower1 []int
var tower2 []int
var tower3 []int

func main() {
	initHanoi(10)
	solveHanoi()
}

// initialize the first tower with the given discs
func initHanoi(numberOfDiscs int) {
	for i := numberOfDiscs; i > 0; i-- {
		add(i, 1)
	}
}

// moves all the discs from first tower to the third one
func solveHanoi() {
	firstTowerDiscs := len(tower1)

	mov(firstTowerDiscs, 1, 3)
}

func mov(discs int, from, to int) {
	if discs == 1 {
		movSingle(from, to)
		return
	}

	tmp := calcTempTower(from, to)

	mov(discs-1, from, tmp)
	movSingle(from, to)
	mov(discs-1, tmp, to)
}

func movSingle(from, to int) {
	d := remove(from)
	add(d, to)

	logAction("move [%d] from '%d' to '%d'\n", d, from, to)
}

// tower1: 5, 4, 3, 2, [1]

func remove(from int) int {
	ptower := pickTower(from)
	tower := *ptower

	if len(tower) == 0 {
		panic(fmt.Sprintf("'tower%d' is empty", from))
	}

	disc := tower[len(tower)-1]
	*ptower = tower[:len(tower)-1]
	return disc
}

func add(discNo int, to int) {
	ptower := pickTower(to)
	tower := *ptower

	if len(tower) > 0 && tower[len(tower)-1] < discNo {
		panic(
			fmt.Sprintf("can't put disc of a bigger size '%d' onto disc of a smaller size '%d', tower '%d'",
				discNo, tower[len(tower)-1], to))
	}

	*ptower = append(tower, discNo)
}

func pickTower(no int) *[]int {
	switch no {
	case 1:
		return &tower1
	case 2:
		return &tower2
	default:
		return &tower3
	}
}

func calcTempTower(from, to int) int {
	if from == to {
		panic(fmt.Sprintf("'from' and 'to' can't be equal: '%d'", from))
	}

	if from < 1 || from > 3 {
		panic("'from' should be '1', '2' or '3'")
	}

	if to < 1 || to > 3 {
		panic("'to' should be '1', '2' or '3'")
	}

	if from == 1 {
		if to == 2 {
			return 3 // 1->2 => 3
		}

		return 2 // 1->3 => 2
	}

	if from == 2 {
		if to == 1 {
			return 3 // 2->1 => 3
		}

		return 1 // 2->3 => 1
	}

	// from == 3
	if to == 1 {
		return 2 // 3->1 => 2
	}

	return 1 // 3->2 => 1
}

var actionNo int

func logAction(f string, a ...interface{}) {
	actionNo++
	actionNoText := fmt.Sprintf("%3d: ", actionNo)
	fmt.Printf(actionNoText+f, a...)
}
