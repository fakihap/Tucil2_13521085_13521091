package main

import (
	"fmt"
	"strconv"
)

type InputHandler struct{}

func NewInputHandler() *InputHandler {
	return new(InputHandler)
}

func (i InputHandler) readUserConfig() (int, int, float64, float64) {
	var n, dimension int
	var lowerBound, upperBound float64

	for {
		n = i.GetInt("\nEnter number of points : ")

		if n >= 2 {
			break
		}

		fmt.Println("Masukan harus >= 2 !")
	}

	for {
		dimension = i.GetInt("Enter dimension : ")

		if dimension >= 1 {
			break
		}

		fmt.Println("Masukan harus >= 1 !")
	}

	lowerBound = i.GetFloat64("Enter lower bound : ")

	upperBound = i.GetFloat64("Enter upper bound : ")

	if lowerBound > upperBound {
		fmt.Println("\nNilai batas lower lebih tinggi dari upper!")
		fmt.Println("Menukar nilai...\n")

		lowerBound, upperBound = upperBound, lowerBound
	}

	fmt.Println()

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

func (i InputHandler) GetInt(msg string) int {
	var temp string
	var res int64
	var err error

	for {
		fmt.Printf(msg)
		fmt.Scan(&temp)
		res, err = strconv.ParseInt(temp, 10, 64)

		if err == nil {
			break
		}

		fmt.Println("Masukan salah!")
	}

	return int(res)
}

func (i InputHandler) GetFloat64(msg string) float64 {
	var temp string
	var res float64
	var err error

	for {
		fmt.Printf(msg)
		fmt.Scan(&temp)
		res, err = strconv.ParseFloat(temp, 64)

		if err == nil {
			break
		}

		fmt.Println("Masukan salah!")
	}

	return res
}
