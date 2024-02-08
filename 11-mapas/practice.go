package main

import "fmt"

func main() {
	repertorio := make(map[string][]string)
	repertorio["vallenato"] = []string{"La casa en el aire", "La despedida", "La brasilera"}
	repertorio["colombiana"] = []string{"Hurí", "Las Acacias", "Espumas"}
	fmt.Println(repertorio)
}
