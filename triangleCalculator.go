package main

import (
	"fmt"
	"math"
)

func toRadians(angle float64) float64 {
	return angle * (math.Pi / 180)
}

func calculateSide(length1, length2, angle float64) float64 {
	// using the law of cosines to find the missing side
	// use this...: a2 = b2 + c2 − 2bc cosA
	// where a is the missing side, and A is the angle

	// note that math.Cos uses radians
	// radians := toRadians(angle)  // * _test.go requires the angle to be in radians

	length3 := math.Sqrt((length1 * length1) + (length2 * length2) - (2 * length1 * length2 * math.Cos(angle)))
	//Insert the code here

	//Do not remove this line
	fmt.Println("The 3rd length of the triange is", length3)
	return length3
}

func main() {
	var length1 float64
	var length2 float64
	var angle float64

	fmt.Println("Enter the first length of the triangle: ")
	fmt.Scanln(&length1)
	fmt.Println("Enter the second length of the triangle: ")
	fmt.Scanln(&length2)
	fmt.Println("Enter the angle in degrees between the two lengths: ")
	fmt.Scanln(&angle)

	calculateSide(length1, length2, angle)
}
