package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

type AreaFinder interface {
	Area() float64
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

// Perimeter
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

type PerimeterFinder interface{ Perimeter() float64 }

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintShapeStats(x ShapeStatsFinder) {
	PrintArea(x)      // => x = interface{ Area() float64 }
	PrintPerimeter(x) // => x = interface{ Perimeter() float64 }
}

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShapeStats(c)

	r := Rectangle{Length: 10, Breadth: 12}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShapeStats(r)

}
