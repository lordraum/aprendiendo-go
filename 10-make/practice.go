package main

import "fmt"

func main() {
	colors := make([]string, 0, 5)
	colors = append(colors, "Red")
	colors = append(colors, "Blue")
	colors = append(colors, "Yellow")
	colors = append(colors, "Orange")
	colors = append(colors, "Purple")
	fmt.Println(colors)
}
