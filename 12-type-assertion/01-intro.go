package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

func main() {
	// var x interface{}
	var x any
	x = 100
	x = "Veniam exercitation velit duis proident sint."
	x = true
	x = 99.99
	x = Product{Id: 100, Name: "Pen", Cost: 10}
	fmt.Println(x)

	x = 200
	// x = "Duis ad ipsum adipisicing consectetur veniam amet anim laboris fugiat et duis nulla."
	// y := x.(int) * 2

	// type assertion (if)
	if val, ok := x.(int); ok {
		y := val * 2
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// type assertion using type-switch
	// x = 100
	// x = "Veniam exercitation velit duis proident sint."
	// x = true
	// x = 99.99
	// x = Product{Id: 100, Name: "Pen", Cost: 10}
	x = struct{}{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case bool:
		fmt.Println("x is a bool, !x = ", !val)
	case float64:
		fmt.Println("x is a float64, x * 0.99 =", val*0.99)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case Product:
		fmt.Println("x is a product, x :", val.Format())
	default:
		fmt.Println("Unknown type")
	}

}
