package main

import "fmt"

type person struct {
	id             int64
	name, lastname string
	age            int
	weight, height float64
}

type personsdb struct {
	store map[int64]person
}

func main() {
	db := personsdb{}

	fmt.Println("db content on start:")

	for _, p := range getAllPersons(&db) {
		fmt.Println(personToString(p))
	}

	fmt.Println()

	addPerson(&db, newPerson(1, "Rauf", "Ismayilov", 37, 70, 185))
	addPerson(&db, newPerson(2, "Amil", "Mammadov", 36, 90, 181))
	addPerson(&db, newPerson(3, "Dilber", "Babayeva", 21, 55, 175))
	addPerson(&db, newPerson(4, "Ismayil", "Ismayilzade", 14, 45, 165))

	fmt.Println()

	personsRauf := getPersonsByName(&db, "Rauf")

	fmt.Println("Rauf's person object", personToString(personsRauf[0]))

	fmt.Println()

	fmt.Println("db content before delete of Rauf:")

	for _, p := range getAllPersons(&db) {
		fmt.Println(personToString(p))
	}

	deletePerson(&db, personsRauf[0].id)

	fmt.Println()
	fmt.Println("db content after delete of Rauf:")

	for _, p := range getAllPersons(&db) {
		fmt.Println(personToString(p))
	}
}

func newPerson(id int64, name, lastname string, age int, weight, height float64) person {
	return person{
		id:       id,
		name:     name,
		lastname: lastname,
		age:      age,
		weight:   weight,
		height:   height,
	}
}

func newPersonsDB() *personsdb {
	return &personsdb{store: map[int64]person{}}
}

func addPerson(db *personsdb, p person) bool {
	initDBIfRequired(db)

	_, ok := db.store[p.id]

	if ok {
		return false
	}

	db.store[p.id] = p

	return true
}

func updatePerson(db *personsdb, name string, p person) bool {
	initDBIfRequired(db)

	_, ok := db.store[p.id]

	if !ok {
		return false
	}

	db.store[p.id] = p

	return true
}

func getPerson(db *personsdb, id int64) (person, bool) {
	initDBIfRequired(db)

	p, ok := db.store[id]

	return p, ok
}

func getPersonsByName(db *personsdb, name string) []person {
	initDBIfRequired(db)

	result := []person{}

	for _, p := range db.store {
		if p.name == name {
			result = append(result, p)
		}
	}

	return result
}

func deletePerson(db *personsdb, id int64) bool {
	initDBIfRequired(db)

	_, ok := db.store[id]

	if !ok {
		return false
	}

	delete(db.store, id)

	return true
}

func getAllPersons(db *personsdb) []person {
	initDBIfRequired(db)

	result := []person{}

	for _, p := range db.store {
		result = append(result, p)
	}

	return result
}

//
//
//
//
//
//
//
func personsDBToString(db *personsdb) string {
	output := "{"

	for _, p := range db.store {
		if output != "{" {
			output += ","
		}

		output += personToString(p)
	}

	output += "}"

	return output
}

func personToString(p person) string {
	return fmt.Sprintf(
		"{name:%s, lastname: %s, age: %d, weight: %f, height: %f}",
		p.name, p.lastname, p.age, p.weight, p.height)
}

func initDBIfRequired(db *personsdb) {
	if db.store == nil {
		db.store = map[int64]person{}
	}
}
