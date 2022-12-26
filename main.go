package main

import (
	"fmt"
	mat "github.com/shelik/mentoring-test/materials"
)

func main() {

	order := Order{
		material: mat.GetMaterial(),
		surface:  GetSurface(),
	}

	fmt.Println("Стоимость работ:")
	fmt.Println(order.CalculateCost())
	//a bank aok rabota

}
