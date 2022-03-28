package main

import "fmt"

func vote(x int, y int, z int) int {
	if x+y+z > 1 {
		return 1
	}

	return 0
}

func main() {
	fmt.Println(vote(0, 0, 1))
}
