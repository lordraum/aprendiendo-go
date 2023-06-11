package main

import "fmt"

var isUser string

var dato string = "cualquier cosa"

// Formatear cadena y guardar
var mensaje string = fmt.Sprint(dato, " y otra cosa")

func main() {
	//hola := "Hello"
	//mundo := "World"

	//name := "Fer"
	//age := 40

	// Impresión simple
	//fmt.Print(hola + " " + mundo)

	// Impresión con salto de línea
	//fmt.Println(hola)
	//fmt.Println(mundo)

	// Imprimir y formatear texto
	//fmt.Printf("Mi nombre es %s y tengo %d años", name, age)
	// Cuándo no se sabe el tipo de dato que se va a recibir en el format se usa el vervo de formato %v
	// La función fmt.Sprintf() permite formatear y guardar el dato en variable

	// Conocer el tipo de dato
	// fmt.Printf("%T /n", name)

	// Recibir datos datos por consola
	fmt.Println("¿Es usuario?")
	fmt.Scanln(&isUser)

	if isUser == "Si" {
		fmt.Println("Si es usuario")
	} else {
		fmt.Println("No es usuario")
	}

	fmt.Println(mensaje)
}

// https://youtu.be/R_Ib7FRQUrA?t=1813
