package main

import (
	"fmt"
)

type InputHandler struct{}

func NewInputHandler() *InputHandler {
	return new(InputHandler)
}

func (i InputHandler) readUserConfig() (int, int, float64, float64) {
	var n, dimension int
	var lowerBound, upperBound float64

	fmt.Printf("\nEnter number of points : ")
	fmt.Scan(&n)

	fmt.Printf("Enter dimension : ")
	fmt.Scan(&dimension)

	fmt.Printf("Enter lower bound : ")
	fmt.Scan(&lowerBound)

	fmt.Printf("Enter upper bound : ")
	fmt.Scan(&upperBound)

	return n, dimension, lowerBound, upperBound
}

func (i InputHandler) askToVisualize() bool {
	var visualize string

	fmt.Printf("\nVisualize your result? (y/n) : ")
	fmt.Scan(&visualize)

	for visualize != "y" && visualize != "n" {
		fmt.Printf("\nInvalid input. Visualize? (y/n) : ")
		fmt.Scan(&visualize)
	}

	if visualize == "y" {
		return true
	} else {
		return false
	}
}

func (i InputHandler) PrintLine(msg string) {
	fmt.Println(msg)
}
