package main

import "fmt"

func inc(num int) {
	num++
	fmt.Printf("Si incremento el valor de v sería igual a : %d  \n", num)
}

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

// var p *int = &v
