package main

import "fmt"

func main() {
	a := 1
	x := 7
	y := 12
	mensaje := ""

	// if

	if a > 2 {
		mensaje = " El número es mayor que 2"
	} else {
		mensaje = " El número no es mayor que 2"
	}

	// else if
	if x < 0 {
		fmt.Println("x es negativo")
	} else if x == 0 {
		fmt.Println("x es cero")
	} else {
		fmt.Println("x es positivo")
	}

	// if anidados
	if y > 0 {
		if y < 10 {
			fmt.Println("y es un número positivo de un solo dígito")
		} else {
			fmt.Println("y es un número positivo de más de un dígito")
		}
	} else {
		fmt.Println("y no es un número positivo")
	}

	fmt.Println(mensaje)

	// if short statement --> Permite asignar una variable, solo se puede utilizar dentro del scope del if

	if v := a + x; v > 5 {
		fmt.Println(v, "es la suma de", a, "+", x, "y es mayor a 5")
	}
}

// Switch

/*
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
*/

/*
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
*/

/*
Switch sin condición
*/

/* package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
} */
