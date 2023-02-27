package CPairSolver

import "fmt"

type Solver struct {
	points    []Point
	dimension int
}

func NewSolver(points ...Point) *Solver {
	s := new(Solver)

	s.points = points
	s.dimension = len(points)

	return s
}

func getPartition(slice []Point, a, b, axis int) int {
	curPiv := slice[b].GetAxisValue(axis)

	i := a - 1

	for j := a; j < b; j++ {
		if slice[j].GetAxisValue(axis) <= curPiv {
			i = i + 1

			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[b] = slice[b], slice[i+1]

	return i + 1
}

func (s Solver) quickSort(a, b, axis int) {
	if a < b {
		p := getPartition(s.points, a, b, axis)

		s.quickSort(a, p-1, axis)
		s.quickSort(p+1, b, axis)
	}
}

func (s Solver) Sort(axis int) {
	s.quickSort(0, len(s.points)-1, axis)
}

func (s Solver) Print() {
	for _, point := range s.points {
		for _, val := range point.val {
			fmt.Print(val, " ")
		}

		fmt.Println()
	}
}
