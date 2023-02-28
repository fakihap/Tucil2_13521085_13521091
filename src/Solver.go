package main

import (
	"fmt"
	"math"
)

var euclidOpsCount int

type Solver struct {
	points    []Point
	dimension int

	isInitialized bool

	solutionFound  bool
	solutionPoints [2]Point
	solutionDist   float64
}

func NewSolver(points ...Point) *Solver {
	s := new(Solver)

	if len(points) == 0 {
		s.dimension = 0
		s.isInitialized = false
	} else {
		s.dimension = points[0].dimension
		s.isInitialized = true
	}

	s.points = points
	s.solutionFound = false

	return s
}

func (s *Solver) GeneratePoints(nPoints, nDimension int, lowerBound, upperBound float64) {
	var tempPoints []Point

	for i := 0; i < nPoints; i++ {
		tempPoints = append(tempPoints, *NewRandomPoint(nDimension, lowerBound, upperBound))
	}

	s.dimension = tempPoints[0].dimension
	s.isInitialized = true

	s.points = tempPoints
	s.solutionFound = false
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

func (s *Solver) quickSort(a, b, axis int) {
	if a < b {
		p := getPartition(s.points, a, b, axis)

		s.quickSort(a, p-1, axis)
		s.quickSort(p+1, b, axis)
	}
}

func (s *Solver) Sort(axis int) {
	s.quickSort(0, len(s.points)-1, axis)
}

// Print prints points saved on solver
func (s *Solver) Print() {
	for _, point := range s.points {
		for _, val := range point.val {
			fmt.Print(val, " ")
		}

		fmt.Println()
	}
}

func getEuclideanDistance(p1, p2 Point) float64 {
	// assuming both points have the same order of dimension

	var d float64

	euclidOpsCount++

	for i := 0; i < p1.dimension; i++ {
		d += math.Pow(float64(p2.GetAxisValue(i))-float64(p1.GetAxisValue(i)), 2)
	}

	return math.Sqrt(d)
}
func getEuclideanDistanceAxis(p1, p2 Point, axis int) float64 {
	// assuming both points have the same order of dimension

	var d float64

	euclidOpsCount++

	for i := axis; i < p1.dimension; i++ {
		d += math.Pow(float64(p2.GetAxisValue(i))-float64(p1.GetAxisValue(i)), 2)
	}

	return math.Sqrt(d)
}

func getClosestByForce(points ...Point) (Point, Point, float64) {
	delta := getEuclideanDistance(points[0], points[1])
	idA, idB := 0, 1

	if len(points) > 3 {
		fmt.Print("\nWARNING BRUTEFORCE USED FOR A SLICE OF ", len(points), " POINTS", "\n\n")
	}

	for id1, p1 := range points {
		for id2, p2 := range points[id1+1:] {
			if p1.ID == p2.ID {
				continue
			}

			dist := getEuclideanDistance(p1, p2)

			if delta > dist {
				delta = dist
				idA, idB = id1, id2+id1+1
			}
		}
	}

	return points[idA], points[idB], delta
}

func (p Point) getPointsInRange(points []Point, d float64, curr_dimension int) []Point {
	var pointsInRange []Point
	if curr_dimension == 1 {
		for _, other := range points {
			if p.ID == other.ID { // dont compare with itself
				continue
			}
			axis := p.dimension - curr_dimension
			if math.Abs(p.GetAxisValue(axis) - other.GetAxisValue(axis)) <= d {
				pointsInRange = append(pointsInRange, other)
			}
		}
	} else {
		candidates := p.getPointsInRange(points, d, curr_dimension - 1)
		axis := p.dimension - curr_dimension
		for _, candidate := range candidates {
			if getEuclideanDistanceAxis(p, candidate, axis) <= d {
				pointsInRange = append(pointsInRange, candidate)
			}
		}
	}

	return pointsInRange
}

func getClosestPair(P []Point, n int) (Point, Point, float64) {
	if n <= 3 {
		return getClosestByForce(P...)
	}
	mid := n / 2

	var a, b, a1, a2, b1, b2 Point
	var d, dl, dr float64

	a1, b1, dl = getClosestPair(P[:mid], mid) // note that slices arent inclusive on both sides -> [lo:hi)
	a2, b2, dr = getClosestPair(P[mid:], n-mid)

	if dl < dr {
		d = dl
		a, b = a1, b1
	} else {
		d = dr
		a, b = a2, b2
	}

	// TODO : optimize this part
	midX := (P[mid-1].GetAxisValue(0) + P[mid].GetAxisValue(0)) / 2

	var strip_l, strip_r []Point
	for _, point := range P {
		if math.Abs(point.GetAxisValue(0) - midX) < d {
			if point.GetAxisValue(0) < midX {
				strip_l = append(strip_l, point)
			} else {
				strip_r = append(strip_r, point)
			}
		}
	}

	for _, point := range strip_l {
		candidates := point.getPointsInRange(strip_r, d, point.dimension)
		for _, candidate := range candidates {
			dist := getEuclideanDistance(point, candidate)
			if dist < d {
				d = dist
				a, b = point, candidate
			}
		}
	}

	return a, b, d
}

func (s *Solver) Solve() {
	if !s.isInitialized {
		fmt.Println("Solver has not been initialized!")
		return
	}

	euclidOpsCount = 0

	s.solutionPoints[0], s.solutionPoints[1], s.solutionDist = getClosestPair(s.points, len(s.points))
	s.solutionFound = true
}

func (s *Solver) SolveByForce() {
	if !s.isInitialized {
		fmt.Println("Solver has not been initialized!")
		return
	}

	euclidOpsCount = 0

	s.solutionPoints = [2]Point{s.points[0], s.points[1]}
	s.solutionDist = getEuclideanDistance(s.points[0], s.points[1])
	s.solutionFound = true

	for i, _ := range s.points {
		for j, _ := range s.points[i+1:] {
			tempDist := getEuclideanDistance(s.points[i], s.points[j+i+1])

			if s.solutionDist > tempDist {
				s.solutionDist = tempDist
				s.solutionPoints = [2]Point{s.points[i], s.points[j+i+1]}
			}
		}
	}
}

func (s *Solver) Describe() {
	if !s.solutionFound {
		fmt.Println("This solver hasn't found any solution")
		return
	}

	fmt.Println()
	fmt.Println("[Closest Pair]")
	fmt.Println("Dimension:", s.dimension)
	fmt.Println("Point 1")
	fmt.Println("ID:", s.solutionPoints[0].ID)
	fmt.Println("Position:", s.solutionPoints[0].val)
	fmt.Println("Point 2")
	fmt.Println("ID:", s.solutionPoints[1].ID)
	fmt.Println("Position:", s.solutionPoints[1].val)
	fmt.Println()
	fmt.Println("Delta:", s.solutionDist)
	fmt.Println("Number of Operations:", euclidOpsCount)
	fmt.Println()

}

func (s *Solver) getSolutionPoints() [2]Point {
	return s.solutionPoints
}
