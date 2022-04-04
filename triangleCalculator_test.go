package main

import "testing"

func TestTriangleCalculator(t *testing.T) {
	get := calculateSide(10, 10, 45)
	res := 9.74349024921019
	if get != res {
		t.Errorf("CalculateSide() = %f Want %f Got %f", get, res, get)
	}
}
