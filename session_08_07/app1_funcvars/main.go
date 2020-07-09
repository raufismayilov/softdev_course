package main

import "fmt"

type person struct {
	id             int64
	name, lastname string
	age            int

	onChange func(*person, string)
}

func (p *person) setName(name string) {
	if p.name == name {
		return
	}

	p.name = name
	if p.onChange != nil {
		p.onChange(p, "name")
	}
}

func (p *person) setLastName(lastname string) {
	if p.lastname == lastname {
		return
	}

	p.lastname = lastname

	if p.onChange != nil {
		p.onChange(p, "lastname")
	}
}

func (p *person) setAge(age int) {
	if p.age == age {
		return
	}

	p.age = age

	if p.onChange != nil {
		p.onChange(p, "age")
	}
}

func (p *person) setOnChange(onChange func(*person, string)) {
	p.onChange = onChange
}

func main() {
	p1 := person{id: 123456789}
	p2 := person{id: 987654321}

	p1.setOnChange(personOnChangeLogger)
	p2.setOnChange(personOnChangeLogger)

	p1.setName("Rauf")
	p1.setLastName("Ismayilov")
	p1.setAge(37)

	p2.setName("Dilber")
	p2.setLastName("Babayeva")
	p2.setAge(30)
}

func personOnChangeLogger(p *person, changeType string) {
	switch changeType {
	case "name":
		fmt.Printf("[%d] Name has changed: %q\n", p.id, p.name)
	case "lastname":
		fmt.Printf("[%d] Last name has changed: %q\n", p.id, p.lastname)
	case "age":
		fmt.Printf("[%d] Age has changed: %d\n", p.id, p.age)
	}
}

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

func sub(x, y int) int {
	return x - y
}

func div(x, y int) int {
	return x / y
}
