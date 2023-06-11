package main

import "fmt"

func main() {
	v := 4
	var p *int = &v
	*p = 8
	//fmt.Println("valor de v :", v)
	//fmt.Printf("La posición de memoria de de v es: %v \n", &v)
	//fmt.Printf("p apunta hacia: %v cuyo valor es %v", p, *p)
	fmt.Printf("ahora v ha cambiado su valor a: %v  \n", v)

}
