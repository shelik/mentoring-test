package main

import (
	"fmt"

	mat "github.com/shelik/mentoring-test/materials"
	orders "github.com/shelik/mentoring-test/orders"
	shapes "github.com/shelik/mentoring-test/shapes"
)

type deco struct {
	price   float64
	expense float64
}

func newDeco() orders.Materials {
	return &deco{
		price:   2,
		expense: 2,
	}
}

func (d *deco) Expense() float64 {
	return d.expense

}

func (d *deco) Price() float64 {
	return d.price
}

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
	materialer := mat.NewMaterialer()
	shaper.AddShape("Triangle", NewTriangle)
	materialer.AddMaterial("deco", newDeco)

	order := orders.Order{
		Material: materialer.GetMaterial(),
		Surface:  shaper.GetSurface(),
	}

	fmt.Println("Стоимость работ:")
	fmt.Println(order.CalculateCost())

	//bank aok rabota

}
