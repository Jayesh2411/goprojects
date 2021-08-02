package main

import "fmt"

func main() {
	var b ball
	b = values{4.0, 5.0}
	fmt.Println(b.volume())
}

type ball interface {
	volume() float64
}

type values struct {
	radius float64
	height float64
}

func (v values) volume() float64 {
	return v.radius * v.radius * 3.14 * v.height
}
