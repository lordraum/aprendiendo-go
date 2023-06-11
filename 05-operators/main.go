package main

import "fmt"

/*
Operadores aritméticos: +, -, *, /, %.
Operadores de comparación: ==, !=, <, >, <=, >=.
Operadores lógicos: &&, ||, !.
Operadores de bits: &, |, ^, <<, >>.
Operadores de asignación: =, +=, -=, *=, /=, %=, <<=, >>=, &=, ^=, |=.
Otros operadores: . (para acceder a campos o métodos de una estructura o tipo), () (para invocar una función o método).
*/

// Verbos de formato
// Especifican el formato de salida esperado en una cadena de texto
/*
%d: para imprimir enteros con signo.
%f: para imprimir números de punto flotante.
%s: para imprimir cadenas de caracteres.
%t: para imprimir valores booleanos.
%v: para imprimir el valor de cualquier tipo de dato.
*/

func main() {
	a := 5
	b := 3
	sum := a + b

	nombre := "Juan"
	apellido := "Pérez"
	edad := 30

	// Concatenar expresiones con función Sprintf()
	mensaje := fmt.Sprintf("Mi nombre es %s %s y tengo %d años.", nombre, apellido, edad)

	fmt.Println(sum)
	fmt.Println(mensaje)
}
