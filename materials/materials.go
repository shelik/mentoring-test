package materials

import (
	"fmt"

	orders "github.com/shelik/mentoring-test/orders"
)

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
	return 1
}

func (s Satin) Price() float64 {
	return 2
}

func (p Putty) Price() float64 {
	return 12
}

func GetMaterial() orders.Materials {
	var mtype int
	var mat orders.Materials

	fmt.Println("Введите тип материала:\n 1 - Satin\n 2 - Putty\n 3 - Plaster")
	fmt.Scanln(&mtype)

	if mtype == 1 {
		mat = Satin{}
	}
	if mtype == 2 {
		mat = Putty{}
	}
	if mtype == 3 {
		mat = Plaster{}
	}

	return mat
}
