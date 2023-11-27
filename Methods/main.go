package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

type PointSlice []Point

func (p PointSlice) Map(v int) PointSlice {
	for _, val := range p {
		val.X += float64(v)
		val.Y += float64(v)

	}
	return p
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) Scale(val float64) {
	p.X *= val
	p.Y *= val
}

func main() {

	p := Point{1, 2}
	// q := Point{4, 6}

	// fmt.Println(p.Distance(q))
	p.Scale(2.0)
	// fmt.Println(p)

	sl := PointSlice{Point{1, 2}, Point{3, 4}}
	fmt.Println(sl)
	fmt.Println(sl.Map(4))
}
