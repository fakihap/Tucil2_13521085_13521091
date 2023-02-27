package main

import (
	"Tucil2/CPairSolver"
	"fmt"
)

func main() {
	// p := NewPoint(2, 3)

	// fmt.Println(p.GetX())
	// fmt.Println(p.GetY())
	p := CPairSolver.NewPoint(3, []int{1, 2, 3})

	fmt.Println(p.GetAxisValue(0))
}
