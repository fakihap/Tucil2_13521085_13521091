package main

func main() {
	sortTimer := NewExecTimer("QuickSort")
	bruteforceTimer := NewExecTimer("Bruteforce")
	divideAndConquerTimer := NewExecTimer("Divide and Conquer")

	InputHandler := NewInputHandler()
	n, d, lb, ub := InputHandler.readUserConfig()

	// by force
	// solver := NewSolver(*p1, *p2, *p3, *p4, *p5, *p6, *p7, *p8)
	solver := NewSolver()
	solver.GeneratePoints(n, d, lb, ub)

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
	
	// visualize
	if InputHandler.askToVisualize() {
		visualizer := NewVisualizer(solver.points, solver.solutionPoints)
		visualizer.visualize()
	} else {
		InputHandler.PrintLine("\nThank you for using this program.")
	}
}
