package main

var (
	num_of_points int = 0
)

type Point struct {
	ID        int
	dimension int
	val       []int
}

// Constructor
func NewPoint(val ...int) *Point {
	p := new(Point)
	p.dimension = len(val)
	p.val = val

	num_of_points++
	p.ID = num_of_points

	return p
}

func (p Point) GetAxisValue(axis int) int {
	return p.val[axis]
}
