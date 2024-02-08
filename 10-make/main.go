package main

import "fmt"

// Make es una función que se utiliza prncipalmente para crear slice

func main() {
	// Parámetros para crear slice con make
	// slice, len, cap (opcional)
	numeros := make([]int, 0, 3)
	numeros = append(numeros, 1)
	numeros = append(numeros, 2)
	numeros = append(numeros, 3)
	numeros = append(numeros, 4)

	fmt.Println(numeros, len(numeros), cap(numeros))
}
