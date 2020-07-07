package main

import "fmt"

func main() {
	colors := make(map[string]string)

	colors["Red"] = "#FF0000"
	colors["Blue"] = "#0000FF"
	colors["White"] = "#ffffff"

	printMap(colors)
}

func printMap(c map[string]string) {

	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
