package main

import (
	"fmt"
)

func main() {
	p1 := NewRandomPoint(3)
	p2 := NewPoint(2, 0, 0)
	p3 := NewRandomPoint(3)

	s := NewSolver(*p1, *p2, *p3)

	fmt.Println("before")
	s.Print()

	s.Sort(0)

	fmt.Println("after")
	s.Print()

	et := NewExecTimer("Process")

	et.Start()

	pa, pb, dist := GetClosestByForce(*p1, *p2, *p3)

	fmt.Println(pa, pb, dist)

	et.Finish() // on progress
	et.Tell()
}
