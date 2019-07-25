package main

import "fmt"

type colorMap map[string]string

func main() {
	colors := colorMap{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	colors.print()
}

func (c colorMap) print() {
	for k, v := range c {
		fmt.Printf("Hex code for color %s is %s\n", k, v)
	}
}
