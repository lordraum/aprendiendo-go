package main

import (
	"fmt"
	"os"
)

// Go no utiliza excepciones como try catch para el manejo de errores
// Go utiliza valores de errores para identificar posibles problemas en la ejecución (tipo error)
// os.ReadFile() --> Lee archivos

func main() {
	// Utilizar err como variable de error
	texto, err := os.ReadFile("prueba_fake.txt")
	if err != nil {
		fmt.Println("Error cargando el archivo", err)
		return
	}
	fmt.Println(string(texto))
}
