package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	//Delete an item
	delete(colors, "white")
	fmt.Println(colors)
}
