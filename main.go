package main

import (
	"fmt"
)

func main() {

	order := Order{
		material: GetMaterial(),
		surface:  GetSurface(),
	}
	fmt.Println("Стоимость работ:")
	fmt.Println(order.CalculateCost())
	//aaaa
	//ab
}
