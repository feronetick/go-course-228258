package main

import "fmt"

func splitNumber(n int) (r []int) {
	for i := n; i > 0; i /= 10 {
		r = append(r, i%10)
	}

	return r
}

func main() {
	fmt.Println(splitNumber(7456853))
}
