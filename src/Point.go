package main

type Point struct {
	x int
	y int
	z int
}

// Constructor
func NewPoint(_x int, _y int, _z int) *Point {
	p := new(Point)
	p.x = _x
	p.y = _y
	p.z = _z

	return p
}

// Getter
func (p Point) getX() int {
	return p.x
}

func (p Point) getY() int {
	return p.y
}

func (p Point) getZ() int {
	return p.z
}

// Setter
func (p * Point) setX(newX int) {
	p.x = newX;
}

func (p * Point) setY(newY int) {
	p.y = newY;
}

func (p * Point) setZ(newZ int) {
	p.z = newZ;
}