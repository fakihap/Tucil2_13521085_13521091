package main

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/disk"
	"strconv"
)

func main() {
	sortTimer := NewExecTimer("QuickSort")
	bruteforceTimer := NewExecTimer("Bruteforce")
	divideAndConquerTimer := NewExecTimer("Divide and Conquer")

	// get user inputs
	IOHandler := NewIOHandler()
	IOHandler.PrintLine("\nUSER CONFIGURATION\n-----------------------")
	n, d, lb, ub := IOHandler.readUserConfig()

	// construct solver
	solver := NewSolver()
	solver.GeneratePoints(n, d, lb, ub)

	// tell specifications
	IOHandler.PrintLine("COMPUTER SPECIFICATIONS\n-----------------------")
	val, _ := cpu.Info()
	IOHandler.PrintLine("CPU: " + val[0].ModelName)
	// get host
	host, _ := host.Info()
	IOHandler.PrintLine("Host: " + host.Hostname)
	// get memory
	v, _ := mem.VirtualMemory()
	IOHandler.PrintLine("Memory: " + strconv.FormatUint(v.Total/1024/1024, 10) + " MB")
	// get disk
	disk, _ := disk.Usage("/")
	IOHandler.PrintLine("Disk: " + strconv.FormatUint(disk.Total/1024/1024, 10) + " MB\n")

	// sort points by x-axis
	sortTimer.Start()

	solver.Sort(0)

	sortTimer.Finish()
	sortTimer.Tell()

	// with bruteforce
	IOHandler.PrintLine("\nWITH BRUTEFORCE\n-----------------------")
	bruteforceTimer.Start()

	solver.SolveByForce()

	solver.Describe()

	bruteforceTimer.Finish()
	bruteforceTimer.Tell()

	// with divide and conquer
	IOHandler.PrintLine("\nWITH DIVIDE AND CONQUER\n-----------------------")
	divideAndConquerTimer.Start()

	solver.Solve()

	solver.Describe()

	divideAndConquerTimer.Finish()
	divideAndConquerTimer.Tell()

	// visualize if 3 dimensions
	if d == 3 {
		if IOHandler.askToVisualize() {
			visualizer := NewVisualizer(solver.points, solver.solutionPoints)
			visualizer.visualize()
		}
	}

	if d != 3 {
		IOHandler.PrintLine("\nUse 3 dimension input for full experience with the visualizer:D")
	}

	IOHandler.PrintLine("\nThank you for using this program.")
}
