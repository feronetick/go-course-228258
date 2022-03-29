package main

import (
	"fmt"
	"os"
)

func readTask() (interface{}, interface{}, interface{}) // Объявлено в тесте

func itof(i interface{}) (f float64) {
	f, ok := i.(float64)

	if !ok {
		fmt.Printf("value=%v: %T", i, i)
		os.Exit(0)
	}

	return f
}

func operate(a, b float64, op interface{}) (r float64) {
	switch op {
	case "+":
		r = a + b
	case "-":
		r = a - b
	case "*":
		r = a * b
	case "/":
		r = a / b
	default:
		fmt.Printf("неизвестная операция")
		os.Exit(0)
	}

	return r
}

func main() {
	v1, v2, op := readTask()
	fmt.Printf("%.4f", operate(itof(v1), itof(v2), op))
}
