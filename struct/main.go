package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// // alex := person{firstName: "ALex", lastName: "Anderson"}
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	// Example 2
	jim := person{
		firstName: "JIM",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "email.gmail.com",
			zipCode: 94000,
		},
	}

	// fmt.Printf("%+v", jim)
	// jimPointer := &jim
	// jimPointer.updateName("James")
	
	jim.updateName("James")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (d person) print() {
	fmt.Printf("%+v", d)
}