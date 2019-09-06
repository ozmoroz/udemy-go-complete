package main

import "fmt"

/** shape interface */
type shape interface {
	getArea() float64
}

/** triangle type */
type triangle struct {
	height float64
	base   float64
}

func (t *triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

/** square type */
type square struct {
	sideLength float64
}

func (s *square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Printf("Shape area is %v\n", s.getArea())
}

func main() {
	t := triangle{height: 12, base: 7.4}
	printArea(&t)

	s := square{sideLength: 12}
	printArea(&s)
}
