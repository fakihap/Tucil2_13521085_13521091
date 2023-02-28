package main

import (
	"github.com/shirou/gopsutil/v3/cpu"
)

func main() {
	sortTimer := NewExecTimer("QuickSort")
	bruteforceTimer := NewExecTimer("Bruteforce")
	divideAndConquerTimer := NewExecTimer("Divide and Conquer")

	// get user inputs
	InputHandler := NewInputHandler()
	n, d, lb, ub := InputHandler.readUserConfig()

	// construct solver
	solver := NewSolver()
	solver.GeneratePoints(n, d, lb, ub)

	// tell specifications
	val, _ := cpu.Info()
	InputHandler.PrintLine("CPU Used: " + val[0].ModelName)

	// sort points by x-axis
	sortTimer.Start()

	solver.Sort(0)

	sortTimer.Finish()
	sortTimer.Tell()

	// with bruteforce
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

	// visualize if 3 dimensions
	if d == 3 {
		if InputHandler.askToVisualize() {
			visualizer := NewVisualizer(solver.points, solver.solutionPoints)
			visualizer.visualize()
		}
	}

	InputHandler.PrintLine("\nThank you for using this program.")
}
