# Interfaces

Reglas a cumplir para que dos partes trabajen en conjunto. Funciona como un `tipo` que implementa un conjunto de métodos.

Se puede también entender la interfaz como un índice que agrupa métodos, funciones, de `tipos` similares. Por ejemplo la interfaz `Animal` puede amacenar los métodos de los tipos `Perro`, `Gato`, etc.

## Sintáxis

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

// Continúa --> https://www.youtube.com/watch?v=kWTLoUTGqSw

- Ejercicio
    - Crear ejercicio que establezca la interfaz `Shape`
        - Esta interfaz indexará los métodos `Perimeter` del tipo `Circle` y `Area` del tipo `rectangle`