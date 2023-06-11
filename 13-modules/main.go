package main

// terminal go env => Muestra todas las variables de entorno que utilza go

// Paquetes ejecutables => todos los package main, no ejecutables => todos los demás

// Crear archivo go.mod => archivo de configuración principal => go mod init nombre-del-modulo

// Dependencias: Se especifican en el archivo go.mod junto con sus versiones, para agregar una nueva dependencia se usa => go get nombre-del-paquete

// Ejecutar el código del módulo => go run .

// Compilar el módulo => ./nombre-del-ejecutable

// También se puede importar estableciendo los otros archivos como paquete

// package miPaquete => import "ruta/al/miPaquete"

func main() {
	morning()
}
