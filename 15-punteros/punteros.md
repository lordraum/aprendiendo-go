# Golang punteros
Los punteros guardan y leen datos de una posición de memoria.

## Punteros y variables

### Posición de memoria (&)
Al crear una variable, se asigna una posición de memoria, en la cuál se almacena su valor. El caracter ambersand `&` se usa para acceder a la posición de esta variable.

```go
fmt.Printf("La posición de memoria de de v es %v : n/", &v)
// La posición de memoria de de v es: 0xc00009a058
```
### Declaración de puntero y desreferenciación (*)

Se declara con el caracter `*`, al crearse no se asigna una posición de memoria, tiene la capacidad de apuntar hacia una posición de memoria.

```go
v := 4
var p *int = &v
fmt.Printf("p apunta hacia: %v cuyo valor es %v", p, *p)
// p apunta hacia: 0xc00009a058 cuyo valor es 4
```
`El puntero p` referencia a una posición de memoria `al 'desreferenciar' => *p` devuelve el valor de esta posición.

### Cambio de valor en la variable original
Al indicar al puntero que asigne un nuevo valor a lo posición a la cual apunta, cambiará también el valor de la variable original.

```go
*p = 8
fmt.Printf("ahora v ha cambiado su valor a: %v  \n", v)
```
[continúa en](https://www.youtube.com/watch?v=NZZ_Yrdd-2U&t=1709s)

## Funciones

### Parámetros por valor
Es la forma que se usa en la mayoría de lenguajes, no hace uso de punteros. El compilador crea otra posición de memoría para el parámetro, copiando el valor a esa nueva posición, generando un coste computacional.

El valor de la variable original, no cambia.

```go
func inc(num int) {
	num++
	fmt.Printf("Se ha incrementado el valor de v a: %d  \n", num)
}

func main() {
	v := 4
	inc(v)
}
```
### Parámetros por referencia

Se define un puntero cómo Parámetro `num *int` y se trabaja con el dato también como puntero `*p++`(desreferenciación). Esto no crea una nueva posición de memoria para realizar el proceso, haciendo más rápida la compilación.

Al apuntar a la posición de memoria de la variable original el valor de esta se sobreescribirá.

```go
func refInc(num *int) {
	*num++
	fmt.Printf("Esto incrementa el valor de v a : %v  \n", *num)
}

func main() {
	v := 4
	var p *int = &v
	refInc(p)
	//fmt.Printf("Valor de v sigue siendo %d \n", v)
	fmt.Printf("Valor de v se ha incrementado a : %d \n", v)
}
```

Continúa en [lista de clientes  con struct y punteros](https://youtu.be/NZZ_Yrdd-2U?t=1120)

