package main

import (
	"fmt"
	"math"
)

type shape interface{
	area() float64
}

type square struct{
	side float64
}

type rect struct{
	length float64
	breadth float64
}

type circle struct{
	radius float64
}


func (s square) area() float64{
	return s.side * s.side
}


func (r rect) area() float64{
	return r.length * r.breadth
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}


func printArea(s shape){
	fmt.Println(s.area())
}


func totalArea(shapes ...shape) float64{
	var a float64
	for _, s := range shapes{
		a += s.area()
	}
	return a
}


func init(){
	sq := square{10}
	re := rect{12, 23}
	c := circle{17}
	// fmt.Println("Square:", sq.area())
	// fmt.Println("Rect: ", re.area())
	// fmt.Println("Circle: ", c.area())

	printArea(sq)
	printArea(re)
	printArea(c)

	fmt.Println("Total Area: ", totalArea(sq, re, c))

}