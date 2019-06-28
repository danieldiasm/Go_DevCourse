package main

func main() {
	//newCards := newDeck()
	//newCards.saveToFile("My_Cards")

	meinCards := newDeckFromFile("My_Cards")
	meinCards.shuffle()
	meinCards.print()
}
