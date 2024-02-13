package main

import (
	"fmt"
	"math"
)

func main() {
	myCircle := Circle{25}
	myRectangle := Rectangle{10, 5}
	fmt.Println("El perímetro del círculo es:")
	printShapeProperties(myCircle)
	fmt.Println("El área del rectángulo es:")
	printShapeProperties(myRectangle)
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	b float64
	h float64
}

type Shape interface {
	area() float64
	perimeter() float64
}

func (c Circle) perimeter() float64 {
	return 2 * (c.radius * math.Pi)
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.b + r.h)
}

func (r Rectangle) area() float64 {
	return r.b * r.h
}

func printShapeProperties(s Shape) {
	fmt.Printf("Perímetro: %f\n", s.perimeter())
	fmt.Printf("Área: %f\n", s.area())
}
