package main

func main() {
	sortTimer := NewExecTimer("QuickSort")
	bruteforceTimer := NewExecTimer("Bruteforce")
	divideAndConquerTimer := NewExecTimer("Divide and Conquer")

	// p1 := NewRandomPoint(4)
	// p2 := NewPoint(2, 0, 0, 9)
	// p3 := NewRandomPoint(4)
	// p4 := NewRandomPoint(4)
	// p5 := NewRandomPoint(4)
	// p6 := NewRandomPoint(4)
	// p7 := NewPoint(3, 99, 1, 4)
	// p8 := NewPoint(-12, -90, 2, 0)

	// by force
	// solver := NewSolver(*p1, *p2, *p3, *p4, *p5, *p6, *p7, *p8)
	solver := NewSolver()
	solver.GeneratePoints(10000, 10)

	sortTimer.Start()

	// fmt.Println("before")
	// s.Print()

	solver.Sort(0)

	// fmt.Println("after")
	// s.Print()

	sortTimer.Finish()
	sortTimer.Tell()

	bruteforceTimer.Start()

	solver.SolveByForce()

	solver.Describe()

	bruteforceTimer.Finish()
	bruteforceTimer.Tell()

	// with divide and conquer

	// s2 := NewSolver()
	// s2.GeneratePoints(10, 3)

	// s2.Sort(0)

	divideAndConquerTimer.Start()

	solver.Solve()

	solver.Describe()

	divideAndConquerTimer.Finish()
	divideAndConquerTimer.Tell()

}
