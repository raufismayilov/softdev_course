package main

import "fmt"

type person struct {
	name, lastname string
	age            int
	weight, height float64
}

func main() {
	employees := []person{
		newPerson("Rauf", "Ismayilov", 37, 70, 185),
		newPerson("Amil", "Mammadov", 36, 90, 181),
		newPerson("Dilber", "Babayeva", 21, 51, 175),
		newPerson("Ismayil", "Ismayilzade", 14, 45, 165),
	}

	showPersons(employees)
}

func newPerson(name, lastname string, age int, weight, height float64) person {
	return person{
		name:     name,
		lastname: lastname,
		age:      age,
		weight:   weight,
		height:   height,
	}
}

func showPersons2(names, lastnames []string, ages []int) {
	for i := names {
		fmt.Printf("%s %s %d\n", names[i], lastnames[i], ages[i])
	}
}

func showPersons(persons []person) {
	for _, p := range persons {
		showPerson(p)
	}
}

func showPerson(p person) {
	fmt.Printf("%15s %15s %4d %4.00f %4.00f\n", p.name, p.lastname, p.age, p.weight, p.height)
}
