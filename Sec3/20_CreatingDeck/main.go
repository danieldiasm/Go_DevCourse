package main

import "fmt"

func main() {
	myCards := newDeck()
	fmt.Println("\nFull deck:")
	myCards.print()

	hand, remaining := deal(myCards, 5)

	fmt.Println("\nHand cards:")
	hand.print()
	fmt.Println("\nRemaining cards:")
	remaining.print()

	// greetings := "Hi There!"
	// fmt.Println([]byte(greetings))
}
