package main

import "fmt"

// Se declaran con la palabra clave var
// Es opcional declarar el tipo de dato, si no se declara go utiliza la inferencia de datos para asignarlo, es buena práctica utilizar la inferencia siempre que se pueda

var nombre string = "Fernando"
var edad int = 39
var a string = "prueba"

// las constantes se declaran con la palabra resevada const

const port = 8090

func main() {
	// Variable inicializada :=
	lastName := "Gómez"
	dineroRestante := 7.52
	buleano := 5 > 3
	fmt.Println(nombre, lastName, edad, buleano, port)
	fmt.Println(buleano)
	// Imprimir el tipo de dato de una variable
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", dineroRestante)
	fmt.Printf("%T\n", buleano)
	fmt.Printf("%T\n", port)
}
