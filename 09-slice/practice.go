package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	subNumbers := numbers[:2]
	colors := []string{"rojo", "verde", "azúl", "naranja"}
	newColors := append(colors, "amarillo")
	fmt.Println(numbers, subNumbers, colors, newColors)
}
