package main

import "fmt"

type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func PrintArea(s Shape) {
    fmt.Printf("Area: %f\n", s.Area())
}

func main() {
    rectangle := Rectangle{Width: 5, Height: 3}
    circle := Circle{Radius: 2.5}

    PrintArea(rectangle)
    PrintArea(circle)
}
