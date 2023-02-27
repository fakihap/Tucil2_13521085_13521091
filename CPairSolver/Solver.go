package CPairSolver

import (
	"fmt"
	"math"
)

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

// sort utility
// we use quicksort
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

// Print prints points saved on solver
func (s Solver) Print() {
	for _, point := range s.points {
		for _, val := range point.val {
			fmt.Print(val, " ")
		}

		fmt.Println()
	}
}

func getDelta(p1, p2 Point) float64 {
	// assuming both points have the same order of dimension

	var d float64

	for i := 0; i < p1.dimension; i++ {
		d += math.Pow(float64(p2.GetAxisValue(i)-p1.GetAxisValue(i)), 2)
	}

	return math.Sqrt(d)
}

func GetClosestByForce(points ...Point) (Point, Point, float64) {
	delta := getDelta(points[0], points[1])
	idA, idB := 0, 1

	if len(points) > 3 {
		fmt.Print("\nWARNING BRUTEFORCE USED FOR A SLICE OF ", len(points), " POINTS", "\n\n")
	}

	for id1, p1 := range points {
		for id2, p2 := range points[id1:] {
			if p1.ID == p2.ID {
				continue
			}

			dist := getDelta(p1, p2)

			if delta > dist {
				delta = dist
				idA, idB = id1, id2
			}
		}
	}

	return points[idA], points[idB], delta
}
