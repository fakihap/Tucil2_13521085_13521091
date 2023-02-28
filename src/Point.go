package main

import (
	"math/rand"
	"time"
)

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

func NewRandomPoint(dimension int, lowerBound int, upperBound int) *Point {
	rand.Seed(time.Now().UnixNano())

	p := new(Point)
	p.dimension = dimension

	p.val = make([]int, dimension)

	for i, _ := range p.val {
		// TODO :
		// there is still chance for random generated points to be having same position
		p.val[i] = rand.Intn(upperBound - lowerBound + 1) + lowerBound // [-100, 100]
	}

	num_of_points++
	p.ID = num_of_points

	return p
}

func (p Point) GetAxisValue(axis int) int {
	return p.val[axis]
}

func (p1 Point) Equals(p2 Point) bool {
	return p1.dimension == p2.dimension && p1.ID == p2.ID
}