package main

import "math"

var k, p, v float64 // объявлено в самом тесте

func T() float64 {
	return 6 / W()
}

func W() float64 {
	return math.Sqrt(k / M())
}

func M() float64 {
	return p * v
}
