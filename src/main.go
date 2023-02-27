package main

import (
	"fmt"
)

func main() {
	et := NewExecTimer("Bruteforce")
	et2 := NewExecTimer("Divide and Conquer")

	p1 := NewRandomPoint(4)
	p2 := NewPoint(2, 0, 0, 9)
	p3 := NewRandomPoint(4)
	p4 := NewRandomPoint(4)
	p5 := NewRandomPoint(4)
	p6 := NewRandomPoint(4)
	p7 := NewPoint(3, 99, 1, 4)
	p8 := NewPoint(-12, -90, 2, 0)

	// by force
	s := NewSolver(*p1, *p2, *p3, *p4, *p5, *p6, *p7, *p8)

	fmt.Println("before")
	s.Print()

	s.Sort(0)

	fmt.Println("after")
	s.Print()

	et.Start()

	s.SolveByForce()

	if s.solutionFound {
		fmt.Println(s.solutionPoints, s.solutionDist)
	}

	et.Finish()
	et.Tell()

	// with divide and conquer

	s2 := NewSolver(*p1, *p2, *p3, *p4, *p5, *p6, *p7, *p8)

	s2.Sort(0)

	et2.Start()

	s2.Solve()

	if s2.solutionFound {
		fmt.Println(s2.solutionPoints, s2.solutionDist)
	}

	et2.Finish()
	et2.Tell()

}
