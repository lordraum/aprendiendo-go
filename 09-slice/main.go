package main

import "fmt"

func main() {
	// Un slice es como un array pero con la longitud no definida
	nums := []int{1, 2, 3}

	colores := [...]string{"Azul", "blanco", "verde", "Rojo", "Naranja"}
	// Crear slice a partir de un array
	subArray := colores[3:]

	// Añadir elementos a un slice con append()
	other_nums := append(nums, 4)

	// Popiedades de los slice
	// Puntero => Donde empieza
	// Longitud
	// Capacidad

	fmt.Println(nums, subArray, other_nums)
	fmt.Printf("len: %v cap: %v, %p /n", len(nums), cap(nums), nums)
}
