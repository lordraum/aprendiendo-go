package main

import "fmt"

func main() {
	// Parámetros para crear slice con make
	// slice, len, cap (opcional)
	numeros := make([]int, 0, 3)

	fmt.Println(numeros, len(numeros), cap(numeros))
}
