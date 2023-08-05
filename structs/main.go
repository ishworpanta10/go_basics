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

	fmt.Printf("%+v \n", alex)
	newPerson := person{
		firstName: "Ishwor",
		lastName:  "Panta",
		age:       28,
		detail:    contactInfo{email: "ishwor@gmail.com", phone: 9842345689}}

	fmt.Println("Full User Detail", newPerson)
	fmt.Println(newPerson.firstName, "is", newPerson.age, "years of old.")

	// updating user in structs

}
