package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	// newPerson := person{"Ishwor", "Panta", 28}

	var alex person
	alex.firstName = "Alex"
	alex.age = 30

	fmt.Printf("%+v \n", alex)
	newPerson := person{firstName: "Ishwor", lastName: "Panta", age: 28}

	fmt.Println("Full User Detail", newPerson)
	fmt.Println(newPerson.firstName, "is", newPerson.age, "years of old.")

	// updating user in structs

}
