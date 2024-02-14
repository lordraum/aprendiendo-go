package main

import (
	"fmt"
	"math"
)

func main() {
	myCircle := Circle{25}
	myRectangle := Rectangle{10, 5}
	myTriangle := Triangle{10, 5, 6, 8}
	fmt.Println("Perímetro y area del círculo")
	printShapeProperties(myCircle)
	fmt.Println("Perímetro y area del rectángulo")
	printShapeProperties(myRectangle)
	fmt.Println("Perímetro y area del triángulo")
	printShapeProperties(myTriangle)
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	b float64
	h float64
}

type Triangle struct {
	b  float64
	h  float64
	bh float64 // l1
	hc float64 // l2
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

func (t Triangle) area() float64 {
	return t.b * t.h / 2
}

func (t Triangle) perimeter() float64 {
	return t.bh + t.hc + t.b
}

func (r Rectangle) area() float64 {
	return r.b * r.h
}

func printShapeProperties(s Shape) {
	fmt.Printf("Perímetro: %f\n", s.perimeter())
	fmt.Printf("Área: %f\n", s.area())
}
