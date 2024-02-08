package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func prod(x, y int) (result int) {
	result = x * y
	return
}

func main() {
	fmt.Println(sum(2, 3))
	fmt.Println(prod(2, 3))
}
