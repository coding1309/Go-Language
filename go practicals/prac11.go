package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

const pi = 3.14159

func (c Circle) Area() float64 {
	return pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * pi * c.radius
}

type Triangle struct {
	base, height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (t Triangle) Perimeter() float64 {
	return t.base + 2*math.Sqrt(t.height*t.height+t.base*t.base/4)
}
func printInfo(s Shape) {
	fmt.Println("Area =", s.Area())
	fmt.Println("Perimeter =", s.Perimeter())
}

func main() {
	r := Rectangle{width: 4, height: 5}
	fmt.Println("Rectangle :")
	printInfo(r)
	fmt.Println("\n")

	c := Circle{radius: 5}
	fmt.Println("Circle :")
	printInfo(c)
	fmt.Println("\n")

	t := Triangle{base: 3, height: 4}
	fmt.Println("Triangle :")
	printInfo(t)
	fmt.Println("\n")
}
