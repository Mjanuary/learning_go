package main

import "fmt"

func main() {

	// var colors map[string]string
	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red": "#uewtw",
	// 	"green": "sdfdfs",
	// }

	colors["color"] = "red"
	colors["background"] = "Green"
	// fmt.Println(colors)

	delete(colors, "color")

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Value for", key, "is", value)
	}
}