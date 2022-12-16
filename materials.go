package main

import "fmt"

type Satin struct { //satingips
	price   float64 //2$ - kg
	expense float64 //1.5kg - metr
}

type Putty struct { //po derevu shpaklevka
	price   float64 //12$ - kg
	expense float64 //1kg - metr
}

type Plaster struct { //
	price   float64 //1$ - kg
	expense float64 //1kg - m2
}

func (s Satin) Expense() float64 {
	return 1.5

}

func (p Putty) Expense() float64 {
	return 1
}

func (d Plaster) Expense() float64 {
	return 1
}

func (d Plaster) Price() float64 {
	return 25
}

func (s Satin) Price() float64 {
	return 35
}

func (p Putty) Price() float64 {
	return 45
}

func GetMaterial() Materials {
	var mtype int
	var mat Materials

	fmt.Println("Введите тип материала: (1. Satin. )")
	fmt.Scanln(&mtype)

	if mtype == 1 {
		mat = Satin{}
	}

	return mat
}
