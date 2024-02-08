package main

import "fmt"

func main() {
	a := 1
	x := 7
	y := 12
	mensaje := ""

	// if

	if a > 2 {
		mensaje = " El número es mayor que 2"
	} else {
		mensaje = " El número no es mayor que 2"
	}

	// else if
	if x < 0 {
		fmt.Println("x es negativo")
	} else if x == 0 {
		fmt.Println("x es cero")
	} else {
		fmt.Println("x es positivo")
	}

	// if anidados
	if y > 0 {
		if y < 10 {
			fmt.Println("y es un número positivo de un solo dígito")
		} else {
			fmt.Println("y es un número positivo de más de un dígito")
		}
	} else {
		fmt.Println("y no es un número positivo")
	}

	fmt.Println(mensaje)
}
