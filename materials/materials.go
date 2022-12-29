package materials

import (
	"fmt"

	orders "github.com/shelik/mentoring-test/orders"
)

type Material struct { //satingips
	price   float64 //2$ - kg
	expense float64 //1.5kg - metr
}

type MaterialConstructor func() orders.Materials
type MaterialsCollection map[string]MaterialConstructor

type Materialer struct {
	materials MaterialsCollection
}

func NewMaterialer() *Materialer {

	return &Materialer{
		materials: MaterialsCollection{
			"Satin":   NewSatin,
			"Putty":   NewPutty,
			"Plaster": NewPlaster,
		},
	}
}

func NewSatin() orders.Materials {
	return &Material{
		price:   2.0,
		expense: 1.5,
	}
}

func NewPutty() orders.Materials {
	return &Material{
		price:   12,
		expense: 1,
	}
}

func NewPlaster() orders.Materials {
	return &Material{
		price:   1,
		expense: 1,
	}
}

func (m *Material) Expense() float64 {
	return m.expense

}

func (m *Material) Price() float64 {
	return m.price
}

func (m *Materialer) AddMaterial(name string, ctr MaterialConstructor) {
	m.materials[name] = ctr
}

func (m *Materialer) GetMaterial() orders.Materials {

	type SelectMaterial map[int]string

	sm := SelectMaterial{}
	n := 1

	var fstr string
	for key := range m.materials {
		fstr = fmt.Sprintf("%s\n %d - %s", fstr, n, key)
		sm[n] = key
		n++
	}

	var stype int

	fmt.Printf("Введите тип Материала:%s\n", fstr)
	fmt.Scanln(&stype)

	KeyMaterials := sm[stype]
	ValueMaterials := m.materials[KeyMaterials]

	return ValueMaterials()

}
