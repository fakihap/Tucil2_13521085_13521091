package CPairSolver

type Point struct {
	x int
	y int
}

// Constructor
func NewPoint(_x int, _y int) *Point {
	p := new(Point)
	p.x = _x
	p.y = _y

	return p
}

// Getter
func (p Point) getX() int {
	return p.x
}

func (p Point) getY() int {
	return p.y
}

// Setter
func (p * Point) setX(newX int) {
	p.x = newX;
}

func (p * Point) setY(newY int) int {
	p.y = newY;
}