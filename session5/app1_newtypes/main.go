package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type intptr *int

var (
	names     []string
	lastnames []string
	ages      []int
	genders   []string

	// persons []persion { name, lastname, age, gender }
)

type person struct {
	name, lastname string
	age            int
	gender         string
}

var persons []person

func main() {
	readData("input.txt")

	// for i := 0; i < len(names); i++ {
	// 	fmt.Printf("%10s %15s %5d %s\n", names[i], lastnames[i], ages[i], genders[i])
	// }

	for _, p := range persons {
		fmt.Printf("%10s %15s %5d %s\n", p.name, p.lastname, p.age, p.gender)
	}
}

func readData(fileName string) error {
	f, err := os.Open(fileName)

	if err != nil {
		return err
	}

	s := bufio.NewScanner(f)

	for s.Scan() {
		row := s.Text()

		// name<TAB>last name<TAB>age
		// divide and split by <TAB>
		// [0] - name
		// [1] - last name
		// [2] - age
		p := personFromRow(row)

		persons = append(persons, p)
	}

	return nil
}

func personFromRow(row string) person {
	tokens := strings.Split(row, ",")

	name, lastname, age, gender := tokens[0], tokens[1], tokens[2], tokens[3]
	ageNumber, _ := strconv.Atoi(age)

	return person{
		name:     name,
		lastname: lastname,
		age:      ageNumber,
		gender:   gender,
	}
}
