package main

import "fmt"

func main() {
	/*
		var pen struct {
			id   int
			name string
			cost float64
		}

		// fmt.Printf("%#v\n", pen)
		// fmt.Printf("%+v\n", pen)

		pen.id = 100
		pen.name = "pen"
		pen.cost = 10
	*/

	var pen struct {
		id   int
		name string
		cost float64
	} = struct {
		id   int
		name string
		cost float64
	}{
		id:   100,
		name: "pen",
		cost: 10,
	}

	// fmt.Printf("%+v\n", pen)
	fmt.Println(Format(pen))
}

func Format(p struct {
	id   int
	name string
	cost float64
}) string {
	return fmt.Sprintf("id = %d, name = %q, cost = %0.2f", p.id, p.name, p.cost)
}
