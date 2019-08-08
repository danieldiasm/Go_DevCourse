package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	//sb := spanishBot{}

	printGreeting(eb)
	// printGreeting(sb)

}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

//We can here remove the value for the receiver (englishBot)
//Since we are using nothing from it! We can also leave it there... (eb englishBot)
func (englishBot) getGreeting() string {
	//VERY CUSTOM LOGIC FOR GENERATING AN EN GREETING
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
