package main

import (
	"fmt"
)

func calculateSide(length1, length2, angle float64) float64 {

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
	fmt.Println("Enter the angle between the two lengths: ")
	fmt.Scanln(&angle)

	calculateSide(length1, length2, angle)
}
