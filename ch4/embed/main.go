package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type wheel struct {
	Circle
	Spokes int
}

func main() {
	var w wheel
	w = wheel{Circle: Circle{Point: Point{X: 8, Y: 8}, Radius: 5}, Spokes: 20}
	fmt.Printf("%#v\n", w)

	w.X = 42
	fmt.Printf("%#v\n", w)
}
