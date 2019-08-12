package main

import "fmt"

type shape interface {
	getArea() float64
}

//Implementation of Square and GetArea
type square struct {
	side float64
}

func (sq square) getArea() float64 {
	return (sq.side * sq.side)
}

//Implementation of Triangle and GetArea
type triangle struct {
	height, base float64
}

func (tr triangle) getArea() float64 {
	return (tr.height * tr.base) / 2
}

//Print the area of shape using the interface
func printArea(s shape) {
	fmt.Println(s)
	fmt.Println(s.getArea())
}

func main() {
	mySquare := square{side: 5}
	printArea(mySquare)

	myTriangle := triangle{height: 10, base: 20}
	printArea(myTriangle)
}
