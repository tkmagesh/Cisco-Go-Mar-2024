package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

type Dummy struct {
	Id int
}

type PerishableProduct struct {
	// Dummy
	Product
	Expiry string
}

// factory for PerishableProduct
func NewPerishableProduct(id int, name string, cost float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

func main() {
	var grapes PerishableProduct = PerishableProduct{
		Product: Product{
			Id:   100,
			Name: "Arabian Grapes",
			Cost: 50,
		},
		Expiry: "2 Days",
	}
	fmt.Printf("%#v\n", grapes)
	/*
		fmt.Println(grapes.Product.Id)
		fmt.Println(grapes.Product.Name)
		fmt.Println(grapes.Product.Cost)
	*/
	/*
		grapes.Id = 100
		grapes.Name = "Arabian Grapes"
		grapes.Cost = 50
		grapes.Expiry = "2 Days"
	*/
	fmt.Println(grapes.Id)
	fmt.Println(grapes.Name)
	fmt.Println(grapes.Cost)
	fmt.Println(grapes.Expiry)

	milk := NewPerishableProduct(101, "Milk", 40, "1 Day")
	fmt.Println(milk)

	fmt.Printf("%T\n", milk)
}
