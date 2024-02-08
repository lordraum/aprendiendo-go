package main

import "fmt"

const nombre = "Fernando"
const apellido string = "Gómez"

func main() {
	edad := 40
	// fmt.Printf("Hola mi nombre es %s %s y tengo %d años \n", nombre, apellido, edad)
	fmt.Println("Hola mi nombre es", nombre, apellido, "y tengo", edad, "años")
}
