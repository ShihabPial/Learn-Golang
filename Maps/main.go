package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":    "#ff0000",
		"yellow": "#ffff00",
		"blue":   "#0066ff",
	}
	colors["white"] = "ffffff"

	vehicles := make(map[string]string)
	vehicles["Toyota"] = "Camry"
	vehicles["Audi"] = "S7"

	//delete(colors, "blue")

	printMap(colors)
	//fmt.Println(vehicles)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
