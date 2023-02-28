package main

func main() {
	sortTimer := NewExecTimer("QuickSort")
	bruteforceTimer := NewExecTimer("Bruteforce")
	divideAndConquerTimer := NewExecTimer("Divide and Conquer")

	InputHandler := NewInputHandler()
	n, d, lb, ub := InputHandler.readUserConfig()

	// by force
	solver := NewSolver()
	solver.GeneratePoints(n, d, lb, ub)

	sortTimer.Start()

	solver.Sort(0)

	sortTimer.Finish()
	sortTimer.Tell()

	bruteforceTimer.Start()

	solver.SolveByForce()

	solver.Describe()

	bruteforceTimer.Finish()
	bruteforceTimer.Tell()

	// with divide and conquer
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
