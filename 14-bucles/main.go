package main

import "fmt"

var acc int = 0

// var int result2

func main() {
	nums := [4]int{1, 2, 3, 4}

	// Bucle for
	/* for i := 1; i <= 4; i++ {
		fmt.Println(i * 2)
	} */

	// bucle for range
	for _, v := range nums {
		acc += v
	}

	fmt.Println(acc)
	printNumbers(1, 4)
}

// Búcle while
func printNumbers(init, end int) {
	i := init
	for i <= end {
		fmt.Println(i * 2)
		i++
	}
}
