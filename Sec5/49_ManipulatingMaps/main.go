package main

import "fmt"

func main() {
	//There are many ways to start a map

	// var colors map[string]string

	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	//Adding an item
	colors["white"] = "#ffffff"
	fmt.Println(colors)

	//Delete an item
	delete(colors, "white")
	fmt.Println(colors)
}
