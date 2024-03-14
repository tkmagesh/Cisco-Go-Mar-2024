package main

import "fmt"

type Product struct {
	id   int
	name string
	cost float64
}

func main() {

	var pen Product = Product{
		id:   100,
		name: "pen",
		cost: 10,
	}

	// value
	anotherPen := pen // creates a copy
	anotherPen.cost = 90

	fmt.Println(pen.cost)
	fmt.Println(anotherPen.cost)

	// reference (pointer)
	var penPtr *Product
	penPtr = &pen
	// fmt.Println((*penPtr).id)
	fmt.Println(penPtr.id)

	fmt.Printf("Before applying the discount, pen => [ %v ]\n", Format(pen))
	ApplyDiscount(&pen, 10)
	fmt.Printf("After applying the discount, pen => [ %v ]\n", Format(pen))
}

func Format(p Product) string {
	return fmt.Sprintf("id = %d, name = %q, cost = %0.2f", p.id, p.name, p.cost)
}

func ApplyDiscount(p *Product, discountPercentage float64) {
	p.cost = p.cost * ((100 - discountPercentage) / 100)
}
