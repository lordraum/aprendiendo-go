package main

import "fmt"

// Map => Colección desordenada tipo clave/valor
// Crear con make(map[typeKey]typeValue)
/*
// Crear con asignación
m := map[TipoClave]TipoValor{
  clave1: valor1,
  clave2: valor2,
  clave3: valor3,
}

*/
// borrar =>  delete(map, key)

// Se puede usar arrays, slices, etc como valor

func main() {
	dias := make(map[int]string)

	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[3] = "Miercoles"
	dias[4] = "Jueves"
	dias[5] = "Viernes"
	dias[6] = "Sábado"
	dias[7] = "Domingo"

	delete(dias, 5)

	notas_estudiantes := make(map[string][]int)

	notas_estudiantes["Alex"] = []int{4, 5, 4}
	notas_estudiantes["Javier"] = []int{3, 2, 4}

	fmt.Println(dias, notas_estudiantes, notas_estudiantes["Alex"][0])
}
