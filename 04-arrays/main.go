package main

import "fmt"

// var nombreArray [longitud]tipoDeDatos, los arrays solo pueden contener un tipo de dato.

var miArray [5]int

func main() {

	// Asignar valores (Deben ser declarados en la función main)
	miArray[0] = 1
	miArray[1] = 2
	miArray[2] = 3
	miArray[3] = 4
	miArray[4] = 5

	// Crear y declarar array
	miArray2 := [5]int{6, 7, 8, 9, 10}

	//sustituir valor array
	miArray[0] = 7

	// Crear arreglo sin definir la longitud
	colores := [...]string{"Azul", "blanco", "verde", "Rojo", "Naranja"}

	// Crear arreglo con los índices sin definirlos todos
	monedas := [...]string{0: "Dolar", 2: "Peso", 4: "Euro"}

	// Crear subArray (slice)
	subArray := colores[:3]

	fmt.Println(miArray, miArray2, colores, monedas, subArray)
	fmt.Printf("%T /n", subArray)
}
