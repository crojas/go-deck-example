package main

import (
	"fmt"
)

func main() {
	//var colors map[string]string
	//colors := make(map[string]string)
	//colors["white"] = "#ffffff"
	//fmt.Println(colors)
	// delete(colors, "white")

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}
