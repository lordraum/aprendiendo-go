package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	start := time.Now()

	wg := &sync.WaitGroup{}

	for i := range 10 {
		wg.Add(1)
		go showGoroutine(i, wg)
	}

	wg.Wait()
	duration := time.Since(start).Milliseconds()
	fmt.Println("Duración es:", duration, "ms")
}

func showGoroutine(id int, wg *sync.WaitGroup) {
	delay := rand.Intn(500)
	fmt.Printf("Goroutine #%d width %dms\n", id, delay)
	time.Sleep(time.Millisecond * time.Duration(delay))
	wg.Done()
}

// Goroutines
// Una de las formas de manejar la concurrencia en go
// Se declara con la palabra reservada 'go'
// Al declarar una función como Go rutina se ejecutará primero que las demás funciones

// Ejemplo time rand --> Sin la Goroutine se ejecuta en orden secuencial, con la Goroutine se ejecuta con respecto al tiempo.

// sync/waitGroup
// --> wg.Wait
// Esto hace que espere a que se ejecuten todas las GoRoutines
// --> wg.add(item)
// Añade una goRoutine al waitGroup
//wg.done()
// Hace conteo regresivo del waitGroup a medida que se van ejecutando las Goroutines

// Indicar que una goRoutine finaliza
// Añadir un argumento tipo sync.WaitGroup a la Goroutine y añadirlo en la ejecución (en este caso búcle for)

// Solucionar problema de que una Goroutine puede estar necesitando un dato de otra

/* https://youtu.be/hnBy2kA3MJU?t=553 */

// Estudiar módulos --> math/rand, time
