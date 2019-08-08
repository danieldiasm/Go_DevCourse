package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	//can be replaced by, no need to specify a field name:
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@email.com",
			zipCode: 94000,
		},
	}

	//jimPointer := &jim
	jim.print()
	jim.pointlessupdateName("JACK")
	jim.print()
	jim.updateName("John")
	jim.print()
}

//a function that accepts a struct on its receiver
func (p person) print() {
	fmt.Println(p)
}

//Check this one, no modifications happen
func (p person) pointlessupdateName(newFirstName string) {
	(p).firstName = newFirstName
	p.print()
}

//To modify the value, have a pointer is needed
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
