package main

import (
	"fmt"
	"os"
)

func scanInt() (i int) {
	if _, err := fmt.Scan(&i); err != nil {
		fmt.Printf("Invalid input: %s", err)
		os.Exit(1)
	}

	return i
}

func pluralize(num int, one string, three string, five string) string {
	num %= 100

	if num >= 11 && num <= 19 {
		return five
	}

	num %= 10

	if num == 1 {
		return one
	}

	if num > 0 && num <= 4 {
		return three
	}

	return five
}

func main() {
	n := scanInt()

	fmt.Printf("%d %s", n, pluralize(n, "korova", "korovy", "korov"))
}
