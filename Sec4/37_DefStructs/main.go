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
	//This way of defining a struct rely only on the order
	//of the fields, it has to be used under caution though
	//alex := person{"Alex", "Anderson"}

	//This is a second way of declaring a type
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//fmt.Println(alex)

	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%v", alex)
}
