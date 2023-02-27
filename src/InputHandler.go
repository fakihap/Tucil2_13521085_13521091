package main

import (
	"fmt"
)

func readPoints() []Point {
	var points []Point
	var n int

	fmt.Printf("\nEnter number of points : ")
	fmt.Scan(&n)

	fmt.Printf("\nFor each point, write each coordinates seperated by a space\nExample : 1 2 3\n\n")
	for i := 0; i < n; i++ {
		var x, y, z int
		fmt.Printf("Point %d : ", i+1)
		fmt.Scan(&x, &y, &z)
		points = append(points, *NewPoint(x, y, z))
	}

	return points
}

func printPoints(points []Point) {
	for i, point := range points {
		fmt.Printf("Point %d : (%d, %d, %d)\n", i, point.GetAxisValue(0), point.GetAxisValue(1), point.GetAxisValue(2))
	}
}
