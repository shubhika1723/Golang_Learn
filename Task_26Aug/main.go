package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func LargestArea(shapes []Shape) Shape {
	if len(shapes) == 0 {
		return nil
	}
	largest := shapes[0]
	for _, s := range shapes {
		if s.Area() > largest.Area() {
			largest = s
		}
	}
	return largest
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 7},
		Rectangle{Width: 3, Height: 8},
	}

	for _, shape := range shapes {
		fmt.Printf("Area: %.2f, Perimeter: %.2f\n", shape.Area(), shape.Perimeter())
	}

	largest := LargestArea(shapes)
	fmt.Printf("\nShape with largest area: Area=%.2f, Perimeter=%.2f\n", largest.Area(), largest.Perimeter())
}
