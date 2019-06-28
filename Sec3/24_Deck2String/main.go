package main

import "fmt"

func main() {
	//newCards := newDeck()
	//newCards.saveToFile("My_Cards")

	meinCards := newDeckFromFile("My_Cards")
	fmt.Println(meinCards)
}
