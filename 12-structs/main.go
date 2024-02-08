package main

import "fmt"

// las structs crean 'tipos' personalizados que permiten almacenar diferentes tipos de datos

// Crear struct
type User struct {
	ID   int
	name string
	age  int
}

// Agregar método (función a un struct)
func (u User) say_hello() {
	fmt.Println("Hola me llamo", u.name)
}

func main() {
	// Instanciar struct
	Martha := User{1, "Martha", 20}
	// fmt.Println(Martha)
	fmt.Printf("%+v", Martha) // Imprime key:value
	// Invocar el método de la instancia
	Martha.say_hello()
}
