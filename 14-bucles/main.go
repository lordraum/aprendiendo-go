package main

import "fmt"

var acc int = 0

// var int result2

func main() {
	nums := [4]int{1, 2, 3, 4}

	// Bucle for
	for i := 1; i <= 4; i++ {
		fmt.Println(i * 2)
	}

	// bucle for range
	for _, v := range nums {
		acc += v
	}

	fmt.Println(acc)
}

// Función con bucle for donde itera según una condición, no requiere variable iteradora (i)

func Between(a, b int) (c []int) {
	for a <= b {
		c = append(c, a)
		a++
	}
	return
}
