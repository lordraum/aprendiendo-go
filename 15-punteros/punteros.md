# Golang punteros
Los punteros guardan y leen datos de una posición de memoria.

*// archivo => punteros_y_variables.go*

## Posición de memoria (&)
Al crear una variable, se asigna una posición de memoria, en la cuál se almacena su valor. El caracter ambersand `&` se usa para acceder a la posición de esta variable.

```go
fmt.Printf("La posición de memoria de de v es %v : n/", &v)
// La posición de memoria de de v es: 0xc00009a058
```
## Declaración de puntero y desreferenciación (*)

Se declara con el caracter `*`, al crearse no se asigna una posición de memoria, tiene la capacidad de apuntar hacia una posición de memoria.

```go
v := 4
var p *int = &v
fmt.Printf("p apunta hacia: %v cuyo valor es %v", p, *p)
// p apunta hacia: 0xc00009a058 cuyo valor es 4
```
`El puntero p` referencia una posición de memoria `al 'desreferenciar' => *p` devuelve el valor de esta posición.

## Cambio de valor en la variable original
Al indicar al puntero que asigne un nuevo valor a lo posición a la cual apunta, cambiará también el valor de la variable original.

```go
*p = 8
fmt.Printf("ahora v ha cambiado su valor a: %v  \n", v)
```
[continúa en](https://www.youtube.com/watch?v=NZZ_Yrdd-2U&t=1709s)

