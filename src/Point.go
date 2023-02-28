package main

import (
	"math/rand"
)

var (
	num_of_points int = 0
)

type Point struct {
	ID        int
	dimension int
	val       []float64
}

// Constructor
func NewPoint(val ...float64) *Point {
	p := new(Point)
	p.dimension = len(val)
	p.val = val

	num_of_points++
	p.ID = num_of_points

	return p
}

func NewRandomPoint(dimension int, lowerBound float64, upperBound float64) *Point {
	p := new(Point)
	p.dimension = dimension

	p.val = make([]float64, dimension)

	for i, _ := range p.val {
		p.val[i] = rand.Float64()*(upperBound-lowerBound) + lowerBound // [-100, 100]
	}

	num_of_points++
	p.ID = num_of_points

	return p
}

func (p Point) GetAxisValue(axis int) float64 {
	return p.val[axis]
}

func (p1 Point) Equals(p2 Point) bool {
	return p1.dimension == p2.dimension && p1.ID == p2.ID
}
