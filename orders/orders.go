package orders

type Shape interface {
	Area() float64
}

type Materials interface {
	Expense() float64
	Price() float64
}

type Order struct {
	Material Materials
	Surface  Shape
}

func (o *Order) CalculateCost() float64 {
	return o.Surface.Area() * o.Material.Expense() * o.Material.Price()
}
