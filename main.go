package main

import (
	"Tucil2/CPairSolver"
	"fmt"
)

func main() {
	// p := NewPoint(2, 3)

	// fmt.Println(p.GetX())
	// fmt.Println(p.GetY())
	p1 := CPairSolver.NewPoint(1, 2, 3)

	//fmt.Println(p.GetAxisValue(0))
	p2 := CPairSolver.NewPoint(2, 0, 0)
	p3 := CPairSolver.NewPoint(-1, 999, 999)

	s := CPairSolver.NewSolver(*p1, *p2, *p3)

	fmt.Println("before")
	s.Print()

	s.Sort(0)

	fmt.Println("after")
	s.Print()
}
