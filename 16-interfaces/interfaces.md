# Interfaces

Indice que agrupa métodos, funciones de `tipos` similares. Por ejemplo la interfaz `Animal` puede amacenar los métodos de los tipos `Perro`, `Gato`, etc.

## Sintáxis

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```
## Crear método para usar la interfaz

```go
func printShapeProperties(s Shape) {
	fmt.Printf("Perímetro: %f\n", s.perimeter())
	fmt.Printf("Área: %f\n", s.area())
}

```

Cualquier `tipo` que tenga un método indexado en la `interfaz` podrá ser usado como el tipo de la interfaz.

Ejemplo

- Interfaz `Figura` --> perimeter(), area() --> método imprimirFigura()
- Tipo `Círculo`, tipo `Rectángulo` --> perimeter(), area()
    - Ambos tipos podrán ser usados como tipo `Shape` en el método de `imprimir figura`

```go
fmt.Println("El perímetro del círculo es:")
printShapeProperties(myCircle)
```

Un ejemplo de uso práctico para una interfaz es para crear los métodos que se necesiten para interactuar con la BBDD, esto facilita el desarrollo independientemente del gestor de BBDD que se utilice.

- Ejercicio
    - Crear ejercicio que establezca la interfaz `Shape`
        - Esta interfaz indexará los métodos `Perimeter` del tipo `Circle` y `Area` del tipo `rectangle`
        - Implementar el tipo triángulo.

