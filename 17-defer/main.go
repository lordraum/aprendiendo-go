package main

import "fmt"

func main() {
	// defer retrasa la ejecución de una función
	//defer fmt.Println("Esto se imprimirá al final")
	//fmt.Println("Esto se iimprimirá después")

	multiples_defer()
}

func multiples_defer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	// output 3, 2, 1
}

func argumentos_a_defer() {
	a := 10
	defer fmt.Println("El valor de a es:", a)
	a = 20
	// input 10
	/* Esta capacidad es útil para garantizar que las funciones diferidas utilicen los valores correctos al ejecutarse */
}
