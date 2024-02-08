package main

import "fmt"

var user = "Fer"
var isUser bool

func main() {

	if user == "Fer" {
		fmt.Println("Usuario correcto")
	} else if user == "Valcho" {
		fmt.Println("Vete de aquí Chalvo!!!")

	} else {
		fmt.Println("Ingresa un usuario correcto")
	}

}
