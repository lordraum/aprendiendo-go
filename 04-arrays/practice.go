package main

import "fmt"

var numbers [4]int

var otherNumbers [3]int

func main() {

	letters := [3]string{"a", "b", "c"}

	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4

	fmt.Println(numbers, letters)
}
