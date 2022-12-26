package main

import (
	"fmt"

	mat "github.com/shelik/mentoring-test/materials"
	orders "github.com/shelik/mentoring-test/orders"
	shapes "github.com/shelik/mentoring-test/shapes"
)

func main() {

	order := orders.Order{
		Material: mat.GetMaterial(),
		Surface:  shapes.GetSurface(),
	}

	fmt.Println("Стоимость работ:")
	fmt.Println(order.CalculateCost())

	//bank aok rabota

}
