package main

import (
	"fmt"
	"strconv"
)

type IOHandler struct{}

func NewIOHandler() *IOHandler {
	return new(IOHandler)
}

func (i IOHandler) readUserConfig() (int, int, float64, float64) {
	var n, dimension int
	var lowerBound, upperBound float64

	for {
		n = i.GetInt("Enter number of points : ")

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

func (i IOHandler) askToVisualize() bool {
	var visualize string

	fmt.Printf("\nVisualize your result? (y/n) : ")
	fmt.Scan(&visualize)

	for visualize != "y" && visualize != "n" {
		fmt.Printf("\nInvalid input. Visualize? (y/n) : ")
		fmt.Scan(&visualize)
	}

	if visualize == "y" {
		fmt.Println("\nVisualizing... (Ctrl-C to exit))")
		return true
	} else {
		return false
	}
}

func (i IOHandler) PrintLine(msg string) {
	fmt.Println(msg)
}

func (i IOHandler) GetInt(msg string) int {
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

func (i IOHandler) GetFloat64(msg string) float64 {
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
