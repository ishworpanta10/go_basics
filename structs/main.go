package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int

	detail contactInfo
}

type contactInfo struct {
	email string
	phone int
}

func main() {

	// newPerson := person{"Ishwor", "Panta", 28}

	var alex person
	alex.firstName = "Alex"
	alex.age = 30
	alex.detail.email = "alex@gmail.com"

	// fmt.Printf("%+v \n", alex)
	alex.print()
	newPerson := person{
		firstName: "Ishwor",
		lastName:  "Panta",
		age:       28,
		detail:    contactInfo{email: "ishwor@gmail.com", phone: 9842345689}}

	// fmt.Println("Full User Detail", newPerson)
	newPerson.print()
	fmt.Println(newPerson.firstName, "is", newPerson.age, "years of old.")

	// updating user in structs
	// newPerson.updateFirstName("Messi")
	// first name is not updated , lets go to pointer in  Go
	// newPerson.print()

	// Pointer in GO

	// newUserPointer := &newPerson // no need to assign this as go can do itself
	//if we have pointer type receiver function so it update that person

	newPerson.updateFirstName("Lionel")

	newPerson.print()

}

func (pointerToPerson *person) updateFirstName(firstName string) {
	(*pointerToPerson).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

/// value type in go
// int, float, string, bool, structs => use pointers to change these values : copies new values to new memory location

/// Reference type in go
// slices, maps, channels, pointers, functions => it just refrence to same memory loaction
