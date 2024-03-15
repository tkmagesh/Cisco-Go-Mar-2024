package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func main() {

	var pen Product = Product{
		Id:   100,
		Name: "pen",
		Cost: 10,
	}

	// fmt.Println(Format(pen))
	fmt.Println(pen.Format())

	fmt.Printf("Before applying the discount, pen => [ %v ]\n", pen.Format())
	// ApplyDiscount(&pen, 10)
	pen.ApplyDiscount(10)
	fmt.Printf("After applying the discount, pen => [ %v ]\n", pen.Format())

}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
