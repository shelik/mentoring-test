package shapes

import (
	"fmt"
	"math"

	orders "github.com/shelik/mentoring-test/orders"
)

type Square struct {
	height float64
	width  float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

type ShapeConstructor func() orders.Shape
type ShapesCollection map[string]ShapeConstructor

type Shaper struct {
	shapes ShapesCollection
}

func NewShaper() *Shaper {
	return &Shaper{
		shapes: ShapesCollection{
			"Circle":    NewCircle,
			"Square":    NewSquare,
			"Rectangle": NewRectangle,
		},
	}
}

func (s *Shaper) AddShape(name string, ctr ShapeConstructor) {
	s.shapes[name] = ctr
}

func NewCircle() orders.Shape {

	var c Circle

	fmt.Println("Enter radius:")
	fmt.Scanln(&c.radius)

	return &c
}

func NewSquare() orders.Shape {

	var s Square

	fmt.Println("Enter height:")
	fmt.Scanln(&s.height)
	fmt.Println("Enter width:")
	fmt.Scanln(&s.width)

	return &s
}

func NewRectangle() orders.Shape {

	var r Rectangle

	fmt.Println("Enter height:")
	fmt.Scanln(&r.height)
	fmt.Println("Enter width:")
	fmt.Scanln(&r.width)

	return &r
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r *Rectangle) Area() float64 {
	return r.height * r.width
}

func (s *Square) Area() float64 {
	return s.height * s.width
}

func (s *Shaper) GetSurface() orders.Shape {

	type SelectSurface map[int]string

	ss := SelectSurface{}
	n := 1

	var fstr string
	for key := range s.shapes {
		fstr = fmt.Sprintf("%s\n %d - %s", fstr, n, key)
		ss[n] = key
		n++
	}

	var stype int

	fmt.Printf("Введите тип плоскости:%s\n", fstr)
	fmt.Scanln(&stype)

	KeyShapes := ss[stype]
	ValueShapes := s.shapes[KeyShapes]

	return ValueShapes()
}
