package main

import (
	"fmt"
)

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// colors["black"] = "#000000"
	// delete(colors, "white")

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
