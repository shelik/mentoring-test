package orders

type Shape interface {
	Area() float64
}

type Materials interface {
	Expense() float64
	Price() float64
}

type Order struct {
	material Materials
	surface  Shape
}

func (o *Order) CalculateCost() float64 {
	return o.surface.Area() * o.material.Expense() * o.material.Price()
}
