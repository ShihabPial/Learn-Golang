package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	// //Ways to call a struct
	// peter := person{"Peter", "Porker"}
	// miles := person{firstName: "Miles", lastName: "Morales"}
	// var gwen person

	// //Updating a struct
	// gwen.firstName = "Gwen"
	// gwen.lastName = "Stacy"

	// fmt.Println(peter)
	// fmt.Println(miles)
	// fmt.Println(gwen)

	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 20000,
		},
	}
	jim.updateName("Pam")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
