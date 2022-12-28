package main

import (
	"fmt"

	mat "github.com/shelik/mentoring-test/materials"
	orders "github.com/shelik/mentoring-test/orders"
	shapes "github.com/shelik/mentoring-test/shapes"
)

type Triangle struct {
	height float64
	width  float64
}

func NewTriangle() orders.Shape {

	var t Triangle

	fmt.Println("Enter height:")
	fmt.Scanln(&t.height)
	fmt.Println("Enter width:")
	fmt.Scanln(&t.width)

	return t
}

func (r Triangle) Area() float64 {
	return r.height * r.width
}

func main() {
	shaper := shapes.NewShaper()

	shaper.AddShape("Triangle", NewTriangle)

	order := orders.Order{
		Material: mat.GetMaterial(),
		Surface:  shaper.GetSurface(),
	}

	fmt.Println("Стоимость работ:")
	fmt.Println(order.CalculateCost())

	//bank aok rabota

}
