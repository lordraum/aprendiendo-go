package main

import "fmt"

func saludar(nombre string) {
	fmt.Println("Hello", nombre, "and welcome")
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) (result int) {
	result = a - b
	return
}

func main() {
	saludar("Fer")
	fmt.Println(sumar(3, 5))
	fmt.Println(restar(3, 5))

}

// Las funciones públicas se declaran con CapitalCase, las privadas cons lowercase
// públicas --> Son accesibles para otros módulos
// privadas --> No son accesibles
