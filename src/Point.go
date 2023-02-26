package main

var (
	num_of_points int = 0
)

type Point struct {
	ID int
	dimension int
	val []int
}

// Constructor
func NewPoint(dimension int, val []int) *Point {
	p := new(Point)
	p.dimension = dimension
	p.val = val

	num_of_points++
	p.ID = num_of_points

	return p
}