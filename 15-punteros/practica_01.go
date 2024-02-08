package main

import "fmt"

func main() {
	name := "Fernando"
	var p_name *string = &name
	addLastNameAndReplace(p_name, "Gómez")
	fmt.Println(name)
}

func addLastName(firstname *string, lastname string) string {
	return *firstname + " " + lastname
}

func addLastNameAndReplace(firstname *string, lastname string) string {
	*firstname = *firstname + " " + lastname
	return *firstname
}
