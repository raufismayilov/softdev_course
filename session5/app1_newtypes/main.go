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

	p := person{name: "Eldar", lastname: "Novruzov", age: 31, gender: "M"}

	// some code in the middle
	//
	//
	// end

	pp := &p

	// pp is a pointer to p which is person

	fmt.Println("name is: ", pp.name)
}

func readData(fileName string) error {
	f, err := os.Open(fileName)

	if err != nil {
		return err
	}

	s := bufio.NewScanner(f)

	for s.Scan() {
		row := s.Text()

		// name,last name,age
		// divide and split by ','
		// [0] - name
		// [1] - last name
		// [2] - age
		// [3] - gender
		p := personFromRow(row)

		persons = append(persons, p)
	}

	return nil
}

// Ismayil,Ismayilzade,14,M
// [0] - Ismayil
// [1] - Ismayilzade
// [2] - 14
// [3] - M
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

	// mytype{ field1: value1, field2: value2, field3: value3 }
	/* mytype {
			field1: value1,
			field2: value2,
			field3: value3,
	}
	*/

	// return person {value1, value2, value3}
	// return person { lastname, name, age, gender }

	// return person { name: name, lastname: lastname }
}
