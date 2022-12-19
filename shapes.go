package main

import (
	"fmt"
	"math"
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

func NewCircle() Circle {
	var c Circle
	fmt.Println("Enter radius:")
	fmt.Scanln(&c.radius)

	return c
}

func NewSquare() Square {
	var s Square
	fmt.Println("Enter height:")
	fmt.Scanln(&s.height)
	fmt.Println("Enter width:")
	fmt.Scanln(&s.width)

	return s
}

func NewRectangle() Rectangle {
	var r Rectangle
	fmt.Println("Enter height:")
	fmt.Scanln(&r.height)
	fmt.Println("Enter width:")
	fmt.Scanln(&r.width)

	return r
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (s Square) Area() float64 {
	return s.height * s.width
}

func GetSurface() Shape {
	var stype int
	var surf Shape

	fmt.Println("Введите тип плоскости:\n 1 - Circle\n 2 - Square\n 3 - Rectangle")
	fmt.Scanln(&stype)

	if stype == 1 {
		surf = NewCircle()
	}
	if stype == 2 {
		surf = NewSquare()
	}
	if stype == 3 {
		surf = NewRectangle()
	}

	return surf
}
